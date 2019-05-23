+++
title = 'Catalyst から紐解く Enigma ①'
tags = ['blockchain', 'enigma']
date = '2017-11-29T01:59:10+09:00'
+++

Bitcoin さんに SegWit が導入されたこともあり、Lightning Network や Raiden に代表されるような Layer 2 関連のトピックはもちろんのこと、atomic swap 関連のトピックも盛り上がりを見せている今日この頃。DEX なども含め、来年はオフチェーンなソリューションやクロスチェーンなソリューションが一気に加速する年になりそうな雰囲気がプンプンですね。理由はさておき、このエントリでは数あるトピックの中から Enigma プロジェクトをピックアップし、Enigma プロトコルを実装する最初のアプリケーションである Catalyst から Enigma を紐解いていこうかなと思います。

<!--more-->

## 前置き

[Enigma プロジェクト](https://www.enigma.co) 自体は 2015 年の時点で公開されており、自分も去年 [ホワイトペーパー](https://www.enigma.co/enigma_full.pdf) をざっと読んで「ははーん」と思って放置していたのですが、今年に入って Catalyst という 1st アプリケーションの提案と併せて ENG トークンの ICO が行われるなど、新しい動きが見え始めてきました（昨今の ICO ブームに乗らねばという気持ちも少なからずあったのでしょう）し、知らない間に Medium にも情報が増えていた模様です。ということで、改めて現在の Enigma プロジェクトの動向をきちんと把握しておこうかなと思ったというわけです。

で、このエントリで何をするかというと、まずは技術的な詳細にはあまり踏み込まず、Catalyst のホワイトペーパーのイントロや結論あたりをざっくり日本語に訳して、プロジェクトの全体像や Catalyst の位置付けを再確認していこうかなと。モチベーションが続いたら技術面にも踏み込んでいこうかなと思います。タイトルに ① とか書いてしまったのはその辺りを加味してですが、まあ、正直 ② が生まれるかに関しては不安でいっぱいです。

それでは、以下、Catalyst のホワイトペーパーである [Enigma Catalyst: A machine-based investing platform and infrastructure for crypto-assets](https://www.enigma.co/enigma_catalyst.pdf) から、

- Abstract
- I. INTRODUCTION
- IV. CONCLUSIONS

を抜粋して訳していこうと思います。結構意訳してるところや専門用語の訳し方が間違ってるところなどあるかと思いますが、ご容赦ください。

---

## Abstract

> Inspired by the rapid growth and proliferation of crypto-assets, we propose Catalyst – the first investment platform that enables developers to build, test, and execute micro crypto-funds. Through Catalyst, developers can access Enigma’s decentralized data marketplace protocol [15] and consume valuable crypto-data that can be used in their strategies. Catalyst is therefore the first application to be deployed on top of the Enigma protocol.

暗号資産の急速な成長と普及に触発された我々は、開発者がマイクロ暗号ファンドを構築・テスト・実行できるようにする最初の投資プラットフォームである Catalyst を提案する。Catalyst を通じて、開発者は Enigma 分散型データ市場プロトコルにアクセスし、投資戦略の中で使用可能な貴重な暗号データを利用することができる。すなわち、Catalyst は Enigma プロトコル上に展開される最初のアプリケーションである。

---

## I. INTRODUCTION

> Algorithmic trading and machine learning are proving to be disruptive trends in investment management. From 2009 to 2015 alone, the amount of assets under management (AUM) by quantitative hedge funds grew at a rate of 14% year-over-year, nearly double the 8% year-over-year growth of assets managed by traditional hedge funds. The traditionally opaque and secretive asset management industry is also being challenged by more egalitarian access to financial data, which has successfully enabled the development of crowd-sourced investment strategies. Moreover, the barriers to enter algorithmic trading are swiftly being dismantled, offering new investment opportunities to a burgeoning open source community of developers, quants, traders, and investors.

アルゴリズム取引と機械学習は、投資管理において破壊的なトレンドであることが証明されつつある。2009 年から 2015 年にかけて、定量戦略を採用するヘッジファンドの運用資産（AUM）の規模は前年比 14 % ずつ増加している。これは、伝統的なヘッジファンドの運用資産の年次成長率である 8 % のほぼ 2 倍の数字である。伝統的で不透明かつ秘密主義な資産運用を脅かすものはこれだけではない。金融データへのアクセスがより平等となることで、クラウドソース型の投資戦略をうまく開発することが可能となったのである。さらに、アルゴリズム取引を行うための障壁は急速に取り除かれつつあり、急成長するオープンソースコミュニティの開発者、金融アナリスト、トレーダー、投資家に対して新しい投資機会を提供している。

> Following the rising demand for crypto-currencies, we believe an interesting opportunity arises: algorithmic trading on crypto-assets. To be fair, many exchanges offer the ability to place orders through RESTful APIs, permitting users to run their trading algorithms locally. However, traders are currently forced to develop the infrastructure for development, testing, and deployment of their trading strategies. These systems involve an inordinate amount of complexity, data curation, and otherwise impose a significant barrier to safely begin experimentation with algorithmic trading of crypto-currencies.

暗号通貨の需要が高まるにつれ、おもしろいチャンスが生まれるだろう。暗号通貨のアルゴリズム取引である。公平を期すため、多くの取引所は RESTful な API を通じて注文を行えるようにし、ユーザーがローカル環境で独自の取引アルゴリズムを実行することを許容している。しかし、トレーダーたちは独自の取引戦略を開発し、テストし、展開するためのインフラストラクチャの開発を余儀なくされている。これらのシステムは過度な複雑さやデータキュレーションを伴い、暗号通貨のアルゴリズム取引を安全に実験開始するにあたって、大きな障壁となっている。

> Like many who are passionate about the opportunities in the crypto-space, our mission is to increase the adoption of crypto-assets. We are building a tool that makes it easier to make educated investment decisions in crypto-assets, based on a data-driven approach. Catalyst is a set of applications and the infrastructure to drive better investment strategies, hence increasing the adoption of crypto-assets.

暗号通貨界隈におけるチャンスに熱烈な多くの人々と同様、我々のミッションは暗号資産をより広く普及させることである。我々は、データ駆動型のアプローチに基づき、知識・経験に基づいた暗号資産の投資判断をより簡単に行えるようにするツールを構築している。Catalyst はより優れた投資戦略を推進するための一連のアプリケーションであり、インフラストラクチャである。これにより、暗号資産が広く普及する。

> More importantly, we see Catalyst as the first application to be deployed on the data marketplace protocol we laid out in our previous work ([15], [16]), which we recently revisited in [17]. Our overarching goal is to create a decentralized, open and secure data marketplace protocol for the web, that is set to change how data is aggregated, shared and monetized.

さらに重要なのは、我々が以前の研究の中で提案したデータ市場プロトコルの上に展開される最初のアプリケーションが Catalyst であるということである。これについては、先日 [Towards a Decentralized Data Marketplace — Part 2](https://blog.enigma.co/towards-a-decentralized-data-marketplace-part-2-1362c8e11094) にて再検討を行った。我々の最も重要な目標は、分散型でオープンかつセキュアなデータ市場プロトコルを web 上に構築することである。これにより、データの集約・共有・マネタイズ方法が変化する。

### A. Related Work

> Investing in crypto-assets, namely applications and exchanges that facilitate trading, is a fast growing area in the blockchain space. ICONOMI is a centralized crypto-investment platform, where a user invests through the service in a crypto-index-fund that tracks multiple assets. Prism, backed by Shapeshift, operates using a semi-centralized model of a similar concept. The user deposits funds into a smart contract, that replicates a Contract for Difference, and specifies the assets it wishes to simulate holding. As the market-maker, Shapeshift holds the real assets on behalf of the user, and allows the user to withdraw the assumed returns on their virtual portfolio. In both cases, custody of the true underlying assets remains in the control of a single entity users must trust.

暗号資産への投資、つまり、取引を促進するアプリケーションや取引所への投資は、ブロックチェーン分野で急速に拡大している。ICONOMI は、中央集権的な暗号投資プラットフォームであり、ユーザーはサービスを通して複数の資産を追跡できる暗号インデックスファンドに投資を行う。ShapeShift によって運営される Prism は、同様のコンセプトを半中央集権的なモデルを用いて実践している。ユーザーは差金決済取引（CFD）を再現するスマートコントラクトに資金を預け、シミュレーション保有したい資産を指定する。マーケットメーカーとして、ShapeShift はユーザーに代わって実際の資産を保持し、ユーザーが仮想ポートフォリオで定められたリターンを引き出せるようにする。ICONOMI も ShapeShift も、裏付けとなっている資産の管理が単一機関の制御下にあり、ユーザーはこれらの機関を信頼しなければならない。

> Recently announced decentralized on-chain investment solutions such as 0x, Melonport and Bancor, face a performance issue that limits their utility in trading applications. Since all transactions require on-chain settlement, the speed of these systems lag behind that of centralized ones. In addition, since funds will be locked up for a longer period of time, on-chain settlement may lead to liquidity problems. Bancor attempts to overcome this concern with an automated market-maker function that is not based on supply and demand, but reportedly these are targeted for niche currencies that are not frequently traded [3]. Another limitation of the aforementioned on-chain protocols is that they currently only support ERC20 [4] compatible tokens, which leaves out more than half of the crypto-assets in circulation.

0x、Melonport、Bancor といった、最近発表された分散型のオンチェーン投資ソリューションは、パフォーマンス上の問題に直面しており、これにより取引アプリケーションでの利用が制限されてしまう。全てのトランザクションがオンチェーンでの決済を必要とするため、これらのシステムの速度は中央集権型のシステムよりも遅い。加えて、資産は長期間ロックアップされてしまうため、オンチェーン決済は流動性に関する問題を引き起こす可能性がある。Bancor は、需要と供給に基づかない自動化されたマーケットメーカー機能によってこの懸念を克服しようとしているが、これらは頻繁に取引されないニッチな通貨を対象としたものである。前述したオンチェーンプロトコルのもう 1 つの制限は、現在 ERC20 に準拠したトークンのみをサポートしていることである。これは、流通する暗号資産の半分以上を考慮していないということを意味する。

> Finally, while not developed for the crypto-market, Quantopian is the leading platform that lowered the barriers to become a quantitative trader by providing a tool that enables developers to build, test and execute trading strategies. Based on how successful this product has been in traditional markets, we are expanding on the existing work of Quantopian to enable developers to create successful crypto-asset trading strategies.

最後に、暗号市場のために開発されたものではないが、Quantopian は、開発者が取引戦略を構築・テスト・実行するためのツールを提供することで、定量的なトレーダーになるにあたっての障壁を引き下げたプラットフォームの筆頭である。開発者が暗号資産取引戦略をうまく構築できるようにするため、我々は Quantopian が従来の市場でどのように成功したかを参考に、Quantopian が成したことを拡張している。

### B. Our Contributions

> Addressing the aforementioned challenges, we propose Catalyst, an investment platform that allows anyone to build their own crypto hedge-fund.

前述した挑戦に取り組むため、誰もが独自の暗号ヘッジファンドを構築できる投資プラットフォーム Catalyst を提案する。

> Our main contribution is creating the first application to be built on top of decentralized data marketplace protocol, where data is exchanged and monetized in a peer-to-peer network.

我々の主要な貢献は、分散型データ市場プロトコルの上に最初のアプリケーションを実装することである。そこでは、P2P ネットワークを介してデータが交換・マネタイズされる。

> A second, related contribution is that of standardizing data for the Blockchain ecosystem. Currently, given that the ecosystem surrounding crypto-markets is still in its early days, relevant data sources are scarce and fragmented. We attempt to improve upon the existing status quo by identifying several key data-sets that we intend to curate and make available to anyone using our platform. More importantly, given the open nature of the data marketplace protocol, we believe that the long-tail of data aggregated by the community will quickly surpass in size any central repository that exists today.

第 2 の関連する貢献は、ブロックチェーンのエコシステムにおけるデータを標準化することである。現在、暗号市場を取り巻くエコシステムはまだ初期段階であるため、関連するデータソースは乏しく、断片化している。我々は、鍵となるいくつかのデータセットを特定、キュレートし、我々のプラットフォームを利用する誰もが利用可能にすることで、現状を改善しようとしている。さらに重要なのは、データ市場プロトコルがオープンなものであった場合、コミュニティによって収集されたデータのロングテールが、現存する中央集権型のリポジトリのサイズをすぐに超えてしまうということである。

> A third contribution relates to a proposed architecture for a decentralized crypto exchange that does not require a custodian. While not our primary focus, as we were developing Catalyst we noticed how unscalable, not to mention insecure, existing exchanges are. We therefore decided to propose a better infrastructure for the community, with the hope that this idea will lead to further research on the subject. Our proposed solution can operate as an extension to existing off-chain payment networks built on bidirectional payment channels and hashed timelock contracts (HTLCs), such as the Lightning or proposed Raiden network. This design allows users to make fast, cross-chain transfers while maintaining full custody of their assets.

第 3 の貢献は、管理者を必要としない分散型暗号取引のために提案したアーキテクチャに関するものである。これは主な関心事ではないが、我々は Catalyst の開発中に、既存の取引所がいかにスケーラブルでないか気づいた（安全でないことは言うまでもなく）。よって、我々はコミュニティのためによりよいインフラストラクチャを提案することを決めた。このアイデアが、このテーマに関するさらなる研究のきっかけとなることを願っている。我々が提案したソリューションは、Lightning network や Raiden network のような双方向ペイメントチャネルや HTLCs（Hashed TimeLock Contracts）によって構築される既存のオフチェーン決済ネットワークの拡張として機能することができる。この設計によって、ユーザーは自身の資産を完全に管理しながら迅速なクロスチェーン取引を行うことが可能となる。

> Order books are maintained by a permissionless network of liquidity providers, each of which spans multiple, individual payment networks. To begin trading, users open payment channels with a chosen liquidity provider, in the currencies they wish to trade. Orders are then submitted to the liquidity provider that a trader chooses, and matched with an online counterparty. Finally, the assets are exchanged atomically by executing a single, cross-chain payment, routed through the liquidity provider.

オーダーブックは、参加自由な流動性プロバイダのネットワークによって維持管理され、それぞれが複数の独立した決済ネットワークにまたがっている。取引を始めるにあたり、ユーザーは選択した流動性プロバイダとの間に取引したい通貨のペイメントチャネルを開く。注文は選択された流動性プロバイダに提出され、オンラインな取引相手とマッチングされる。最終的に、流動性プロバイダによってルーティングされた単一のクロスチェーン決済を実行することによって、資産はアトミックに取引される。

> Finally, Catalyst attempts to make algorithmic trading accessible for developers, by providing a complete toolchain that makes developing and testing trading strategies easy. Our toolchain will be open sourced and accessible both locally, or through a web IDE pre-loaded with all the dependencies. Aligned with our mission to increase adoption of crypto-assets, we will, over time, enable investors to pick winning strategies and invest in them. This marketplace of trading strategies will not only provide non-developers an interesting investment vehicle, but also allow the best quants to run their own micro hedge fund.

最後に、Catalyst は取引戦略を簡単に開発・テストするための完全なツールチェーンを提供することで、全ての開発者がアルゴリズム取引を利用できるようにする。我々のツールチェーンはオープンソースとなり、ローカルからも、全ての依存関係がプレロードされた web IDE からもアクセス可能となる予定である。暗号資産を広く普及させるというミッションに従い、我々は、徐々にではあるが、投資家が勝ち目のある戦略を選択し、暗号資産に投資できるようにしていく。この取引戦略市場は、非開発者に興味深い投資手段を提供するだけでなく、超一流の金融アナリストが独自のマイクロヘッジファンドを運用できるようにもするだろう。

---

## IV. CONCLUSIONS

> As the market surrounding crypto-assets is expanding, so should the investment tools and underlying financial infrastructure. In this paper, we have demonstrated Catalyst – a platform that provides the tools and data necessary to quickly build your own crypto hedge-fund.

暗号資産を取り巻く市場が拡大するにつれて、投資ツールや基盤となる金融インフラも拡大するはずである。本稿では、独自の暗号ヘッジファンドを迅速に構築するために必要なツールとデータを提供するプラットフォームである Catalyst を紹介した。

> Catalyst is the first application to make use of the Enigma decentralized data marketplace protocol. It is our hope that the adoption of Catalyst by developers and quants, would create a demand for proper, standardized crypto-data. These data-sets are likely to become of significant use to traders, researchers, journalists and anyone else who wishes to analyze the blockchain ecosystem from a data-driven perspective. Moreover, the adoption of Catalyst also implies that the Enigma data marketplace would become a vibrant peer-to-peer data exchange, paving the way for it to become an indispensable network of information for the web, that can change the way people aggregate, share and monetize their data.

Catalyst は、Enigma 分散型データ市場プロトコルを使用する最初のアプリケーションである。我々の望みは、Catalyst が開発者や金融アナリストに採用されることにより、正式に標準化された暗号データの需要が生まれることである。これらのデータセットは、トレーダーや研究者、ジャーナリスト、そしてデータ駆動型でブロックチェーンのエコシステムを分析したい誰しもにとって、重要な意味を持つようになるだろう。さらに、Catalyst の採用は、Enigma データ市場が活気ある P2P データ取引所となることも意味する。web にとって不可欠な情報ネットワークとなるための下地をつくりつつ、人々がデータを集約・共有・マネタイズする方法が変えることができる。

> Finally, we have also proposed a framework for building a decentralized crypto exchange protocol. We feel that this technical contribution could help the community in moving towards a more secure and scalable solution.

最後に、分散型暗号取引プロトコルを構築するためのフレームワークも提案した。この技術的貢献は、コミュニティがよりセキュアでスケーラブルなソリューションに移行する助けとなるだろう。

---

以上です。ざっくりまとめると、

- Catalyst は、暗号資産の投資プラットフォームであり、Enigma プロトコルを実装した最初のアプリケーションである
- Catalyst を利用することで、独自の暗号ヘッジファンドを簡単に構築し、データ駆動型の投資戦略を実践することができる
- Catalyst を用い、Enigma プロトコルに基づいたデータ市場を活性化していくことで、人々がデータを集約・共有・マネタイズする方法を変えていく

こんな感じでしょうか。今日はこれでおやすみなさい。
