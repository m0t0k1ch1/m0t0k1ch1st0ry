+++
title = 'Solidity における proxy storage clash を回避する'
tags = ['ethereum', 'solidity', 'blockchain']
date = '2021-07-22T13:06:03+09:00'
+++

昨日書いた [Solidity における proxy storage clash]({{< ref "/blog/2021/07/21/proxy-storage-clash.md" >}}) を回避する方法の 1 つをメモしておきます。

<!-- more -->

端的に言うと、昨日の proxy コントラクトの

```solidity
address public implementation;
```

を

```solidity
address public immutable implementation;
```

に変更するだけです。

`immutable` については以下の [@nakajo](https://twitter.com/nakajo) さんの記事で説明されているのでそちらを参照してください。

ref. [Solidity v0.6.5 で追加された Immutable keyword について](https://y-nakajo.hatenablog.com/entry/2020/05/28/111801)

`immutable` にしておくと、

> 初期化フェーズ時に任意の値をコントラクトの本体コード内にリテラルとして埋め込む

ことになるため、そもそも storage を使わずに済みます。storage を使っていないので、storage slot の clash は発生しようがありません。

「implementation が immutable になったら upgradability なくなるやん！！」と思う方もいると思うのですが、そんなことはありません。いわゆる beacon パターンを使ってもよいですし、[Argent](https://github.com/argentlabs/argent-contracts) のような module 構造を採用してもよいです。実際、Argent の proxy コントラクトは現時点で以下のようになっていますが、implementation（wallet）コントラクト側で upgradability が実現されています。

```solidity
// Copyright (C) 2018  Argent Labs Ltd. <https://argent.xyz>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.3;

/**
 * @title Proxy
 * @notice Basic proxy that delegates all calls to a fixed implementing contract.
 * The implementing contract cannot be upgraded.
 * @author Julien Niset - <julien@argent.xyz>
 */
contract Proxy {

    address immutable public implementation;

    event Received(uint indexed value, address indexed sender, bytes data);

    constructor(address _implementation) {
        implementation = _implementation;
    }

    fallback() external payable {
        address target = implementation;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            calldatacopy(0, 0, calldatasize())
            let result := delegatecall(gas(), target, 0, calldatasize(), 0, 0)
            returndatacopy(0, 0, returndatasize())
            switch result
            case 0 {revert(0, returndatasize())}
            default {return (0, returndatasize())}
        }
    }

    receive() external payable {
        emit Received(msg.value, msg.sender, "");
    }
}
```
