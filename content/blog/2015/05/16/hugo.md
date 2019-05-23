+++
title = "Hugo ＋ wercker ＋ GitHub Pages でブログ管理"
tags = [ "hugo", "wercker" ]
date = "2015-05-15T23:18:39+09:00"
+++

[Octopress](http://octopress.org) ＋ [GitHub Pages](https://pages.github.com) で管理してた当ブログを [Hugo](http://gohugo.io) ＋ [wercker](http://wercker.com) ＋ [GitHub Pages](https://pages.github.com/) で管理するようにしましたメモ。最近プライベートで新しいもの触れてなかったので結構楽しかった。

<!--more-->

## 移行後のソース

特に隠すものとかないので全部公開してます。`config.toml` や `wercker.yml` が同じことしようとしてる方の助けになれば幸いです。

{{< github "m0t0k1ch1" "m0t0k1ch1st0ry" >}}

## Hugo に移行した理由

他のエントリでも言われてることがほとんどだけど、一応。

### 記事生成が速い

評判通り速い。もう Octopress には戻れない。

``` sh
$ time hugo
0 draft content
0 future content
52 pages created
8 paginator pages created
17 tags created
0 categories created
in 305 ms
hugo  0.24s user 0.09s system 87% cpu 0.383 total
```

### 構成がシンプル

基本的にはこんな感じ。わかりやすい。

``` sh
$ tree -L 1
.
├── archetypes
├── config.toml
├── content
├── data
├── layouts
├── public
├── static
└── themes
```

自分は `content/blog` 以下に記事を配置して、`static/my-images` 以下に画像を配置してる（`static/images` だと theme が使ってるディレクトリ名と被っちゃうので）。

### Golang で書かれてる

いざとなればソースコード読んでがんばれる。

## 移行作業

既にいろんな人がエントリ書いてるし公式のドキュメントも十分揃ってるので、あんまり詳しくは書かない。

### 記事をもってくる

`content` 以下にぶちこむ。

### 記事フォーマットを調整する

まずはみんながやってるメタデータの日付フォーマット変換。

``` sh
$ find . -type f | xargs sed -i '' -e 's/date: \([0-9]\{4\}-[0-9]\{2\}-[0-9]\{2\}\).*$/date: \1/g'
```

あとは、これと同じ手法を使って地味にいろいろ記事フォーマットを調整した。具体的には、

- 記事のメタデータ部分を調整
  - 要らない要素を削除
  - YAML 形式から TOML 形式（Hugo のデフォルト）にがんばって置換（`MetaDataFormat` を指定することはできるので、別に TOML 形式にこだわる必要はないけど、郷に入っては郷に従え思想なので）
- `<!-- more -->` だと automatic summary split が動かなかったので `more` の左右のスペースを削除

などなど。

#### 移行前のメタデータ部分

``` txt
---
layout: post
title: "..."
description: "..."
date: YYYY-MM-DD hh:mm
comments: true
categories: [ category1, ... ]
---
```

#### 移行後のメタデータ部分

``` txt
+++
date = "YYYY-MM-DD"
tags = [ "tag1", ... ]
title = "..."
+++
```

### theme を調整する

theme を何にするかは悩みに悩んだ。これだ！！！みたいなのがなくて心が折れそうになったけど、[vienna](https://github.com/keichi/vienna) を fork してちょこちょこいじることでなんとかなった。

手を動かしたのは、記事のリスト表示でページャー使われてなかったので使うようにした（1ページ7記事）くらい。この theme でいくぞ！！！って決めてからは速かった。

あと、ローカルで build すると

``` txt
ERROR: 2015/05/15 Site's .BaseUrl is deprecated and will be removed in Hugo 0.15. Use .BaseURL instead.
```

こんなエラーがいくつか出たのでなおしたかったけど、なおしちゃうと wercker で [hugo-build](https://app.wercker.com/#applications/54a7744c6b3ba8733de4dcde/tab/details) がコケて死ぬので、なおさずにそのままにしておいた。

## wercker で build ＋ deploy

Octopress 時代は `bundle exec rake gen_deploy` で build ＋ deploy してたけど、Hugo 自体にはそういう仕組みはないので、wercker を使ってみることにした。

[公式で紹介されてるやり方](http://gohugo.io/tutorials/automated-deployments) の通りにセットアップすると、GitHub への push にフックして自動的に build ＋ deploy が走るようになる。便利。


## 結果

エントリを公開する手順は以下のような感じになった。

### 新しいエントリの雛形をつくる

``` sh
$ hugo new poyo.md
```

で、中身を書く。

### ローカルで確認

`http://127.0.0.1:1313` でプレビューできる。ファイル更新すると勝手にリロードまでしてくれる。便利。

``` sh
$ hugo server --watch
```

### build ＋ deploy

``` sh
$ git add .
$ git commit -m 'add poyo.md'
$ git push origin master
```

あとは wercker がよしなにしてくれる。便利。
