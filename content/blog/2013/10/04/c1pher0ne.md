+++
date = "2013-10-04"
tags = [ "api", "c1pher0ne", "perl" ]
title = "Twitter bot にリマインダー機能を実装してみた"
+++

![c1pher0ne on](/img/entry/c1pher0ne_on.png)

表題の通り、[先日のブログ](http://m0t0k1ch1st0ry.com/blog/2013/09/30/c1pher0ne) でご紹介した [01](http://twitter.com/c1pher0ne) にリマインダー機能を実装しました。  
これ、誰でも使えるので、みなさんもお構いなくお使いください（とか言うと一瞬でバグを見つけられそうですが）。

<!--more-->

<br />
## どうやって使うの？？

* `#remind` というハッシュタグと `****/**/** **:**` というフォーマットでリマインドしてほしい日時を末尾に添えて、[@c1pher0ne](http://twitter.com/c1pher0ne) にメンションをとばします
* 成功していると、リプライが返ってきます
* 具体的には以下のような感じです

<blockquote class="twitter-tweet"><p><a href="https://twitter.com/m0t0k1ch1">@m0t0k1ch1</a> 御意！2013/10/2 21:00になったらリマインドするね</p>&mdash; 01 (@c1pher0ne) <a href="https://twitter.com/c1pher0ne/statuses/385343920579674112">October 2, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<br />
* すると、指定した日時にメンションを飛ばしてくれます
* この日、僕はこのリマインドのおかげで無事に「[地獄でなぜ悪い](http://play-in-hell.com)」を観ることができました

<blockquote class="twitter-tweet"><p><a href="https://twitter.com/c1pher0ne">@c1pher0ne</a> ok</p>&mdash; 色即是空 (@m0t0k1ch1) <a href="https://twitter.com/m0t0k1ch1/statuses/385375627399806977">October 2, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<br />
* ちなみに、寝てるときに頼んでも嫌々ながら引き受けてくれます

<blockquote class="twitter-tweet"><p><a href="https://twitter.com/m0t0k1ch1">@m0t0k1ch1</a> ...もー！起こさないでよ！2013/10/4 10:00 ね！はいはい！おやすみ！</p>&mdash; 01 (@c1pher0ne) <a href="https://twitter.com/c1pher0ne/statuses/385819054293794816">October 3, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<blockquote class="twitter-tweet"><p><a href="https://twitter.com/c1pher0ne">@c1pher0ne</a> お…押忍…</p>&mdash; 色即是空 (@m0t0k1ch1) <a href="https://twitter.com/m0t0k1ch1/statuses/385932525358428160">October 4, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<br />
* 過去へのリマインドを登録しようとすると現実が見れます

<blockquote class="twitter-tweet"><p><a href="https://twitter.com/m0t0k1ch1">@m0t0k1ch1</a> 過去には戻れないよ！現実を見て！</p>&mdash; 01 (@c1pher0ne) <a href="https://twitter.com/c1pher0ne/statuses/386051581470208000">October 4, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<br />
## どうやってやってるの？？

* 瀕死状態だったさくらの VPS を復活させ、cron で「メンションを解析して保存するバッチ」と「保存したデータをツイートするバッチ」を走らせています
* ※追記
  * 現在は [AnyEvent::Twitter::Stream](http://search.cpan.org/~miyagawa/AnyEvent-Twitter-Stream-0.26/lib/AnyEvent/Twitter/Stream.pm) を使っています

<br />
## 珍しく告知

ちなみに、上に貼ったツイートの中で出てきた「サムライのやつ」っていうのは [これ](http://kiban.doorkeeper.jp/events/5291) です。  
10/12（土）に、大先輩の [@fujiwara](https://twitter.com/fujiwara) さんと共に発表させていただきます。

恥かかないようにがんばりますので、興味のある方は是非ご参加ください。
