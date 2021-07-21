+++
title = 'Go で好きな画像を LGTMize する'
tags = ['go']
date = '2015-03-14'
images = ['img/entry/zushi-lgtm.jpg']
+++

{{< github "m0t0k1ch1" "lgtmize" >}}

普段から多用している `:ok_woman:` に少し飽きてきたので、そろそろしめやかに LGTM 童貞を卒業しようと思い、つくってみた。

<!--more-->

## 使い方

[README](https://github.com/m0t0k1ch1/lgtmize/blob/master/README.md) だけだとちょっとよくわからないので補足。

まずは `go get` する。

```sh
$ go get github.com/m0t0k1ch1/lgtmize
```

あとは `lgtmize` するだけ。

[LGTM.in/g](http://www.lgtm.in) 然り、インターネット上にカジュアルにアップされてる LGTM 画像みたいなことをしてしまうと著作権的にアウトな気がするので、今回は自分で撮った逗子の写真で試してみる。

{{< figure src="/img/entry/zushi.jpg" >}}

まさに涅槃。住みたい。

これを `zushi.jpg` みたいな名前で適当なディレクトリに置いて、以下を実行する。

```sh
$ lgtmize /path/to/zushi.jpg
```

すると、同じディレクトリに `zushi-lgtm.jpg` という名前で、以下のような画像が生成される。

{{< figure src="/img/entry/zushi-lgtm.jpg" >}}

あとはなんとかして GitHub まで持っていく。Chrome 拡張とかまでつくる元気はなかった。とはいえ、GitHub が提供してるインターフェースでコメントに画像貼るのも意外と簡単なので、とりあえずはローカルに LGTM 画像フォルダをつくってがんばろうと思う。実際にやってみてめんどくさかったらなんか考える。

## ちなみに

- フォントは大好きな [M+](http://mplus-fonts.sourceforge.jp) を使わせていただきました。こういうハイクオリティなフォントがフリーソフトウェアとして公開されてるのは本当に助かります。
- ホントは [かなえちゃん](https://www.tumblr.com/search/%E7%A5%9E%E5%B4%8E%E3%81%8B%E3%81%AA%E3%81%88) で実践したい。著作権〜〜〜。
