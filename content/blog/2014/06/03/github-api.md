+++
date = "2014-06-03"
tags = [ "api", "golang" ]
title = "Golang で GitHub の Markdown API をたたく"
+++

Google 先生が [github.com/google/go-github/github](https://github.com/google/go-github) なるものを提供してくれていたのでものすごく簡単にできましたメモ。

<!--more-->

<br />
## たたく API

[https://developer.github.com/v3/markdown](https://developer.github.com/v3/markdown)

<br />
## たたいてみる

``` go
package main

import (
    "github.com/google/go-github/github"
    "io/ioutil"
)

func main() {
    md, err := ioutil.ReadFile("sample.md")
    if err != nil {
        panic(err)
    }

    client := github.NewClient(nil)

    html, _, err := client.Markdown(string(md), nil)
    if err != nil {
        panic(err)
    }

    ioutil.WriteFile("sample.html", []byte(html), 0644)
}
```

上記のコードを使って、例えば以下の Markdown を変換してみると…

``` md
# Header 1

## Header 2

* List
* List
* List

## Header 2

> Blockquote
> [Link](http://m0t0k1ch1st0ry.com)
```

期待通り、以下のような HTML が出力される。

``` html
<h1>
<a name="user-content-header-1" class="anchor" href="#header-1"><span class="octicon octicon-link"></span></a>Header 1</h1>

<h2>
<a name="user-content-header-2" class="anchor" href="#header-2"><span class="octicon octicon-link"></span></a>Header 2</h2>

<ul>
<li>List</li>
<li>List</li>
<li>List</li>
</ul><h2>
<a name="user-content-header-2-1" class="anchor" href="#header-2-1"><span class="octicon octicon-link"></span></a>Header 2</h2>

<blockquote>
<p>Blockquote
<a href="http://m0t0k1ch1st0ry.com">Link</a></p>
</blockquote>
```

めでたし！
