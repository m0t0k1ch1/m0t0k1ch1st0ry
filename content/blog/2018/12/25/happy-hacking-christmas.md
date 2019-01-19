+++
date = "2018-12-25T01:54:05+09:00"
tags = [ "ethereum", "solidity", "blockchain" ]
title = "Happy Hacking Christmas"
+++

この記事は [ex-KAYAC Advent Calendar 2018](https://qiita.com/advent-calendar/2018/ex-kayac) の 25 日目の記事です。恐縮ですが、私 [@m0t0k1ch1](https://twitter.com/m0t0k1ch1) が大トリを務めさせていただきます。メリークリスマス。

<!--more-->

![santa claus](/img/entry/santa-claus.jpg)

<br />
## はじめに

とりあえず、軽く自己紹介を。自分は 2013 新卒としてカヤックに入社し、ソーシャルゲーム事業部で 3 年ほどサーバーサイドエンジニア「ダンゲル」としてお世話になりました。カヤックに在籍していたときのことは [退職してしばらく経った後で書いたブログ](https://m0t0k1ch1st0ry.com/blog/2017/01/30/kayac) でも触れていますので、今回はこの辺で。現在は大阪で Blockchain をベースとした [cryptoeconomic](https://m0t0k1ch1st0ry.com/blog/2018/08/05/cryptoeconomics) なプロトコルの R&D に取り組んでいます。技術的なトピックとしては特に [Plasma](https://scrapbox.io/sivira-plasma) を中心にリサーチしていますが、来年は R だけではなく D でも成果を出していきたい所存です。その際はよろしくお願い致します。

あ、あと、最初にお伝えしておきたいことがもう 1 つ。

__この第 1 回 ex-KAYAC Advent Calendar 開催を記念して、ex-KAYAC の皆さんにはちょっとイイことが起こるかもしれません。__

これに関する続報は、この Advent Calendar と同時に立ち上がった ex-KAYAC な Slack チャンネルに流れると思いますので、チャンネルに join したい方は、[@m0t0k1ch1](https://twitter.com/m0t0k1ch1) もしくは [@ngystks](https://twitter.com/ngystks) まで気軽にお声がけください。

<br /><br /><br />
...さて、前置きはこのくらいにして、本題に移っていきたいと思います。

<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">街でかかるクリスマスソングが全てこれに変われば解脱が捗るのに、と毎年思う &gt;&gt; Utada / Merry Christmas Mr. Lawrence - FYI <a href="https://t.co/38IUI9EFmG">https://t.co/38IUI9EFmG</a></p>&mdash; m0t0k1ch1 (@m0t0k1ch1) <a href="https://twitter.com/m0t0k1ch1/status/1073265831847485442?ref_src=twsrc%5Etfw">2018年12月13日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

そうです、今日はクリスマスです。自分がこのブログを書いている今この瞬間（12/24 深夜）も、世界中のサンタクロース達は大忙しでしょう。

ということで、今日は、自分もそんなサンタクロース達の力になるべく、分散型サンタクロースを使ってプレゼント配りのお手伝いをしようと思います。

<br />
## 分散型サンタクロース（という名の [鎖野郎](https://twitter.com/leo_hio/status/1021570150888222720) 諸兄姉への挑戦状）

🎅 [SantaClaus contract](https://ropsten.etherscan.io/address/0x05d9cbee05e82d492ad66842fc7c0cb363b384ea#code)

^ こちらが、Ethereum（Ropsten testnet）上にデプロイされた分散型サンタクロースです。かわいいですね。

もちろん、かわいいだけではありません。この SantaClaus contract は __[SantaClausToken（SCT）](https://ropsten.etherscan.io/token/0xa9b76b79e3254d7835401a8b43af2fac93a83f2d)__という特別なトークンをプレゼントとして配ることができるのです。

しかし、この SantaClausToken（SCT）、誰もがもらえるわけではありません。

正確に言うと、

__Ethereum（Ropsten testnet）上にクリスマスを題材にしたある種のパズルが構築されており、このパズルを解くことができた方のみ、SantaClausToken（SCT）が獲得できるようになっています。__

パズルに関連する contract の概要を以下にまとめます。

- __Letter__（低難度パズル）
  - 💌 [Letter contract](https://ropsten.etherscan.io/address/0xbade12c0bd7943a066e77f0466d529d78d2f70db#code)
  - SantaClaus へのお手紙です
  - SantaClausToken をもらうには、チップを包んで封をする必要があります
- __ChristmasStocking__（中難度パズル）
  - 🧦 [ChristmasStockin contractg](https://ropsten.etherscan.io/address/0x408f56c4541bd00ec836102d06f7ee6a2a820678#code)
  - SantaClaus がプレゼントを入れるための靴下です
  - SantaClausToken をもらうには、この中に賄賂を仕込む必要があります
- __ChristmasTree__（高難度パズル）
  - 🎄 [ChristmasTree contract](https://ropsten.etherscan.io/address/0x0a97246d46703f72b5c34828f80171f005f66c60#code)
  - お祈りをしたり、飾りつけをしたりすることができる、クリスマスツリーです
  - SantaClausToken をもらうには、たくさんたくさんお祈りをする必要があります
- __SantaClaus__
  - 🎅 [SantaClaus contract](https://ropsten.etherscan.io/address/0x05d9cbee05e82d492ad66842fc7c0cb363b384ea#code)
  - 分散型サンタクロースです
  - 条件を満たした良い子にだけ、SantaClausToken（SCT）をプレゼントしてくれます
- __SantaClausToken（SCT）__
  - 💎 [SantaClausToken contract](https://ropsten.etherscan.io/address/0xa9b76b79e3254d7835401a8b43af2fac93a83f2d#code)
  - パズルを解いた証です
  - 中身は、SantaClaus しか transfer できない ERC20 トークンです

Letter と ChristmasStocking と ChristmasTree がパズル要素を持った contract で、SantaClaus はパズルが解けたかどうかのチェックと報酬である SantaClausToken（SCT）の付与を担う contract です。

パズルを全て解いた状態で SantaClaus contract の `requestToken()` function（下記）を実行することで、SantaClausToken（SCT）が獲得できます。

``` solidity
function requestToken() public {
  require(_letter.isSealed(msg.sender));
  require(_christmasStocking.balanceOf(msg.sender) > 0);
  require(_christmasTree.powerOf(msg.sender) > 99999999);
  require(_token.balanceOf(msg.sender) == 0);

  // Congratulations!!
  _token.mint(msg.sender, 1);
}
```

なお、全 contract のソースコードはこちらに置いておきました。

<div class="github-card" data-user="m0t0k1ch1" data-repo="happy-hacking-christmas"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

パズルに関連する contract は、全て数十行程度でとてもシンプルなので、Solidity 初心者の方でも問題なく理解できるレベルかと思います。

また、不要な気遣いかもしれませんが、念のため、パズルに関連する各 contract に簡単にアクセスするためのスクリプトの雛形を [こちら](https://github.com/m0t0k1ch1/happy-hacking-christmas/blob/master/scripts/sample.js) に置いておきました。使い方は [README](https://github.com/m0t0k1ch1/happy-hacking-christmas/blob/master/README.md) に記載していますので、必要であればお使いください。

<br /><br /><br />
...さて、あまり語り過ぎてもパズルがつまらなくなってしまうので、説明はこの辺にしておきましょう。

なお、テストも兼ねて自分は既にクリア済みですので、「どうやっても SCT 獲得できない！！詰んでる！！！」という大事故は回避できているはずです。自分が予想していないようなハックが発生する可能性はありますが、それはそれでまた一興でしょう。

また、このパズルをつくった意図やパズルの解法については、来年、年始の休みが明けた頃に改めてブログにまとめようと思いますが、その時点で誰もクリアできていなかったらちょっと考えます（実際、そうなってもおかしくない程度の難易度かもしれません）。

<br /><br /><br />
...はい！ということで、鎖野郎諸兄姉の皆様、是非とも奮ってご参加いただき、分散型サンタクロースから SantaClausToken（SCT）を勝ち取ってください！！

（記事前半はそんな主旨ではなかった気がしますが、気にしない！）

<div style="width:100%;height:0;padding-bottom:42%;position:relative;"><iframe src="https://giphy.com/embed/3o85xAojNshmzlySyc" width="100%" height="100%" style="position:absolute" frameBorder="0" class="giphy-embed" allowFullScreen></iframe></div><p><a href="https://giphy.com/gifs/luke-skywalker-good-luck-force-be-with-you-3o85xAojNshmzlySyc">via GIPHY</a></p>

あ、パズルが解けた方は、[@m0t0k1ch1](https://twitter.com/m0t0k1ch1) に一声かけていただけると嬉しみです。是非、感想などお聞きしてみたいです。

<br />
## ※追記（2018-12-27）

3 人目の SCT ホルダーである yohei さんが、SCT ホルダーしか名前を刻めない contract をつくってくださいました。

[TokenHolders contract](https://ropsten.etherscan.io/address/0x5a8b024f544ed745afc4d980e403c5e0967e62df#code)

SCT を獲得した方は、是非こちらに名前を刻んでいただけたらと思います！

自分の方でも、電子署名と Twitter の OAuth を使ってアドレスと Twitter アカウントを紐付けて、SCT ホルダーの Twitter アカウントを可視化する仕組みをつくろうかなあと思っていたりしますので、完成したあかつきには、またこのブログと GitHub repo でお知らせしようと思います。

<br />
## ※追記（2019-01-17）

続きを書きました。解答について言及している部分もあるので、「まだ見たくない！」という方はご注意を。

[Happy Hacking Christmas の解答と狙い](https://m0t0k1ch1st0ry.com/blog/2019/01/17/after-happy-hacking-christmas)

<br />
## 最後に

最後になってしまいましたが、今回の ex-KAYAC Advent Calendar 2018 を立ち上げ（させ）た身として、参加してくださった皆さんにお礼を申し上げたいと思います。

__突拍子もない企画に乗っかっていただき、本当にありがとうございました 🙏__

まさにカヤックの文化の 1 つである「乗っかる」が体現された第 1 回 ex-KAYAC Advent Calendar だったように思います。

実は、この企画自体は 1 年前に [@ryusukefuda のツイート](https://twitter.com/ryusukefuda/status/940185306220847104) を見つけたときに自分の中では開催することを心に決めており、予定通り 1 年の時を経て [当該ツイートをほじくり返した](https://twitter.com/m0t0k1ch1/status/1058232279963656192) ことによって立ち上がりました。

裏方として動いてくださった [@ryusukefuda](https://twitter.com/ryusukefuda) と [@ngystks](https://twitter.com/ngystks) の両氏、ありがとうございました 🙏

ちなみに、自分が ex-KAYAC Advent Calendar を開催してみたかった主な理由は以下です。

- 今も変わらずカヤックのことが好き（まだ恩返しができていないという気持ちもある）なので、これからも何かしら繋がりを持っていたいと思っていた
- 自分と近しい気持ちを持った ex-KAYAC な方はそれなりにいるのではないかと思ったので、そういう方々が気軽に集まって繋がれる場をつくってみたかった

実際、ex-KAYAC な Slack チャンネルが開設されたり、何人かの方から Twitter でフォローしていただいたり、、新しい繋がりが生まれるきっかけはいくつかつくれたかなと思います。また、この記事をきっかけにしてさらに繋がりが増えたら嬉しいなと思いますし、同じような気持ちの方が多ければ、ゆるやかに ex-KAYAC 文化を育んでいけたらよいなとも思います。さらに言うならば、個人的には、[DAO](https://en.wikipedia.org/wiki/Decentralized_autonomous_organization) の文脈でカヤックと何かできたらいいなと企んでいたりもします。

...と、そろそろ夜が明けてしまいそうなので、、今年の自分が良い子だったと信じて、急いでサンタさんに [お願い](http://amzn.asia/hYPtxfi) をしてから眠りにつこうと思います。

ex-KAYAC Advent Calendar、来年も続くといいな。それでは皆さん、よいお年を〜👋
