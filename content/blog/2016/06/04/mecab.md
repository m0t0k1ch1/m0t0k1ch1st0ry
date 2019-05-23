+++
title = "MeCab のユーザー辞書に単語を追加してみる"
tags = ["mecab"]
date = "2016-06-05T00:29:34+09:00"
+++

いろいろあって、とりあえず [MeCab](http://taku910.github.io/mecab) を触ってみるかという気持ちになったので、[こちらのエントリ](https://blog.apar.jp/linux/2748) を参考にして触ってみた。

<!--more-->

## install

``` sh
$ brew install mecab
$ brew install mecab-ipadic
```

### dicdir 確認

``` sh
$ mecab-config --dicdir
/usr/local/lib/mecab/dic
```

### libexecdir 確認

``` sh
$ mecab-config --libexecdir
/usr/local/Cellar/mecab/0.996/libexec/mecab
```

## ユーザー辞書に単語を追加する

### before

「ブラックフェニックス」も「日本酒」も知らないので、分割されてしまう。

``` sh
$ echo 'ブラックフェニックスは最高の日本酒である' | mecab
ブラック        名詞,一般,*,*,*,*,ブラック,ブラック,ブラック
フェニックス    名詞,固有名詞,一般,*,*,*,フェニックス,フェニックス,フェニックス
は      助詞,係助詞,*,*,*,*,は,ハ,ワ
最高    名詞,一般,*,*,*,*,最高,サイコウ,サイコー
の      助詞,連体化,*,*,*,*,の,ノ,ノ
日本    名詞,固有名詞,地域,国,*,*,日本,ニッポン,ニッポン
酒      名詞,接尾,一般,*,*,*,酒,シュ,シュ
で      助動詞,*,*,*,特殊・ダ,連用形,だ,デ,デ
ある    助動詞,*,*,*,五段・ラ行アル,基本形,ある,アル,アル
EOS
```

### ユーザー辞書の保存先を準備

``` sh
$ mkdir /usr/local/lib/mecab/dic/userdic
```

### ユーザー辞書に追加する単語を csv 化

「ブラックフェニックス」と「日本酒」を教える。今回は試しにやってみたいだけなのでコストは 1。

``` sh
$ touch sake.csv
$ echo 'ブラックフェニックス,,,1,名詞,一般,*,*,*,*,ブラックフェニックス,ブラックフェニックス,ブラックフェニックス' >> sake.csv
$ echo '日本酒,,,1,名詞,一般,*,*,*,*,日本酒,ニホンシュ,ニホンシュ' >> sake.csv
```

### csv からユーザー辞書を生成

以下を実行すると、`/usr/local/lib/mecab/dic/userdic/sake.dic` が生成される。

``` sh
$ /usr/local/Cellar/mecab/0.996/libexec/mecab/mecab-dict-index \
-d /usr/local/lib/mecab/dic/ipadic \
-u /usr/local/lib/mecab/dic/userdic/sake.dic \
-f utf-8 \
-t utf-8 \
sake.csv
```

### ユーザー辞書を読み込む

`/usr/local/etc/mecabrc` に以下を追加する。

``` txt
userdic = /usr/local/lib/mecab/dic/userdic/sake.dic
```

### after

「ブラックフェニックス」と「日本酒」が分割されなくなった。

``` sh
$ echo 'ブラックフェニックスは最高の日本酒である' | mecab
ブラックフェニックス    名詞,一般,*,*,*,*,ブラックフェニックス,ブラックフェニックス,ブラックフェニックス
は      助詞,係助詞,*,*,*,*,は,ハ,ワ
最高    名詞,一般,*,*,*,*,最高,サイコウ,サイコー
の      助詞,連体化,*,*,*,*,の,ノ,ノ
日本酒  名詞,一般,*,*,*,*,日本酒,ニホンシュ,ニホンシュ
で      助動詞,*,*,*,特殊・ダ,連用形,だ,デ,デ
ある    助動詞,*,*,*,五段・ラ行アル,基本形,ある,アル,アル
EOS
```
