+++
date = "2014-09-08"
tags = [ "golang" ]
title = "ksatriya update log"
+++

前に勢いで書いちゃった [ksatriya](https://github.com/m0t0k1ch1/ksatriya)（[こちら](http://m0t0k1ch1st0ry.com/blog/2014/08/16/ksatriya) をご参照ください）、自分で使うにあたってちょっと機能追加したくなったので追加した。

<!--more-->

## 追加した機能

結果として、

- before / after hook を controller ごとに定義可能
- HTML をレンダリングする際にベースとなるテンプレートを controller ごとに定義可能

になった。

手を動かす前は「これくらい簡単やろ〜」と思っていた。でも、これだけをやるために結構いろんな学びがあったし、実際、コードもかなり改修した。

あと、せっかくなので [https://github.com/unrolled/render](https://github.com/unrolled/render) に頼っていたレンダリング部分も自前で実装してみた。結果、標準パッケージ以外への依存が1つ減って幸せになれた。

## 使い方

現状の ksatriya を使ったサンプルアプリケーションは [ここ](https://github.com/m0t0k1ch1/ksatriya-sample) に置いてあるけれど、今後も仕様は変わっていくと思うので、現状の `ksatriya.Controller` の使い方だけ以下にメモしておこうと思う。後で見なおして思い出に浸ったりしたい。

まず、`ksatriya.Controller` を埋め込んだ `Controller` を定義する

``` go
import (
    "net/http"

    "github.com/m0t0k1ch1/ksatriya"
)

type Controller struct {
    *ksatriya.Controller
}

func NewController() *Controller {
    c := &Controller{ksatriya.NewController()}

    c.AddBeforeFilter(c.Before)   // add before hook
    c.AddAfterFilter(c.After)     // add after hook

    c.GET("/", c.Index)
    c.GET("/user/:name", c.User)

    return c
}

func (c *Controller) Before(ctx *ksatriya.Context) {
    ctx.SetTmplDirPath("app/view")
    ctx.SetBaseTmplPath("layout.html")   // set base template
}

func (c *Controller) After(ctx *ksatriya.Context) {
    ctx.RenderArgs["title"] = "ksatriya-sample"
}

func (c *Controller) Index(ctx *ksatriya.Context) {
    ctx.HTML(http.StatusOK, "index.html", nil)
}

func (c *Controller) User(ctx *ksatriya.Context) {
    name := ctx.Param("name")
    ctx.HTML(http.StatusOK, "user.html", ksatriya.RenderArgs{
        "name": name,
    })
}
```

で、以下のようにする。

``` go
k := ksatriya.New()
k.RegisterController(NewController())
k.Run(":8080")
```

以上。

## 今後

自分がほしいな〜と思っていた最低限の機能は揃った気がするので、そろそろテストを書いていきたい。

{{< tweet 507920947249049600 >}}

「例の Gunosy のやつ」っていうのは [こちら](http://gunosygo.connpass.com/event/8485)。楽しみ。
