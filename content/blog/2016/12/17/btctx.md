+++
title = 'Bitcoin の transaction を Go で解読する'
tags = ['bitcoin', 'blockchain', 'go']
date = '2016-12-17T02:08:17+09:00'
+++

Go で Bitcoin の transaction のバイト列をちまちま読む package を結構前に勉強も兼ねて描き始めていたのだが、これが普段ちょこちょこ使うようになって整ってきた。最近筆不精になってるし、せっかくなのでブログのエントリというカタチでも整えてみることにした。

<!--more-->

## つくったもの

{{< github "m0t0k1ch1" "btc" >}}

Bitcoin の transaction は [こういう構造](https://en.bitcoin.it/wiki/Protocol_documentation#tx) になっているので、これを愚直に読んだり書いたりしている。エンディアンが揃っていないところがあったりで嵌ったりした。近くに有識者がいたので助かったけど、1 人でやってたら諦めてたかもしれない。

## できること

ざっくりとは以下の 2 つ。

- hex な transaction をいい感じの struct に変換する
- struct を hex な transaction として書き出す

txid がわかってれば、RPC 叩いて getrawtransaction して decoderawtransaction すればいい感じに返ってくるんだけど、ちょっとこみ入ったことをしようと思うと、こっちでバイト列から Go の struct に変換できたり、そこから hex に戻せると便利ではある。

## 使い方

[Testnet に取り込まれた実際の transaction](http://tbtc.blockr.io/tx/info/d7a4684b71776c8c96edd670a9d0c61d03c293f4c6266b70ff5030b2c4f0bdfe) に対して使ってみる。

``` go
package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/m0t0k1ch1/btc"
)

func main() {
    txhexOrigin := "0100000001ce3cf2e2b334e7e9fa84619469d9edc49368c2f752ea30fb48b080fc794f6d56010000006a473044022065fe1ea4e94a9b44fb62c2b874b63a947504273a60b99b8f7bbf77b4db9331b002205559d8ee93cf341d75866f9eb912af05904fb6eed7372a837308c4e37f3ab58f012103bae5f04799c40862358560e42e441c3080b997a3dec161dd40395e992362bfc9feffffff0200f2052a010000001976a914cbc222711a230ecdd9a5aa65b61ed39c24db2b3488acc08d931a1d0000001976a914426c1ad9fa94f9ea3e6f9248b8bff6768e3ac8c488ac951a1000"

    tx, err := btc.NewTxFromHex(txhexOrigin)
    if err != nil {
        log.Fatal(err)
    }

    b, err := json.Marshal(tx)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(b))

    txhex, err := tx.ToHex()
    if err != nil {
        log.Fatal(err)
    }

    if txhex == txhexOrigin {
        log.Println("match!")
    }
}
```

途中で出力される json は以下のようになる。

``` json
{
  "version": 1,
  "txIns": [
    {
      "hash": "566d4f79fc80b048fb30ea52f7c26893c4edd969946184fae9e734b3e2f23cce",
      "index": 1,
      "sigScript": {
        "hex": "473044022065fe1ea4e94a9b44fb62c2b874b63a947504273a60b99b8f7bbf77b4db9331b002205559d8ee93cf341d75866f9eb912af05904fb6eed7372a837308c4e37f3ab58f012103bae5f04799c40862358560e42e441c3080b997a3dec161dd40395e992362bfc9",
        "asm": "3044022065fe1ea4e94a9b44fb62c2b874b63a947504273a60b99b8f7bbf77b4db9331b002205559d8ee93cf341d75866f9eb912af05904fb6eed7372a837308c4e37f3ab58f01 03bae5f04799c40862358560e42e441c3080b997a3dec161dd40395e992362bfc9"
      },
      "sequence": 4294967294
    }
  ],
  "txOuts": [
    {
      "value": 5000000000,
      "pkScript": {
        "hex": "76a914cbc222711a230ecdd9a5aa65b61ed39c24db2b3488ac",
        "asm": "OP_DUP OP_HASH160 cbc222711a230ecdd9a5aa65b61ed39c24db2b34 OP_EQUALVERIFY OP_CHECKSIG"
      }
    },
    {
      "value": 124999929280,
      "pkScript": {
        "hex": "76a914426c1ad9fa94f9ea3e6f9248b8bff6768e3ac8c488ac",
        "asm": "OP_DUP OP_HASH160 426c1ad9fa94f9ea3e6f9248b8bff6768e3ac8c4 OP_EQUALVERIFY OP_CHECKSIG"
      }
    }
  ],
  "lockTime": 1055381
}
```

## 今後

直近だと、txout のスクリプトの中に入ってる公開鍵のハッシュを address として解釈するのはやっときたい。また、今は標準的な P2PKH しか解釈できないので、少なくとも P2SH な multisig くらいは対応しておきたい。あとは解釈するだけじゃなくて生成・バリデーションするのもラクにできたら嬉しい。以前に Testnet で以下のようなこみ入ったスクリプトの実験をしたりしていて、そんときのデバッグがしんどかったというのがある。

``` txt
OP_DUP 0371aaa70b225a097c615038cdc0fec8b850be37437f6b5ae2c5ecaf463ee30ed6 OP_EQUAL OP_IF OP_DUP OP_HASH160 62b0beae2b5abaccfab8c9d551f3bb2aae289891 OP_EQUALVERIFY OP_CHECKSIG OP_VERIFY OP_SHA256 6c87d5434a635e1470ff3c38956e7962ab1bf036e716b5a0636ab930faf6ce3d OP_EQUAL OP_ELSE 2 03f566e563224460fb7ec66f81dd67d974090343ddeb22321cdab562cdac593d4b 0371aaa70b225a097c615038cdc0fec8b850be37437f6b5ae2c5ecaf463ee30ed6 2 OP_CHECKMULTISIG OP_ENDIF
```

あとは署名。RPC の signrawtransaction は標準的な transaction にしか署名できないので、これもこみ入ったことをやろうとすると融通が効かない。https://github.com/piotrnar/gocoin とかは署名部分も実装してるので、参考にしながら自分でも実装してみようと思う。頭で分かってんのと実装レベルで手を動かすとでは身につき方が違うっぺさ。
