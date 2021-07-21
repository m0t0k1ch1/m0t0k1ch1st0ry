+++
title = 'native meta transaction が扱える ERC20 トークンの簡易実装'
tags = ['ethereum', 'solidity', 'blockchain']
date = '2019-03-09T18:02:51+09:00'
+++

ここ最近話題になっている [ERC1776](https://github.com/ethereum/EIPs/issues/1776) で標準化されようとしている native meta transaction について把握するため、native meta transaction が扱える ERC20 トークンの簡易実装までやってみましたメモ。

<!--more-->

## meta transaction とは？

端的に言うと、ETH（gas）を保有していなくても Ethereum 上で transaction を発行できるようにするための仕組みです。これを実現する際、一般的には、利用するエンドユーザーが meta transaction 対応したコントラクトアカウントを保有している必要がありますが、[ERC1776](https://github.com/ethereum/EIPs/issues/1776) は、EOA しか保有していないエンドユーザーでも meta transaction を実行できるようにするための標準規格です。これが標準化されて普及すると、EOA しか保有していないエンドユーザーであっても transaction 手数料を支払うことなくコントラクトを実行することが可能となるため、DApps の利用ハードルがグッと下がります。

ただ、1 つ注意点があります。native meta transaction の場合、エンドユーザーが保有しているアカウントは EOA なので、コントラクトアカウントを活用した meta transaction のように、エンドユーザーが保有するアカウントを経由して transaction を実行することはできません。つまり、native meta transaction を処理するコントラクトが外部コールをしたとしても、その実行先であるコントラクトの `msg.sender` が EOA のアドレスになることはないので、`msg.sender` を用いて認証を行うようなコントラクトの実行を目的とした場合、役に立ちません。まあ、EOA ベースなので当たり前と言えば当たり前なのですが。

より詳細に知りたい方は、[ERC1776](https://github.com/ethereum/EIPs/issues/1776) や、そこに記載されているリンクを辿るとよいかなと思います。また、meta transaction 自体は数年前から議論されているアイデアなので、ググると色々情報は出てきます。

## 実装

[ERC1776](https://github.com/ethereum/EIPs/issues/1776) は複数の ERC が絡んでいて少し複雑かつまだドラフト段階なので、native meta transaction の基本原理を把握したいだけの人（数日前の自分）が軽い気持ちで首を突っ込むと、それなりに骨が折れます。ということで、今回は native meta transaction の基本原理の把握に特化して、表題の通りのものを実装してみました（[ERC1776](https://github.com/ethereum/EIPs/issues/1776) に準拠しているわけではないのでご注意を）。数日前の自分のような方の手助けとなれば幸いです。

今回実装した諸々は [こちら](https://github.com/m0t0k1ch1/sandbox/tree/master/ethereum/native-meta-transfer) に置いておきましたが、そんなに量はないので、コントラクトとテストをここにも記載しておきます。説明するよりもソースコードを読んでもらった方が理解が捗ると思います。なお、実装には [truffle](https://github.com/trufflesuite/truffle) を利用しています。

`frm` が実行したい操作（`frm` から `to` への MT 譲渡）を行う transaction を `relayer` がブロードキャストして gas を負担する代わりに `frm` から MT を徴収している、辺りがポイントかなと思います。なお、今回は 1 contract 内で完結する固定の 1 操作（`transfer`）だけが実行できるような実装ですが、ここをより汎用的に実装することは可能です。というか、どちらかと言うと本来はそうあるべきでしょう。

```solidity
pragma solidity >=0.4.21 <0.6.0;

import "openzeppelin-solidity/contracts/cryptography/ECDSA.sol";
import "openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";

contract MetaToken is ERC20, ERC20Detailed {
  using ECDSA for bytes32;

  mapping (address => uint256) private _nonces;

  constructor(uint256 supply) ERC20Detailed("MetaToken", "MT", 18) public {
    _mint(msg.sender, supply);
  }

  function nonceOf(address owner) public view returns (uint256) {
    return _nonces[owner];
  }

  function metaTransfer(
      address frm,
      address to,
      uint256 amount,
      uint256 fee,
      uint256 nonce,
      address relayer,
      bytes memory sig
  ) public returns (bool) {
    require(msg.sender == relayer, "wrong relayer");
    require(nonceOf(frm) == nonce, "invalid nonce");
    require(balanceOf(frm) >= amount.add(fee), "insufficient balance");

    bytes32 hash = metaTransferHash(frm, to, amount, fee, nonce, relayer);
    address signer = hash.toEthSignedMessageHash().recover(sig);
    require(signer == frm, "signer != frm");

    _transfer(frm, to, amount);
    _transfer(frm, relayer, fee);
    _nonces[frm]++;

    return true;
  }

  function metaTransferHash(
      address frm,
      address to,
      uint256 amount,
      uint256 fee,
      uint256 nonce,
      address relayer
  ) public view returns (bytes32) {
    return keccak256(
        abi.encodePacked(
            address(this),
            "metaTransfer",
            frm,
            to,
            amount,
            fee,
            nonce,
            relayer
        )
    );
  }
}
```

```js
const MetaToken = artifacts.require("MetaToken");
const BN = web3.utils.BN;

contract("MetaToken", async (accounts) => {
  it("transfer", async () => {
    let metaToken = await MetaToken.deployed();

    let frm = accounts[0];
    let to = accounts[1];
    let amount = new BN("100");

    let balanceOfFromBefore = await metaToken.balanceOf(frm);
    let balanceOfToBefore = await metaToken.balanceOf(to);

    await metaToken.transfer(to, amount, { from: frm });

    let balanceOfFromAfter = await metaToken.balanceOf(frm);
    let balanceOfToAfter = await metaToken.balanceOf(to);

    assert.isTrue(balanceOfFromBefore.sub(amount).eq(balanceOfFromAfter));
    assert.isTrue(balanceOfToBefore.add(amount).eq(balanceOfToAfter));
  });

  it("metaTransfer", async () => {
    let metaToken = await MetaToken.deployed();

    let frm = accounts[0];
    let to = accounts[1];
    let amount = new BN("100");
    let fee = new BN("1");
    let nonce = new BN("0");
    let relayer = accounts[2];

    let hash = await metaToken.metaTransferHash(
      frm,
      to,
      amount,
      fee,
      nonce,
      relayer
    );
    let sig = await web3.eth.sign(hash, frm);

    let balanceOfFromBefore = await metaToken.balanceOf(frm);
    let balanceOfToBefore = await metaToken.balanceOf(to);
    let balanceOfRelayerBefore = await metaToken.balanceOf(relayer);

    await metaToken.metaTransfer(frm, to, amount, fee, nonce, relayer, sig, {
      from: relayer,
    });

    let balanceOfFromAfter = await metaToken.balanceOf(frm);
    let balanceOfToAfter = await metaToken.balanceOf(to);
    let balanceOfRelayerAfter = await metaToken.balanceOf(relayer);

    assert.isTrue(
      balanceOfFromBefore.sub(amount).sub(fee).eq(balanceOfFromAfter)
    );
    assert.isTrue(balanceOfToBefore.add(amount).eq(balanceOfToAfter));
    assert.isTrue(balanceOfRelayerBefore.add(fee).eq(balanceOfRelayerAfter));
  });
});
```

## 補足

繰り返しになりますが、今回の実装はかなり単純化されたものです。[ERC1776](https://github.com/ethereum/EIPs/issues/1776) を見てもらえばわかると思いますが、実稼働を想定する場合は考えるべきことが増えます。例えば以下などです。

- [ERC223](https://github.com/ethereum/EIPs/issues/223) や [ERC777](https://github.com/ethereum/EIPs/issues/777) の場合、received hook 機構を考慮した実装
- 柔軟なトークン手数料計算
  - transaction 実行時のトークン価格を加味した計算
  - 実際に消費した gas の量を加味した計算
- クライアント（ex. ウォレット）が扱いやすいインターフェース

この辺りのことを踏まえながら [ERC1776](https://github.com/ethereum/EIPs/issues/1776) がどのような仕様に着地するのか、楽しみです。
