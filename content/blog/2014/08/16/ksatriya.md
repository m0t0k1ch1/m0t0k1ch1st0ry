+++
title = "Golang で ksatriya という薄い WAF を書いてみた"
tags = ["golang"]
date = "2014-08-16"
+++

{{< github "m0t0k1ch1" "ksatriya" >}}

気分転換も兼ねて、Golang で薄い WAF を書いてみた。「書いた」というより「つぎはぎした」と表現した方が正しい感じのものですが、そこはご了承ください。

<!--more-->

## 動機

ちょっと前から Golang の WAF を触る機会が増えてきて、[Revel](https://github.com/revel/revel)、[Kocha](https://github.com/naoina/kocha)、[Martini](https://github.com/go-martini/martini)、[beego](https://github.com/astaxie/beego) など、有名どころの WAF は触ってみたけれど、個人で簡単なうぇっぶアプリケーションを書きたくなったときに、「これや！」っていうのがなかった。最近だと [Gin](https://github.com/gin-gonic/gin) が気になっているけれど、v1.0 になるまでにまだやることが結構残っている模様なので、とりあえずは待ち状態。ということで、勉強がてら自分で薄いやつを書いてみるのもおもしろいかもと思い始めた。

## 方針

自分のニーズをものすごく端的に表現すると、こんな感じ。

- ベースは薄い
- でも拡張性はほしい

思想的には [Martini](https://github.com/go-martini/martini) がそれを満たしているんだけど、[Three reasons you should not use Martini](http://stephensearles.com/?p=254) や [My Thoughts on Martini](http://codegangsta.io/blog/2014/05/19/my-thoughts-on-martini) などでも議論されているように、Golang らしくないところがいくつかある。確かに、inject パッケージを使ってなんでもかんでもよしなに handler の引数として受け取れてしまうのは結構気持ち悪かったりした。

そんなこんなで目をつけたのが [Negroni](https://github.com/codegangsta/negroni)。これは [Martini](https://github.com/go-martini/martini) の作者の方が批判へのカウンターとして提案しているライブラリ。[Negroni](https://github.com/codegangsta/negroni) 自体は薄く、middleware を実装することで拡張していく。なので、WAF ではないのだけれど、WAF っぽく機能拡張していくことは可能となっている。これを使って、

- ルーティングとレンダリングを担う薄い WAF をつくって、[Negroni](https://github.com/codegangsta/negroni) の middleware として動かす
- 拡張性はサードパーティの middleware に任せる

ようにすればニーズが満たせるのでは？？と思った。

で、書いたのが [ksatriya](https://github.com/m0t0k1ch1/ksatriya)。

## 勉強になったこと

結果的にルーティングとレンダリングは既存のパッケージに任せているので、自分で書いたところはほとんどない。。唯一、「こんな書き方できるのか〜」と思ったのが以下。

``` go
func (k *Ksatriya) Handle(method, path string, handler HandlerFunc) {
    k.Router.Handle(method, path, func(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
        c := NewContext(w, req, Params{params}, k.Renderer)
        handler(c)
    })
}
```

こう書くことで、ルーターには `httprouter.Handle` として handler を登録しつつも、その中で `ksatriya.Context` をつくって、それを `ksatriya.HandlerFunc` に渡して処理を行うようにできる。こうすることで、独自の context を引数にして handler を定義していくことができるので拡張しやすくなるし、毎回 `func Handler(w http.ResponseWriter, req *http.Request, params httprouter.Params) { ... }` みたいに冗長な引数を書く必要がなくなる。

ま、これも [Gin](https://github.com/gin-gonic/gin) を参考にして書いたんですが笑

## 今後

とりあえず思いつきで書いちゃったので、もうちょっと考えてリファクタリングする。あとは個人でつくろうと思っているものがあるので、それのベースとして使いつつ改善していければいいかなあと思う。

[ksatriya](https://github.com/m0t0k1ch1/ksatriya) 自体のコードは大したことないけれど、書く過程でいろいろなパッケージのソースコード読めたので勉強になった。こんな感じでうまいこと目的をつくって人の書いたコードを読むようにしていきたい。
