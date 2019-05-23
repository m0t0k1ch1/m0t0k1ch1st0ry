+++
title = 'トークンエコノミーとの関係性から考える cryptoeconomics'
tags = ['cryptoeconomics', 'blockchain']
date = '2018-07-28T00:59:36+09:00'
+++

最近は [Scrapbox](https://scrapbox.io/m0t0k1ch1) に引きこもってインプットに特化しがちだったので、、そろそろ日向に出て、少しずつ cryptoeconomist 的なアウトプットを出していこうかなと思います。

<!--more-->

いきなりですが、表題で言及している「cryptoeconomics とトークンエコノミーの関係性」について、自分の解釈をできるだけ端的に綴ってみると、以下のようになります。

__「cryptoeconomics」は、その名の通り、cryptography（暗号学）と economics（経済学）を組み合わせた新しい「学問」のこと。「トークンエコノミー」は、その名の通り、トークンが循環する新しい「経済圏」のこと。つまり、トークンエコノミーは、cryptoeconomics という学問から生まれた 1 つの経済の在り方を形容しているに過ぎません。また、トークンという概念も、cryptoeconomics を考える材料の 1 つに過ぎません。__

その名の通りのことを何を今更と感じる方もいらっしゃるとは思いますが、、そもそも cryptoeconomics という言葉自体に馴染みのない方も結構多いのではないでしょうか？

さて。誰が流行らせたのかは知らないのですが、日本の Blockchain 界隈では「トークンエコノミー」という言葉をよく見かけるような気がしますよね。一方、世界に目を向けてみると、「token economy」という言葉を見かけない代わりに、似たような響きの「cryptoeconomics」という言葉をよく見かけます。勢いのある Ethereum 界隈について軽く調べてみるだけでも、最先端にいるプレイヤーである [Vitalik](https://twitter.com/vitalikbuterin) や [Vlad](https://twitter.com/vladzamfir) や [Karl](https://twitter.com/karl_dot_tech) などが cryptoeconomist の筆頭であり、みんな cryptoeconomics について議論している様子が見てとれます。

- [Introduction to Cryptoeconomics - Vitalik Buterin](https://www.youtube.com/watch?v=pKqdjaH1dRo)
- [What Is Cryptoeconomics - Vlad Zamfir](https://www.youtube.com/watch?v=9lw3s7iGUXQ)
- [Programmable Incentives: Intro to Cryptoeconomics - Karl Floersch](https://www.youtube.com/watch?v=-alrVUv6E24)

cryptoeconomics の定義はさておき、自分は昨年末 〜 今年の頭頃にかけてこの状況を認識し始めまして、そこでまず感じたことは「ん？日本、焦点ズレてる。。。？」でした。ズレていることが良いのか悪いのかはさておき、上記のような最先端のプレイヤー達が cryptoeconomics について盛んに議論しているのに、日本でそんな言葉は耳に入ってこないぞと（当時は自分の観測範囲も狭かったですが）。その代わり、トークンエコノミーという似たような言葉は耳に入ってくるぞと。今でもこの状況は大きく変化していないと感じます（とか言うと日本の数少ない cryptoeconomist 勢に怒られそうですが、、敢えて言っておきます）。

別にトークンエコノミーを真っ向からディスりたいわけではなく（[PoliPoli](https://www.polipoli.work) とか、めっちゃ応援してまっせ！！）、単純に上記のような状況にモヤモヤしていたわけです。もし cryptoeconomics が重要な概念だとすると（後述しますが、自分は非常に重要だと考えています）、日本においてトークンエコノミーという言葉が先行して広まってしまったがゆえに cryptoeconomics と向き合う人が生まれにくくなっているのであれば、それは大いなる機会損失を生んでしまっているのではなかろうかと。

実際、例えば最近の Ethereum 関連で議論が盛んなトピックで言うと、[Casper](https://ethresear.ch/c/casper) のプロトコル設計も cryptoeconomics の範疇ですし、[Plasma](https://ethresear.ch/c/plasma) のプロトコル設計も cryptoeconomics の範疇と言えるでしょう。また、これらは今後の Ethereum の発展にクリティカルに関与する極めて重要なプロトコルであり、日本でも興味を持っている方々は多いと思います（かく言う自分も Plasma については相当量のリサーチを行い、その重要性を肌で感じている身です）。が、これらをトークンエコノミーの範疇と言うのはちょっとイメージが違うような感じがします。

__このような状況を受けて、まずは cryptoeconomics というキーワードの存在だけでもいいから、多くの人に認知してもらう必要性を感じ続けてきました。この記事を書いているのも、その認知が目的です。__

特に今、日本の Blockchain 界隈では、[CryptoAge](https://twitter.com/cryptoage_) を筆頭に若い世代のエネルギーが渦巻いています。これから世界と向き合いながら新しい時代をつくっていく彼ら彼女らの世代が、上の世代が流行らせた言葉で機会損失する可能性があるとするなら、それはあかんやろという想いもあります。どうせみんなで暴れるなら Vitalik 達と同じフィールドで暴れようぜ？と。まあ、自分もまだ 20 代で人のこと言うてる場合ではないので、もちろんプレイヤーとして負けないように精進しますが。

最後に。

__「cryptoeconomics」は、Blockchain が示した可能性を集約した言葉だと考えています。敢えて振り切って例えるなら「合理性のデザイン」、極論「物理法則のデザイン」と比喩することもできると思います。これは、「え？トークンを使ったインセンティブ設計の話でしょ？」という表面的なニュアンスを汲み取って思考停止するにはあまりにも勿体ないテーマですし、学術的な研究対象としても十分に興味深いテーマだと考えています。__

おそらく自分が今の状態で大学生に戻ったら、cryptoeconomics を軸にして Ph.D. を取りに行くと思います。実際、今、自分は、ゲーム理論の枠組みを用いて複雑ネットワークの形成過程をモデル化する network formation game の考え方を、cryptoeconomics 的なプロトコル設計に応用できないか考えています。

ここまで読んで、おい！結局 cryptoeconomics って何なんだ！教えろ！と思ってくださった、あなた。ありがとうございます。あなたの存在でこの記事の目的は最低限達成です。

自分が cryptoeconomics について考えていることはおいおい別でアウトプットしていこうと思いますので、お待ちください。それでは今回はこの辺で。

## ※追記（2018-08-11）

2018-08-05 に続きを書きました。

ref. [cryptoeconomics: crypto-backed mechanism design]({{< ref "/blog/2018/08/05/cryptoeconomics.md" >}})
