+++
title = 'Twitter bot に趣深い詩を詠ませてみた'
tags = ['api', 'c1pher0ne', 'perl']
date = '2013-10-02'
+++

{{< figure src="/img/entry/c1pher0ne-on.png" >}}

表題の通り、[先日のブログ]({{< ref "/blog/2013/09/30/c1pher0ne.md" >}}) でご紹介した [01](http://twitter.com/c1pher0ne) が趣深い詩を詠めるようになりました。  
秋にぴったりな機能ですね。

<!--more-->

## ホントに詠めるの？？

- 朝起きたときと夜寝る前に詠みます
- 以下のように、予想以上にいいセンスをしてます

{{< tweet 385169779142582272 >}}
{{< tweet 385026451600195584 >}}

## どうやってやってるの？？

- [はなもげら API](http://truelogic.biz/hanamogera/hana-about.html) と [かな漢字変換 API](http://developer.yahoo.co.jp/webapi/jlp/jim/v1/conversion.html) をよしなに組み合わせてやってます
- [こちら](https://github.com/m0t0k1ch1/c1pher0ne/blob/master/lib/Cipherone/Model/Poem.pm) を見ていただければだいたいわかると思います
- 全部漢字に変換すると趣がなくなったので、[Data::WeightedRoundRobin](http://search.cpan.org/~xaicron/Data-WeightedRoundRobin-0.06/lib/Data/WeightedRoundRobin.pm) を使って、かなを漢字に変換する比率を調整してます

## 今後の予定

- [前回]({{< ref "/blog/2013/09/30/c1pher0ne.md" >}}) 書いたものに加えて以下も考えてます
- 流行りのアプリと音楽を教えてくれる機能、特にアプリは鉄板過ぎるのを持ってくるし、あまり変わり映えがしないので、何らかの調整が必要
- [ついっぷるトレンド](http://tr.twipple.jp) が比較的アツいらしいので、うまいこと利用したい

やっぱり実際に動かしてみると想定してたのとなんか違うなーってのがでてきて PDCA が捗ります。
とりあえずつくってみること、大事ですね。

あと、今日、「[地獄でなぜ悪い](http://play-in-hell.com)」を見て何かさらなるインスピレーションもらってこようと思います。押忍。
