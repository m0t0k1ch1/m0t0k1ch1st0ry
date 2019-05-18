+++
date = "2017-12-11T03:47:23+09:00"
tags = [ "ico", "raiden", "ethereum", "solidity", "blockchain" ]
title = "大人の事情には目を瞑って、Raiden の ICO について考える"
+++

この記事は [Ethereum Advent Calendar 2017](https://qiita.com/advent-calendar/2017/ethereum) の 11 日目の記事です。表題の通り、この記事では、今年行われた [Raiden](https://raiden.network) の ICO と向き合います。敢えて Raiden そのものとも向き合いません。ご承知ください。

<!--more-->

![raiden_1](/img/entry/raiden_1.png)

※この記事は ICO について取り扱いますが、ICO に関する詳しい説明は省きます。ICO に関して知識の浅い方は、現状、各メディアから発行されている記事の海に飛び込むよりも、まずは [増島先生の考察](https://www.scribd.com/document/362902074/Initial-Coin-Offering-from-Japanese-legal-and-practical-perspectives) にしっかりと目を通すのがよいかと思います。

## 0. 目次

- 1. 前置き
  - 1-1. なぜこの記事を書こうと思ったか
  - 1-2. Raiden の ICO に着目した理由
  - 1-3. 大人の事情に目を瞑った理由
- 2. そもそもダッチオークションとは？
- 3. Raiden のダッチオークションの仕組み
  - 3-1. Raiden のダッチオークションそのものについて
  - 3-2. Gnosis のダッチオークションとの違いについて
- 4. Raiden のダッチオークションについて考える
  - 4-1. 民主的なトークンの価値決定プロセス
  - 4-2. 難解な仕組み
  - 4-3. 重要な要素の不確定性
- 5. まとめ
- 6. 最後に
- 7. 参考

## 1. 前置き

### 1-1. なぜこの記事を書こうと思ったか

この質問に答える前に、まず、ICO に対する自分のスタンスをある程度明確にしておこうかなと思います（とは言え、これが本題ではないので、ほどほどに）。

まず、大前提として、自分は ICO を肯定的に捉えています。しかし、今の ICO は様々な課題を抱えているというのも事実です。

ref. [10億集めたICOが何もプロダクトをローンチできない理由](https://btcnews.jp/2ojkkse512376)

上記の記事で言及されているような ICO の暗黒面は、多くのメディアで取り上げられている課題であるため、この記事のテーマからは外します。その代わり、この記事で着目したいのは、

__クラウドセールのスキーム__

です。個人的には、これがまだ未熟であることも問題なのではないかと考えています。

もう少し具体的に言うと、

- プロジェクト主導でトークンを「売るだけ」のクラウドセールが多いのではないか
- クラウドセールのスキーム自体ももっと進化すべきなのではないか

と考えており、もう少し具体的に言うと、

- もっとクラウドセールの民主性を高めるべきではないか
  - コミュニティの価値向上をプロジェクトの価値向上の主軸に据え、支援者を良い意味で共犯者として巻き込んだコミュニティの形成を促進するようなクラウドセールができないか
- せっかくスマートコントラクトがあるのだから、コードによるガバナンスをもっとうまく活用したクラウドセールを行うべきではないか
  - 精神論ではなく、このガバナンスによって ICO の暗黒面を正すべきではないか

という感じです。ちょっと今回の記事が扱う範囲を超えた内容も含まれてはいますし、「いきなりそんなこと言われても。。」感も否めませんが、要するに、

__「Blockchain というトラストレスで分散化された世界にふさわしい、もっとイケてる ICO ってできないんすかね？」__

ということを悶々と考えている、くらいのイメージで捉えていただければと思います。

そして、この「もっとイケてる ICO」について想いを馳せるため、

__今年行われた ICO の中で、そのクラウドセール形式が最も印象的だった Raiden の ICO__

に着目し、その仕組みや性質について考えてみたいと思った、というのがこの記事を書こうと思った動機となります。

もちろん、Raiden のクラウドセールが自分の理想形というわけではないのですが、テーマとして扱う価値は十分にあると考えています。

### 1-2. Raiden の ICO に着目した理由

自分にとって、なぜ Raiden のクラウドセールが最も印象的なスキームだったのかと言うと、

__民主的なトークン価値（時価総額）決定プロセスとして、初めてダッチオークション形式のクラウドセールが機能したと感じたから__

です。

ダッチオークション形式のクラウドセールについては後ほど詳しく説明し、それに対する自分の考えも述べます。これが、この記事の本題となります。

なお、Raiden よりも前にダッチオークション形式のクラウドセールを実施したプロジェクトとしては [Gnosis](https://gnosis.pm) が挙げられますが、こちらは開始 10 分で終了、売れ残った大量のトークンが Gnosis チームに引き渡されたということもあり、結果的にはうまく機能しなかったという認識です（自分が知らないだけで他にも事例があるかもしれませんので、知っている方は是非教えてください）。

### 1-3. 大人の事情に目を瞑った理由

本題に入る前にもう 1 つ。表題に入れた「大人の事情には目を瞑って」というフレーズの意図についても言及しておこうと思います。

端的に述べると、

__ICO の意義や必然性ではなく、クラウドセールのスキームに着目したい__

ということです。

ご存知の方も多いと思いますが、Raiden の ICO は [コミュニティ内で議論を呼び](https://www.reddit.com/r/ethereum/comments/71r6bw/to_the_raiden_devs_why_do_you_need_an_ico) 、「[Vitalik](https://twitter.com/VitalikButerin) が反対した」という内容の記事もいくつか公開されたため、どうしてもそのイメージが強くなってしまいました。表面的な情報を見て、Raiden の ICO に対して、なんとなく「悪」というイメージを抱いてしまった方も多かったのではないでしょうか。また、このイメージが先行してしまったことで、ICO のスキームまで踏み込んで詳細に情報収集した人も少なかったのではないでしょうか。

実際、Vitalik は「反対した」というよりも、[「自身の立場を明確にした」](https://twitter.com/VitalikButerin/status/911217371158683649) という印象です。以下のツイートで、do not blame them とも言っています。

{{< tweet 911300771819352064 >}}

ref. [Vitalik Buterin Responds To Raiden ICO](https://www.ethnews.com/vitalik-buterin-responds-to-raiden-ico)

確かに、Raiden が ICO を実施する意義や必然性に対して疑問の声が上がるのは然りだと思いますし、議論をすることは重要だと思います。が、前述の通り、この記事では敢えてその点には触れず、自分にとって非常に印象的だったクラウドセールのスキームにフォーカスし、技術面にも踏み込みながら向き合っていきたいと思います。

前置きが長くなりましたが、以下、本題に入っていきます。

## 2. そもそもダッチオークションとは？

[Wikipedia](https://ja.wikipedia.org/wiki/%E7%AB%B6%E5%A3%B2) によると、

> 通常のオークションとは逆に、価格が順番に下がっていく。売り手が設定する最高価格から順番に価格を下げていき、買い手は適当なところで入札し、その時点の価格で落札が行われる。取引のスピードが高速化できるので、様々な市場で採用されている。また、バナナの叩き売りもこの一種である。

とのことです。冒頭に記載されている通り、

__通常のオークションとは逆に、価格が順番に下がっていく__

という点が最たる特徴と言えるでしょう。

「早く入札しないと、先に誰かに落札されてしまう」といった心理状態を利用して入札を誘発し、それによって取引の高速化を図る、というのがキモのようです。

Gnosis や Raiden は、この仕組みをトークンのクラウドセールに応用したというわけです。しかし、トークンのクラウドセールの場合、誰か 1 人が入札したら終了というわけではないので、少し異なる性質も持ち合わせていると言えるでしょう。以下でその仕組みを詳しく見ていきます。

## 3. Raiden のダッチオークションの仕組み

何はともあれ、Raiden のダッチオークションがどのような仕組みなのかを正しく把握するところから始めようと思います。

実際に使われたコントラクトは [こちら](https://etherscan.io/address/0xb5e5585d0057501c91c48094029a6f4fb10b5a01#code) です。必要に合わせてご参照ください。

まずは、`DutchAuction` コントラクトの冒頭に記載されているコメントを見てみます。

> @title Dutch auction contract - distribution of a fixed number of tokens using an auction.
> The contract code is inspired by the Gnosis auction contract. Main difference is that the auction ends if a fixed number of tokens was sold.

ざっくり訳すと、

__ダッチオークションコントラクト：一定数のトークンをオークションを利用して分配する。このコントラクトは Gnosis のオークションコントラクトにインスパイアされたものであるが、主たる違いは、一定数のトークンが売れた時点でオークションが終了することである__

という感じでしょうか。

以下、このコメントを前半部分と後半部分に分け、それぞれ詳しく説明していきたいと思います。

### 3-1. Raiden のダッチオークションそのものについて

「一定数のトークンをオークションを利用して分配する」とさらっと書かれていましたが、これはどういうことでしょうか？

[Raiden のクラウドセール用サイト](https://token.raiden.network) に表示されている、ダッチオークションの結果を示す図のスクショを撮ってきました。これをベースに説明していこうと思います。

![raiden_2](/img/entry/raiden_2.png)

図の各軸は以下を表しています。

- 横軸：時間
- 縦軸：ETH の量（対数軸となっていることに注意）

ダッチオークションは、図の左端の時点（10/19）で開始され、白い縦線が入っている時点（11/1）で終了しました。そして、この 10 日程度の開催期間でなにやら 2 つの値が変動していることが分かります。

まず、図の上方から垂れ下がってきている曲線についてです。これは、発行される RDN トークンの時価総額の変化を表しています。これが徐々に低下しているのは、Raiden のダッチオークションに以下のようなルールがあるからです（もちろん、これはスマートコントラクトとして定義されています）。

- RDN トークンの価格（ETH/RDN）は時間とともに低下する
  - ダッチオークション開始時は、ありえないくらい高価格である
- RDN トークンの販売量は一定である

なお、Raiden のクラウドセールでは、総トークン発行量である 1 億 RDN の半分である 5,000 万 RDN が販売されました（残り半分は Raiden チームの取り分）。そのため、上図における RDN トークンの時価総額は、

__総トークン販売量（5,000 万で一定） × RDN トークンの価格（時間とともに低下）__

で計算されています。上記ルールを加味すると、RDN トークンの時価総額が時間とともに下がっていくことは簡単に理解できるかと思います。

次に、図の下方からせり上がってきている面グラフについてです。これは、Raiden チームがクラウドセールによって調達した ETH の総量となります。クラウドセールの進行とともに入札者（入札額）は増えていくため、この値が増加していくというのも簡単に理解できるかと思います。

ただ、ここで 1 点注意しなければならないのは、入札者は「10 RDN 買います！」と言って参加しているのではなく、「10 ETH 出します！」と言って参加しているということです。すなわち、

__宣言しているのは入札額であり、実際にどれだけの RDN トークンが手に入るかはダッチオークションが終了するまでわからない__

ということです。

では、ダッチオークションはいつ終了するのでしょうか？上図がそのまま解答となっているのですが、文章にすると、

__RDN トークンの時価総額と、Raiden チームが調達した ETH の総量が等しくなったとき__

となります。この条件を満たした時点において、調達した ETH の総量を総トークン販売量（5,000 万で一定）で割った値が最終的な RDN トークンの価格となり、入札額と最終的な RDN トークンの価格に基づいてトークン分配が行われます。

なお、手に入る RDN トークンの量は、入札額を最終的な RDN トークンの価格で割った値となります。ここで重要なのは、

__結果的に、いつ入札しても RDN トークンの価格は変わらない__

ということです。オークション開始直後に入札しても、オークション終了ギリギリで入札しても同じです。よって、オークションが終了するタイミングを予想しつつ、自分が「これなら買ってもいい」と思えるラインまで価格が下がってきたら入札する（追加で入札することもできます）、というのが入札者の基本的な動きとなりそうです。入札によって、手に入る RDN トークンの最小量が保証されます。

長々とした説明になってしまいましたが、、重要なポイントは、「RDN トークンの時価総額の下降曲線はあらかじめ Raiden チームによって決められたものだが、調達される ETH の総量がどのようなペースで増加していくかは入札者の動きに委ねられている」ということです。端的に言うと、

__入札者側がダッチオークションの終了タイミングを決定する権利を持っている__

ということです。

当然、入札者（入札額）の増加ペースが芳しくなければ、Raiden チームが調達できる ETH の総量はどんどん下がっていくわけです。が、Raiden チームはこれを許容してこのスキームをチョイスしたということでしょう。

こういったダッチオークションの性質に関する考察は後ほど行うとして、ここでは [実際のコントラクト](https://etherscan.io/address/0xb5e5585d0057501c91c48094029a6f4fb10b5a01#code) まで踏み込んでその仕組みを確認してみようと思います。

まずは入札用の `bid` function です。

``` solidity
/// @notice Send `msg.value` WEI to the auction from the `msg.sender` account.
/// @dev Allows to send a bid to the auction.
function bid()
    public
    payable
    atStage(Stages.AuctionStarted)
{
    require(msg.value > 0);
    require(bids[msg.sender] + msg.value <= bid_threshold || whitelist[msg.sender]);
    assert(bids[msg.sender] + msg.value >= msg.value);

    // Missing funds without the current bid value
    uint missing_funds = missingFundsToEndAuction();

    // We require bid values to be less than the funds missing to end the auction
    // at the current price.
    require(msg.value <= missing_funds);

    bids[msg.sender] += msg.value;
    received_wei += msg.value;

    // Send bid amount to wallet
    wallet_address.transfer(msg.value);

    BidSubmission(msg.sender, msg.value, missing_funds);

    assert(received_wei >= msg.value);
}
```

冒頭にある以下の処理は基本的なバリデーションです。

``` solidity
require(msg.value > 0);
require(bids[msg.sender] + msg.value <= bid_threshold || whitelist[msg.sender]);
assert(bids[msg.sender] + msg.value >= msg.value);
```

`bid_threshold`（2.5 ETH）以上の入札については事前にホワイトリストへの登録申請を行う必要があったので、それに関する処理も含まれています。

基本的なバリデーションに続いて、`missingFundsToEndAuction` function が実行され、`missing_funds` が計算されています。

``` solidity
// Missing funds without the current bid value
uint missing_funds = missingFundsToEndAuction();

// We require bid values to be less than the funds missing to end the auction
// at the current price.
require(msg.value <= missing_funds);
```

`missing_funds` は、RDN トークンの時価総額と調達した ETH の総量の差分であり、「あとこれだけ ETH を調達したらオークション終了だよ」ということを意味する値となります。`missing_funds` の計算の後、 入札額がこの値を超えていないかについても確認が行われています。実質、この確認処理がオークションの終了を司ることになります。

入札が正常であることが確認できたら、入札者のアドレスと入札額を紐づけて `bids` に保存（入札は複数回可能であるため、加算している）し、受け取った ETH の総量である `received_wei` にも入札額分を加算しています。

``` solidity
bids[msg.sender] += msg.value;
received_wei += msg.value;
```

この後、入札額分の ETH を Raiden チームのウォレットに送り、イベント（レシートのようなもの）を発行、念押しのバリデーションが行われています。

以上が入札処理のざっくりとした解説となります。意外とシンプルでした。が、肝心のトークンの価格計算処理がまだ登場していません。これはどうなっているのでしょうか。実際のコードを見てみます。

``` solidity
/*
 *  Private functions
 */

/// @dev Calculates the token price (WEI / RDN) at the current timestamp
/// during the auction; elapsed time = 0 before auction starts.
/// Based on the provided parameters, the price does not change in the first
/// `price_constant^(1/price_exponent)` seconds due to rounding.
/// Rounding in `decay_rate` also produces values that increase instead of decrease
/// in the beginning; these spikes decrease over time and are noticeable
/// only in first hours. This should be calculated before usage.
/// @return Returns the token price - Wei per RDN.
function calcTokenPrice() constant private returns (uint) {
    uint elapsed;
    if (stage == Stages.AuctionStarted) {
        elapsed = now - start_time;
    }

    uint decay_rate = elapsed ** price_exponent / price_constant;
    return price_start * (1 + elapsed) / (1 + elapsed + decay_rate);
}
```

ちょっと複雑そうです。しかし、`price_start` と `price_exponent` と `price_constant` は定数であり、`decay_rate` は `elapsed` から計算されています。すなわち、この式の挙動を想像するには、`elapsed` に着目すればよさそうです。

`elapsed` は `now - start_time` で定義されていますので、ダッチオークション開始時のタイムスタンプと現在のブロックのタイムスタンプの差分を意味します。ということは、ダッチオークションが進行するにつれて、`elapsed` は大きくなっていきます。そうすると、`(1 + elapsed) / (1 + elapsed + decay_rate)` は小さくなっていきます。すなわち、この計算式は、

__オークション開始時からの経過時間を利用して、トークンの価格を徐々に引き下げている__

ということになります。前述した通り、これはダッチオークションの特徴的なルールです。

ダッチオークションの終了後には、以下の `finalizeAuction` function が実行されます。

``` solidity
/// @notice Finalize the auction - sets the final RDN token price and changes the auction
/// stage after no bids are allowed anymore.
/// @dev Finalize auction and set the final RDN token price.
function finalizeAuction() public atStage(Stages.AuctionStarted)
{
    // Missing funds should be 0 at this point
    uint missing_funds = missingFundsToEndAuction();
    require(missing_funds == 0);

    // Calculate the final price = WEI / RDN = WEI / (Rei / token_multiplier)
    // Reminder: num_tokens_auctioned is the number of Rei (RDN * token_multiplier) that are auctioned
    final_price = token_multiplier * received_wei / num_tokens_auctioned;

    end_time = now;
    stage = Stages.AuctionEnded;
    AuctionEnded(final_price);

    assert(final_price > 0);
}
```

上でも出てきた `missing_funds` が 0 であることを確認し、最終的なトークンの価格である `final_price` を決定しています。`final_price` は、時間の経過につれて値が変動する `calcTokenPrice` function を用いず、調達した ETH の総量を総トークン販売量で割ることで計算しています。

さらにこの後、別の function によって、入札額と `final_price` に基づいたトークンの分配処理が行われます。ここで、総トークン販売量である 5,0000 万 RDN が全て入札者に分配されます。これが、`DutchAuction` コントラクトの冒頭に記載されていた「一定数のトークンをオークションを利用して分配する」という行為に該当します。

分配処理を行う function についての説明は割愛しますが、詳しく知りたい方は、是非 [実際のコントラクト](https://etherscan.io/address/0xb5e5585d0057501c91c48094029a6f4fb10b5a01#code)を見てみてください。なお、コントラクトで実際にどのような値が使われたかも [こちら](https://etherscan.io/address/0xb5e5585d0057501c91c48094029a6f4fb10b5a01#readContract) から確認できます。例えば `final_price` は 2190632199853296 となっていますので、最終的な RDN トークンの価格は約 0.0022 ETH/RDN だったということがわかります。

### 3-2. Gnosis のダッチオークションとの違いについて

`DutchAuction` コントラクトの冒頭には、「一定数のトークンが売れた時点でオークションが終了すること」が、Gnosis のダッチオークションとの主な違いであると記載されていましたが、Gnosis のダッチオークションの終了条件はどうなっていたのでしょうか？実際にコントラクトを見て確認してみようと思います。

Gnosis のダッチオークションで実際に使われたコントラクトは [こちら](https://etherscan.io/address/0x1d0dcc8d8bcafa8e8502beaeef6cbd49d3affcdc#code) のようです。今回も入札用の `bid` function から見ていこうと思います。

``` solidity
/// @dev Allows to send a bid to the auction.
/// @param receiver Bid will be assigned to this address if set.
function bid(address receiver)
    public
    payable
    isValidPayload
    timedTransitions
    atStage(Stages.AuctionStarted)
    returns (uint amount)
{
    // If a bid is done on behalf of a user via ShapeShift, the receiver address is set.
    if (receiver == 0)
        receiver = msg.sender;
    amount = msg.value;
    // Prevent that more than 90% of tokens are sold. Only relevant if cap not reached.
    uint maxWei = (MAX_TOKENS_SOLD / 10**18) * calcTokenPrice() - totalReceived;
    uint maxWeiBasedOnTotalReceived = ceiling - totalReceived;
    if (maxWeiBasedOnTotalReceived < maxWei)
        maxWei = maxWeiBasedOnTotalReceived;
    // Only invest maximum possible amount.
    if (amount > maxWei) {
        amount = maxWei;
        // Send change back to receiver address. In case of a ShapeShift bid the user receives the change back directly.
        if (!receiver.send(msg.value - amount))
            // Sending failed
            throw;
    }
    // Forward funding to ether wallet
    if (amount == 0 || !wallet.send(amount))
        // No amount sent or sending failed
        throw;
    bids[receiver] += amount;
    totalReceived += amount;
    if (maxWei == amount)
        // When maxWei is equal to the big amount the auction is ended and finalizeAuction is triggered.
        finalizeAuction();
    BidSubmission(receiver, amount);
}
```

基本的な処理フローは Raiden のものと同じ、かつコメントが細かく入っているので、追加で説明することもあまりなさそうですが、一応、軽く説明します。

コードの前半部分には、総トークン発行量の 90 %（`MAX_TOKENS_SOLD` に対応）を超える量のトークンが配布されないようにするための処理が含まれています。これは、Gnosis チームの取り分を確保するためのものでしょう。これを加味して計算された `maxWei` が、Raiden のダッチオークションにおける `missing_funds` と対応します。コードの中盤では、`maxWei` を超えた分の返金処理なども定義されています。入札が正常であった場合に辿り着く後半部分では、入札者のアドレスと入札額を紐づけて `bids` に保存し、受け取った ETH の総量である `totalReceived` に入札額分を加算しています。

そして、最後に以下のようなダッチオークションの終了判定が行われています。

``` solidity
if (maxWei == amount)
    // When maxWei is equal to the big amount the auction is ended and finalizeAuction is triggered.
    finalizeAuction();
```

すなわち、一定量の ETH を調達した時点でダッチオークションが終了するということになります。これは、

__調達額のキャップが存在する__

ということを意味します。

また、上記 `bid` function の中身が実行される前には、`timedTransitions` modifier に定義された処理も実行されており、この中で `finalizeAuction` function が実行されるパターンもあります。そのため、調達額がキャップに達しなくとも、以下の条件に当てはまった場合にはオークションが終了することになります。

具体的には、以下のコードの 1 つ目の if 文です。

``` solidity
modifier timedTransitions() {
    if (stage == Stages.AuctionStarted && calcTokenPrice() <= calcStopPrice())
        finalizeAuction();
    if (stage == Stages.AuctionEnded && now > endTime + WAITING_PERIOD)
        stage = Stages.TradingStarted;
    _;
}
```

条件式の中で呼ばれている `calcTokenPrice` function と `calcStopPrice` function の定義は以下のようになっています。Raiden の `calcTokenPrice` function とは異なるロジックで計算されているようです。

``` solidity
/// @dev Calculates stop price.
/// @return Returns stop price.
function calcStopPrice()
    constant
    public
    returns (uint)
{
    return totalReceived * 10**18 / MAX_TOKENS_SOLD + 1;
}

/// @dev Calculates token price.
/// @return Returns token price.
function calcTokenPrice()
    constant
    public
    returns (uint)
{
    return priceFactor * 10**18 / (block.number - startBlock + 7500) + 1;
}
```

ダッチオークションであるがゆえに token price は時間とともにその値が小さくなっていく一方、その時点での調達額基準で計算される stop price は入札があるごとにその値が大きくなっていきます。すなわち、例えば時間軸を横軸にとってこの 2 つの値をプロットすることで描かれるグラフは、どこかのタイミングで交わることになります。そして、その交差点がオークションの終了を意味するわけです。

オークションの終了条件が満たされた場合に実行される `finalizeAuction` function の内容は以下のようになっています。

``` solidity
/*
 *  Private functions
 */
function finalizeAuction()
    private
{
    stage = Stages.AuctionEnded;
    if (totalReceived == ceiling)
        finalPrice = calcTokenPrice();
    else
        finalPrice = calcStopPrice();
    uint soldTokens = totalReceived * 10**18 / finalPrice;
    // Auction contract transfers all unsold tokens to Gnosis inventory multisig
    gnosisToken.transfer(wallet, MAX_TOKENS_SOLD - soldTokens);
    endTime = now;
}
```

`totalReceived` と `finalPrice` に基づいて入札者に対して分配するトークンの総量を計算し、売れ残った分を Gnosis チームのウォレットに移しています。

以上を踏まえて Raiden のダッチオークションとの違いをまとめると、以下のようになります。

- トークンの価格を決定するロジックが異なる
- ダッチオークションの終了条件が異なる
  - 調達額にキャップが存在し、キャップに到達した時点でダッチオークションは終了する
  - キャップに到達しなかった場合の終了条件も異なる（ref. `timedTransitions` modifier）
- トークンに売れ残り分が存在する
  - 売れ残り分は Gnosis チームに引き渡される

これらの中でも、

__大量のトークンが売れ残り、それが全て Gnosis チームに引き渡される可能性がある__

というのは大きな問題であると考えられます。

ちなみに、実際にはどの程度のトークンが売れ残ったのでしょうか？[実際のコントラクトにおける値](https://etherscan.io/address/0x1d0dcc8d8bcafa8e8502beaeef6cbd49d3affcdc#readContract) を基準に計算してみます。売れ残り分を計算するのに必要な値は以下のようになっていました。

- `MAX_TOKENS_SOLD`：9000000000000000000000000
- `totalReceived`（`ceiling`）：250000000000000000000000
- `finalPrice`：596975325019899178

`finalizeAuction` function に従って売れ残り分をざっくり計算すると、以下のようになります。

$$
9000000 \times 10^{18} - (\frac{250000 \times 10^{18} \times 10^{18}}{596975325019899178}) \approx 8580000 \times 10^{18}
$$

クラウドセールでの総トークン販売量は `MAX_TOKENS_SOLD` に対応する \\(9000000 \times 10^{18}\\) ですので、これに対する割合を計算すると、以下のようになります。

$$
\frac{8580000 \times 10^{18}}{9000000 \times 10^{18}} \approx 0.95
$$

なんと、総トークン販売量の約 95 % が売れ残り、それが全て Gnosis チームに引き渡されたということになります。Gnosis のダッチオークションは開始 10 分で調達額のキャップに達してしまい、トークンの価格低下がほぼ起こらなかったため、このような残念な結果となってしまったのです。

なお、前述した通り、Raiden のダッチオークションでは総トークン販売量が全て入札者に対して分配されるため、この売れ残り問題は発生しません。これは大きな違いと言えるでしょう。

## 4. Raiden のダッチオークションについて考える

ここまでで見てきた仕組みを踏まえつつ、一般的なクラウドセールと比較しながら、Raiden が実施したダッチオークション形式のクラウドセールの性質について考えてみようと思います。なお、ここで言う「一般的なクラウドセール」は、以下のようなものを指すこととします。

- 一定期間、プロジェクト側が決めた価格で ERC20 トークンを販売する形式
  - トークンの価格は、販売期間を通じて一定、もしくは一定期間ごとに上昇する

トークンの価格が一定期間ごとに上昇するパターンは、「1 週目は 1 トークンあたり 0.01 ETH、2 週目は 1 トークンあたり 0.02 ETH、3 週目は、、」といったイメージです。

それでは、以下、ダッチオークション形式のクラウドセールの性質について、自分が重要だと感じたポイントを 3 つに分けて説明しながら、自分の考えを述べていこうと思います。

### 4-1. 民主的なトークンの価値決定プロセス

前述した通り、ダッチオークションは入札者側が終了タイミングを決定する権利を持っており、終了した時点でトークンの価格が決定します。これは、一般的なクラウドセールと比較すると、

__民主的なトークン価値（時価総額）決定プロセスが実現している__

と言えるのではないでしょうか。また、これにより、

__不当なトークン価格になり難い__

と言えるかもしれません。ある程度著名なプロジェクトが発行するトークンであっても、取引所に上場したタイミングでの売り抜けによって価格が急降下するパターンは多々見受けられますが、民主的なプロセスを通じて決定した価格であれば、上場した途端に大きく価格変動する可能性は低そうに思えます。

実際、執筆時点における [CoinMarketCap](https://coinmarketcap.com/currencies/raiden-network-token) のチャートを見る限り、RDN トークンの価格変動は比較的落ち着いているように思えます。

![raiden_3](/img/entry/raiden_3.png)

民主的なプロセスやそれによって生まれるこのような性質は、今後の ICO にも活かされていくべきと考えています。

また、プロジェクトオーナーと支援者はともにトークン保有者であり、プロジェクトに貢献することで、その報酬をトークンの長期的な価値向上として享受することができます。この観点で双方のインセンティブは一致しているという前提に立てば、支援者側にもプロジェクトの成長に貢献する動機があると言えます。クラウドセールにうまく民主性を盛り込むことで、早期にこの動機に基づいた貢献を促すことができるかもしれません。

と、ここまで良いことばかりのように見えますが、そうでもありません。例えば、以下のような問題点が挙げられます。

- 民主的に決定されたとは言えど、それが適正価格なのかどうかはわからない
  - クラウドセール時点で完成したプロダクトが存在しない場合、トークンの利用価値を予想することは困難であり、投機目的での参加もかなりの割合を占めると考えられる
- ダッチオークションにサクラを紛れ込ませることができる

前者に関しては、何をもって適正とするかで大きく意見が分かれそうな気もしますが、、ここでは目を瞑ります。

なお、ここで述べた内容については、Gnosis のダッチオークションについて言及している以下のブログ記事の中で指摘されている内容と重複する部分もありますので、こちらも併せてご覧いただくのがよいかと思います。

ref. [GnosisのICOモデルは伝説を残すのか？](http://coinandpeace.hatenablog.com/entry/gnosis_ICO)

### 4-2. 難解な仕組み

うん。難しいですよね。この記事の説明に対して「へ？」ってなっている方も多いと思いますが、自分も最初に [Raiden チームの説明](https://medium.com/@raiden_network/the-raiden-token-auction-explained-1cc0c7946b26) を読んだとき、「へ？」ってなりました。

「え、そんな感想を考察と言われても。。。」と思われるかもしれませんので、もう少し踏み込んで説明していきます。

仕組みが難しい場合、以下のようなデメリットがあると考えられます。

- スマートコントラクトにバグが混入する可能性が高まる
- クラウドセールに参加するハードルが上がる
- 仕組みを理解しないままクラウドセールに参加する人が出てくる

これは結構痛いのではないでしょうか。

1 つ目は我々エンジニアが歯を食いしばってなんとか解決するとしても、2 つ目と 3 つ目は、我々が直接解決できる問題ではありません。

加えて、今後より広く ICO が普及していくとしたら、リテラシーの低い参加者の数もどんどん増えてくると考えられます。

__どれだけトークンの時価総額決定プロセスに高尚な民主性があったとしても、「うーん。よく分からないから参加しません」「参加はするけど、仕組みはよくわかってません」という人が大多数であれば、わざわざダッチオークション形式のクラウドセールを選択する価値が薄れてしまう__

と考えられます。また、このような状況で実施したとしても、ダッチオークションが期待通りに機能する可能性は低いでしょう。

個人的には、4-1. で述べたような民主性を保ちながらも、よりシンプルな仕組みで稼働するスキームが必要なのではないかと考えています。

### 4-3. 重要な要素の不確定性

重要な要素とは、最終的なトークンの価格のことです。これは参加者にとって「気になる」値でしょう。ダッチオークションが終了するまでこれが不確定であるということは、

__「気になる」状態がダッチオークションの実施期間中継続している可能性が高い__

と言えるのではないでしょうか。

これをポジティブに捉えると、参加者にとってクラウドセールが自分ごと化されやすいこの期間を活かして、

- 参加者を巻き込んだマーケティングができないか
- コミュニティ形成のきっかけを創出し、これをうまくプロジェクトの価値向上に繋げられないか

などを議論することもできそうです。また、

__この不確定性は、エンターテイメント性と言い換えることができる__

かもしれません。ちょっと不謹慎かもしれませんが、クラウドセールがエンターテイメント化することによって、今まで参加していなかった層が参加するきっかけが生まれる可能性もあるかなと考えています。

この辺りのことを加味すると、巨額の資金調達だけに留まらない ICO の在り方というのも見えてくるのではないでしょうか。

## 5. まとめ

- 自分の ICO に対するスタンスを踏まえつつ、Raiden の ICO（特にそのクラウドセールのスキーム）に着目した理由について説明しました
- 実際のスマートコントラクトの内容を踏まえつつ、Raiden が実施したダッチオークション形式のクラウドセールの仕組みについて説明し、Gnosis のそれと比較することで理解を深めました
- ダッチオークション形式のクラウドセールについて、自分が重要だと感じた 3 つの性質について説明するとともに、自分の考えを述べました

## 6. 最後に

詐欺的なプロジェクトであっても巨額の資金調達が行えてしまう、そんな異常な ICO バブルも各国における規制の影響などによって落ち着きを見せ始めている昨今、ICO を取り巻く状況は刻々と変化しています。しかし、ICO がイノベーションであるということに変わりはないと思います。

そんな中、金融庁の英断によって世界に先駆けて仮想通貨に関する法整備が成された日本こそ、ICO というイノベーションが健全に進化するための土壌が整った国と言えるのではないでしょうか。また、日本にいる我々がそれについて議論することは非常に有意義なことではないでしょうか。この記事が、そんな議論のきっかけになれば幸いです。

中国のリープフロッグをきちんと受け止めつつ、次は日本がリープフロッグする番です。イノベーションの芽を枯らせてしまわぬよう、自分も 1 プレイヤーとして、引き続き ICO と向き合っていこうと思います。

## 7. 参考

- [Initial Coin Offering - from Japanese legal and practical perspectives](https://www.scribd.com/document/362902074/Initial-Coin-Offering-from-Japanese-legal-and-practical-perspectives)
- [10億集めたICOが何もプロダクトをローンチできない理由](https://btcnews.jp/2ojkkse512376)
- [To the Raiden Devs: Why do you need an ICO, seriously?](https://www.reddit.com/r/ethereum/comments/71r6bw/to_the_raiden_devs_why_do_you_need_an_ico)
- [Vitalik Buterin Responds To Raiden ICO](https://www.ethnews.com/vitalik-buterin-responds-to-raiden-ico)
- [競売](https://ja.wikipedia.org/wiki/%E7%AB%B6%E5%A3%B2)
- [GnosisのICOモデルは伝説を残すのか？](http://coinandpeace.hatenablog.com/entry/gnosis_ICO)
- [The Raiden Network Token Auction Explained](https://medium.com/@raiden_network/the-raiden-token-auction-explained-1cc0c7946b26)
