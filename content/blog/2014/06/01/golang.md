+++
title = 'クロージャにメソッドを実装してみる - Go'
tags = ['go']
date = '2014-06-01'
+++

そろそろスピリチュアルなエントリから脱却して Go について書こうと思います。最近仕事で少し触っていたり、昨日「[Go Conference 2014 spring](http://connpass.com/event/6370) 」にも参加してきたりで Go に対するモチベーションが上がっているので、少しずつアウトプットしていこうかと。

<!--more-->

今回はクロージャ周りで「お。。そうか〜こんなこともできるんか〜」っと思ったことについてメモ。

## 実際にやってみる

- クロージャもカスタム定義型にキャストしたりすればメソッド（インターフェース）を実装することができる

``` go
package main

import (
    "fmt"
)

type MyFunction func() string

func (f MyFunction) String() string {
    return "poyo"
}

func main() {
    function := func() string {
        return "piyo"
    }
    fmt.Println(function)   // 0x22d0
    fmt.Println(function()) // piyo

    myFunction := MyFunction(function)
    fmt.Println(myFunction)   // poyo
    fmt.Println(myFunction()) // piyo
}
```

- 上記の例ではクロージャを `fmt.Stringer` インターフェースを実装した `MyFunction` にキャストすることで、`fmt.Println` で出力できるようにしている
- こんなパターンどこで使うの…？と思ったりもするけれど、net/http の `http.HandleFunc` がそれっぽいことをしている

``` go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "piyo!")
    })
    http.ListenAndServe(":4000", nil)
}
```

- 上記の例では `http.HandleFunc` の第二引数としてクロージャを渡しているが、これはさらに `http.ServeMux.HandlerFunc` に渡されて `http.HandlerFunc` にキャストされる
- `http.HandlerFunc` は `ServeHTTP` というメソッドを実装しているので、渡したクロージャは結果的に `ServeHTTP` を実装したことになる

## まとめ

絶対使いこなせないけど、知ってるといいことあってほしい
