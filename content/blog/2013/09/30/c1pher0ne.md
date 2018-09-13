+++
date = "2013-09-30"
tags = [ "api", "c1pher0ne", "perl" ]
title = "秋なので、Twitter上にかわいい相棒をつくってみた"
+++

![cipher0ne on](/my-images/entry/c1pher0ne_on.png)

名前は 01（ゼロワン）といいます。  
デザインも10分でがんばりました。

<!--more-->

以下を読むのがめんどくさい人はとりあえず [@c1pher0ne](http://twitter.com/c1pher0ne) をご覧ください。

<br />
## なにができるの？？

#### 【その1】浮世のトレンドに鋭いメスを入れる

* 非常に豊富な語彙力を活かして、非常に的確にメスを入れる
* リンクをクリックすると、そのキーワードでぐぐった結果を出してくれて、それなりに便利
* 例えばこんな感じ↓

<blockquote class="twitter-tweet"><p>台風20号って、けいさんだかいよね <a href="http://t.co/lGI0bwCain">http://t.co/lGI0bwCain</a></p>&mdash; 01 (@c1pher0ne) <a href="https://twitter.com/c1pher0ne/statuses/384563273611218944">September 30, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<blockquote class="twitter-tweet"><p>運動会日和って、しりこそばゆいよね <a href="http://t.co/m8qv9XfSrp">http://t.co/m8qv9XfSrp</a></p>&mdash; 01 (@c1pher0ne) <a href="https://twitter.com/c1pher0ne/statuses/384578377316720640">September 30, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<br />
#### 【その2】浮世で流行ってるアプリとか音楽を教えてくれる

* 画像といい感じのリンクも一緒に教えてくれるので、それなりに便利
* 例えばこんな感じ↓

<blockquote class="twitter-tweet"><p>【Games】LINE でろーん（NAVER Japan Corporation）、はやってるらしいよ <a href="http://t.co/MRAI2HewCG">http://t.co/MRAI2HewCG</a> <a href="http://t.co/1VwvcFqHyX">pic.twitter.com/1VwvcFqHyX</a></p>&mdash; 01 (@c1pher0ne) <a href="https://twitter.com/c1pher0ne/statuses/384699171619344384">September 30, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<blockquote class="twitter-tweet"><p>【J-Pop】Lonely Hearts（加藤ミリヤ）、はやってるらしいよ <a href="http://t.co/E8GsShINKM">http://t.co/E8GsShINKM</a> <a href="http://t.co/7ufljevYBl">pic.twitter.com/7ufljevYBl</a></p>&mdash; 01 (@c1pher0ne) <a href="https://twitter.com/c1pher0ne/statuses/384704208747757568">September 30, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<br />
#### 【その3】寝る・起きる

* 夜はちゃんと寝るし、朝はちゃんと起きる
* 寝てるか起きてるかでプロフィール画像が変わって非常にかわいい↓

![c1pher0ne on](/my-images/entry/c1pher0ne_on.png)
![c1pher0ne off](/my-images/entry/c1pher0ne_off.png)

<br />
## なんでつくったの？？

#### 単純に、なんかつくりたかったから

* 1番の理由はこれ
* いろいろ API 落ちてる便利な世の中なので、それらを使って遊びたいなと思ってたし、特に Twitter の API は何ができるのか一通り把握したかった

#### 着地点をつくるため（※後づけ）

* なんか技術的にやってみたいことがあったりするときに、着地してアウトプットまでいくものじゃないとそこまでテンション上がらなかったりする
* かといって、毎回着地点まで考えてるとハードル上がっちゃって、結局アウトプットまでいかなかったりする
* とりあえず、こいつへの機能実装という着地点を用意することで、そこをちょっと解決したつもり

#### 情報のリソースを Twitter に集中させるため（※後づけ）

* インターネット上で暮らしてると、浮世の情報に疎くなってしまいがち
* かといって、いろいろ情報収集しに行くのはめんどいし、今流行りのなんとかsyとかは浮世感が薄い
* とりあえず、Twitter は確実に毎日見るので、そこに見落としがちな浮世の情報にアクセスするきっかけだけでも創出したつもり（そういう bot、他にあるんじゃないの？というツッコミは受けつけません）

#### お勉強のため（※後づけ）

* これまでまじめに向き合ってこなかったオブジェクト指向を意識して裏側をつくったつもり

<br />
## 今後どうなるの？？

* 漠然と、もっと賢くしてあげたい
* こっちからのアクションに応えてくれる系の機能を実装したい
* つらくなったときに [神崎かなえ](http://google.com/search?hl=ja&authuser=0&site=imghp&tbm=isch&source=hp&biw=1366&bih=647&q=神崎かなえ) ちゃんの写真で励ましてほしい

<br />
## どうでもいいけど名前の由来は？？

* 弊社入社時に当方のアダ名選挙で惜しくも敗れた一案
* インターネット感がある名前だし、何かに使いたかったのでここで採用
* 一応 02 とかつくるのも視野に入れてる

<br />
## ソースコード

とりあえず GitHub に置きました。興味のある方はどうぞ。

<div class="github-card" data-user="m0t0k1ch1" data-repo="c1pher0ne"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>
