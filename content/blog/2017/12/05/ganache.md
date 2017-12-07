+++
date = "2017-12-05"
tags = [ "ethereum", "truffle", "ganache" ]
title = "Ganache で始める Ðapp 開発"
+++

この記事は、[Ethereum Advent Calendar 2017](https://qiita.com/advent-calendar/2017/ethereum) の 5 日目の記事です（ちなみに、まだまだ空いている日があります！この記事を読んでいる方の中に「参加してもいいよ！」という方がいらっしゃったら是非ご参加ください！Ethereum 盛り上げましょう〜〜！！）。

<!-- more -->

この記事では、11 月に [Truffle Suite](https://github.com/trufflesuite) の仲間に加わったばかりのローカル開発用ツール [Ganache](http://truffleframework.com/ganache) について紹介し、これを利用した簡単な Ðapp 開発を実践してみようと思います。この記事を読んで、Ganache を利用した Ðapp 開発のイメージを掴んでいただければ幸いかなと思います。

<br />
## Ganache について簡単に紹介

[公式サイト](http://truffleframework.com/ganache) から、紹介文を引用します。

> ONE CLICK BLOCKCHAIN
>
> Quickly fire up a personal Ethereum blockchain which you can use to run tests, execute commands, and inspect state while controlling how the chain operates.

ざっくり訳すと、「ワンクリックブロックチェーン：個人用の Ethereum ブロックチェーンを素早く起動し、チェーンの動作を制御しながら、テストやコマンドを実行したり、状態を調査したりすることができます」といったところでしょうか。

もう少し具体的な情報として、公式サイトで言及されている特徴をざっと列挙してみます。

- Javascript 実装の Ethereum を内蔵（ローカルクライアントのインストール不要）
- アドレス、秘密鍵、トランザクション、残高など、アカウントの状態が閲覧可能
- レスポンスやデバッグ用の情報を含んだログ出力
- 開発時のニーズに合わせてブロックタイムを調整可能
- Block Explorer を内蔵
- Byzantium に対応

GUI ベースで諸々の情報閲覧や簡単なマイニング制御が行えるのはお手軽だなという感じです。インストールも楽ちんなので、起動するまでに嵌りどころはほぼないです。

実際に触ってみた感じとしても、「Ethereum 詳しくないけど気になるな、ちょっと触ってみたいな」という開発者が Ethereum の挙動を把握するにはうってつけなツールかなと思いました。もちろん、[Truffle](http://truffleframework.com) とともに、ローカル環境での本格的なスマートコントラクト開発の強い味方になってくれるとも思います。

<br />
## ダウンロードして起動してみる

まず、[公式サイト](http://truffleframework.com/ganache) から Ganache をダウンロードして起動します。執筆時点でのバージョンは 1.0.1 でした。

起動すると、以下のような感じです。100 ETH 保有しているアカウントが 10 匹登録されていました。

![ganache_1](/my-images/entry/ganache_1.png)

<br />
## JSON-RPC で送金してみる

初期設定だと、JSON-RPC サーバーが 7545 番ポートで起動しています。ここに HTTP 経由でリクエストを送信し、簡単な動作確認をしてみます。

まず、新しくアカウントを登録します。

``` sh
$ curl -X POST http://127.0.0.1:7545 --data '{"jsonrpc":"2.0","method":"personal_newAccount","params":["pass"],"id":0}'
```
``` json
{
  "id": 0,
  "jsonrpc": "2.0",
  "result": "0x05eee23f682718f129719df9d0d0254542c6a10e"
}
```

0x05eee23f682718f129719df9d0d0254542c6a10e というアカウントが登録されました。このアカウントに対し、初期状態で登録されていたアカウントの 1 匹である 0x627306090abaB3A6e1400e9345bC60c78a8BEf57 から 1 ETH を送ってみます。

``` sh
$ curl -X POST http://127.0.0.1:7545 --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from":"0x627306090abaB3A6e1400e9345bC60c78a8BEf57","value":"0x0de0b6b3a7640000","to":"0x05eee23f682718f129719df9d0d0254542c6a10e"}],"id":0}'
```
``` json
{
  "id": 0,
  "jsonrpc": "2.0",
  "result": "0xc838f6cb6881028723afd611d0b0de648115d7746afa5d5ccbae76c7a3ee0d1f"
}
```

初期設定だと AUTOMINE が有効になっており、なんらかのトランザクションが送信されるまでブロックはマイニングされないようですが、トランザクションが送信されると瞬時にマイニングされます（設定画面から AUTOMINE を無効にし、ブロックタイムを指定してマイニングを行うことも可能です）。

マイニングされたブロックや、それに含まれるトランザクションは GUI からも確認することができます。便利ですね。

![ganache_2](/my-images/entry/ganache_2.png)
![ganache_3](/my-images/entry/ganache_3.png)

JSON-RPC で送信先のアカウントの残高も確認してみます。

``` sh
$ curl -X POST http://127.0.0.1:7545 --data '{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x05eee23f682718f129719df9d0d0254542c6a10e", "latest"],"id":1}'
```
``` json
{
  "id": 1,
  "jsonrpc": "2.0",
  "result": "0x0de0b6b3a7640000"
}
```

0x0de0b6b3a7640000 は 1 ETH（1,000,000,000,000,000,000 wei）に相当するので、送金は正常に完了したようです。

<br />
## コントラクトをデプロイしてみる

Advent Calendar の 2 日目の記事「[Truffle で始める Ethereum 入門 - ERC20 トークンを作ってみよう](https://qiita.com/amachino/items/8cf609f6345959ffc450)」で紹介された ERC20 トークンをそのまま Ganache にデプロイしてみようと思います。

実際のコードは GitHub に置いておきましたので、必要であればご参照ください。

<div class="github-card" data-user="m0t0k1ch1" data-repo="ERC20-token-sample"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

[マイグレーションファイルの作成](https://qiita.com/amachino/items/8cf609f6345959ffc450#%E3%83%9E%E3%82%A4%E3%82%B0%E3%83%AC%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%AE%E4%BD%9C%E6%88%90) までは 2 日目の記事と同じですが、次の手順である「コントラクトのデプロイ」が異なります。

コントラクトは Ganache にデプロイする必要がありますので、`truffle.js` を以下のように修正し、development ネットワークとして Ganache を登録します（Ganache の初期設定に合わせています）。

``` js
module.exports = {
  networks: {
    development: {
      host: 'localhost',
      port: 7545,
      network_id: 5777
    }
  }
};
```

準備ができたら、以下を実行してデプロイします。

``` sh
$ truffle migrate --network development
```

<pre>
Using network 'development'.

Running migration: 1_initial_migration.js
  Deploying Migrations...
  ... 0xeb26d390a2848ae97f8166fc9046933bb2d20418872d7b5581d7e51b1790776b
  Migrations: 0xf12b5dd4ead5f743c6baa640b0216200e89b60da
Saving successful migration to network...
  ... 0x5ba32299b5c8f9cdedd35f33d4fe32810050ffad5eabf572cf700ac4f9a35133
Saving artifacts...
Running migration: 2_deploy_my_token.js
  Deploying MyToken...
  ... 0xf8dc0bb97cc48a42b578a7365d493bd5d589a938b8e8584781098abf88dec594
  MyToken: 0xf25186b5081ff5ce73482ad761db0eb0d25abfbf
Saving successful migration to network...
  ... 0x9a573474527bea6504b9900d07c49c2962d5c64aec96780b2f2c0ba9ab6740c1
Saving artifacts...
</pre>

無事デプロイできたようです。MyToken に対応するアドレスは 0xf25186b5081ff5ce73482ad761db0eb0d25abfbf となりました。

実行されたトランザクションは Ganache の GUI からも確認することができます。

![ganache_4](/my-images/entry/ganache_4.png)

<br />
## コントラクトを実行してみる

以下を実行してコンソールを起動します。

``` sh
$ truffle console --network development
```

コンソールが起動したら、2 日目の記事の [デプロイされたコントラクトの確認](https://qiita.com/amachino/items/8cf609f6345959ffc450#%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E3%81%95%E3%82%8C%E3%81%9F%E3%82%B3%E3%83%B3%E3%83%88%E3%83%A9%E3%82%AF%E3%83%88%E3%81%AE%E7%A2%BA%E8%AA%8D) と同様の手順でコントラクトの動作確認を行うことができます。

もちろん、transfer などの各種 function 実行時に発行されたトランザクションも Ganache の GUI から確認することができます。

<br />
## まとめ

- Ethereum のローカル開発用ツール [Ganache](http://truffleframework.com/ganache) を紹介しました
- 簡単な動作確認として、JSON-RPC による送金を行ってみました
- Ganache を利用した Ðapp 開発のイメージを掴んでいただくべく、[Truffle](http://truffleframework.com) を利用したコントラクトのデプロイと動作確認を行ってみました

<br />
## 豆知識

この記事のタイトルでもそうですが、「なぜ Dapp（Decentralized Application）を Ðapp と表記するのだろう？」と疑問に思った方は以下をご覧ください。

[Why do people often write "Ðapp" instead of "Dapp" ?](https://www.reddit.com/r/ethereum/comments/5blnhv/why_do_people_often_write_%C3%B0app_instead_of_dapp)

最もポイントの高いコメントから、その理由を引用しておきます。cool なだけじゃないんやで。

> it's because that letter is called ETH in greek. Technically Ðapp can be read as Ethapp.
> Also it looks cool.

<br />
## 参考

- [Truffle で始める Ethereum 入門 - ERC20 トークンを作ってみよう](https://qiita.com/amachino/items/8cf609f6345959ffc450)
- [Ethereum のローカル開発環境 Ganache を使ってみる](https://qiita.com/kyrieleison/items/8ef926faa4defa8fe930)
