+++
title = 'Scalatra + Slick + Scalate でつくるサンプルうぇっぶアプリケーション'
tags = ['scala', 'scalatra', 'slick']
date = '2013-11-16'
+++

現状、Scala でカジュアルにうぇっぶアプリケーションをつくるにはこの構成が一番キマってる感じがするので、サンプルうぇっぶアプリケーションをつくってみました。3つともイニシャルが S という点も素晴らしいですね。

<!--more-->

## モチベーション

[ISUCON3](http://isucon.net) のお題アプリの Scala 移植を Scalatra + Slick でこそこそやっているのですが、諸々ハマりまくって全然進まないので、先に Scalatra と Slick の基本的な使い方を把握した方がよいのでは？？と今更ながら思いました（頭悪い）。きちんと動くサンプルアプリケーションが1つできれば今後の自分と世界のためになりますしね！！

あと、たまには表も書きたいよってことで、Scalatra にデフォルトで組み込まれてる Scalate に Furatto を組み合わせて使ってみました。

## 使ったものについてのさくっとした説明

### Scalatra（framework）

- 公式は [こちら](http://www.scalatra.org)
- Scala 版 Sinatra
- 前に書いた [Scalatra を Jetty で standalone deploy してみる]({{< ref "/blog/2013/10/14/scalatra.md" >}}) で雰囲気伝わるかと

### Slick（ORM）

- 公式は [こちら](http://slick.typesafe.com)
- ORM というか、modern database query and access library
- `JOIN` とか `WHERE` が Scala のコレクションを扱う感覚で書けてすごくモダン
- 神エントリー：[Slick 1.0.0 Documentationを翻訳した](http://qiita.com/krrrr38/items/488ffc49a01cca8425f8)

### Scalate（template engine）

- 公式は [こちら](http://scalate.fusesource.org)
- Scalatra にデフォルトで組み込まれてる
- いろんなテンプレートをサポートしてる
- 今回は HTML に Scala を普通に埋め込んでいける感じで書ける [SSP](http://scalate.fusesource.org/documentation/ssp-reference.html) を採用

### Furatto（CSS framework）

- 公式は [こちら](http://icalialabs.github.io/furatto)
- もう Bootstrap にも飽きてきたので
- Bootstrap ほどお節介でなく、結構使いやすい印象

## ソースコード

{{< github "m0t0k1ch1" "scalatra-slick" >}}

## どんなうぇっぶアプリケーションなんです？？

トレーナーを登録して、好きなポケモンを捕まえたり逃したりできるうぇっぶアプリケーションです。なんと、お気に入りのポケモンを決めることも可能です。

デモとか詳しい説明は割愛します。routing を定義してある `src/main/scala/com/github/m0t0k1ch1/slick/Slick.scala` を見れば雰囲気伝わると思います。

DB 周りついてですが、schema は `src/main/scala/com/github/m0t0k1ch1/slick/schema` 以下に、schema にマッピングする model は `src/main/scala/com/github/m0t0k1ch1/slick/model` 以下に置きました。なお、これらをどこに置くべきかの明確な指針を見つけられなかったので、適当です。悪しからず。

## 雑感

### Slick がナウい

新感覚でした。コードを書いている際、DB にアクセスしてごにゃごにゃして…みたいな感覚が薄く、Scala のコードとの親和性が高くつくられている印象を受けました。例えば以下のような感じ。

``` scala
val trainerPokemons = for {
  tp <- TrainerPokemons if tp.trainerId === trainerId
  p  <- Pokemons if p.id === tp.pokemonId
} yield (p, tp.id, tp.isFavorite)
```

ただ、JOIN だけは裏でそれなりにゴツいクエリが生成されているようなので、これはまた別のエントリーでまとめたいと思っています。

### Scala について

今はモダンな Perl を書いている（つもり）とはいえ、PHP からプログラミングを始めた自分にとっては、静的型付けかつ純粋なオブジェクト指向の言語はとても勉強になります。考え方の幅がばりばり広げられる感覚です。

Scala らしいスタイルでコードを書くよう努力していますが、まだ `var` に溢れたミュータブルで副作用のある世界に逃げこみたくなります。が、バナージを見習って逃げずに引き続き闘い続けます。押忍。
