+++
date = "2018-08-05T00:26:07+09:00"
tags = [ "cryptoeconomics", "economics", "blockchain" ]
title = "cryptoeconomics: crypto-backed mechanism design"
+++

[これ](https://m0t0k1ch1st0ry.com/blog/2018/07/28/cryptoeconomics) の続きです。

<!--more-->

前回の記事は「まず cryptoeconomics という言葉を認知してほしい」という意図で書いたので、敢えてその定義について踏み込んだことは書きませんでした。謎を謎のまま放置して終わるという、自分の好きな園子温監督もよく使う手です。が、これは映画ではないので、きちんと続きを書こうと思います。

![venus](/my-images/entry/venus.png)

<br />
## 0. 目次

- 1. おさらい
- 2. 結論
- 3. メカニズムデザイン
- 4. Bitcoin に学ぶ
- 5. cryptoeconomics
  - 5-1. 自分の解釈
  - 5-2. Vitalik と Vlad の意見
  - 5-3. さらなる深みへの誘い
- 6. 結び
- 7. 参考文献

<br />
## 1. おさらい

前回の記事では、

> cryptoeconomics は cryptography（暗号学）と economics（経済学）を組み合わせた新しい「学問」のこと。

<!-- -->
> 「cryptoeconomics」は、Blockchain が示した可能性を集約した言葉だと考えています。敢えて振り切って例えるなら「合理性のデザイン」、極論「物理法則のデザイン」と比喩することもできると思います。

などとは書きましたが、この説明だけでは cryptoeconomics がどんなことをする学問なのかよくわかりません。今回は、それを明確にしていこうと思います。

もちろん、これは絶対的な正解があるような話ではありません。この先の話については、鵜呑みにせず、異論・違和感があれば是非教えてください。ちなみに、自分は物理系出身でエンジニアになった人間なので、特に経済学的な観点でのツッコミがあると化学反応が起こるんじゃないかなと予想しています。

実際、前回の記事を書いた直後、[@YuKimura45z](https://twitter.com/YuKimura45z) さんが経済学的な観点から用語の整理をした記事を彗星の如く書いてくださりました。もうこれだけでも前回の記事を書いた甲斐があったなと。。ありがとうございます🙏 内容としても、自分の目線ではこういう表現はできないなあと思うことが多く、非常に助かりました。

ref. [トークンエコノミーと BFT エコノミーの比較](https://yu-kimura.jp/2018/07/29/token-bft-economy)

記事の中では、トークンエコノミーと BFT エコノミーの違いが綺麗に整理されており、

> cryptoeconomics は BFT エコノミーを研究する分野といったほうが正しく、トークンエコノミーとの対比はできない気がする。

<!-- -->
> トークンエコノミーはポジティブなインセンティブを、BFT エコノミーはネガティブなインセンティブを与えることが目的だということです。

と記述されています。非常にわかりやすいですし、トークンエコノミーと BFT エコノミーの定義については概ね同意です。ただ、自分の解釈だと、cryptoeconomics は BFT エコノミーの研究に限ったものではない（後述しますが、BFT エコノミー側の発想の方が重要だとは思います）ので、自分の意見を今回の記事で述べた上で、その辺りについて議論してみたいなと感じました。

ということで、以下、自分の考えを綴っていこうと思います。

少し冗長かつ基本的な説明が多く、cryptoeconomist の皆さんからすると物足りない内容となってしまうかなと思いますが、その辺りはご勘弁ください。

<br />
## 2. 結論

結論を先に書きます。自分が考える cryptoeconomics の定義を一文で表現してみると、以下のようになりました。

__あるシステムに対して、暗号学によって裏づけられた経済的合理性に従った選択をプレイヤーが重ねた場合に、そのシステムが自律的に所望の性質を具備・維持できるようなプロトコルをデザインすること。__

はい、非常にわかりにくい難解な文章です。以降、これの意味を掴んでいただけるように解説をしていこうと思います。

なお、先に言っておくと、この考えも生粋のオリジナルという訳ではないです。これまで [Vitalik](https://twitter.com/vitalikbuterin) をはじめとした何人かの cryptoeconomist の考えに触れながら、自分が腹落ちしていることを紡いでいった結果こうなった、という感じです。おそらく、最も近しいのは、[L4](https://l4.ventures) の co-founder である [Josh Stark](https://twitter.com/0xstark) の考えかなと思います。彼の考えは以下の記事にまとまっているので、気になった方はこちらにも目を通していただければと思います。

ref. [Making Sense of Cryptoeconomics](https://medium.com/l4-media/making-sense-of-cryptoeconomics-c6455776669)

<br />
## 3. メカニズムデザイン

前述した自分の定義について直接言及する前に、もっと大まかなイメージを掴んでいただいた方が伝わりやすいかなと思うので、まずは既存の学問分野で、自分の定義する cryptoeconomics と類似しているものについて軽く触れてみたいと思います。

それは、__メカニズムデザイン__ です。

まずは [Wikipedia](https://ja.wikipedia.org/wiki/%E3%83%A1%E3%82%AB%E3%83%8B%E3%82%BA%E3%83%A0%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3) さんから引用します。

> メカニズムデザイン（英: mechanism design）とは経済学の一分野である。資源配分や公共的意思決定などの領域で実現したい目標が関数の形で与えられたとき、その目標が自律的/分権的に実現できるようなルール（「メカニズム」とか「ゲームフォーム」とも呼ばれる）を設計することを目指している。言い換えれば、与えられた関数が要求する目標を、各プレイヤーの誘因を損なうことなく実現できるようなゲームを設計することをメカニズムデザインでは目指している。メカニズムデザインは経済学のなかでも特に社会選択理論および非協力ゲーム理論、さらには契約理論やマーケットデザインと密接な関係を持つ。

噛み砕くと、「こんな風になってほしいな〜」という目標の状態を決めて、その状態が自律的に実現されるようなゲームのルールをデザインすることを目指す経済学の一分野、という感じになるかなと思います。ルールが満たすべき基本的な性質としては [incentive compatiblity](https://en.wikipedia.org/wiki/Incentive_compatibility) などが重要ですが、今回はあまり深く突っ込まないでおこうと思います。

なお、ここで言う「ゲーム」は、[ゲーム理論](https://ja.wikipedia.org/wiki/%E3%82%B2%E3%83%BC%E3%83%A0%E7%90%86%E8%AB%96) 的な意味でのゲームであり、スマホゲームやボードゲームといった娯楽目的のゲームに限らないあらゆる戦略的状況を指します。戦略的状況というのは、簡単に言うと「自分の行動だけでなく他のプレイヤー達の行動も結果に影響を与えるような状況」のことです。そして、そんな状況の中で各プレイヤーはどのように意思決定すべきか？その結果どうなるのか？そんなことを考えるのがゲーム理論です。

では、同じくゲームを研究対象とするメカニズムデザインはゲーム理論とどう違うのでしょうか。端的に言うと、その違いは __アプローチの方向__ です。

ゲーム理論が

__与えられたゲームについて、プレイヤーの合理的な戦略やその結果について分析を行うこと__

を主目的とするのに対して、メカニズムデザインは、

__プレイヤーの合理性を仮定した場合に所望の結果が得られるようなゲームを設計すること__

を主目的とします。これは、ある特定の与えられたゲームからスタートするゲーム理論とは逆向きのアプローチなので、逆ゲーム理論と呼ばれることもあるようです。

メカニズムデザインの事例としては、Web 広告の RTB に広く利用されている [Vickrey auction](https://en.wikipedia.org/wiki/Vickrey_auction)（封印入札型のセカンドプライスオークション ）が有名です。これは、「各入札者が自分の評価値を正直に入札するのが最適な戦略となる」ようなオークションの仕組みとしてデザインされたそうです。言い換えれば、「各入札者が自分の評価値を正直に入札するのが最適な戦略となる」という、オークションに具備させたい「所望の性質」からスタートし、それを満たすようなルール（入札は封印型で行い、最高の入札額で入札した人が二番目に高い入札額で落札する）をデザインした事例と言えます。

さて、メカニズムデザインのイメージはなんとなく掴んでいただけたかと思いますので、一旦、冒頭に記載した自分の定義を再掲します。

> あるシステムに対して、暗号学によって裏づけられた経済的合理性に従った選択をプレイヤーが重ねた場合に、そのシステムが自律的に所望の性質を具備・維持できるようなプロトコルをデザインすること。

今改めて読むと、言っていることはメカニズムデザインとほぼ同じということが分かっていただけるのではないでしょうか。おそらく、大きく異なるのは「暗号学によって裏づけられた」という部分のみです。確かに、メカニズムデザインの説明の中に暗号学の話は出てきていません。

この部分については、Bitcoin に目を向けるとその意味が自然と分かると思います。

<br />
## 4. Bitcoin に学ぶ

言わずもがな、Bitcoin のネットワークはマイナーが中心となって支えています。マイナーはマイニング報酬を目当てにマイニングをしています。報酬がなければマイニングに参加する人はほとんどいないでしょう。

まず、この観点から、Bitcoin は経済的インセンティブを利用したプロトコルだと言えます。

しかし、そもそもマイニング報酬に価値を感じてもらえなければインセンティブとしては機能しないでしょう。報酬がデータならば、誰でもコピペや改ざんが可能なものであってはいけませんし、攻撃に対して脆弱なシステムに保持されていてもいけません。これらは、報酬が満たすべき要件と言えます。

Bitcoin は、これらの要件を満たすために、（広義の）暗号技術を利用しています。

最も特徴的なのは、暗号学的ハッシュ関数の性質を利用してブロックを発掘するための総当たり計算に費やされた計算量（work）を証明し、最も work が費やされた（最も長い）チェーンを正として採用していく仕組みでしょう。すなわち、PoW（Proof of Work）をベースとしたナカモトコンセンサスです。これによって、ハッシュレートが大きくなるほど 51% 攻撃のコストが高まるため、マイナーのネットワークが拡大するほど不正なチェーンを伸ばし続けることが困難となります。余談ですが、この記事を書いている 2018 年 8 月 5 日時点では、51% 攻撃を行うために [約 90 億ドル分のハードウェアと約 620 万ドルの電気代（1 日あたり）が必要](https://gobitcoin.io/tools/cost-51-attack) なようです。

他にも、UTXO（unspent transaction output）を使用する際には、楕円曲線暗号を利用した電子署名によってその所有者であることを証明させることで、他人が勝手に UTXO を使用することを防いでいます（これは Bitcoin スクリプトの 1 パターンでしかありませんが）。また、前述したハッシュは、ブロックチェーンの順序（過去に蓄積したデータの整合性）を保証するためにも利用されています。

__すなわち、Bitcoin はそのプロトコルの中核に暗号技術が複合的に組み込まれており、これらに由来する堅牢なセキュリティによってマイニング報酬がインセンティブとして成立していると解釈できます。そして、このインセンティブに対してマイナーが合理的に反応することでネットワークが維持されているとも言えるでしょう。__

まさに、このようなプロトコルをデザインすることが cryptoeconomics です。

<br />
## 5. cryptoeconomics

<br />
### 5-1. 自分の解釈

もう一度、冒頭に記載した自分の定義を振り返ってみます。

> あるシステムに対して、暗号学によって裏づけられた経済的合理性に従った選択をプレイヤーが重ねた場合に、そのシステムが自律的に所望の性質を具備・維持できるようなプロトコルをデザインすること。

もう説明することはほとんどないかなと思いますが、要点だけ整理しておきます。

- 基本的にはメカニズムデザインと同様、所望の結果から逆算してルールをデザインする
- 所望の結果へと導くために、暗号学に基づいて生み出したインセンティブを利用する

これをさらに端的に表現するならば、

__crypto-backed mechanism design__

と言ってもよいかなと思います（この記事のタイトルです）。現状、自分なりに cryptoeconomics を解釈した結論がこれであり、これが __Satoshi Nakamoto によって示された最も重要なコンセプト__ だと考えています。

また、前回の記事で、

> トークンという概念も、cryptoeconomics を考える材料の 1 つに過ぎません。

とは書きましたが、cryptoeconomics を考える上で Blockchain という概念すらも必須ではないということに注意してください。もちろん、Bitcoin や Ethereum を始めとしたパブリックな Blockchain は cryptoecnomic なプロダクトの代表と言えますが、例えば、[IOTA](https://www.iota.org) や [Byteball](https://byteball.org) といった [DAG](https://ja.wikipedia.org/wiki/%E6%9C%89%E5%90%91%E9%9D%9E%E5%B7%A1%E5%9B%9E%E3%82%B0%E3%83%A9%E3%83%95) ベースのプロダクトも cryptoeconomics の範疇です（が、これらが本当に所望の性質を具備・維持できるプロトコルなのかは別の話です）。

あと、細かいニュアンスでまだ伝えられていないものがあるとすると、「経済的インセンティブ」ではなく「経済的合理性」と記載したあたりでしょうか。

「経済的インセンティブ」と言ってしまうと、どうしても、ポジティブな動機づけ（好ましい行動を合理的にする）のイメージが先行してしまいます。一方、ネガティブな動機づけ（好ましくない行為を非合理的にする）も同等もしくはそれ以上に重要です。自分としては、これらを組み合わせてプロトコルをデザインしていく必要があると考えているので、定義としてはこの両面性を表現したかったわけです。

自分がネガティブな動機づけを重要視している理由については、同様の見解を述べている記事を見つけたので、それを引用しようと思います。これまた [@YuKimura45z](https://twitter.com/YuKimura45z) さんの記事です笑

ref. [トークンエコノミーは分散化できるか](https://yu-kimura.jp/2018/08/03/decentralized-token-economy)

> BFT エコノミーの場合、不正な挙動（ビザンチンなふるまい）をネガティブインセンティブによって防げればよいわけですから、不正でも正常でもない無関心・無関与に困ることはありません。一方でポジティブインセンティブによって行動を誘導したい（ALIS なら良い記事を書くなど）トークンエコノミーにおいては、無関心・無関与でいられると困ります。行動を誘導したいわけですから笑

所望の性質を具備・維持するにあたり、ポジティブな動機づけが要となる場合は無関心・無関与の人がいると困るわけですが、ネガティブな動機づけが要となる場合は無関心・無関与の人がいても特に困らないわけです。これは cryptoeconomic なプロトコルをデザインする上で非常に重要な視点だと思います。

<br />
### 5-2. Vitalik と Vlad の意見

さて、自分の解釈についての説明はだいたい終わったので、別の意見にも触れてみようと思います。

まずは [Vitalik](https://twitter.com/vitalikbuterin)。以下の資料の中からスライドを 1 枚引用します。

ref. [Blockchain and Smart Contract Mechanism Design Challenges](https://fc17.ifca.ai/wtsc/Vitalik%20Malta.pdf)

![vitalik_1](/my-images/entry/vitalik_1.png)

大枠の考え方は自分の定義とそこまで変わらないと思うのですが、いくつか明確に異なる点もあります。

まず、__to achieve information security goals__ と表現しているところに色があります。自分の定義では、「所望の性質を具備・維持できるような」と、より抽象的に表現していますが、それよりも範囲が絞られています。確かに、既存の cryptoeconomic なプロダクトは、セキュリティに関する要件が所望の性質として非常に重要になっている、というか、ほとんどの所望の性質の土台となっていると思うので、これは極めて妥当な色づけだと思います。が、自分としては、まだセキュリティに限定してよいという確信が持てなかったので、抽象的な表現にとどめています。

経済的インセンティブにも __defined inside a system__ と色づけしています。これは、裏を返せば「システム外のインセンティブに依存してはいけない」ということなので、プロトコルの設計方針に大きく影響する要件と言えるでしょう。とはいえ、インセンティブに限らずシステム外の何かを利用する場合、内と外との境界線部分の設計は高難度になりがち（ex. 分散型オラクル）なので、これも妥当な色づけだなと思います。実際、Bitcoin がインセンティブとして使っているのは、システム内で定義された UTXO であり、ここに美しさを感じている方も多いでしょう。自分としても、ここに関しては同意なので、定義の中に「自律的に」という言葉を含めました。

また、時間軸を過去と未来に分けて語っている点も気になります。ここは確かにわかりやすい切り分けだとは思うのですが、敢えて切り分けた深い意図があるとするなら、自分はまだそれを掴めていません。暗号学ありきでトークンが存在していて、それがインセンティブとして利用できると考えると、暗号学は未来にも関与していると言えるのではないか？と思ってしまいます。ここについて、何か意見がある方は是非教えてください。

次は [Vlad](https://twitter.com/VladZamfir)。彼の 2015 年のツイートを引用します。

前提として、Vitalik も先ほどのスライドの下部で credit to Vlad Zamfir for this characterization と述べているので、そもそもこの 2 人の意見の根本は同じなのだと思います。

<blockquote class="twitter-tweet" data-lang="ja"><p lang="ca" dir="ltr">cryptoeconomics is economics for cryptographers, not cryptography for economists.</p>&mdash; Vlad Zamfir (@VladZamfir) <a href="https://twitter.com/VladZamfir/status/559181087658631168?ref_src=twsrc%5Etfw">2015年1月25日</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

おい。カッコよすぎか。2015 年やぞ。どういうことやねん。

すみません。取り乱しました。この発言の意図を汲み取るには、ここで言うところの cryptographer と economist の目的をどう解釈するかが重要だと思います。

まず、cryptographer の目的について。例えば、先に述べた Vitalik の意見にもあるように、information security が実現されているシステムの構築が cryptographer の目的だとすると、これはしっくりです。彼の目線から見たら、__cryptographer がその目的達成の手段として economic incentive という新しい武器を手にした__ というイメージなのかなと思います。

次に、economist の目的について。これについては、前述した Josh の記事にちょうどいい記述があったので引用します。

ref. [Making Sense of Cryptoeconomics](https://medium.com/l4-media/making-sense-of-cryptoeconomics-c6455776669)

> Economics is the study of choice: how people and groups of people respond to incentives.

Vlad も economist の目的をこう捉えているのであれば、しっくりです。また、ここで言うところの economics にメカニズムデザインは含まれていないと考えてよさそうです。

さらに興味深い文章が続くので、それも引用しておきます。

> The invention of cryptocurrency and blockchain technology does not require a new theory of human choice — the humans haven’t changed. Cryptoeconomics is not the application of macroeconomic and microeconomic theory to cryptocurrency or token market

確かに、cryptoeconomics をこう勘違いされたら困ります笑

こういった economist 視点での解釈がズレているのは多くの同意が得られると思うので、議論の余地はほぼないと思うのですが、自分が難しいなと感じるのは、__cryptographer 視点で解釈するか、mechanism designer 視点で解釈するか__ です。前述した通り、information seculity という重要かつ明確なゴールを持った前者が極めて妥当だなとは思うのですが、、今回は敢えて以下のように呟いておきました。

<blockquote class="twitter-tweet" data-partner="tweetdeck"><p lang="en" dir="ltr">cryptoeconomics is cryptography for mechanism designer</p>&mdash; m0t0k1ch1 (@m0t0k1ch1) <a href="https://twitter.com/m0t0k1ch1/status/1025865045056671744?ref_src=twsrc%5Etfw">August 4, 2018</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

<br />
### 5-3. さらなる深みへの誘い

長々とポエムのような解説を連ねてはきましたが、ここまでで言及できているのは cryptoeconomics の入口部分に過ぎません。

例えば、実際に cryptoeconomic なプロトコルをデザインする場合、その安全性や攻撃耐性についての多角的な評価は必要不可欠と言えます。この評価に関しては、[__griefing factor__](https://ethresear.ch/t/a-griefing-factor-analysis-model/2338)（他プレイヤーが失う利得と自分がそのために支払うコストの比率）や [__P + ε attack__](https://blog.ethereum.org/2015/01/28/p-epsilon-attack)（賄賂によってプレイヤーの支配戦略を変えてしまう攻撃）などの重要な概念をベースにした議論が既に繰り広げられており、例えば [__初期状態がナッシュ均衡かどうかに着目した安全性評価__](https://github.com/zack-bitcoin/amoveo/blob/master/docs/design/cryptoeconomics.md) などが提案されています。この記事では深く突っ込みませんが、個々のトピックについて深掘りするだけで余裕で 1 記事書けてしまうでしょう。

また、そもそも論ではありますが、__人間は [ホモ・エコノミクス](https://ja.wikipedia.org/wiki/%E7%B5%8C%E6%B8%88%E4%BA%BA) ではありません。__いくら綿密に合理性をデザインしたとしても、それに従って素直に動くような生き物ではないと仮定するのが妥当でしょう。cryptoeconomics と向き合う際、このことも重々頭に入れておく必要があります。

このような文脈で、cryptoeconomics の発展には [__行動経済学__](https://ja.wikipedia.org/wiki/%E8%A1%8C%E5%8B%95%E7%B5%8C%E6%B8%88%E5%AD%A6) の知見が必要だと主張している記事もあります。その記事から図を 1 つ引用します。

ref. [Behavioral Crypto-Economics: The Challenge and Promise of Blockchain Incentive Design](https://medium.com/berlin-innovation-ventures/behavioral-crypto-economics-6d8befbf2175)

![bce](/my-images/entry/bce.png)

記事の中では、automatibility（インセンティブに従うために人間の手作業がどれだけ必要か）と size of action space（報酬を最大化するためにできる行動がどれだけあるか）という 2 つの軸で既存の cryptoeconomic なプロダクトが評価されており、それを踏まえて

__図の左下に位置する Bitcoin が成功したからといって、全てが成功するわけではない。もっと人間の非合理性と向き合おう。インセンティブの力を過信してはいけない。__

といった主張が述べられています。

と、ここだけ抜き出すとネガティブに聞こえてしまいますが、著者の cryptoeconomics に対するスタンスはどちらかと言うとポジティブです。その可能性を示唆しながらも冷静に現状整理と課題提起を行っており、個人的にも共感できる内容が多かったので、気になった方は是非記事を読んでみてほしいなと思います。

<br />
## 6. 結び

以上、メカニズムデザインの考え方をベースにして自分なりの cryptoeconomics の解釈について述べた上で、Vitalik や Vlad の意見と比較しながら考察を行ってみました。また、いくつかの発展的なトピックについて、その触りをご紹介し、cryptoeconomics の深さを伝えようと試みました。冒頭でも述べた通り、異論・違和感があれば是非教えていただければと思います。

最後に。

今の自分の主たる興味の対象についても軽く言及しておこうと思います。それは、

__「所望の性質」をどう定義するか__

です。

自分の解釈に従うのであれば、どんなプロトコルをデザインするにあたっても、この定義から始めないといけません。そして何より、__所望の性質には、設計者の哲学を含める余地が大いにある__ と考えています。実現したい世界の姿そのものと言っても過言ではないでしょう。

前回の記事でも少し言及しましたが、自分はこれを [__複雑ネットワーク__](https://ja.wikipedia.org/wiki/%E8%A4%87%E9%9B%91%E3%83%8D%E3%83%83%E3%83%88%E3%83%AF%E3%83%BC%E3%82%AF) の状態としてうまく定義できないか、悶々と考えています。複雑ネットワークの形成過程をゲーム理論の枠組みを用いてモデル化しようとする __ネットワーク形成ゲーム（network formation game）__ という試みは既に存在していますので、これを cryptoeconomics に応用するのは __理論的には__ 可能だと考えています（もちろん、ゆくゆくは社会実装までもっていきたいです）。この辺りの話についても考えがまとまってきてはいますので、徐々にアウトプットしていければと思います。

それでは今回はこの辺で。Vitalik という猛虎の威を借りて締めくくらせていただきます。

ref. [Blockchain and Smart Contract Mechanism Design Challenges](https://fc17.ifca.ai/wtsc/Vitalik%20Malta.pdf)

![vitalik_2](/my-images/entry/vitalik_2.png)

<br />
## 7. 参考文献

- [トークンエコノミーと BFT エコノミーの比較](https://yu-kimura.jp/2018/07/29/token-bft-economy)
- [Making Sense of Cryptoeconomics](https://medium.com/l4-media/making-sense-of-cryptoeconomics-c6455776669)
- [トークンエコノミーは分散化できるか](https://yu-kimura.jp/2018/08/03/decentralized-token-economy)
- [Blockchain and Smart Contract Mechanism Design Challenges](https://fc17.ifca.ai/wtsc/Vitalik%20Malta.pdf)
- [Behavioral Crypto-Economics: The Challenge and Promise of Blockchain Incentive Design](https://medium.com/berlin-innovation-ventures/behavioral-crypto-economics-6d8befbf2175)

本文中では引用しませんでしたが、cryptoeconomic なメカニズムデザインの考え方についてもっと深く知りたい方は、以下の記事に進んでいただくとよいかなと思います。ちょうど今、自分もこの記事の内容を消化しようと奮闘しているところです。

ref. [A Crash Course in Mechanism Design for Cryptoeconomic Applications](https://medium.com/blockchannel/a-crash-course-in-mechanism-design-for-cryptoeconomic-applications-a9f06ab6a976)
