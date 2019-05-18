+++
date = "2013-10-14"
tags = [ "scala", "scalatra" ]
title = "Scalatra を Jetty で standalone deploy してみる"
+++

しばらくの間、ポケモンとともにカロス地方に旅立つことになりそうなので、帰って来れなくなる前に途中書きだった [Scalatra](http://www.scalatra.org) の話をまとめました。

<!--more-->

ちなみに、、  
今作のポケモンにはフレンドサファリなる神施設が存在するようなので、フレンドがつがつ増やしたい所存です。基本どなたでも welcome ですので、「フレンドなってあげるよ！」という心優しいお方は [@m0t0k1ch1](https://twitter.com/m0t0k1ch1) まで気軽にご連絡ください〜。

m0t0k1ch1 のフレンドコードは 4441-9522-5868 です。

それでは、気をとりなおさずに本題に移ります。

## Scalatra って？

公式には↓こう書いてあります

> Scalatra is a simple, accessible and free web micro-framework.

Scala 版 [Sinatra](http://www.sinatrarb.com) 的なやつです。軽量感が漂ってます。

先日 [Play](http://www.playframework-ja.org/) を触ってみて、個人で何かこそこそつくるのには若干重たい感じがしたので、こちらを触ってみることにしました。

## 準備

- [Installation | Scalatra](http://www.scalatra.org/2.2/getting-started/installation.html) を参考にして、Scala とか Java とか Conscript とか giter8 とかをインストールする
  - [Conscript](https://github.com/n8han/conscript)：Scala で書かれたプログラムをインストールしたりアップデートしたりするツール
  - [giter8](https://github.com/n8han/giter8)： GitHub 上で公開されているテンプレートからアプリケーションの雛形を作成するツール

## とりあえず動かす

ここも [First steps | Scalatra](http://www.scalatra.org/2.2/getting-started/first-project.html) を見ればだいたいわかっちゃいます…が、ここからは備忘録の意味も込めてやったことを書いていきます。

- まず g8 コマンドでアプリケーションの雛形を作成

``` sh
$ g8 scalatra/scalatra-sbt
```

上記のコマンドを実行するといろいろ聞かれますが、何を入力すればよいかは [First steps | Scalatra](http://www.scalatra.org/2.2/getting-started/first-project.html) にちゃんと書いてあります。とりあえず、今回は以下のように入力したとして話を進めます。

``` txt
organization [com.example]:       com.k1ch1
package [com.example.app]:        com.k1ch1.app
name [My Scalatra Web App]:       k1ch1
scalatra_version [2.2.1]:         （そのまま）
servlet_name [MyScalatraServlet]: K1ch1
scala_version [2.10.2]:           2.10.1
version [0.1.0-SNAPSHOT]:         （そのまま）
```

- ビルドする
  - `./sbt` すると、ごごごっといろいろダウンロードされるので、ちょっとびっくりする
  - `container:start` すると、8080 番ポートでアプリケーションが立ち上がる
  - `~;copy-resources;aux-compile` しておくと、ファイルをいじったときに自動でアプリケーションを再起動してくれる

``` sh
$ cd k1ch1
$ chmod u+x sbt
$ ./sbt
```

``` txt
> container:start
> ~;copy-resources;aux-compile
```

- ブラウザで見てみる

``` txt
> browse
```

見れた！！

## Jetty で standalone deployment

deploy できなきゃ意味ない！ってことで deploy してみます。

- Jetty って？？？
  - 軽量な Java の Web サーバーらしい
  - WebSocket などもサポートしている模様
  - 個人的には懐かしくてたまらない某猫氏と双璧を成す存在らしい

[Standalone deployment | Deployment | Scalatra](http://www.scalatra.org/2.2/guides/deployment/standalone.html) を見てもいまいちよくわかりませんでしたが、[scalatra で作ったプロジェクトを stand alone で使用できるように](http://takuya71.hatenablog.com/entry/2013/06/22/180808) を見てなんとかなりました。

### JettyLauncher.scala を作成

``` scala
package com.k1ch1.app  // remember this package in the sbt project definition
import org.eclipse.jetty.server.Server
import org.eclipse.jetty.servlet.{DefaultServlet, ServletContextHandler}
import org.eclipse.jetty.webapp.WebAppContext
import org.scalatra.servlet.ScalatraListener

object JettyLauncher { // this is my entry object as specified in sbt project definition
  def main(args: Array[String]) {
    val port = if(System.getenv("PORT") != null) System.getenv("PORT").toInt else 8080

    val server = new Server(port)
    val context = new WebAppContext()
    context setContextPath "/"
    context.setResourceBase("src/main/webapp")
    context.addEventListener(new ScalatraListener)
    context.addServlet(classOf[DefaultServlet], "/")

    server.setHandler(context)

    server.start
    server.join
  }
}
```

### build.scala を一部変更

- 変更前

``` scala
"org.eclipse.jetty" % "jetty-webapp" % "8.1.8.v20121106" % "container",
```

- 変更後

``` scala
"org.eclipse.jetty" % "jetty-webapp" % "8.1.8.v20121106" % "container;compile",
```

### plugins.sbt に以下を追加

``` scala
addSbtPlugin("com.eed3si9n" % "sbt-assembly" % "0.9.0")
```

### build.sbt を作成

``` scala
import AssemblyKeys._

import sbtassembly.Plugin._

seq(assemblySettings: _*)

test in assembly := {}

mergeStrategy in assembly <<= (mergeStrategy in assembly) { (old) =>
  {
    case PathList("META-INF", xs @ _*) => MergeStrategy.discard
    case _ => MergeStrategy.first
  }
}
```

### jar ファイルを作成

- 下記を実行すると、`target/scala-2.10` 以下に `k1ch1-assembly-0.1.0-SNAPSHOT.jar` というのができる

``` sh
$ ./sbt clean assembly
```

### 実行

- 下記を実行すると、8080 番ポートでアプリケーションが立ち上がるので、ブラウザで見てみる

``` sh
$ java -jar target/scala-2.10/k1ch1-assembly-0.1.0-SNAPSHOT.jar
```

見れた！！

## まとめ

- まとめてみると意外に簡単だったけど、何が起こってるかは全然理解していない
- ので、もっと触っていきます
- 次は [Slick](http://slick.typesafe.com) を組み込んでみる予定
