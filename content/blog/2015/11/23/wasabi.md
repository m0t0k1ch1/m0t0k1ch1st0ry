+++
title = 'なんでもルーレットしてくれる Slack bot'
tags = ['go']
date = '2015-11-23T19:06:57+09:00'
+++

自分のプロジェクトのエンジニア間では、「誰がやるか決めようぜ！！！」みたいなことがときたま発生するので、以下のような感じのことができる Slack bot をつくった。

{{< figure src="/img/entry/wasabi.png" >}}

<!--more-->

名前はわさび寿司的な意味で wasabi。

{{< github "m0t0k1ch1" "wasabi" >}}

## 構成

今年の抱負の実現のために Go で API サーバーを立てる練習がしたかったので、それなりにきちんとつくった（若干の牛刀感は否めない）。特徴はざっくりと以下のような感じ。

- ベースは拙作の [potto](https://github.com/m0t0k1ch1/potto)
  - `/`：Slack の Outgoing WebHook をハンドリング
  - `/stats`：[golang-stats-api-handler](https://github.com/fukata/golang-stats-api-handler) を使って runtime の情報を吐く
- graceful な感じ
  - [manners](https://github.com/braintree/manners) による graceful shutdown
  - [go-server-starter](https://github.com/lestrrat/go-server-starter) による graceful restart
- ストレージは Redis
  - ルーレットの要素は set で管理
  - channel ごとにルーレットの要素を管理

## 実際にサーバーで動かすにあたって

### ミドルウェアのセットアップ

[itamae](https://github.com/itamae-kitchen/itamae) でシュッと実行できるようにした。

itamae はシンプルで好き。最近趣味でなんかするときはよく使わせてもらっている。

### daemonize

普段は daemontools を使っているけど、ちょっと前に ISUCON の勉強で supervisor を触っていたので、試しに使ってみるかってことで使ってみた。

が、

- yum でシュッと3.x系が入らない
- 2.x 系と 3.x 系の config の違い

あたりに翻弄されてだいぶ時間を喰われてしまった。。

3.x 系をシュッとセットアップする itamae の cookbook はできたのでまあよしとするかな。。。

## 感想

- 自分のつくったもの（[ksatriya](https://github.com/m0t0k1ch1/ksatriya) や [potto](https://github.com/m0t0k1ch1/potto)）を使って実際にコードを書くと「これはイケてないわあ。。。」みたいなとこがどんどん見えて、がんがん使う側視点でリファクタリングが進むのでよい
- 目標であった「今年の抱負の実現のために Go で API サーバーを立てる練習がしたかった」については、1つ自分の中でテンプレ的なものができたので達成

wasabi は練習台なので、年末焦燥感駆動でもうちょいがんばる🙏

{{< tweet 668673216928309248 >}}
