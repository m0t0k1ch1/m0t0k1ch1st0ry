+++
title = 'Go で RSA 署名'
tags = ['go']
date = '2014-08-18'
+++

暗号化と署名の違いを最近知った若輩者なのですが、、ちょっと仕事でこのあたりのことを扱う必要が出てきたので、表題の通りのことをやってみました。

<!--more-->

## 暗号化と署名について

[暗号化と署名は対称じゃないよという話](http://www.machu.jp/diary/20080302.html#p01) あたりに目を通すとよいかと思います。こちらのエントリは結構前のものですが、とてもわかりやすかったです。

## 秘密鍵と公開鍵を準備する

``` sh
$ openssl genrsa 1024 > private-key.pem
$ openssl rsa -pubout < private-key.pem > public-key.pem
```

## 実際に署名をつくって検証してみる

以下のコードで署名の作成と検証が実行できます。key へのパスはよしなに置換してください。

``` go
package main

import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha512"
    "crypto/x509"
    "encoding/pem"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
)

func readPrivateKey(path string) (*rsa.PrivateKey, error) {
    privateKeyData, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }

    privateKeyBlock, _ := pem.Decode(privateKeyData)
    if privateKeyBlock == nil {
        return nil, errors.New("invalid private key data")
    }
    if privateKeyBlock.Type != "RSA PRIVATE KEY" {
        return nil, errors.New(fmt.Sprintf("invalid private key type : %s", privateKeyBlock.Type))
    }

    privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
    if err != nil {
        return nil, err
    }

    return privateKey, err
}

func readPublicKey(path string) (*rsa.PublicKey, error) {
    publicKeyData, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }

    publicKeyBlock, _ := pem.Decode(publicKeyData)
    if publicKeyBlock == nil {
        return nil, errors.New("invalid public key data")
    }
    if publicKeyBlock.Type != "PUBLIC KEY" {
        return nil, errors.New(fmt.Sprintf("invalid public key type : %s", publicKeyBlock.Type))
    }

    publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
    if err != nil {
        return nil, err
    }

    publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
    if !ok {
        return nil, errors.New("not RSA public key")
    }

    return publicKey, nil
}

func main() {
    // read private key
    privateKey, err := readPrivateKey("/path/to/your/private-key.pem")
    if err != nil {
        log.Fatal(err)
    }

    // precompute
    //   ref. http://golang.org/pkg/crypto/rsa/#PrivateKey.Precompute
    //   Precompute performs some calculations that speed up private key operations in the future.
    privateKey.Precompute()

    // validate
    //   ref. http://golang.org/pkg/crypto/rsa/#PrivateKey.Validate
    //   Validate performs basic sanity checks on the key.
    //   It returns nil if the key is valid, or else an error describing a problem.
    if err := privateKey.Validate(); err != nil {
        log.Fatal(err)
    }

    // read public key
    publicKey, err := readPublicKey("/path/to/your/public-key.pem")
    if err != nil {
        log.Fatal(err)
    }

    // generate token hash from token
    hasher := sha512.New()
    hasher.Write([]byte("token"))
    hasher.Write([]byte("salt"))
    tokenHash := hasher.Sum(nil)

    // sign
    signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA512, tokenHash, nil)
    if err != nil {
        log.Fatal(err)
    }

    // verify
    if err := rsa.VerifyPSS(publicKey, crypto.SHA512, tokenHash, signature, nil); err != nil {
        log.Fatal(err)
    }

    log.Println("success")
}
```
