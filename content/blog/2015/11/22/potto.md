+++
title = 'Slack の Outgoing WebHook をハンドリングするくん'
tags = ['go']
date = '2015-11-22T17:32:45+09:00'
+++

Go で API サーバー立てる練習として、Slack の Outgoing WebHook を捕まえる対話式の bot をつくっていて、その bot 固有の機能でない部分は切り離せるわ〜〜〜と思って、えいやっと書いた。

{{< github "m0t0k1ch1" "potto" >}}

<!--more-->

README にもある通り、以下のような感じで使える。前につくった IRC 用の [ape](https://github.com/m0t0k1ch1/ape) に似たようなインターフェース。

```go
package main

import (
    "strings"

    "github.com/m0t0k1ch1/potto"
)

func Ping(ctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
    return potto.NewResponse("pong"), nil
}

func Say(ctx potto.Ctx, args potto.ActionArgs) (*potto.Response, error) {
    text := strings.Join(args, " ")
    return potto.NewResponse(text), nil
}

func main() {
    p := potto.New()
    p.AddAction("ping", Ping)
    p.AddAction("say", Say)
    p.Run(":8080")
}
```

実際に対話すると、以下のような感じ。

{{< figure src="/img/entry/potto.png" >}}

かわいい。

一応、拙作の [ksatriya](https://github.com/m0t0k1ch1/ksatriya) をベースにしている。ksatriya の context をどうやって拡張するかだいぶ悶々と悩んでいたけど、ksatriya 側も若干更新して、ここが妥協点か。。。というところには落ち着いた気がする。

ksatriya も、これ以上は仕様レベルでアップデートすることないなあという感じになってきたので、テスト書いていきたい（遅い）。と思ったら、

{{< tweet 668300561863675904 >}}

という気持ちになったので、HTML のレンダリング機能とはさようならした。合掌。
