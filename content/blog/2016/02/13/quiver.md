+++
title = 'Quiver でメモ管理改善'
tags = ['others']
date = '2016-02-13T19:11:50+09:00'
+++

[Quiver: The Programmer's Notebook](https://itunes.apple.com/jp/app/quiver-programmers-notebook/id866773894?mt=12)

Twitter のタイムラインで [こちらのエントリ](http://namaraii.com/archives/39626) を見かけて、これはよさそうと思って衝動買いしてちょっと試してみたところ、自分のニーズを全て満たしていると言っても過言ではないくらいよかった。ので、この感動を同じニーズを持った人達に伝えるべく、軽くレポートしてみようと思う。

<!--more-->

{{< figure src="/img/entry/quiver.png" >}}

## Quiver 移行前のメモ管理

- Dropbox に markdown 形式で保存
- emacs で編集
- [Marked 2](http://marked2app.com) を emacs から起動して閲覧

という感じで落ち着いていた。

が、いまいちしっくりきてなかったというか。メモ管理に最適化されてる感じは薄かった。シェルの世界から出ないでよいのでラクだなあというくらい。

## Quiver の特徴

最たる特徴は以下。

- 1つの Note は複数の Cell で構成されている
- 現状、Cell は以下の5種類
  - Text Cell
  - Code Cell
  - Markdown Cell
  - LaTex Cell
  - Diagram Cell

その他の特徴・機能については以下で自分のニーズと対応させながら。

また、以下には書いていないが、プレゼンテーションモードもあるらしい。機会があれば是非使ってみたい。

## 自分のニーズと Quiver

メモ帳に対する自分のニーズとそれを実現している Quiver のスペックを対応させて書いていこうと思う。同じようなニーズを持っている方は是非使ってみてほしい。

### markdown で書ける

言うまでもなく Markdown Cell を使えばできる。

### コードが美しく syntax highlight される

Markdown Cell のコードブロックや、Code Cell を使うと美しく syntax highlight される。theme も30種類くらいあって、素晴らしい。ちなみに、上の画像は Tomorrow Night Bright。

### メモをカテゴライズできる

Notebook を使うことで Note をカテゴライズできる。また Note には Tag をつけることができ、Tag ベースでの絞り込みなども可能。

### 必要な情報にすぐアクセスできる

キーワードで全文検索ができるし、上記の通り、そもそも Notebook や Tag で Note を管理できるので、ごちゃごちゃにならない。Note のソートもできる。

### データを複数デバイスで共有できる

クラウド上に Library を保存するだけ。自分は Dropbox 上に置いた。

### 普段慣れ親しんだキーバインディングで使いたい

Preferences >> Cells から、各 Cell に対してキーバインディングを設定できるようになっており、なんと、emacs と vim のキーバインディングが使える。これにはびっくり。全て emacs のキーバインディングにして使っております。

### 自分が使わない機能はいらない

必要最小限、シンプルであってほしい。贅沢言ってるのは承知の上。が、個人的な感覚として、Quiver には不必要な設定項目・拡張機能がほとんどなく、Preferences がとてもシンプルでわかりやすかった。

## まとめ

Quiver、すごくいい。現状のメモ管理に納得できていない方（特にエンジニア）は是非。
