+++
date = "2013-04-25"
tags = [ "chef", "ruby" ]
title = "vs Chef 第1回戦"
+++

二流シェフ目指してがんばるぞう。

<!--more-->

<br />
## Chef って？

* 「サーバーの状態を管理し収束させるためのフレームワーク」
* Chef 以外にも Puppet というやつもいるよ
* Facebook 社が使っておられます
* ローカルの開発環境セットアップにも使える！！
* 利用形態は大きく2つ
 * Chef Server ＋ Chef Client（大規模環境管理）
 * Chef Solo（単独のコマンドとして Chef を実行）
* recipe：「コード化された手順書」「サーバーの状態」
* cookbook：「特定の recipe に必要なデータやファイルをまとめる入れ物」
* repository：「Chef の実行に必要な一連のファイルをまとめる入れ物」
* repository > cookbook > recipe

<br />
## 購入すると物凄く捗る

* [入門 Chef Solo](http://www.amazon.co.jp/%E5%85%A5%E9%96%80Chef-Solo-Infrastructure-Code-ebook/dp/B00BSPH158)

<br />
## ローカルの初期状態

* rbenv を用いて ruby 1.9.3.p392 をインストール済み
* chef 11.4.4 をインストール済み
* 初期設定も終わってたっぽい
* 終わってなければ以下を実行

``` sh
$ gem install chef
$ knife configure
```

<br />
## knife-solo 0.3.0 のインストール

* knife-solo 0.2.0 に注意：[knife-solo 0.2.0 で rsync エラーによって苦しまないためのたったひとつの方法](http://tk0miya.hatenablog.com/entry/2013/04/18/011339)
* ということで、knife-solo 0.3.0 をインストールするために奔走
* gem で knife-solo 0.3.0 が入る時代になっていれば無問題
* knife-solo をインストールするディレクトリを作成
* 当方、ここで入れた knife-solo にはパス通してますが、通ってない前提で書き続けます

``` sh
$ mkdir ~/knife-solo
$ cd knife-solo
```

* bundle 初期化

``` sh
$ bundle init
```

* `Gemfile` を編集する

``` sh
$ source "https://rubygems.org"
$ gem 'knife-solo', '0.3.0.pre3'
```

* knife-solo をインストール

``` sh
$ bundle install --path=vendor/bundle --binstubs
```

* `.chef/knife.rb` に以下を追加（knife-solo 0.3.0 のときだけ）

``` ruby
knife[:solo_path] = '/tmp/chef-solo'
```

* repository をつくる

``` sh
$ ~/knife-solo/bin/knife solo init your-repo
```

* 先人たちの cookbook を手に入れるための準備

``` sh
$ cd your-repo
$ git init
$ git add .
$ git commit -m 'first commit'
```

<br />
## 先人たちの cookbook を手に入れてみる

* [Opscode Community](http://community.opscode.com/cookbooks) にユーザー登録
* 秘密鍵をダウンロードして `~/.chef/username.pem` にパーミッション 600 で保存
* 試しに yum の cookbook をダウンロード

``` sh
$ knife cookbook site vendor yum -o cookbooks
```

* とりあえず、参考になりそうな以下をダウンロード

``` sh
$ knife cookbook site vendor nginx -o cookbooks
$ knife cookbook site vendor mysql -o cookbooks
$ knife cookbook site vendor daemontools -o cookbooks
$ knife cookbook site vendor perlbrew -o cookbooks
$ knife cookbook site vendor jenkins -o cookbooks
```

* サードパーティーの cookbook を扱うためのツールとして [Berkshelf](http://berkshelf.com/) や [librarian](https://github.com/applicationsonline/librarian) というものもあるらしいが、なんかそういうのあんまり好きじゃないので今回は使わない

<br />
## とりあえず nginx の cookbook をつくって cook してみる

* cookbook の雛形をつくる

``` sh
$ knife cookbook create nginx -o site-cookbooks
```

* recipe を書く
  * package とか service とか template は Chef が提供する [resource](http://docs.opscode.com/resource.html)
  * user と group は attributes を参照する

``` ruby
package 'nginx' do
  action :install
end

service 'nginx' do
  supports :status => true, :restart => true, :reload => true
  action [ :enable, :start ]
end

template 'nginx.conf' do
  path '/etc/nginx/nginx.conf'
  source 'nginx.conf.erb'
  owner node['nginx']['user']
  group node['nginx']['group']
  mode 0644
  notifies :reload, 'service[nginx]'
end
```

* port を指定できるようにして template を書く

``` nginx
user             nginx;
worker_processes 1;
error_log        /var/log/nginx/error.log;
pid              /var/run/nginx.pid;

events {
    worker_connections 1024;
}

http {
    include      /etc/nginx/mime.types;
    default_type application/octet-stream;

    server {
        listen      <%= node['nginx']['port'] %>;
        server_name localhost;
        location / {
            root  /usr/share/nginx/html;
            index index.html index.htm;
        }
    }
}
```

* attributes にデフォルト値を書いておく

``` ruby
default['nginx']['user']  = 'user-name'
default['nginx']['group'] = 'user-group'
default['nginx']['port']  = '80'
```

* JSON を作成
  * Chef Solo 実行時に渡す変数の値やどの recipe を実行するかを設定する
  * JSON はダブルクォーテーションじゃないとだめ

``` json
{
    "run_list" : [
        "nginx"
    ]
}
```

* 対象サーバーに Chef をインストール
  * 対象サーバー：CentOS
  * SSH の設定は `~/.ssh/config` でうまいことやってる前提

``` sh
$ ~/knife-solo/bin/knife solo prepare host-name
```

* cook する（cook するユーザーはパスワードなしで sudo 実行可能にしておく）

``` sh
$ ~/knife-solo/bin/knife solo cook host-name
```

<br />
## 参考

* [Chef Solo の正しい始め方](http://tsuchikazu.net/chef_solo_start/)
* [macにknife-soloをインストール](http://devlog.mitsugeek.net/entry/2013/03/25/mac%E3%81%ABknife-solo%E3%82%92%E3%82%A4%E3%83%B3%E3%82%B9%E3%83%88%E3%83%BC%E3%83%AB(prepare%E3%81%BE%E3%81%A7%E7%A2%BA%E8%AA%8D%EF%BC%89)
* [[Mac]『入門Chef Solo』を読んで試してみた](http://blog.hello-world.jp.net/?p=461)
