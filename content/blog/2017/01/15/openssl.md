+++
title = "OpenSSL を使って Bitcoin 用の鍵ペアを生成する"
tags = [ "openssl", "bitcoin" ]
date = "2017-01-15T04:11:05+09:00"
+++

Bitcoin 用の鍵ペアをローカルでシュッとつくりたくなって有識者に話してみたら、openssl でつくれるのではという話を聞いたのでやってみた。

<!--more-->

調べてみると、[OpenSSL のコマンドラインツール](https://wiki.openssl.org/index.php/Command_Line_Elliptic_Curve_Operations) の中に `ecparam` と `ec`というのがあって、これで楕円曲線を扱える模様。Bitcoin が楕円曲線暗号に使用しているパラメータは [secp256k1](https://en.bitcoin.it/wiki/Secp256k1) なので、まずはこれが使えるか調べてみる。

``` sh
$ openssl ecparam  -list_curves | grep secp256k1
```

``` txt
  secp256k1 : SECG curve over a 256 bit prime field
```

使えるっぽい。

さらに調べてみると、[どストライクな回答](http://bitcoin.stackexchange.com/questions/5198/for-a-non-technical-person-how-do-i-generate-a-ecdsa-key-pair-easily) があったので、これの通りにやってみる。ホントに使う秘密鍵はインターネット上に晒すなど言語道断であるが、今回はただの検証用なので晒します。

``` sh
$ openssl ecparam -genkey -name secp256k1 -out privkey.pem
$ cat privkey.pem
```

``` txt
-----BEGIN EC PARAMETERS-----
BgUrgQQACg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIEvg5+xDEQsHfz+UcPAZWzyoi7VbXz4kH9h8GIIohP/koAcGBSuBBAAK
oUQDQgAE1vrgGI7c+Wq2ofem3+PkGQs0Tnt8MQdHDtNCRaD6b331GnpwHThfVwQ0
/SwC7v48uvwxQgW5hLSWo3XDtAAt4Q==
-----END EC PRIVATE KEY-----
```

``` sh
$ openssl ec -in privkey.pem -outform DER | tail -c +8 | head -c 32 | xxd -p -c 32
```

``` txt
read EC key
writing EC key
4be0e7ec43110b077f3f9470f0195b3ca88bb55b5f3e241fd87c18822884ffe4
```

``` sh
$ openssl ec -in privkey.pem -pubout -outform DER | tail -c 65 | xxd -p -c 65
```

``` txt
read EC key
writing EC key
04d6fae0188edcf96ab6a1f7a6dfe3e4190b344e7b7c3107470ed34245a0fa6f7df51a7a701d385f570434fd2c02eefe3cbafc314205b984b496a375c3b4002de1
```

できた。

これらをいつも見てる Base58 エンコードされた状態にするには、

- [Base58Check encoding](https://en.bitcoin.it/wiki/Base58Check_encoding)
- [Technical background of version 1 Bitcoin addresses](https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses)
- [List of address prefixes](https://en.bitcoin.it/wiki/List_of_address_prefixes)

この辺りを参考に素直に順番に処理していけばよい。
