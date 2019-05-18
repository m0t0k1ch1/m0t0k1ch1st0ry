+++
date = "2013-11-13"
tags = [ "scala", "scalatra", "slick" ]
title = "Scalatra のアクションで Int が返らなくてはまった"
+++

表題の通りです。[Scalatra](http://www.scalatra.org) で [Slick](http://slick.typesafe.com) を使っててはまりました。

<!--more-->

## はまりポイント

``` scala
get("/") {
  db withSession {
    Pokemons.insert(303, "Mawile")
  }
}
```

^ みたいな、Slick を使って insert するだけの簡単なお仕事をするアクションをブラウザから叩くと、データは挿入されるのにページのロードが全然終わらない。。[公式](http://www.scalatra.org/2.2/guides/persistence/slick.html) にサンプルとして載ってるレベルなのに。。つらい。

## insert の返り値ってなんなの？？

当然ここを疑います。  
[こちらの記事](http://xerial.org/scala-cookbook/recipes/2013/02/01/reflection) を参考にして xerial-lens で返り値の型を出力してみると、`Int` の 1 でした。

## どういうこと？？

公式の [Actions](http://www.scalatra.org/2.2/guides/http/actions.html) の Default behavior によると、`Int` は `Any` に該当するはずで、その `Any` はどうなるかというと、、

> For any other value, if the content type is not set, it is set to text/plain. The value is converted to a string and written to the response's writer.

とあります。なので、`String` に変換されると思っていたのですが、それが間違いでした。

Scalatra の [ソースコード](https://github.com/scalatra/scalatra) をたどってみると、現時点（2.3 系）の `scalatra/core/src/main/scala/org/scalatra/ScalatraBase.scala` の `renderPipeline` に以下のようなコードが。。

``` scala
case status: Int => response.status = ResponseStatus(status)
```

HTTP のレスポンスコードに変換されてるや〜ん。すっきり〜。  
でも仕様としてどうなんでしょこれ？？
