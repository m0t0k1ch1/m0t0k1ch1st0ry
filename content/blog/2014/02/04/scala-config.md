+++
title = "Scala で config ファイルを使う"
tags = ["scala"]
date = "2014-02-04"
+++

最近ブログ更新が滞っていてよろしくないので、地味なことでもアウトプットしていこうと思います。

<!--more-->

表題の件について、[Scalatra](http://www.scalatra.org) で Twitter の API を使った地味なアプリケーションをつくっていて、consumer key とか access token とかはどこに書けばよいのでしょう〜と思って調べた結果、とりあえず [com.typesafe.conf](http://search.maven.org/#artifactdetails%7Ccom.typesafe%7Cconfig%7C1.2.0%7Cbundle) を使うことにしました。また忘れた頃に使いそうなので、使い方をメモ。

## 使い方

`src/main/resources/application.conf` に

``` scala
twitter {
    consumerKey       = "**********"
    consumerSecret    = "**********"
    accessToken       = "**********"
    accessTokenSecret = "**********"
}
```

って感じで書いて、

``` scala
import com.typesafe.config.ConfigFactory

val conf        = ConfigFactory.load
val consumerKey = conf.getString("twitter.consumerKey")
```

的な感じで使います。デフォルトで `application.conf` を見に行ってくれます。  
シンプルですね！！！！！

久々に Scala 書いたら感覚が失われ過ぎててつらいので、動くものつくって動かさなきゃなと思います。
