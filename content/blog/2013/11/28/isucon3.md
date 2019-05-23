+++
title = "ISUCON3 本戦のお題アプリを Scala に移植しました"
tags = ["isucon", "scala", "scalatra", "slick"]
date = "2013-11-28"
+++

あのアツい闘いから早3週間。。底辺スタッフとして運営に関わらせていただいた m0t0k1ch1 です（[ISUCON3 底辺スタッフの記録]({{< ref "/blog/2013/11/10/isucon3.md" >}})）。勉強目的でこそこそ進めていた本戦お題アプリの Scala 移植がとりあえず終わったので、学んだことを簡単にまとめるついでにソースコードを公開したいと思います。

<!--more-->

…とはいえ現状、

- [テスト](https://github.com/kayac/isucon3/blob/master/final/webapp/perl/t/01_webapp.t) 通った
- ブラウザで一通り正常に動作した

このような段階でして、まだベンチはかけておりません。。悪しからず。。

勿論、ベンチかけてスコア出してごにゃごにゃみたいなエントリも熱が冷めないうちに書きたいなと思っています。ベンチかけてから全部まとめて書けや！クソが！という声も聞こえてきそうですが、[予選お題アプリの scala 移植情報](https://dl.dropboxusercontent.com/u/261418/scala_at_isucon3/index.html) も公開されておりますので、出る幕が無くなる前に一旦アウトプットしておこうと思います。。

## 使ったもの

- Scalatra（framework）
- Slick（ORM）

Scalatra ？ Slick ？という方は [こちらのエントリ]({{< ref "/blog/2013/11/16/scalatra.md" >}}) をご参照ください。

## ソースコード

{{< github "m0t0k1ch1" "isucon3-scala" >}}

`src/main/scala/com/github/m0t0k1ch1/isucon/Isucon.scala` にほとんどの実装が書いてあります。

## 学び

### Option と match による null ハンドリング

[ScalaのOptionステキさについてアツく語ってみる](http://yuroyoro.hatenablog.com/entry/20100710/1278763193) に書かれていることを実際にコード書いて体感した感じです。非常に勉強になりました。

簡単にまとめると、Java では `null` を扱うにあたって

- `null` チェックをいちいちやらないといけない
- `null` チェックを怠ると実行時に `NullPointerException`（通称：ぬるぽ）に苛まれる
- `null` チェックの不備をコンパイル時に見つけることができない

以上のような注意点があるのですが、Scala では「あー、これ `null` かもなー」っていうやつを `Option` でくるんで `match` でハンドリングすることによってコンパイル時に `null`（実際には `None`）チェック的なことを行うことができるので、積極的にやっていきましょうという話です。

例として、各 API の頭で行われる「`api_key` から `user` を取得する filter 的な処理」の部分を抜粋して載せておきます。

``` scala
def getUserContainer: Option[User] = {
  db withSession {
    val apiKey = Option(request.getHeader("X-API-KEY")) match {
      case Some(v) => Some(v)
      case None    => cookies.get("api_key")
    }
    val userContainer = apiKey match {
      case Some(v) => Query(Users).filter(_.apiKey === v).firstOption
      case None    => None
    }
    userContainer
  }
}

def getUser: User = {
  val userContainer = getUserContainer
  if (userContainer.isEmpty) halt(400)
  userContainer.get
}
```

### 関数型の恩恵

関数型言語である Scala において、関数は第一級オブジェクトです。オブジェクトなので、変数に入れたり関数の引数として渡したりすることができます。

この性質についてですが、今回の Scala 移植を通じて、個人的には以下のような恩恵があるのかなと感じました。

- 処理を細分化しやすく、細分化した結果、コードの見通しが良くなる
- 関数は副作用がないように書くのが基本なので、各関数の独立性が高く、テストしやすい

また、Scala の `if` は Perl などで言うところの三項演算子に近いので、Scala では `if` をがつがつネストして命令型のスタイルでコードを書いていくっていうのがものすごくやりにくいんです。このような性質からも、処理を細分化して独立性の高い関数に落としこんでいかざるを得ない感じが出ているなあと思いました。

## 課題

### GET /timeline

`get("/timeline")` の中で呼んでいるこいつ。

``` scala
def getTimeline(userId: Int, latestEntryContainer: Option[String]): List[Entry] = {
  val end = new Timestamp(now.getTime + timeout * 1000)

  var entries: List[Entry] = Nil

  val loop = new Breaks
  loop.breakable {
    while (now.before(end)) {
      entries = latestEntryContainer match {
        case Some(v) => getLatestEntriesAgain(userId, v.toInt)
        case None    => getLatestEntriesFirstTime(userId)
      }
      if (entries.nonEmpty) loop.break
      Process("sleep ${interval}") !
    }
  }

  entries
}
```

ここは非常につらい実装になっております。まず、`var` 使ってますし、Scala らしからぬ `while` と `break` も登場しております。加えて `Process("sleep ...") !` などというものまで。。絶対にもっと Scala らしく書ける気がしているのですが、、未熟者にはまだその道は見えぬようです。

こういった無作法も許してくれている Scala の慈悲深さに感謝しつつ、引き続きコップ本で精進します。

## ※追記

課題として挙げた `getTimeline` について、再帰を用いた書き方をご提案いただきましたので記載させていただきます。なお、GitHub にあげているソースコードもこちらに置換しています。

``` scala
def getTimeline(userId: Int, latestEntryContainer: Option[String]): List[Entry] = {
  val end = new Timestamp(now.getTime + timeout * 1000)

  @annotation.tailrec
  def loop: List[Entry] = {
    if (now.before(end)) {
      val entries = latestEntryContainer match {
        case Some(v) => getLatestEntriesAgain(userId, v.toInt)
        case None    => getLatestEntriesFirstTime(userId)
      }
      if (entries.nonEmpty) {
        entries
      } else {
        Thread.sleep(interval * 1000)
        loop
      }
    } else Nil
  }

  loop
}
```

## 最後に

当方、Scala は完全に独学でやっておりますゆえ、「そこ！全然イケてないで！！」的なご指摘を心よりお待ちしております。
