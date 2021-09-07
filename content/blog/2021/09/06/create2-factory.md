+++
title = '複数 EVM チェーンでコントラクトのアドレスを揃えるためにやっておくとよいこと'
tags = ['ethereum', 'solidity', 'blockchain']
date = '2021-09-06T15:25:18+09:00'
+++

複数 EVM チェーンで複数コントラクトを展開するようなプロジェクトの場合、各 EVM チェーンで以下のような準備をしておくとよいと思っていますというメモです。

<!-- more -->

1. nonce: 0 のまっさらな EOA を用意する
2. 各チェーンにおいて、上記 EOA から nonce: 0 の tx で以下のような factory コントラクトをデプロイしておく

```solidity
pragma solidity ^0.8.7;

contract Factory {
    event Created(address indexed addr);

    function create(
        bytes memory code,
        uint256 salt,
        bytes calldata data
    ) external returns (address) {
        address addr;

        assembly {
            addr := create2(0, add(code, 0x20), mload(code), salt)
            if iszero(extcodesize(addr)) {
                revert(0, 0)
            }
        }

        if (data.length > 0) {
            (bool success, ) = addr.call(data);
            if (!success) {
                assembly {
                    returndatacopy(0, 0, returndatasize())
                    revert(0, returndatasize())
                }
            }
        }

        emit Created(addr);

        return addr;
    }
}
```

同一アカウントから同一 nonce の tx でデプロイすれば、まず、複数チェーンでこの factory コントラクトのアドレスを揃えることができます。あとは、このコントラクトの `create` を同一の `code` と `salt` で叩けば、CREATE2 によってコントラクト（`code`）がデプロイされるので、そのコントラクトのアドレスも複数チェーンで揃えることができます。

1 コントラクトデプロイすれば ok みたいなプロジェクトの場合は CREATE2 を使わずに（factory コントラクトのアドレスを揃えたのと同じように）デプロイ元アカウントとその nonce を揃えたらよいかなと思うんですが、冒頭に記載したように複数コントラクトをデプロイしていくようなプロジェクトの場合、nonce をぴったり合わせながらアカウント運用するのはミスりそうですし、かと言ってデプロイ元アカウントを毎回切り替えて nonce: 0 で揃えたりするのも面倒そうと思うので、このような factory コントラクトを 1 つ立てておくと便利かなと思います。

また、第三引数の `data` による初期化処理パートは別にあってもなくてもよいのですが、例えば、初期 owner に `msg.sender` を設定するような一般的な ownable なコントラクトを `create` した場合、初期 owner が（いわゆる `transferOwnership` を call できない）factory コントラクトになると詰んでしまうので、そういったケースをカバーするために一応入れています。[こういう感じ](https://github.com/m0t0k1ch1/hardhat-sample/blob/dfb8447c6aabaab6515ddf56f4452ee1b98d8aba/test/Factory.test.ts) で使ってください。
