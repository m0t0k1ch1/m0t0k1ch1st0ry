+++
title = 'EOS な serialization を Go で'
tags = ['eos', 'go', 'blockchain']
date = '2021-01-05T16:41:07+09:00'
+++

[これ]({{< ref "/blog/2020/08/31/eos-serialization.md" >}}) と同じことを Go でやってみましょうの回。

<!--more-->

EOS Canada の

{{< github "eoscanada" "eos-go" >}}

を使って実装します。

```go
package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	eos "github.com/eoscanada/eos-go"
)

func main() {
	var buf bytes.Buffer
	encoder := eos.NewEncoder(&buf)

	n := eos.AN("motokichi")
	a, err := eos.NewAssetFromString("21000000 POYO")
	if err != nil {
		log.Fatal(err)
	}
	var i uint64 = 1231006505
	s := "poyo!"
	b := false

	if err := encoder.Encode(n); err != nil {
		log.Fatal(err)
	}
	if err := encoder.Encode(a); err != nil {
		log.Fatal(err)
	}
	if err := encoder.Encode(i); err != nil {
		log.Fatal(err)
	}
	if err := encoder.Encode(s); err != nil {
		log.Fatal(err)
	}
	if err := encoder.Encode(b); err != nil {
		log.Fatal(err)
	}

	hashed := hash(buf.Bytes())

	fmt.Println(hex.EncodeToString(hashed))
}

func hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}
```

実行してみます。

```sh
$ go run main.go
```

```txt
095207f69473db5e4dd7329d3f7838bf3664a6ca9465fe6cf33e52afcf5a61ef
```

無事前回と同じ結果が出力されました。

おしまい。
