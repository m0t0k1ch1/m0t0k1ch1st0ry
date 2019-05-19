+++
date = "2014-06-09"
tags = [ "golang" ]
title = "net/http の動きを少しだけ追ってみた - Golang"
+++

Golang の標準パッケージである net/http を使えば簡単に HTTP サーバーを立てることができる。とは言うものの、自分はそのへんが実際どうなってるのか全然わかってない。つらい。ということで、Golang の勉強も兼ねて net/http の動きを少しだけ追ってみることにした。

<!--more-->

まず、net/http を用いたよく見かけるサンプルコードを書いてみる。

``` go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func poyo(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "poyo!")
}

func main() {
    http.HandleFunc("/", poyo)
    if err := http.ListenAndServe(":9090", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

やっていることはシンプルで、`http.HandleFunc` と `http.ListenAndServe` だけ。今回はこれらについてコードを追ってみることにした。

## 1. http.HandleFunc

最初に呼んでるのがこれ。

``` go
http.HandleFunc("/", poyo)
```

中身はこう。

``` go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
```

ちなみに、`DefaultServeMux` は以下のようになっている。

``` go
type ServeMux struct {
    mu    sync.RWMutex
    m     map[string]muxEntry  // ルーティングルール
    hosts bool                 // host の情報が含まれているルールが存在するかどうか
}

type muxEntry struct {
    explicit bool     // 正確にマッチするか
    h        Handler  // 対応する handler
    pattern  string   // マッチングさせるパターン
}

func NewServeMux() *ServeMux { return &ServeMux{m: make(map[string]muxEntry)} }

var DefaultServeMux = NewServeMux()
```

そもそも「mux ってなんやの…」ってレベルな自分。どうやら multiplexer の略らしい。

> ServeMux is an HTTP request multiplexer.
> It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

適当に訳すと、「入力された URL にマッチするパターンを登録されているパターン群の中から探して、それに対応する handler を呼び出す」という感じだと思う。

### 1-1. http.ServeMux.HandleFunc

``` go
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    mux.Handle(pattern, HandlerFunc(handler))
}
```

ここで `handler` が強制的に `HandlerFunc` に変換されている。これは [以前のエントリ]({{< ref "blog/2014/06/01/golang.md" >}}) にも書いた。

### 1-2. http.ServeMux.Handle

ルーティングルールを追加している。

``` go
func (mux *ServeMux) Handle(pattern string, handler Handler) {
    mux.mu.Lock()
    defer mux.mu.Unlock()

    if pattern == "" {
        panic("http: invalid pattern " + pattern)
    }
    if handler == nil {
        panic("http: nil handler")
    }
    if mux.m[pattern].explicit {
        panic("http: multiple registrations for " + pattern)
    }

    mux.m[pattern] = muxEntry{explicit: true, h: handler, pattern: pattern}

    if pattern[0] != '/' {
        mux.hosts = true
    }

    n := len(pattern)
    if n > 0 && pattern[n-1] == '/' && !mux.m[pattern[0:n-1]].explicit {
        path := pattern
        if pattern[0] != '/' {
            path = pattern[strings.Index(pattern, "/"):]
        }
        mux.m[pattern[0:n-1]] = muxEntry{h: RedirectHandler(path, StatusMovedPermanently), pattern: pattern}
    }
}
```

お。。っと思ったのは、`pattern` のおしりが `/` のときに、`explicit` が `false` の301リダイレクトが暗黙的にルーティングルールに追加されるということ。もちろん、`pattern` が `/` のときも追加される。こいつはどういう働きをするかというと、こういうことらしい。

> Note that since a pattern ending in a slash names a rooted subtree, the pattern "/" matches all paths not matched by other registered patterns, not just the URL with Path == "/".

実際にどうやってるかはもっと後で出てくる。

## 2. http.ListenAndServe

次に呼んでいるのがこれ。

``` go
http.ListenAndServe(":9090", nil)
```

中身はこう。

``` go
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}
```

もうちょい先の話だけど、`handler` が `nil` のときは、`DefaultServeMux` が `handler` として使われる。

### 2-1. http.Server.ListenAndServe

`net.Listen` で指定したポートを監視し、`http.Server.Serve` を呼んでいる。

``` go
func (srv *Server) ListenAndServe() error {
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    l, e := net.Listen("tcp", addr)
    if e != nil {
        return e
    }
    return srv.Serve(l)
}
```

### 2-2. http.Server.Serve

`for` ループを起動して、リクエストを待ち受けている。リクエストが来ると `net.Listener.Accept` がそれを受けてコネクションを設立し、リクエストに対して goroutine を1つ立ち上げるようになっている。

``` go
func (srv *Server) Serve(l net.Listener) error {
    defer l.Close()
    var tempDelay time.Duration
    for {
        rw, e := l.Accept()
        if e != nil {
            if ne, ok := e.(net.Error); ok && ne.Temporary() {
                if tempDelay == 0 {
                    tempDelay = 5 * time.Millisecond
                } else {
                    tempDelay *= 2
                }
                if max := 1 * time.Second; tempDelay > max {
                    tempDelay = max
                }
                log.Printf("http: Accept error: %v; retrying in %v", e, tempDelay)
                time.Sleep(tempDelay)
                continue
            }
            return e
        }
        tempDelay = 0
        c, err := srv.newConn(rw)
        if err != nil {
            continue
        }
        go c.serve()
    }
}
```

### 2-3. http.conn.serve

こいつはちょっと長かったので「う。。」ってなった。ということで、わかりやすくするためにいろいろ省きまくってみた。

``` go
func (c *conn) serve() {
    for {
        // 省略
        w, err := c.readRequest()
        // 省略
        serverHandler{c.server}.ServeHTTP(w, w.req)
        // 省略
    }
}
```

`http.conn.readRequest` でリクエストの内容を読み込み、`http.serverHandler.ServeHTTP` を呼んでいることがわかる。

### 2-4. http.serverHandler.ServeHTTP

``` go
type serverHandler struct {
    srv *Server
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
    handler := sh.srv.Handler
    if handler == nil {
        handler = DefaultServeMux
    }
    if req.RequestURI == "*" && req.Method == "OPTIONS" {
        handler = globalOptionsHandler{}
    }
    handler.ServeHTTP(rw, req)
}
```

上にもちらっと書いたように、`handler` が `nil` のときは `DefaultServeMux` が使われる。

### 2-5. http.ServeMux.ServeHTTP

最初に登録した `handler` の中からリクエストに対応するものを探して、そいつの `ServeHTTP` を呼んでいる。　

``` go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
    if r.RequestURI == "*" {
        if r.ProtoAtLeast(1, 1) {
            w.Header().Set("Connection", "close")
        }
        w.WriteHeader(StatusBadRequest)
        return
    }
    h, _ := mux.Handler(r)
    h.ServeHTTP(w, r)
}
```

これで終わり。

### 2-6. おまけ

URL に対応する `handler` を探す過程は以下のような感じ。`http.PathMatch` の中で、上にもちらっと書いた以下のルールが適用されている。

> Note that since a pattern ending in a slash names a rooted subtree, the pattern "/" matches all paths not matched by other registered patterns, not just the URL with Path == "/".

``` go
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
    if r.Method != "CONNECT" {
        if p := cleanPath(r.URL.Path); p != r.URL.Path {
            _, pattern = mux.handler(r.Host, p)
            url := *r.URL
            url.Path = p
            return RedirectHandler(url.String(), StatusMovedPermanently), pattern
        }
    }

    return mux.handler(r.Host, r.URL.Path)
}
```

``` go
func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
    mux.mu.RLock()
    defer mux.mu.RUnlock()

    if mux.hosts {
        h, pattern = mux.match(host + path)
    }
    if h == nil {
        h, pattern = mux.match(path)
    }
    if h == nil {
        h, pattern = NotFoundHandler(), ""
    }
    return
}
```

``` go
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
    var n = 0
    for k, v := range mux.m {
        if !pathMatch(k, path) {
            continue
        }
        if h == nil || len(k) > n {
            n = len(k)
            h = v.h
            pattern = v.pattern
        }
    }
    return
}
```

``` go
func pathMatch(pattern, path string) bool {
    if len(pattern) == 0 {
        return false
    }
    n := len(pattern)
    if pattern[n-1] != '/' {
        return pattern == path
    }
    return len(path) >= n && path[0:n] == pattern
}
```

## まとめ

今回はさらっと表面を舐めただけっぽいので、時間があるときにもう少し深入りしてみたい。で、ゆくゆくは俺々 WAF を思い描けるくらいになりたい。

## 参考

- [https://github.com/astaxie/build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang)
