+++
date = "2015-11-22T17:32:45+09:00"
tags = [ "golang" ]
title = "Slack の Outgoing WebHook をハンドリングするくん"
+++

Golang で API サーバー立てる練習として、Slack の Outgoing WebHook を捕まえる対話式の bot をつくっていて、その bot 固有の機能でない部分は切り離せるわ〜〜〜と思って、えいやっと書いた。

<div class="github-card" data-user="m0t0k1ch1" data-repo="potto"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

<!--more-->

README にもある通り、以下のような感じで使える。前につくった IRC 用の [ape](https://github.com/m0t0k1ch1/ape) に似たようなインターフェース。

``` go
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

![potto](/my-images/entry/potto.png)

かわいい。

一応、拙作の [ksatriya](https://github.com/m0t0k1ch1/ksatriya) をベースにしている。ksatriya の context をどうやって拡張するかだいぶ悶々と悩んでいたけど、ksatriya 側も若干更新して、ここが妥協点か。。。というところには落ち着いた気がする。

ksatriya も、これ以上は仕様レベルでアップデートすることないなあという感じになってきたので、テスト書いていきたい（遅い）。と思ったら、

<blockquote class="twitter-tweet" data-partner="tweetdeck"><p lang="ja" dir="ltr">ksatriya、きちんとテスト書くかあと思ったけど HTML 吐いてるとこめんどくさそう。。。と思ったし、golang で HTML 吐くのつらそうなので text と json だけ吐くようにしよ</p>&mdash; m0t0k1ch1 (@m0t0k1ch1) <a href="https://twitter.com/m0t0k1ch1/status/668300561863675904">November 22, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

という気持ちになったので、HTML のレンダリング機能とはさようならした。合掌。

