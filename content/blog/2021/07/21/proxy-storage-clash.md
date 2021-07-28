+++
title = 'Solidity における proxy storage clash'
tags = ['ethereum', 'solidity', 'blockchain']
date = '2021-07-21T21:29:15+09:00'
+++

スマートコントラクトの upgradability を実現するためによく使われるいわゆる proxy パターンを実装する際に必ず回避しないといけない proxy storage clash を実際に起こしてみましょうの回。

<!-- more -->

まず、適当なコントラクトを用意します。

```solidity
pragma solidity ^0.8.6;

contract Account {
    address public owner;

    function setOwner(address newOwner) external {
        owner = newOwner;
    }
}
```

次に、適当な proxy コントラクトを用意します。

```solidity
pragma solidity ^0.8.6;

contract Proxy {
    address public implementation;

    constructor(address impl) {
        implementation = impl;
    }

    fallback() external payable {
        _delegate(implementation);
    }

    receive() external payable {
        _delegate(implementation);
    }

    function _delegate(address impl) internal {
        assembly {
            calldatacopy(0, 0, calldatasize())
            let result := delegatecall(gas(), impl, 0, calldatasize(), 0, 0)
            returndatacopy(0, 0, returndatasize())
            switch result
            case 0 {
                revert(0, returndatasize())
            }
            default {
                return(0, returndatasize())
            }
        }
    }
}
```

挙動を検証するために [@nomiclabs/hardhat-waffle](https://www.npmjs.com/package/@nomiclabs/hardhat-waffle) を使ってテストコードを書いてみます。

```js
const { expect } = require("chai");

describe("Proxy", () => {
  let Account;
  let account;

  let Proxy;
  let proxy;

  let owner;
  let other;

  beforeEach(async () => {
    Account = await ethers.getContractFactory("Account");
    account = await Account.deploy();
    await account.deployed();

    Proxy = await ethers.getContractFactory("Proxy");
    proxy = await Proxy.deploy(account.address);
    await proxy.deployed();

    [owner, other] = await ethers.getSigners();
  });

  it("storage clash", async () => {
    expect(await proxy.implementation()).to.equal(account.address);

    await owner.sendTransaction({
      to: proxy.address,
      data: Account.interface.encodeFunctionData("setOwner", [other.address]),
    });

    expect(await proxy.implementation()).to.equal(other.address);
  });
});
```

このテストを例えば

```sh
$ npx hardhat test
```

などで実行すると、問題なく通ります。

が、このテストが通ったということは、delegatecall された `setOwner` の引数である `newOwner` によって proxy コントラクトの `implementation` が上書きされた、すなわち proxy storage clash が発生したということです。相当ヤバいです。

一般的な回避方法が知りたい方は [The State of Smart Contract Upgrades](https://blog.openzeppelin.com/the-state-of-smart-contract-upgrades) を参照してください。スマートコントラクトの upgradability に関する基礎がまとまった良記事です。
