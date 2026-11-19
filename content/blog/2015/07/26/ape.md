---
authors: ['m0t0k1ch1']
date: '2015-07-26T21:36:12+09:00'
tags: ['go']
title: 'Go で IRC bot をつくりやすくする'
---

Perl で IRC bot を書くとき、[UnazuSan](https://github.com/Songmu/p5-UnazuSan) の `on_command` を使うのに慣れちゃっていたので、Go でも同じ雰囲気で書きたいなあと思って、とりあえず骨組みだけ書いてみた。誰か既に似たようなのを書いていそうな気もする。。

<!--more-->

{{< github "m0t0k1ch1" "ape" >}}

README にも書いた通り、こんな感じで使います。

```go
package main

import (
    "log"
    "strings"

    "github.com/m0t0k1ch1/ape"
)

func main() {
    con := ape.NewConnection("pooh", "pooh")
    con.UseTLS = true
    con.Password = "XXXXX"
    if err := con.Connect("irc.example.com:6667"); err != nil {
        log.Fatal(err)
    }

    con.RegisterChannel("#poyo")

    con.AddAction("piyo", func(e *ape.Event) {
        con.SendMessage("poyo")
    })

    con.AddAction("say", func(e *ape.Event) {
        con.SendMessage(strings.Join(e.Command().Args(), " "))
    })

    con.AddAction("🙏", func(e *ape.Event) {
        con.SendMessage("解脱")
        con.Part(con.Channel())
        con.Join(con.Channel())
        con.SendMessage("輪廻転生")
    })

    con.Loop()
}
```

すると、こんな感じでやりとりできます。

```txt
21:26:57 m0t0k1ch1: pooh: piyo
21:26:57 pooh: poyo
21:27:04 m0t0k1ch1: pooh: say 🙏
21:27:04 pooh: 🙏
21:27:11 m0t0k1ch1: pooh: 🙏
21:27:11 pooh: 解脱
21:27:11 pooh has left ()
21:27:11 pooh has joined (~pooh@example.com)
21:27:12 pooh: 輪廻転生
```

合掌。
