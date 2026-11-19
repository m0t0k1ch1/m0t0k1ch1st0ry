---
title: '最近の Hugo + wercker + GitHub pages で嵌ったこと'
tags: ['hugo', 'wercker']
date: '2017-01-14T12:29:55+09:00'
---

[このブログを Hugo + wercker + GitHub pages で管理するようにした]({{< ref "/blog/2015/05/16/hugo.md" >}}) のが結構昔だったので、box が未だにクラシックな `wercker/default` だった。最新版に合わせとこっかなと軽い気持ちで application を登録しなおしたら盛大に嵌ったのでメモ。

<!--memo-->

## 嵌ったこと

大きくは 2 つ。

- git push をトリガーにした pipeline で branch の指定ができない？
  - 前までは ignore branch 的な設定があった気がする
  - トリガーより後ろの pipeline については branch の指定は可能
- theme として git submodule を使っている場合はきちんと初期化してあげないといけない？
  - 前までは `wercker.yml` に初期化処理書いてなくても build できてたんだけどな。。。逆になんでだったんだろう

後者は `git submodule update --init --recursive` とかを `wercker.yml` に書いてあげればいいとして、前者はどうするのがベターなんだろうという感じ。前までは master branch に push したら gh-pages branch に `public` 以下だけを push（deploy）していて、ignore branch 的な設定で gh-pages branch への push には反応しないようにしていた。

ignore branch 的な設定ができないと、gh-pages branch への push（deploy）にも反応して、無駄に build が実行されてしまう。で、`public` 以下には `wercker.yml` が存在しないので、その build がこけてしまうという。つらい。

## 解決策

無理やり感は否めないが、

- build の手前に適当な pipeline を入れる
- build 時に `public` 以下に `wercker.yml` をコピーする
- トリガーより後ろの pipeline は master branch のときのみ実行する

ようにして、gh-pages branch に反応した場合は build の手前で workflow が終了するようにした。wercker 上では以下のような感じ。

{{< figure src="/img/entry/wercker.png" >}}

そもそも workflow が走ってること自体がもやっとするのは変わらないけど、こけなくはなった。

最終的な `wercker.yml` は以下。もっとスマートな解決策をご存知の方いらっしゃいましたら、ご教示ください 🙏

```yaml
box: golang
pre:
  steps:
    - script:
        name: output message
        code: |
          echo 'start to build & deploy'
build:
  steps:
    - script:
        name: initialize git submodules
        code: |
          git submodule update --init --recursive
    - arjen/hugo-build:
        version: "0.18"
        flags: -v --buildDrafts=true
    - script:
        name: copy wercker.yml
        code: |
          cp wercker.yml public
deploy:
  steps:
    - lukevivier/gh-pages:
        token: $GIT_TOKEN
        domain: m0t0k1ch1st0ry.com
        basedir: public
```
