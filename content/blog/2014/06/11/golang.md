+++
title = 'nil だと思ってやつが nil じゃなかった - Go'
tags = ['go']
date = '2014-06-11'
+++

Go の `nil` で完全に嵌ったのでメモ。

<!--more-->

## 嵌りポイント

例えば、[Revel](http://revel.github.io) の `validator.go` の中にこんな関数がいる。

```go
func (r Required) IsSatisfied(obj interface{}) bool {
    if obj == nil {
        return false
    }

    if str, ok := obj.(string); ok {
        return len(str) > 0
    }
    if b, ok := obj.(bool); ok {
        return b
    }
    if i, ok := obj.(int); ok {
        return i != 0
    }
    if t, ok := obj.(time.Time); ok {
        return !t.IsZero()
    }
    v := reflect.ValueOf(obj)
    if v.Kind() == reflect.Slice {
        return v.Len() > 0
    }
    return true
}
```

よくある validation なんだけど、これで完全に嵌った。わかりやすいように、上記から一部抜粋してサンプルコードを書いてみる。

```go
package main

import (
    "fmt"
)

func IsSatisfied(obj interface{}) bool {
    if obj == nil {
        return false
    }
    return true
}

func main() {
    fmt.Println(IsSatisfied(nil)) // false

    var x *int = nil
    fmt.Println(IsSatisfied(x)) // true
}
```

`IsSatisfied(x)` は `true`。ん…？  
`x` は `nil` じゃ、、ない？？なんでや〜〜〜と 1 時間くらい悶えていた。

## 検証してみる

```go
package main

func IsNil(obj interface{}) bool {
    if obj == nil {
        return true
    }
    return false
}

func main() {
    println(IsNil(nil)) // true

    var x *int = nil
    println(IsNil(x)) // false
}
```

なるほど。

```go
package main

import (
    "reflect"
)

func IsNil(obj interface{}) bool {
    if obj == nil || reflect.ValueOf(obj).IsNil() {
        return true
    }
    return false
}

func main() {
    println(IsNil(nil)) // true

    var x *int = nil
    println(IsNil(x)) // true
}
```

なるほど。

## どういうこと？？

上記の例で、`obj` は当然 `interface{}` として扱われる。また、interface 変数は「型」と「値」の情報を持っており、それらがともに設定されていないときに限り `nil` として扱われるらしい。[公式の FAQ](http://golang.org/doc/faq#nil_error) にも以下のような記載があった。

> An interface value is nil only if the inner value and type are both unset, (nil, nil). In particular, a nil interface will always hold a nil type. If we store a pointer of type *int inside an interface value, the inner type will be *int regardless of the value of the pointer: (\*int, nil). Such an interface value will therefore be non-nil even when the pointer inside is nil.

上記の例における `x` は「値」の情報を持ってはいないが「型」の情報（`*int`）は持っている。このため、単純に `nil` と比較しても等しくはならなかった、ということ。試しに `reflect.TypeOf` を使って「型」の情報を出力してみると、以下のようになる。

```go
package main

import (
    "fmt"
    "reflect"
)

func PrintType(obj interface{}) {
    fmt.Println(reflect.TypeOf(obj))
}

func main() {
    PrintType(nil) // <nil>

    var x *int = nil
    PrintType(x) // *int
}
```

## まとめ

Go における `nil` の扱い、ちょっと自分の感覚とは違っていたので注意したい。

## 参考

- [絶対ハマる、不思議な nil](http://qiita.com/umisama/items/e215d49138e949d7f805)
