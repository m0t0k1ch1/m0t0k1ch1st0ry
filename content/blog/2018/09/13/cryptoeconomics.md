+++
date = "2018-09-13T03:39:01+09:00"
tags = [ "cryptoeconomics", "economics", "blockchain" ]
title = "cryptoeconomics の「守」"
+++

[この記事](https://m0t0k1ch1st0ry.com/blog/2018/08/05/cryptoeconomics) を書いてから気づけば早 1 ヶ月。お久しぶりな cryptoeconomics 記事です。なお、タイトルにある「守」は、守破離の「守」です。

<!--more-->

![dna](/img/entry/dna.png)

<br />
## はじめに

現在進行形で cryptoeconomics 沼に嵌りながら奮闘する中でですね、もちろん消化できていないことは盛り沢山なのですが、「これは脳内整理も兼ねてブログにアウトプットしてもよかろう」と思うこともポツポツ出てきました。インプット偏重状態は不健全なので、これからしばらくは少しずつアウトプットも進捗させてバランスをとっていこうと思います。

で、今回はまず、そのプロローグ的な位置付けとして、この 1 ヶ月の中で新しく気づいたことの概要を軽くまとめておこうと思います。これらの気づきは、この記事のタイトルに書いた通り、cryptoeconomics の「守」として認識しておくべき内容だと考えており、今後のアウトプットの前提知識としても重要になってくると思うので、先にざっくり整理しておこう、という感じです。

注：[前回の記事](https://m0t0k1ch1st0ry.com/blog/2018/08/05/cryptoeconomics) の内容を前提として書きますので、まだ読まれていない方はこの記事を読む前に是非ご一読ください🙏

<br />
## 重要な気づき

前回の記事を書いてからどんなインプットをしていたのかについては、[Scrapbox](https://scrapbox.io/m0t0k1ch1/cryptoeconomics) や [日々のツイート](https://twitter.com/m0t0k1ch1) からお察しいただければと思うのですが、その中でも重要な気づきに繋がったインプットを 2 つピックアップして、内容と感想を軽くまとめてみます。深堀りに関しては今後のアウトプットの中で進めていければと思います。

双方、cryptoeconomist を名乗りたいのであれば必読と言える記事だと思いますので、cryptoeconomics に興味がある方は一読されることを強くオススメします。

<br />
### A Crash Course in Mechanism Design for Cryptoeconomic Applications

ref. https://medium.com/blockchannel/a-crash-course-in-mechanism-design-for-cryptoeconomic-applications-a9f06ab6a976

これは、前回の記事の最後で

> ちょうど今、自分もこの記事の内容を消化しようと奮闘しているところです。

と言っていた記事ですね。ちなみに、自分の読書メモは [こちら](https://scrapbox.io/m0t0k1ch1/A_Crash_Course_in_Mechanism_Design_for_Cryptoeconomic_Applications) です。

記事の前半部分では、メカニズムデザインの基礎的な概念から、既存のメカニズムの中でも非常に強力な VCG メカニズムに至るまで、その概要が解説されています。後半部分では、Vitalik や Vlad の考えをベースにしながら、メカニズムデザインが crypto 領域でどのように活用できる可能性があるのか？活用するにあたっての勘所は何か？などについて考察されています。

特に重要だと感じたのは、「Open Challenges: Ethereum Foundation」と「Introducing Cooperative Game Theory」のパートです。自分がここで気づいたことは、

__cryptoeconomic なプロトコルの性質（特に安全性）を評価する際、__

- __どのような市場モデルを前提として議論するかは非常に重要であること__
- __協力ゲーム理論を前提とした市場モデルを採用するのが妥当な場合が多いこと__

です。

前者についてもうちょっと踏み込むと、以下のような感じです。

__既存のフォールトトレランスの研究では honest majority model（少なくとも過半数のプレイヤーが正直だという仮定）に基づいて安全性を議論することが一般的だったが、経済学的な側面も有する cryptoeconomic なプロトコルの安全性評価に関しては、このモデルでは不適切な場合が多い。当然、モデリングが不適切だった場合、それを前提とした安全性評価にはほぼ意味がなくなるので、これは非常に重要。__

要するに、市場モデルは cryptoeconomics を議論する際の土台なので、慎重に考えましょうということです。

後者についてももうちょっと踏み込むと、以下のような感じです。

__cryptoeconomic なプロトコルについて考える際、プレイヤーのマイニングパワーやトークン保有量の分布は重要な要素である。また、これらはある程度集権化していることが多い。そのため、これは寡占市場としてモデル化し、協力ゲーム理論に基づいて議論するのが適切である場合が多い。しかし、既存のメカニズムデザイン事例の多くは非協力ゲーム理論に基づいているため、上記のようなモデリングを前提とした設計は非常に困難。強力な VCG メカニズムもここではその力を失う。__

（ちょっと補足をしておくと、協力ゲーム理論は、プレイヤーが提携することを前提とし、どのような提携が構築されるのか？提携に起因する利得はどのように分配すべきか？といった問題を扱う学問で、プレイヤーが独自に意思決定することを前提とする非協力ゲーム理論とはその前提が大きく異なります）

結論、この辺りの話が cryptoeconomics を新しい学問足らしめているのでしょう。既存のツールではそう簡単に倒せない相手だということです。

が、記事の中でも書かれているように、既存のメカニズムデザインや協力ゲーム理論の知見が全く使えないわけではありません。むしろ、これまで数々の研究者が積み上げてきた知見が既に存在するわけです。適当に車輪の再発明をするのではなく、まずはこれらの知見と素直に向き合うのが「守」の姿勢と言えるのではないでしょうか。

<br />
### The History of Casper — Chapter 4

ref. https://medium.com/@Vlad_Zamfir/the-history-of-casper-chapter-4-3855638b5f0e

有名な The History of Casper の chapter 4 です。

The History of Casper は、（現時点では）全 5 chapter で構成されており、Vlad 主観で Casper 誕生の物語が綴られています（あくまで初期の Casper の話であり、最新の Casper に関する知見がまとまっているわけではないので注意してください）。一応、以下にリンクを全部まとめておきます。

- [Chapter 1](https://medium.com/@Vlad_Zamfir/the-history-of-casper-part-1-59233819c9a9)
- [Chapter 2](https://medium.com/@Vlad_Zamfir/the-history-of-casper-chapter-2-8e09b9d3b780)
- [Chapter 3](https://medium.com/@Vlad_Zamfir/the-history-of-casper-chapter-3-70fefb1182fc)
- [Chapter 4](https://medium.com/@Vlad_Zamfir/the-history-of-casper-chapter-4-3855638b5f0e)
- [Chapter 5](https://medium.com/@Vlad_Zamfir/the-history-of-casper-chapter-5-8652959cef58)

読んで感じたことは、「A Crash Course in Mechanism Design for Cryptoeconomic Applications」のパートで書いた気づきとほぼ同じなので、付け加えて書くことはほとんどないのですが、現実の事象をベースにして綴られているので「あー、そんなことあったのね〜」「あー、だからそうなったのね〜」みたいな発見はあり、それは単純におもしろかったです。

<blockquote class="twitter-tweet" data-lang="ja"><p lang="ja" dir="ltr">こんなことあったのね。どの程度の &quot;lightly trolled&quot; だったのかは分からないけど笑、Vlad の気持ち察するわ &gt;&gt; &quot;The History of Casper - Chapter 3&quot; - Vlad Zamfir <a href="https://t.co/W7Mv5e73rk">https://t.co/W7Mv5e73rk</a> <a href="https://t.co/BYjoFADPPD">pic.twitter.com/BYjoFADPPD</a></p>&mdash; m0t0k1ch1 (@m0t0k1ch1) <a href="https://twitter.com/m0t0k1ch1/status/1036953921577738240?ref_src=twsrc%5Etfw">2018年9月4日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

<blockquote class="twitter-tweet" data-lang="ja"><p lang="en" dir="ltr">I totally agree &gt;&gt; &quot;The History of Casper - Chapter 4&quot; - Vlad Zamfir <a href="https://t.co/golrO4Il9m">https://t.co/golrO4Il9m</a> <a href="https://t.co/I3wRXJwleu">pic.twitter.com/I3wRXJwleu</a></p>&mdash; m0t0k1ch1 (@m0t0k1ch1) <a href="https://twitter.com/m0t0k1ch1/status/1036965168335532032?ref_src=twsrc%5Etfw">2018年9月4日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

<blockquote class="twitter-tweet" data-conversation="none" data-lang="ja"><p lang="ja" dir="ltr">初期の Casper は、寡占市場すなわちコンセンサスが協力ゲームであることを仮定し、カルテルに対する検閲耐性を考慮することで形作られていったと。この辺りのことを知らずして cryptoeconomics は語れないな。。。反省である。</p>&mdash; m0t0k1ch1 (@m0t0k1ch1) <a href="https://twitter.com/m0t0k1ch1/status/1037017153352790018?ref_src=twsrc%5Etfw">2018年9月4日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

こんな感じで、cryptoeconomics の源流は Caspter（PoS）の歴史にあるということが改めて認識できました。そういう意味では、これも「守」の物語と言えるかなと思います。

<br />
## 最後に

以上、前回の記事を執筆した時点では気づいていなかった、cryptoeconomics の「守」に関する重要な気づきを簡単にまとめてみました。

今後、cryptoeconomics の範疇は、その源流を知らない人たちによって氾濫していくんだろうなあと思います。HipHop みたいなもんです。別に自分は原理主義者ではないので、それを悪とは思いませんが、守破離の文脈において、源流を知ることは大切です。サンプリングにおける Amen Break リスペクトみたいなもんです。[MIT Cryptoeconomics Lab](https://medium.com/mit-cryptoeconomics-lab) の記事を読みながら（全記事、自分の読書メモは [Scrapbox](https://scrapbox.io/m0t0k1ch1) に転がってますので、興味のある方はご覧ください）、そんなことを考えていたわけです。

まあ、源流を辿り出すとキリはない（結局は程度の問題）ですし、「守」を知らないからこそ生まれる斬新な発想もあるでしょう。自分も基本的にはミーハーです。が、こと cryptoeconomics に関しては本気で取り組んでいきたいので、「守」から向き合っていきたいと考えています。

ということで、自戒の意も込めたプロローグでした。おしまい。

<br />
## 次回予告

さて、ここまでの記事は抽象的な話が中心となってしまっているので、今後のアウトプットでは、できるだけ具体的な理論や事例を絡めていこうと考えています。

気が変わらなければ、次回は、メカニズムデザインにおいて重要な概念である incentive compatibility について、具体的な事例を交えながら整理する予定です。cryptoeconomics というよりは、まずはその土台となるミクロ経済学（市場理論、ゲーム理論、メカニズムデザインなど）に寄り添っていこうというスタンスです。

<br />
## おまけ

<br />
### 「cryptribe」の宣伝

CHIP で月額 100 円のファンクラブ「[cryptribe](http://thechip.in/fanclubs/419)」をつくりまして、ご加入いただいた方を「自分が日々のインプットにコメントをつけてを垂れ流すだけの個人進捗空間 Slack workspace」にご招待しております。

現在、メンバー 3 人でこじんまりと値段相応な感じでまったりやっておりますので、進捗を観察したい方は是非ご加入いただき、m0t0k1ch1 のモチベーションを支えてくださいませ🙏

ありがたいことに「Android 版がリリースされたら加入します」とおっしゃってくださっている方も数人いらっしゃいますので、とりあえずは継続を重視してムリせず運用していく予定です。

<br />
### この記事のもう 1 つの目的

これにて一旦封印解除じゃ👴

<blockquote class="twitter-tweet" data-conversation="none" data-lang="ja"><p lang="ja" dir="ltr">最近アマいので、次何か公に対して cryptoeconomic なブログ書くまで音についてのツイートを禁止とする。巧みなメタファーも禁止。like と RT も禁止。最後に言わせてください。S.W. × ジメサギのアルバム楽しみです。うっ。。。</p>&mdash; m0t0k1ch1 (@m0t0k1ch1) <a href="https://twitter.com/m0t0k1ch1/status/1037735916423372800?ref_src=twsrc%5Etfw">2018年9月6日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

我慢できず、既にこの記事の中に意味不明な記述を盛り込んでしまったことをお許しください。