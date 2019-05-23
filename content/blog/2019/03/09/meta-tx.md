+++
title = 'meta transaction が扱える ERC20 トークンの簡易実装'
tags = ['ethereum', 'solidity', 'blockchain']
date = '2019-03-09T18:02:51+09:00'
+++

ここ最近話題になっている [ERC1776](https://github.com/ethereum/EIPs/issues/1776) で標準化されようとしている meta transaction について把握するため、meta transaction が扱える ERC20 トークンの簡易実装までやってみましたメモ。

<!--more-->

## meta transaction とは？

端的に言うと、ETH（gas）を保有していなくても Ethereum 上で transaction を発行できるようにするための仕組みです。この仕組みが [ERC1776](https://github.com/ethereum/EIPs/issues/1776) によって標準化されて普及すると、ユーザーが Dapps とやりとりするために必須なのは秘密鍵だけとなるので、Dapps の利用ハードルがグッと下がります。また、[ERC1776](https://github.com/ethereum/EIPs/issues/1776) はトークンに関する meta transaction インターフェースの標準規格なので、例えば DEX などはその恩恵を大きく受けることになると予想されます。

より詳細に知りたい方は、[ERC1776](https://github.com/ethereum/EIPs/issues/1776) や、そこに記載されているリンクを辿るとよいかなと思います。また、meta transaction 自体は数年前から議論されているアイデアなので、ググると色々情報は出てきます。

## 実装

[ERC1776](https://github.com/ethereum/EIPs/issues/1776) は複数の ERC が絡んでいて少し複雑かつまだドラフト段階なので、meta transaction の基本原理を把握したいだけの人（数日前の自分）が軽い気持ちで首を突っ込むと、それなりに骨が折れます。ということで、今回は meta transaction の基本原理の把握に特化して、表題の通りのものを実装してみました（[ERC1776](https://github.com/ethereum/EIPs/issues/1776) に準拠しているわけではないのでご注意を）。数日前の自分のような方の手助けとなれば幸いです。

今回実装した諸々は [こちら](https://github.com/m0t0k1ch1/sandbox/tree/master/ethereum/meta-tx) に置いておきましたが、そんなに量はないので、コントラクトとテストをここにも記載しておきます。説明するよりもソースコードを読んでもらった方が理解が捗ると思います。なお、実装には [truffle](https://github.com/trufflesuite/truffle) を利用しています。

`frm` が実行したい操作（`frm` から `to` への MT 譲渡）を行う transaction を `relayer` がブロードキャストして gas を負担する代わりに `frm` から MT を徴収している、辺りがポイントかなと思います。なお、今回は 1 contract 内で完結する固定の 1 操作だけが実行できるような実装ですが、外部コールを実行するような実装にすることも可能です。というか、どちらかと言うと本来はそうあるべきでしょう。

``` solidity
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

``` js
const MetaToken = artifacts.require('MetaToken');
const BN        = web3.utils.BN;

contract('MetaToken', async (accounts) => {
  it('transfer', async () => {
    let metaToken = await MetaToken.deployed();

    let frm    = accounts[0];
    let to     = accounts[1];
    let amount = new BN('100');

    let balanceOfFromBefore = await metaToken.balanceOf(frm);
    let balanceOfToBefore   = await metaToken.balanceOf(to);

    await metaToken.transfer(to, amount, {from: frm});

    let balanceOfFromAfter = await metaToken.balanceOf(frm);
    let balanceOfToAfter   = await metaToken.balanceOf(to);

    assert.isTrue(balanceOfFromBefore.sub(amount).eq(balanceOfFromAfter));
    assert.isTrue(balanceOfToBefore.add(amount).eq(balanceOfToAfter));
  });

  it('metaTransfer', async () => {
    let metaToken = await MetaToken.deployed();

    let frm     = accounts[0];
    let to      = accounts[1];
    let amount  = new BN('100');
    let fee     = new BN('1');
    let nonce   = new BN('0');
    let relayer = accounts[2];

    let hash = await metaToken.metaTransferHash(frm, to, amount, fee, nonce, relayer);
    let sig  = await web3.eth.sign(hash, frm);

    let balanceOfFromBefore    = await metaToken.balanceOf(frm);
    let balanceOfToBefore      = await metaToken.balanceOf(to);
    let balanceOfRelayerBefore = await metaToken.balanceOf(relayer);

    await metaToken.metaTransfer(frm, to, amount, fee, nonce, relayer, sig, {from: relayer});

    let balanceOfFromAfter    = await metaToken.balanceOf(frm);
    let balanceOfToAfter      = await metaToken.balanceOf(to);
    let balanceOfRelayerAfter = await metaToken.balanceOf(relayer);

    assert.isTrue(balanceOfFromBefore.sub(amount).sub(fee).eq(balanceOfFromAfter));
    assert.isTrue(balanceOfToBefore.add(amount).eq(balanceOfToAfter));
    assert.isTrue(balanceOfRelayerBefore.add(fee).eq(balanceOfRelayerAfter));
  });
});
```

## 補足

繰り返しになりますが、今回の実装はかなり単純化されたものです。[ERC1776](https://github.com/ethereum/EIPs/issues/1776) を見てもらえばわかると思いますが、実稼働を想定する場合は考えるべきことが増えます。例えば以下などです。

- [ERC223](https://github.com/ethereum/EIPs/issues/223) や [ERC777](https://github.com/ethereum/EIPs/issues/777) のような fallback 機構を加味した外部コール
- 柔軟なトークン手数料計算
  - transaction 実行時のトークン価格を加味した計算
  - 実際に消費した gas の量を加味した計算
- クライアント（ex. ウォレット）が扱いやすいインターフェース

この辺りのことを踏まえながら [ERC1776](https://github.com/ethereum/EIPs/issues/1776) がどのような仕様に着地するのか、楽しみです。
