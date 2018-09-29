+++
date = "2016-08-28T01:52:33+09:00"
tags = [ "python", "word2vec" ]
title = "Python で「老人と海」を word2vec する"
+++

[これ](http://m0t0k1ch1st0ry.com/blog/2016/07/30/nlp) の続き。今回は [gensim](http://radimrehurek.com/gensim) を使って word2vec できるようにするまで。さくっと試せるよう、wikipedia とかではなくて青空文庫のデータをコーパスにする。ちなみに前回 CaboCha も準備したけど、今回は使わない。

<!--more-->

<br />
## 環境

[ここまで](http://m0t0k1ch1st0ry.com/blog/2016/07/30/nlp) やってある前提。

<br />
## 下準備

``` sh
$ yum install nkf unzip
```

<br />
## gensim をインストール

``` sh
$ pip install gensim
```

バージョンを確認。

``` sh
$ pip list | grep gensim
```

``` txt
gensim (0.13.1)
```

<br />
## 青空文庫からデータを持ってくる

[青空文庫](http://www.aozora.gr.jp) から好きなデータを持ってくる。自分はヘミングウェイの [「老人と海」](http://www.aozora.gr.jp/cards/001847/card57347.html) が好きなので、迷いはない。

__テキストファイル(ルビあり)__ の zip（`57347_ruby_57225.zip`）をダウンロードしてもってきて unzip する。


``` sh
$ unzip 57347_ruby_57225.zip
```

unzip すると、`rojinto_umi.txt` になる。下記のように、カジュアルなサイズ。

``` sh
$ wc rojinto_umi.txt
```

``` txt
   726    807 122222 rojinto_umi.txt
```

nkf で文字コードを見てみる。

``` sh
$ nkf -g rojinto_umi.txt
```

``` txt
Shift_JIS (CR)
```

Shift_JIS だとつらいので、UTF-8 にする。

``` sh
$ nkf -w --overwrite rojinto_umi.txt
```

UTF-8 になったか確認。

``` sh
$ nkf -g rojinto_umi.txt
```

``` txt
UTF-8 (CR)
```

なお、ファイルの頭とケツについてる物語と関係ない説明っぽい部分は消しておいた。

<br />
## MeCab で分かち書きする

日本語は分かち書きしないと word2vec できないので、MeCab を使って分かち書きする。

今回はだいぶ小さいコーパスで word2vec することになるので、語彙数をむやみに増やしたくない。なので、`-Owakati` を使わずに、基本形で分かち書きしてみることにする。

ということで、出力フォーマットを調整して分かち書きする Python スクリプトを描いた。出力フォーマットの調整法については [ここ](https://taku910.github.io/mecab/mecab.html) や [ここ](https://taku910.github.io/mecab/format.html) に記載がある。

``` python
# -*- coding: utf-8 -*-

import MeCab
import sys

tagger = MeCab.Tagger('-F\s%f[6] -U\s%m -E\\n')

fi = open(sys.argv[1], 'r')
fo = open(sys.argv[2], 'w')

line = fi.readline()
while line:
    result = tagger.parse(line)
    fo.write(result[1:]) # skip first \s
    line = fi.readline()

fi.close()
fo.close()
```

このスクリプトを `wakati.py` として保存して、`rojinto_umi.txt` に対して実行する。分かち書き後のデータは `rojinto_umi_wakati.txt` として保存する。

``` sh
$ python wakati.py rojinto_umi.txt rojinto_umi_wakati.txt
```

`rojinto_umi_wakati.txt` の中身を見てみると、想定通り基本形で分かち書きされている。

>　 彼 は 老いる て いる た 。 小さな 船 で メキシコ 湾流 に 漕ぐ 出す 、 独り で 漁 を する て いる た 。 一 匹 も 釣れる ない 日 が 、 既に 八 四 日 も 続く て いる た 。 最初 の 四 〇 日 は 少年 と 一緒 だ た 。 しかし 、 獲物 の 無い まま に 四 〇 日 が 過ぎる と 、 少年 に 両親 が 告げる た 。 あの 老人 は もう 完全 に 「 サラオ 」 だ ん だ よ 、 と 。 サラオ と は 、 すっかり 運 に 見放す れる た という こと だ 。 少年 は 両親 の 言いつける 通り に 別 の ボート に 乗り換える 、 一 週間 で 三 匹 も 立派 だ 魚 を 釣り上げる た 。 老人 が 毎日 空っぽ の 船 で 帰る て くる の を 見る たび に 、 少年 の 心 は 痛む だ 。 彼 は いつも 老人 を 迎える に 行く て 、 巻く た ロープ 、 手鉤 《 ギャフ 》 、 銛 《 もる 》 、 帆 を 巻く つける た マスト など を 運ぶ 手伝い を する の だ た 。 粉 袋 で 継ぎ あて する れる た 帆 は 、 巻き上げる られる て 、 永遠 の 敗北 を 示す 旗印 の よう に 見える た 。

<br />
## word2vec する

分かち書きできたので、word2vec する。

ベクトルの次元数は 100。skip-gram モデルを用い、階層的ソフトマックスで学習させる。ネガティブサンプリングはなし。

ちなみに、インターネット上にある word2vec を試してみました的なエントリのいくつかでは、text8 以外のコーパスを読み込むときにも `word2vec.Text8Corpus` を使っているみたいだけれど、`word2vec.LineSentence` を使って文ごとに読み込まないと、文と文の間で文脈がごちゃごちゃになってしまうような気がするので、今回は後者を採用する。

``` python
# -*- coding: utf-8 -*-

from gensim.models import word2vec
import logging
import sys

logging.basicConfig(format='%(asctime)s : %(levelname)s : %(message)s', level=logging.INFO)

sentences = word2vec.LineSentence(sys.argv[1])
model = word2vec.Word2Vec(sentences,
                          sg=1,
                          size=100,
                          min_count=1,
                          window=10,
                          hs=1,
                          negative=0)
model.save(sys.argv[2])
```

このスクリプトを `train.py` として保存して、`rojinto_umi_wakati.txt` に対して実行する。学習の結果として生成される model は `rojinto_umi.model` として保存する。

``` sh
$ python train.py rojinto_umi_wakati.txt rojinto_umi.model
```

``` txt
2016-08-27 16:49:18,590 : INFO : collecting all words and their counts
2016-08-27 16:49:18,590 : INFO : PROGRESS: at sentence #0, processed 0 words, keeping 0 word types
2016-08-27 16:49:18,607 : INFO : collected 3147 word types from a corpus of 37519 raw words and 634 sentences
2016-08-27 16:49:18,619 : INFO : min_count=1 retains 3147 unique words (drops 0)
2016-08-27 16:49:18,619 : INFO : min_count leaves 37519 word corpus (100% of original 37519)
2016-08-27 16:49:18,629 : INFO : deleting the raw counts dictionary of 3147 items
2016-08-27 16:49:18,629 : INFO : sample=0.001 downsamples 44 most-common words
2016-08-27 16:49:18,629 : INFO : downsampling leaves estimated 21746 word corpus (58.0% of prior 37519)
2016-08-27 16:49:18,629 : INFO : estimated required memory for 3147 words and 100 dimensions: 4720500 bytes
2016-08-27 16:49:18,632 : INFO : constructing a huffman tree from 3147 words
2016-08-27 16:49:18,726 : INFO : built huffman tree with maximum node depth 15
2016-08-27 16:49:18,728 : INFO : resetting layer weights
2016-08-27 16:49:18,776 : INFO : training model with 3 workers on 3147 vocabulary and 100 features, using sg=1 hs=1 sample=0.001 negative=0
2016-08-27 16:49:18,776 : INFO : expecting 634 sentences, matching count from corpus used for vocabulary survey
2016-08-27 16:49:19,871 : INFO : PROGRESS: at 25.96% examples, 26216 words/s, in_qsize 5, out_qsize 0
2016-08-27 16:49:21,016 : INFO : PROGRESS: at 69.94% examples, 33451 words/s, in_qsize 5, out_qsize 0
2016-08-27 16:49:21,563 : INFO : worker thread finished; awaiting finish of 2 more threads
2016-08-27 16:49:21,572 : INFO : worker thread finished; awaiting finish of 1 more threads
2016-08-27 16:49:21,741 : INFO : worker thread finished; awaiting finish of 0 more threads
2016-08-27 16:49:21,741 : INFO : training on 187595 raw words (108660 effective words) took 3.0s, 36682 effective words/s
2016-08-27 16:49:21,742 : WARNING : under 10 jobs per worker: consider setting a smaller `batch_words' for smoother alpha decay
2016-08-27 16:49:21,742 : INFO : saving Word2Vec object under rojinto_umi.model, separately None
2016-08-27 16:49:21,742 : INFO : not storing attribute cum_table
2016-08-27 16:49:21,742 : INFO : not storing attribute syn0norm
```

所要時間 3 秒。

ちなみに、gensim の word2vec の学習部分のコードには Python 実装と Cython 実装があって、デフォルトで Cython 実装の方が使われる。Cython 実装では、GIL をリリースして並列化されていたりするので、Python 実装に比べるとかなり速い。

<br />
## word2vec の結果を確認

とりあえず、指定した単語とコサイン類似度の高い単語をリストアップするスクリプトを描いて実行してみる。

``` python
# -*- coding: utf-8 -*-

from gensim.models import word2vec
import sys

model   = word2vec.Word2Vec.load(sys.argv[1])
results = model.most_similar(positive=sys.argv[2], topn=10)

for result in results:
    print(result[0], '\t', result[1])
```

``` sh
$ python similars.py rojinto_umi.model 人生
```

``` txt
泥棒     0.9779643416404724
不足     0.969954788684845
温存     0.9699369668960571
高値     0.9684933423995972
平気     0.9683408737182617
苦労     0.9680980443954468
明ける   0.9679121971130371
どころか         0.9677099585533142
手間     0.9673588275909424
ソックス         0.9660428762435913
```

「人生」と最もコサイン類似度の高い単語は「泥棒」とのこと。次いで「不足」「温存」。それなりに物語を汲み取れてるのではないだろうか。

ちっちゃいコーパスでもこんな感じで結構楽しい。
