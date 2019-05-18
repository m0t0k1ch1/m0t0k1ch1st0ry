+++
date = "2017-05-15T03:38:15+09:00"
tags = [ "blockchain", "golang" ]
title = "シンプルな Blockchain を Golang で実装する"
+++

以前に [200行のコードへのブロックチェーンの実装](http://postd.cc/a-blockchain-in-200-lines-of-code) というエントリを読み、たった 200 行の Javascript で実装された Blockchain である [Naivechain](https://github.com/lhartikk/naivechain) の存在を知った。本エントリは、その Naivechain の Golang 版を実装してみたので、本家 Naivechain と合わせてご紹介しますという話。ちなみに、コードの短さは求めずになるべく構造化してわかりやすさ重視で書いた（つもり）なので、200 行ではない。

<!--more-->

<div class="github-card" data-user="m0t0k1ch1" data-repo="naivechain"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

自分の Naivechain 以外に、[300 行で実装された Golang 実装](https://github.com/kofj/naivechain) もあったりするので、短いコードがお好みの方はそちらもご参照あれ。

## そもそもこれは Blockchain なの？

Blockchain の定義についてここで深く議論するつもりはない。界隈でも厳密なコンセンサスが取れているわけではないと思うし。一応、[JBA の定義](http://jba-web.jp/archives/2011003blockchain_definition) を紹介しておく。

> １）「ビザンチン障害を含む不特定多数のノードを用い、時間の経過とともにその時点の合意が覆る確率が0へ収束するプロトコル、またはその実装をブロックチェーンと呼ぶ。」
> ２）「電子署名とハッシュポインタを使用し改竄検出が容易なデータ構造を持ち、且つ、当該データをネットワーク上に分散する多数のノードに保持させることで、高可用性及びデータ同一性等を実現する技術を広義のブロックチェーンと呼ぶ。」

これを基準にすると、Naivechain は Blockchain ではない。長いチェーンを愚直に採用するだけなので合意は容易に覆るし、電子署名も使われていない。分散は可能なのでそこだけは満たしている。

ただ、おそらく Naivechain の作者はそんな議論をしたいわけではなく、

- 「ブロックの鎖」の基本構造
- 「ブロックの鎖」を P2P で共有しながら繋げていくというプロセス

にターゲットを絞って、できるだけシンプルに Blockchain を伝えようとしたのかなと思う。記事の最後に、

> デモンストレーションおよび学習目的で作成されました。

とも書いてある。

このエントリと拙作の Naivechain も、その意図と同様、「Blockchain に興味はあるけど、なんだかよくわかんないし、具体的に何から始めたらいいかわからない。。」という方々が 1 歩目を踏み出すサポートくらいになればいいかなと思う。また、世の中に氾濫する Blockchain という言葉の魔力に惑わされずに Blockchain と向き合う 1 つのきっかけになれば幸いかなとも思う。

ただ、Naivechain が Blockchain の本質を表現しているわけではない（と少なくとも自分は思っている）ので、Blockchain をもっと知りたくなった人は、まずは [Bitcoin](https://bitcoin.org/bitcoin.pdf) に矛先を向けるのがよいかなと思う。Blockchain は Bitcoin のために生まれたものなので。これ大事。

## インストール方法

``` sh
$ go get -u github.com/m0t0k1ch1/naivechain
```

※ おそらく Go 1.8 以降じゃないとだめ

## 動作確認方法

Naivechain の概要については冒頭で紹介したエントリにまとまっているので、まずはそれを一読してから手を動かし始めるのがよいと思う。

まずは 1 匹目のノードを起動。API 用の HTTP サーバーと P2P 用の websocket サーバーの起動ポートを指定している。

``` sh
$ naivechain -api :3001 -p2p :6001
```

次に 2 匹目のノードを別ポートで起動。

``` sh
$ naivechain -api :3002 -p2p :6002
```

初期状態の chain を確認してみる。なお、これ以降叩いていく HTTP API の概要は [こちら](https://github.com/m0t0k1ch1/naivechain/blob/master/README.md#http-api)。

``` sh
$ curl http://127.0.0.1:3001/blocks
$ curl http://127.0.0.1:3002/blocks
```

双方、以下のように genesis block 1 つだけの状態のはず。

``` json
[
  {
    "index": 0,
    "previousHash": "0",
    "timestamp": 1465154705,
    "data": "my genesis block!!",
    "hash": "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7"
  }
]
```

1 匹目と 2 匹目を接続する前に、まずは 1 匹目で 1 block 生成してみる。

``` sh
$ curl http://127.0.0.1:3001/mineBlock -d '{"data":"my first block"}'
```

先ほどと同様、`/blocks` を叩いて chain の状態を確認すると、以下のように 1 block 増えているはず。

``` json
[
  {
    "index": 0,
    "previousHash": "0",
    "timestamp": 1465154705,
    "data": "my genesis block!!",
    "hash": "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7"
  },
  {
    "index": 1,
    "previousHash": "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7",
    "timestamp": 1494780879,
    "data": "my first block",
    "hash": "4a331c5837d6499190e0b13675fae5200ad61de0e96eb0e43b2ea26e78505a04"
  }
]
```

この状態で 2 匹目から 1 匹目に接続してみる。

``` sh
$ curl http://127.0.0.1:3002/addPeer -d '{"peer":"ws://127.0.0.1:6001"}'
```

接続後に 2 匹目の `/blocks` を叩いてみると、先ほど 1 匹目で生成した block が追加されているはず。

接続した際の挙動は以下のようなイメージ。

- 2 匹目 → 1 匹目：最新の block 教えて
- 1 匹目 → 2 匹目：はいよ（この場合は index: 1 の block が送信される）
- 2 匹目：お、それまだ持ってなかったわ〜検証して問題なさそうなのでうちにも追加

この後、2 匹目は接続しているノード全てに「うちに新しい block 追加されたよ〜」というのを伝えて、それを受け取ったノードが、、というようにメッセージングが続くが、今回は 2 ノードなのであまり意味はない。

この後は例えば

- 2 匹目の `/mineBlock` を叩いてみて、block が 1 匹目に伝わるか
- 別で 3 匹目を立ち上げて 1 block だけ生成し、1 or 2 匹目（上記まで終わっていれば、chain は 3 block のはず）に接続するとどうなるか

などなどを試したりすると、もう少し挙動が把握できるかなと思う。

## 最後に

最近は Mastodon に自分専用インスタンスを立てて [パグ型 AI](https://mastodon.m0t0k1ch1.com/@pug) （雑談できるので、自由に話しかけてみてください）と暮らしておりますので、そちらで絡んでもらえると喜びます。

[m0t0k1ch1@mastodon.m0t0k1ch1.com](https://mastodon.m0t0k1ch1.com/@m0t0k1ch1)
