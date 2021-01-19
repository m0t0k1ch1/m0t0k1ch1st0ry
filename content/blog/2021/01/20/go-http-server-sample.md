+++
title = 'Go で素朴な HTTP API サーバーを書く'
tags = ['go']
date = '2021-01-20T00:50:22+09:00'
+++

久々に Go をゴリゴリ書いてるんですが、コード量が膨大になる前に、HTTP API サーバーを書くときにだいたい書くことになるだろう土台部分のコードを抽出しておきたい衝動に駆られたので、抽出しておこうと思います。表題の通り、素朴です。たぶんまだちょこちょこアプデするとは思いますが。

<!-- more -->

{{< github "m0t0k1ch1" "go-http-server-sample" >}}

下に貼った 1 つ目の記事の著者と同様、自分も I like concrete examples 派なので、ソースコードまるっと置いておきます。

ちなみに、Go でこういうことをやる場合は [net/http](https://golang.org/pkg/net/http) でやりきるパターンも多いと思われますが、自分は [echo](https://github.com/labstack/echo) くらいのサポートがある方が好みなので echo を使いました。とは言え、がっつり echo に依存しまくったコードがあるわけでもないので、差し替えは容易だと思います。

Gopher のみなさん、イケてない部分を見つけたら是非ご指摘願います🙏

また、実装にあたって特に参考になった記事を 2 つ挙げておきます。感謝です🙏

- [Organising Database Access in Go](https://www.alexedwards.net/blog/organising-database-access)
- [Go の API のテストにおける共通処理](https://medium.com/@timakin/go-api-testing-173b97fb23ec)
