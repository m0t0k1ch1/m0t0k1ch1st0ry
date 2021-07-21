+++
title = 'EOS な serialization を C++ と Node.js で'
tags = ['eos','cpp','nodejs','blockchain']
date = '2020-08-31T20:43:55+09:00'
+++

表題の通りのことをやります。非常に基礎的な処理なのですが、いい感じのコードが見当たらなかったので、自分でテスト用に書いたものを載っけておきます。もっとイケてる方法があったら教えてほしいです。

<!--more-->

## serialize するデータ

EOS でよく使う型を適当に組み合わせます。

|     |  type  |      value      |
| :-: | :----: | :-------------: |
|  n  |  name  |   "motokichi"   |
|  a  | asset  | "21000000 POYO" |
|  i  | uint64 |   1231006505    |
|  s  | string |     "poyo!"     |
|  b  |  bool  |      false      |

## C++（コントラクト側）で実行する場合

```cpp
#include <eosio/asset.hpp>
#include <eosio/crypto.hpp>
#include <eosio/eosio.hpp>

using namespace eosio;

class [[eosio::contract]] test : public contract
{
  public:

    using contract::contract;

    [[eosio::action]]
    void serialize(
      const name&        n,
      const asset&       a,
      const uint64_t     i,
      const std::string& s,
      const bool         b
    ) {
      size_t size = 0;
      size += pack_size(n);
      size += pack_size(a);
      size += pack_size(i);
      size += pack_size(s);
      size += pack_size(b);

      std::vector<char> buf;
      buf.resize(size);

      datastream<char*> ds(buf.data(), buf.size());
      ds << n << a << i << s << b;

      checksum256 hashed = sha256(buf.data(), buf.size());
      print(hashed);
    };
};
```

[EOSIO.CDT v1.7.0](https://github.com/EOSIO/eosio.cdt/releases/tag/v1.7.0) でコンパイルして、[jungle3 にデプロイして、](https://jungle3.bloks.io/account/s11ntest1111)`serialize` action を実行してみます。

```sh
$ cleos --url https://jungle3.cryptolions.io push action s11ntest1111 serialize '["motokichi", "21000000 POYO", 1231006505, "poyo!", false]' -p s11ntest1111@active
```

```txt
executed transaction: 3e2d76342d9ec5451bafbc001ea20b7682f359d0007920d71c1dec1f35da09ec  136 bytes  198 us
#  s11ntest1111 <= s11ntest1111::serialize      {"n":"motokichi","a":"21000000 POYO","i":1231006505,"s":"poyo!","b":0}
>> 095207f69473db5e4dd7329d3f7838bf3664a6ca9465fe6cf33e52afcf5a61ef
```

serialize したデータの SHA-256 ハッシュとして

```txt
095207f69473db5e4dd7329d3f7838bf3664a6ca9465fe6cf33e52afcf5a61ef
```

が出力されました。

## Node.js（クライアント側）で実行する場合

```js
const { Serialize } = require("eosjs");
const ecc = require("eosjs-ecc");
const { TextEncoder, TextDecoder } = require("util");

const buf = new Serialize.SerialBuffer({
  textEncoder: new TextEncoder(),
  textDecoder: new TextDecoder(),
});

const n = "motokichi";
const a = "21000000 POYO";
const i = 1231006505;
const s = "poyo!";
const b = false;

buf.pushName(n);
buf.pushAsset(a);
buf.pushNumberAsUint64(i);
buf.pushString(s);
buf.push(b);

const hashed = ecc.sha256(
  Buffer.from(Serialize.arrayToHex(buf.asUint8Array()), "hex"),
  null
);
console.log(hashed.toString("hex"));
```

実行してみます。

```sh
$ node serialization.js
```

```txt
095207f69473db5e4dd7329d3f7838bf3664a6ca9465fe6cf33e52afcf5a61ef
```

コントラクト上で実行したときと同様の結果となりました。

おしまい。
