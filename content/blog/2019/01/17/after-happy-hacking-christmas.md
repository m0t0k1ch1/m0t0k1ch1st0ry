+++
title = 'Happy Hacking Christmas の解答と狙い'
tags = ['dao', 'ethereum', 'solidity', 'blockchain']
date = '2019-01-17T03:29:24+09:00'
+++

昨年末、[ex-KAYAC Advent Calendar 2018](https://qiita.com/advent-calendar/2018/ex-kayac) の大トリとして執筆した [Happy Hacking Christmas]({{< ref "/blog/2018/12/25/happy-hacking-christmas.md" >}}) という記事の中で出題したスマートコントラクトパズルについて、出題からしばらく時間も経ったので、解答と狙いをまとめておこうと思います。

<!--more-->

## はじめに

本題に入る前に、パズルをプレイしてくださった皆さんに心からお礼申し上げます。

__ありがとうございました 🙏__

内心、出題直後は「つくったはいいが、、果たしてこの面倒な問題に取り組んでくださる方はいるのだろうか。。。？」と、かなり心配していたのですが、その心配もすぐさま流石の [@nakajo](https://twitter.com/nakajo) さんによってかき消され、結果的には（この記事の執筆時点で）見事 7 人の方がパズルをクリアしてくださいました（以下、8 人の SCT ホルダーのうち 1 人は自分）。

{{< figure src="/img/entry/sct-holders.png" >}}

ref. [SantaClausToken Token Holders](https://ropsten.etherscan.io/token/tokenholderchart/0xa9b76b79e3254d7835401a8b43af2fac93a83f2d)

クリアされた方々の中にははじめましての方もいらっしゃいましたし、なんと、普段は Solidity をあまり書いていないような方までいらっしゃいました。そういった、普段自分やクリプト界隈と強い関わりのない方々にもパズルに取り組んでいただけたことをとても嬉しく思います（後述しますが、この企画自体、そういった状況を生むことが狙いの 1 つでもありました）。

## 解答

<div style="width:100%;height:0;padding-bottom:56%;position:relative;"><iframe src="https://giphy.com/embed/3o7TKQTr6B72IuMlfG" width="100%" height="100%" style="width:100%;height:100%;position:absolute" frameBorder="0" class="giphy-embed" allowFullScreen></iframe></div><p><a href="https://giphy.com/gifs/actionaliens-action-bronson-friends-watch-ancient-aliens-3o7TKQTr6B72IuMlfG">via GIPHY</a></p>

自分が実装した「一発で全ての問題を解いて SCT を獲得するスクリプト」を公開しようかと思ったのですが、、やめました！笑

強いリクエストがあれば公開しようと思いますが、まあ、謎は「ほどよく」謎なままの方が魅力を保てることもあるのかなと。テスト程度の感覚で回したスクリプトで、猛者のみが有するパズルクリアの証拠 SCT の価値を無駄に毀損してしまうのも申し訳ないなと思いますし。

ただ、あくまで「ほどよく」でよいと思っています。具体的には、「公式な解答は存在しない状態」でよいのかなと。

というのは、パズルをクリアされた [@sugaret](https://twitter.com/sugaret) さんが [聖夜のsolidityクイズ！Happy Hacking Christmas 奮闘記 〜解答と解説付き〜](https://sugaret.hatenablog.com/entry/2018/12/31/184410) という記事を書いてくださっており、その中でかなり詳細な解説がなされているからこそ言えることでもあります。@sugaret さん、本当にありがとうございます 🙏🙏🙏

また、「公式な解答」をつくりたくないもう 1 つの理由として

__そもそも Blockchain が実現する世界は、良い意味でも悪い意味でも（今の世界より）絶対的正解のない世界だと思っているから__

というのがあります。

自分の解答も完璧なものでないとは思いますが、「公式な解答」という絶対的正解のようなものが存在してしまうことに、なんとなく違和感を覚えたというか、この企画自体の Blockchain らしさが失われてしまうのではないか？と。そう思ったがゆえの上記の判断でもあります。これもこの企画自体のメッセージとして受け取っていただければ幸いです。勝手なこだわりをお許しください 🙏

もちろん、解答を見ようが見まいが、自分で手を動かしながらパズルを解くことに十分意味はあると思いますので、「チャレンジしたけどクリアできなかった。。。もう潔く諦めます！」という方は、是非 @sugaret さんの記事を読みながら再チャレンジしていただければと思います。

## 各パズルコントラクトの狙い

<div style="width:100%;height:0;padding-bottom:56%;position:relative;"><iframe src="https://giphy.com/embed/2wh8A9PH93Ix93b1XC" width="100%" height="100%" style="width:100%;height:100%;position:absolute" frameBorder="0" class="giphy-embed" allowFullScreen></iframe></div><p><a href="https://giphy.com/gifs/rdr2-arthur-morgan-arthurmorgan-2wh8A9PH93Ix93b1XC">via GIPHY</a></p>

各パズルコントラクトについて、「どんなテクニックを要求するものとしたかったのか」という観点でその狙いについてまとめます。解答ではありませんが、パズルを解くためのヒントにはなってしまいますので、「まだ見たくない！！」という方はここをスキップして、興味があれば、最後の「企画自体の狙い」に目を通していただければと思います。

### Letter（低難度パズル）

💌 [Letter contract](https://ropsten.etherscan.io/address/0xbade12c0bd7943a066e77f0466d529d78d2f70db#code)

これは極めて簡単なパズルで、「基本的なトランザクション発行ができること」を要求しています。もう少し具体的には、

- 正しく送金トランザクションを発行できること
- 正しく function を実行できること

を要求していると言えます。

最初から高難度の問題にしてしまうとそもそもの参加ハードルが上がってしまい、後述する「企画自体の狙い」と逆行してしまうので、ここはできるだけ門戸を広く開くことを意識しました。

### ChristmasStocking（中難度パズル）

🧦 [ChristmasStocking contract](https://ropsten.etherscan.io/address/0x408f56c4541bd00ec836102d06f7ee6a2a820678#code)

中難度とは言え、Letter よりはかなり難易度の高いパズルです。このパズルが要求するテクニックは「Reentrancy ができること」です。

ref. [Reentrancy](https://consensys.github.io/smart-contract-best-practices/known_attacks/#reentrancy)

この Reentrancy はとても有名な攻撃手法の 1 つなので、ご存知な方も多いと思いますが、これを攻撃者の立場で実際に実行したことがある方は少ないのではないでしょうか？

今回は Reentrancy の中でも比較的簡単なパターンをパズルとして実装し、それを攻撃者の立場で実行してもらうことで、記事を読んで「へーなるほどねー」と思って終わりがちな Reentrancy の話に実体験を伴わせることを意図しました。

また、後述する「企画自体の狙い」とも関連します。

### ChristmasTree（高難度パズル）

💎 [ChristmasTree contract](https://ropsten.etherscan.io/address/0xa9b76b79e3254d7835401a8b43af2fac93a83f2d#code)

このパズルは、今回の企画を「鎖野郎諸兄姉への挑戦状」として恥じない問題とし、クリアしてくださった方にできるだけ大きな達成感を感じていただくための「おまけ」みたいなものですね笑

このパズルが要求するテクニックは「Storage Manipulation ができること」です。

ref. [Storage Manipulation](https://consensys.github.io/smart-contract-best-practices/known_attacks/#underflow-in-depth-storage-manipulation)

素直な Storage Manipulation は上記リンク先を辿って見よう見まねでやれば成功してしまってつまらないなあと思ったので、少し手を加えて、攻撃対象を array ではなく mapping にしています。こうすることで、storage の構造やハッシュ関数の扱いをきちんと理解していないと解けないパズルにできたかなと思っています。

## 企画自体の狙い

最後に、この企画自体の狙いについて述べて締めくくろうと思います。

個人的に実験してみたかったことや細かい狙いは複数あるのですが、一番の狙いに絞って述べます。一番の狙いは、端的に言うと、

__「伝えたいこと」にゲームという皮を被せることで、より多くの人が興味を持つ機会を提供すること__

でした。

<div style="width:100%;height:0;padding-bottom:42%;position:relative;"><iframe src="https://giphy.com/embed/BtX1KVvkHPp7i" width="100%" height="100%" style="width:100%;height:100%;position:absolute" frameBorder="0" class="giphy-embed" allowFullScreen></iframe></div><p><a href="https://giphy.com/gifs/mission-impossible-ethan-hunt-freakmasks-BtX1KVvkHPp7i">via GIPHY</a></p>

ちょっと脱線してしまうのですが、この企画が生まれたのは、冒頭に記載した通り ex-KAYAC Advent Calendar がきっかけです。この企画を思いつく前は、

__ex-XXX という「ある組織を卒業した人達」も含めて 1 つの [DAO](https://en.wikipedia.org/wiki/Decentralized_autonomous_organization) と考え、既存の「組織」という枠に捉われないような組織の在り方が実現できないだろうか？という問いをテーマにしたポエム色の強い記事__

でも書こうかなと思ってたんですね笑

ex-XXX な Advent Calendar は、これを書くのに適した場でもありました。が、正直、これ、小難しくて読みたいと思う人が少ないだろうなと。クリプト界隈外の Advent Calendar で記事を書き、クリプト界隈に馴染みのない方にも自分のメッセージを伝えられるチャンスなのに、DAO とか言ってる時点で、そういった方々が「自分ごと化」するのをとても難しくしてしまいます。し、これでは、[カヤックの経営理念である「つくる人を増やす」](https://www.kayac.com/vision/vision)も実現できないなと考えました。

そこで目をつけたのが、[The DAO](https://en.wikipedia.org/wiki/The_DAO_(organization)) で実際に行われた Reentrancy ハックです。つまり、今回の企画は、

__Reentrancy を面白そうなゲームに仕立てあげ、DAO を語るにあたっては避けて通れないであろう The DAO の失敗について技術的側面から興味を持ってもらうことができないか？__

という発想から生まれました。

簡単に言うと、パズルで興味を引いて、遊びながら Reentrancy について知ってもらうことで The DAO に興味を持つとっかかりをつくり、そこから DAO という概念にも興味を持ってもらえたらいいな〜と思って企画したということです。

まあ、流石にパズルを公開しただけでそんなにうまく導けるわけはないと思うので、、この記事で The DAO への橋渡しをさせてください。

ref. [The DAO 事件から１年 — 熱狂する ICO バブルと、これからの資金調達手法](https://medium.com/@amachino/the-dao-%E4%BA%8B%E4%BB%B6%E3%81%8B%E3%82%89%EF%BC%91%E5%B9%B4-%E7%86%B1%E7%8B%82%E3%81%99%E3%82%8Bico-%E3%83%90%E3%83%96%E3%83%AB%E3%81%A8-%E3%81%93%E3%82%8C%E3%81%8B%E3%82%89%E3%81%AE%E8%B3%87%E9%87%91%E8%AA%BF%E9%81%94%E6%89%8B%E6%B3%95-48a90c7c20c5)

今回の企画で、少しでも興味を持ってもらえたら幸いです。

{{< figure src="/img/entry/matchstick.jpg" >}}

__2016 年、Reentrancy ハックが起こらず、The DAO という灯火が今も燃え続けていたら、どうなっていただろうか？__

そんな問いを胸に 2019 年も精進していきたいと思います。

やったるで 💪
