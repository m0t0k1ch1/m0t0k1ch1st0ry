+++
date = "2017-12-08T04:45:34+09:00"
tags = [ "ethereum", "truffle", "solidity", "blockchain" ]
title = "truffle で書いたコントラクトを solidity-flattener でがっちゃんこして Etherscan で公開する"
+++

現状、`import` を利用して truffle で書いたコントラクトを [Etherscan](https://etherscan.io) で公開するためには、依存関係にあるコントラクトを全部がっちゃんこしたコードを用意しないといけないのだが、これをつくるのにちょっと手こずったのでメモ。もっとイケてる方法をご存知の方がいらっしゃいましたら何卒ご教示ください。

<!--more-->

## ※追記

[Hi-Ether](https://qiita.com/amachino/items/605ff76209d7193dc92c) コミュニティにて、[truffle-flattener](https://github.com/alcuadrado/truffle-flattener) の存在を教えていただきました。これを使うと、簡単にがっちゃんこしたコードを生成することができます。

``` sh
$ truffle-flattener contracts/AnotherEther.sol > Combined.sol
```

実際に新しいコントラクトを ropsten にデプロイし、がっちゃんこしたコードを登録してみたところ、問題なく検証を通過しました。

https://ropsten.etherscan.io/address/0x1514e87adb657288060d820c6bffa86b70382f6e#code

このとき、solidity-flattener でがっちゃんこした際に発生した「`pragma solidity ^0.4.18;` が `pragma solidity ^0.4.13;` に書き換わってしまう問題」も発生しませんでした。また、後から気づいたのですが、solidity-flattener でがっちゃんこした場合、contract 定義の直上に記載されていたコメントも消えてしまっていたようで。。truffle-flattener の場合は、これも発生しませんでした。総じて、現状では truffle-flattener を使うのがよさそうです。

## 準備

まず、truffle 自体にがっちゃんこする機能はなさそう。[Feature request: Export code for etherscan verification #564](https://github.com/trufflesuite/truffle/issues/564) という issue が立ってはいたが、執筆時点ではまだ未実装。

便利ツールがないか探してみたところ、どうやら [solidity-flattener](https://github.com/BlockCatIO/solidity-flattener) を使えば簡単にがっちゃんこできるらしい。使うには solc が必要とのこと。

しかし、truffle が内部的に利用している [solc-js](https://github.com/ethereum/solc-js) のコマンドラインオプションは solc と互換性がないらしい。[solidity のドキュメント](https://solidity.readthedocs.io/en/latest/installing-solidity.html#npm-node-js) にも以下の記載がある。

> The comandline options of solcjs are not compatible with solc and tools (such as geth) expecting the behaviour of solc will not work with solcjs.

ということで、愚直に solc をインストールすることにした。以下、Ubuntu 16.04 にて。

``` sh
$ sudo add-apt-repository ppa:ethereum/ethereum
$ sudo apt-get update
$ sudo apt-get install solc
```

次に solidity-flattener をインストールする。

``` sh
$ apt-get install python3-pip
$ pip3 install --upgrade pip
$ pip3 install solidity-flattener
```

これで準備完了。

## がっちゃんこする

今回のターゲットはこちら。[既に ropsten にデプロイ済みの ERC20 トークン](https://ropsten.etherscan.io/token/0x480d291d75a79b48bf5da25921a39cbf73fd060b) である。

<div class="github-card" data-user="m0t0k1ch1" data-repo="another-ether"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

時は満ちたり。以下を実行してがっちゃんこする。`--solc-paths` オプションの使い方がポイント。

``` sh
$ solidity_flattener --solc-paths=zeppelin-solidity=$(pwd)/node_modules/zeppelin-solidity/ contracts/AnotherEther.sol --output Combined.sol
```

無事、`Combined.sol` が生成された。

が、ここで 1 つだけ問題が発生した。きちんと原因が調査できていないが、`AnotherEther.sol` では `pragma solidity ^0.4.18;` としていた部分が、`Combined.sol` では `pragma solidity ^0.4.13;` となってしまっていた。これだけ修正。

修正後の `Combined.sol` は以下のようになった。

``` solidity
pragma solidity ^0.4.18;

library SafeMath {
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    if (a == 0) {
      return 0;
    }
    uint256 c = a * b;
    assert(c / a == b);
    return c;
  }

  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    // assert(b > 0); // Solidity automatically throws when dividing by 0
    uint256 c = a / b;
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
    return c;
  }

  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    assert(b <= a);
    return a - b;
  }

  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    assert(c >= a);
    return c;
  }
}

contract ERC20Basic {
  uint256 public totalSupply;
  function balanceOf(address who) public view returns (uint256);
  function transfer(address to, uint256 value) public returns (bool);
  event Transfer(address indexed from, address indexed to, uint256 value);
}

contract ERC20 is ERC20Basic {
  function allowance(address owner, address spender) public view returns (uint256);
  function transferFrom(address from, address to, uint256 value) public returns (bool);
  function approve(address spender, uint256 value) public returns (bool);
  event Approval(address indexed owner, address indexed spender, uint256 value);
}

contract BasicToken is ERC20Basic {
  using SafeMath for uint256;

  mapping(address => uint256) balances;

  /**
  * @dev transfer token for a specified address
  * @param _to The address to transfer to.
  * @param _value The amount to be transferred.
  */
  function transfer(address _to, uint256 _value) public returns (bool) {
    require(_to != address(0));
    require(_value <= balances[msg.sender]);

    // SafeMath.sub will throw if there is not enough balance.
    balances[msg.sender] = balances[msg.sender].sub(_value);
    balances[_to] = balances[_to].add(_value);
    Transfer(msg.sender, _to, _value);
    return true;
  }

  /**
  * @dev Gets the balance of the specified address.
  * @param _owner The address to query the the balance of.
  * @return An uint256 representing the amount owned by the passed address.
  */
  function balanceOf(address _owner) public view returns (uint256 balance) {
    return balances[_owner];
  }

}

contract StandardToken is ERC20, BasicToken {

  mapping (address => mapping (address => uint256)) internal allowed;


  /**
   * @dev Transfer tokens from one address to another
   * @param _from address The address which you want to send tokens from
   * @param _to address The address which you want to transfer to
   * @param _value uint256 the amount of tokens to be transferred
   */
  function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
    require(_to != address(0));
    require(_value <= balances[_from]);
    require(_value <= allowed[_from][msg.sender]);

    balances[_from] = balances[_from].sub(_value);
    balances[_to] = balances[_to].add(_value);
    allowed[_from][msg.sender] = allowed[_from][msg.sender].sub(_value);
    Transfer(_from, _to, _value);
    return true;
  }

  /**
   * @dev Approve the passed address to spend the specified amount of tokens on behalf of msg.sender.
   *
   * Beware that changing an allowance with this method brings the risk that someone may use both the old
   * and the new allowance by unfortunate transaction ordering. One possible solution to mitigate this
   * race condition is to first reduce the spender's allowance to 0 and set the desired value afterwards:
   * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
   * @param _spender The address which will spend the funds.
   * @param _value The amount of tokens to be spent.
   */
  function approve(address _spender, uint256 _value) public returns (bool) {
    allowed[msg.sender][_spender] = _value;
    Approval(msg.sender, _spender, _value);
    return true;
  }

  /**
   * @dev Function to check the amount of tokens that an owner allowed to a spender.
   * @param _owner address The address which owns the funds.
   * @param _spender address The address which will spend the funds.
   * @return A uint256 specifying the amount of tokens still available for the spender.
   */
  function allowance(address _owner, address _spender) public view returns (uint256) {
    return allowed[_owner][_spender];
  }

  /**
   * approve should be called when allowed[_spender] == 0. To increment
   * allowed value is better to use this function to avoid 2 calls (and wait until
   * the first transaction is mined)
   * From MonolithDAO Token.sol
   */
  function increaseApproval(address _spender, uint _addedValue) public returns (bool) {
    allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_addedValue);
    Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

  function decreaseApproval(address _spender, uint _subtractedValue) public returns (bool) {
    uint oldValue = allowed[msg.sender][_spender];
    if (_subtractedValue > oldValue) {
      allowed[msg.sender][_spender] = 0;
    } else {
      allowed[msg.sender][_spender] = oldValue.sub(_subtractedValue);
    }
    Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
    return true;
  }

}

contract AnotherEther is StandardToken {
  string public name = "AnotherEther";
  string public symbol = "AETH";
  uint public decimals = 18;

  function AnotherEther(uint initialSupply) public {
    totalSupply = initialSupply;
    balances[msg.sender] = initialSupply;
  }
}
```

## Etherscan で公開する

ropsten の場合、[ここ](https://ropsten.etherscan.io/verifyContract) で必要事項を入力する。今回は以下のようになった。

- Contract Address：0x480d291d75a79b48bf5da25921a39cbf73fd060b
- Contract Name：AnotherEther
- Compiler：v0.4.18+commit.9cf6e910
- Optimization：Disabled
- Solidity Contract Code：`Combined.sol` の内容をそのままコピペ
- Constructor Arguments ABI-encoded (For contracts that accept constructor parameters)：000000000000000000000000000000000000000000115eec47f6cf7e35000000

ちなみに、constructor には initialSupply として 21,000,000 を渡したので、Constructor Arguments ABI-encoded が上記のようになっている。

検証を無事通過すると、コントラクトコードが公開される。

https://ropsten.etherscan.io/address/0x480d291d75a79b48bf5da25921a39cbf73fd060b#code

めでたしめでたし。
