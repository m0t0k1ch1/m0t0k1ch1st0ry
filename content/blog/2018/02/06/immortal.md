+++
date = "2018-02-06T01:43:28+09:00"
tags = [ "golang" ]
title = "Golang 製の daemon 管理ツール immortal で daemon 管理"
+++

表題の通り、以前に見つけていてずっと触りたかった [immortal](https://github.com/immortal/immortal) という Golang 製の daemon 管理ツールをやっとこさ試すことができたのでメモ。名前がエモくて良いですね。

<!--more-->

<br />
## immortal について

[公式ドキュメント](https://immortal.run/post/how-it-works) には以下のように記載されています。

> immortal runs a command or script detached from the controlling terminal as a Unix daemon, it will supervise and restart the service if it has been terminated.

ざっくり訳すと、

__immortal は、コマンドやスクリプトを制御中のターミナルから Unix daemon としてデタッチする。また、サービスを監視し、それが終了された場合は再起動する。__

という感じでしょうか。

自分にとっての daemon 管理ツールの定番は [daemontools](http://www.emaillab.org/djb/daemontools/daemontools-howto.html) なのですが、これとだいたい同じような挙動をしてくれそうです。

それでは以下、実際に触りつつ基本的な挙動を把握していきます。基本的には [公式ドキュメント](https://immortal.run) を参考にしているので、詳細な説明はそちらをご覧ください。

<br />
## インストールする

今回は Ubuntu 16.04 にインストールしてみます。

``` sh
$ cd /tmp
$ wget https://github.com/immortal/immortal/releases/download/0.18.0/immortal_0.18.0_amd64.deb
...
$ sudo dpkg -i immortal_0.18.0_amd64.deb
...
```

バージョンを確認します。

``` sh
$ immortal -v
```

<pre>
0.18.0
</pre>

[immortal](https://immortal.run/post/immortal) だけでも daemonize は可能なのですが、今回は [immortaldir](https://immortal.run/post/immortaldir) を利用し、設定ファイルベースでの監視まで行ってみようと思います（daemontools の svscan のようなイメージです）。

ということで、設定ファイルを設置するディレクトリ（監視対象）を作成しておきます。

``` sh
$ sudo mkdir /usr/local/etc/immortal
```

また、systemd 経由で制御できるよう、以下の内容で `/etc/systemd/system/immortaldir.service` を作成します。

<pre>
[Unit]
Description=immortaldir
After=network.target

[Service]
ExecStart=/usr/bin/immortaldir /usr/local/etc/immortal 2>&1 | logger -t immortaldir
KillMode=process
Restart=always
Restart=on-failure
Type=simple
User=root

[Install]
WantedBy=multi-user.target
</pre>

これで基本的な設定は終わりです。まだ設定ファイルは何もありませんが、先に監視を開始しておきます。

``` sh
$ sudo systemctl start immortaldir
$ sudo systemctl enable immortaldir
```

<br />
## btcd を daemonize してみる

試しに [btcd](https://github.com/btcsuite/btcd) を daemonize してみます。

btcd を実行するための app ユーザーを作成し、btcd をインストールします。インストール方法については [こちらのエントリ](https://m0t0k1ch1st0ry.com/blog/2018/01/24/lightning-network) をご覧ください。

インストールが終わったら、以下の内容で `/home/app/.btcd/btcd.conf` を作成します。

<pre>
[Application Options]
simnet=1
datadir=/home/app/.btcd/data
logdir=/home/app/.btcd/logs
rpcuser=btcd
rpcpass=btcd
txindex=1
</pre>

今回は daemonize の検証がしたいだけなので、お手軽に simnet で立ち上げます。

次に、immortaldir 用の設定ファイルを `/usr/local/etc/immortal/btcd.yml` で作成します。各設定の意味は [公式ドキュメント](https://immortal.run/post/immortal) に記載されているので、ここでは割愛します。

``` yml
cmd: /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
cwd: /home/app
env:
    HOME: /home/app
log:
    file: /var/log/btcd.log
    age: 86400
    num: 7
    size: 1
    timestamp: true
user: app
```

immortaldir が `/usr/local/etc/immortal/` を監視しているので、自動的にファイルが認識され、btcd が起動します。

<pre>
root      6678  0.0  0.3  47460  6380 ?        Ssl  16:36   0:00 immortal -c /usr/local/etc/immortal/btcd.yml -ctl btcd
app       6682  0.4  0.8 178432 17684 ?        Sl   16:36   0:00  \_ /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

immortal が supervisor となり、子プロセスの btcd を監視しているようです。念のため、btcd に対してコマンドが通るかも確認しておきます。

``` sh
$ btcctl --simnet getinfo
```

``` json
{
  "version": 120000,
  "protocolversion": 70002,
  "blocks": 0,
  "timeoffset": 0,
  "connections": 0,
  "proxy": "",
  "difficulty": 1,
  "testnet": false,
  "relayfee": 0.00001,
  "errors": ""
}
```

daemon の情報は [immortalctl](https://immortal.run/post/immortalctl) を利用することでも確認できます。`status` コマンドを実行してみます。

``` sh
$ sudo immortalctl status
```

<pre>
 PID      Up   Down   Name   CMD
6682   51.0s          btcd   /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

Up で起動からの経過時間が確認できます。

起動が確認できたので、次は徐ろに btcd を終了させてみます。immortal が監視してくれているので、すぐに再起動されるはずです。

``` sh
$ btcctl --simnet stop
```

<pre>
btcd stopping.
</pre>

再起動されているか確認します。

``` sh
$ sudo immortalctl status
```

<pre>
 PID     Up   Down   Name   CMD
6702   2.2s          btcd   /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

PID が新しくなり、Up もリセットされました。

<pre>
root      6678  0.0  0.3  47460  6380 ?        Ssl  16:36   0:00 immortal -c /usr/local/etc/immortal/btcd.yml -ctl btcd
app       6702  0.4  1.1 240796 22720 ?        Sl   16:37   0:00  \_ /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

次に、`halt` コマンドを実行して再起動してみます。なお、`halt` コマンドの挙動は以下のようになっているようです。

> Stop the service by sending a TERM signal and exit the supervisor, if using immortaldir(8) the service will be restarted, first will be stopped and immortaldir(8) will start it up again.

supervisor である immortal も終了させ、immortaldir によって再起動が行われるようです。実際にやってみます。

``` sh
$ sudo immortalctl halt btcd
```

<pre>
PID   Up   Down   Name   CMD
</pre>

``` sh
$ sudo immortalctl status
```

<pre>
 PID     Up   Down   Name   CMD
6736   0.7s          btcd   /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

PID が変わり、Up もリセットされました。immortal の PID も変わっています。

<pre>
root      6732  0.0  0.3 121192  6928 ?        Ssl  16:38   0:00 immortal -c /usr/local/etc/immortal/btcd.yml -ctl btcd
app       6736  0.4  0.8 243968 17784 ?        Sl   16:38   0:00  \_ /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

daemon を停止させたい（再起動もさせたくない）場合は `stop` コマンドを実行します。

``` sh
$ sudo immortalctl stop btcd
```

<pre>
 PID        Up   Down   Name   CMD
6736   1m12.5s          btcd   /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

``` sh
$ sudo immortalctl status
```

<pre>
 PID   Up    Down   Name   CMD
6736        23.2s   btcd   /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

Up が表示されなくなり、代わりに Down（停止からの経過時間）が表示されました。immortal の子プロセスも消えています。

<pre>
root      6732  0.0  0.3 121192  6928 ?        Ssl  16:38   0:00 immortal -c /usr/local/etc/immortal/btcd.yml -ctl btcd
</pre>

もう一度 btcd を起動するには `start` コマンドを実行します。

``` sh
$ sudo immortalctl start btcd
```

<pre>
 PID     Up   Down   Name   CMD
6768   0.0s          btcd   /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

``` sh
$ sudo immortalctl status
```

<pre>
 PID      Up   Down   Name   CMD
6768   23.6s          btcd   /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

再び Up が表示され、immortal の子プロセスも復活したようです。

<pre>
root      6732  0.0  0.3 121192  6928 ?        Ssl  16:38   0:00 immortal -c /usr/local/etc/immortal/btcd.yml -ctl btcd
app       6768  0.7  1.2 186628 25952 ?        Sl   16:40   0:00  \_ /home/app/go/bin/btcd --configfile=/home/app/.btcd/btcd.conf
</pre>

<br />
## まとめ

- Golang 製の daemon 管理ツール [immortal](https://github.com/immortal/immortal) で [btcd](https://github.com/btcsuite/btcd) を daemonize してみました
- [immortalctl](https://immortal.run/post/immortalctl) を用いて daemon に対して基本的な管理操作を実行しつつ、その挙動を確認しました
