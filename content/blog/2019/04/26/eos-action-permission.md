+++
date = "2019-04-26T01:44:37+09:00"
tags = [ "eos", "cpp", "blockchain" ]
title = "EOS における、action 単位の permission 制御の必要性とその実現方法"
+++

EOS 上で複数の contract をまたいだシステムを実装したい場合にぶつかる課題とその解決方法についてのメモ。

<!--more-->

## 課題

横断的（ex. [猫](https://www.cryptokitties.co)を食う[ゾンビ](https://cryptozombies.io)）・縦断的（ex. レイヤー構造化）なニーズの下、別の contract とやりとりする contract を実装したいシーンは往々にして存在する。また、こういった contract は、ユーザーの permission で別の contract の action を実行したい（プロキシのような役割を担いたい）ことが多いと思われる。

このような場合、基本的には下図のような構成になる。

![EOS acction permission 1](/img/entry/eos-action-permission-1.png)

ここで注意する必要があるのは、

__`appaccount11` にデプロイされた contract が `useraccount1` の permission を利用して inline action を実行するためには、利用したい `useraccount1` の permission（ex. `active` permission）に対して、`appaccount11` の `eosio.code` permission が付与されている必要がある__

ということである。これは、裏を返せば、

__`appaccount11` の `eosio.code` permission が付与された `useraccount1` の permission は、`appaccount11` にデプロイされた contract が自由に利用できてしまう__

ということである。

すなわち、この状況下では、`appaccount11` の管理者に `useraccount1` の permission を濫用されてしまう可能性がある。また、`appaccount11` の管理者に悪意がなくとも、デプロイされた contract に脆弱性があった場合の被害拡大に繋がる可能性がある。冒頭で述べたようなニーズを満たすためとはいえ、こういったリスクは極力低減したい。

## 解決方法

上記の問題を解決する 1 つの方法は、

__`appaccount11` の `eosio.code` permission を、それが必要な action の実行時のみ付与し、その action が完了したら即座に外す__

ことである。

このように、action 実行と atomic に permission の着脱を行うことで、上述したリスクを低減することができる。これは、下図のような構成で実現できる。

![EOS acction permission 2](/img/entry/eos-action-permission-2.png)

上図における controller contract のサンプル実装は [こちら](https://github.com/m0t0k1ch1/sandbox/tree/master/eos/action-permission/controller)。実装は非常にシンプルなので、[`controller.cpp`](https://github.com/m0t0k1ch1/sandbox/blob/master/eos/action-permission/controller/controller.cpp) を見るだけで、何をやっているかは把握できるはず。

ユーザーは、この controller contract をあらかじめ `useraccount1` にデプロイしておき、`execute` action を介して別の contract の action を実行すればよい。このとき、`execute` action は、

- `update_auth(auth_before)`：`active` permission に `auth_before` を設定
- `execute_action(acnt, act, data)`：指定した action を実行
- `update_auth(auth_after)`：`active` permission に `auth_after` を設定

というフローで処理を行うため、例えば、`execute` action 実行前の `active` permission が

``` json
{
  "threshold": 1,
  "keys": [{
    "key": "EOS57edFL2dE8sxaVQ6uT7Maizi6bD3zh9moXFjDCA35rCMxNYPyf",
    "weight": 1,
  }],
  "accounts": [{
    "permission": {
      "actor": "useraccount1",
      "permission": "eosio.code"
    },
    "weight": 1
  }],
  "waits": []
}
```

だとすると、`auth_before` を

``` json
{
  "threshold": 1,
  "keys": [{
    "key": "EOS57edFL2dE8sxaVQ6uT7Maizi6bD3zh9moXFjDCA35rCMxNYPyf",
    "weight": 1,
  }],
  "accounts": [{
    "permission": {
      "actor": "useraccount1",
      "permission": "eosio.code"
    },
    "weight": 1
  }, {
    "permission": {
      "actor": "appaccount11",
      "permission": "eosio.code"
    },
    "weight": 1
  }],
  "waits": []
}
```

とし、`auth_after` を

``` json
{
  "threshold": 1,
  "keys": [{
    "key": "EOS57edFL2dE8sxaVQ6uT7Maizi6bD3zh9moXFjDCA35rCMxNYPyf",
    "weight": 1,
  }],
  "accounts": [{
    "permission": {
      "actor": "useraccount1",
      "permission": "eosio.code"
    },
    "weight": 1
  }],
  "waits": []
}
```

とすれば、`appaccount11` の `eosio.code` permission が `useraccount1` の `active` permission に付与された状態をこの action 実行中のみに絞ることができるため、冒頭で述べたようなリスクを低減できる。

もちろん、直接お目当ての action を実行する場合と比べると実行コストが増大していることに注意する必要はある。

## 検証

実際に testnet で検証してみた結果が [こちら](https://kylin.eosx.io/tx/fe1e0fbc4091d8151e53ce1d18ace3b63722150b5afd6e64bccb8961076e0774?listView=traces)。

なお、testnet 上の account と上図に記載された account の対応関係は以下のようになっている。

- `motokichi111`：`useraccount1`
- `proxytest111`：`appaccount11`
- `countertest1`：`appaccount12`

## 補足

上記 controller contract の `execute` action の第 3 引数 `std::vector<char> data` には、実行したい action の引数を適切にエンコードして渡す必要がある。このエンコードを行う方法はいくつかあるが、[cleos](https://developers.eos.io/eosio-cleos/docs) を利用する方法が簡単なので記載しておく。

例えば、`proxytest111` の `increment` action の引数をエンコードしたい場合は以下のようにすればよい。

``` sh
$ cleos --url https://api-kylin.eosasia.one convert pack_action_data proxytest111 increment '{"me":"motokichi111"}'
```

``` txt
1042700d39483395
```
