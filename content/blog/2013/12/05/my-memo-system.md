+++
title = 'Evetnote から脱却して Kobito + Nocs + Dropbox でメモ管理'
tags = ['others']
date = '2013-12-05'
+++

日常生活・仕事ともに、メモツールとしてはずっと [Evernote](http://evernote.com) を使っていたのですが、近頃の自分のニーズとマッチしなくなってきたのでさようなら〜して新しいメモ管理システムを整えました。大したことではありませんが、それなりに納得しているのでさらりとご紹介させていただきます。同じニーズをお持ちの方の参考になれば幸いです。

<!--more-->

## こうしたいねん！！

理想はこんな感じ。

- リストとか見出しとかで構造化してメモ取りたいので、Markdown 形式で書きたい
- コードはいい感じに syntax highlight してほしい
- macbook（複数台）と iPhone でデータを共有したい
- タグとかカテゴリでメモを分類したい

Evernote だと上の 2 つが厳しいですね。。

## 新メモ管理システムについて

### 使ったもの

- [Kobito](http://kobito.qiita.com)
- [Nocs](https://itunes.apple.com/jp/app/nocs-text-editor-dropbox-markdown/id396073482)
- [Dropbox](https://dropbox.com)

### 概要

簡単に言うと、

- Kobito の DB を Dropbox にぶちこんで macbook 複数台で共有
- iPhone で閲覧したいメモは Kobito から Dropbox に書き出して共有（もちろん Kobito の DB とファイル連携する）
- iPhone からは Nocs で Dropbox 内のファイルを閲覧・編集

という感じです。

## Kobito の DB を Dropbox にぶちこむ

シンボリックリンク貼るだけでした。簡単。具体的には以下の手順です。

- Kobito を一度も起動したことがなければ、とりあえず起動する（`~/Library/Kobito/Kobito.db` がつくられる）
- 一旦 Kobito を終了する
- Kobito.db を Dropbox に移してシンボリックリンクを貼る

```sh
$ mkdir ~/Dropbox/Kobito
$ mv ~/Library/Kobito/Kobito.db ~/Dropbox/Kobito
$ ln -s ~/Dropbox/Kobito/Kobito.db ~/Library/Kobito/Kobito.db
```

別に `~/Dropbox/Kobito` はつくる必要ないので、`~/Dropbox` に直接 `Kobito.db` をぶちこんで大丈夫です。自分は `~/Dropbox/Kobito` 内に Kobito から書き出したファイルを整理して管理したかったのでこうしました。

### ※ 注意点 1

Kobito はアプリケーション終了時にデータを DB に保存するようなので、データを共有したいときは Kobito を一旦終了しないといけないようです。これだけは我慢です。

### ※ 注意点 2

Nocs はコードの syntax highlight とかはしてくれません。

個人的に、iPhone では

- Markdown 形式のファイルがそれっぽく見れて、メモの内容が確認できればいい
- 思いついたことをとっさにメモれたらいい

といった感覚なので、そこまで高みは求めなくてもよいかなと思っております。

## まとめ

- とりあえずニーズは全部満たせたのでいい感じ
- もっとイケてる策をご存知の方、是非ご教示願います。
