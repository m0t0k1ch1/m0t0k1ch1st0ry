+++
date = "2016-09-10T01:23:35+09:00"
tags = [ "rpm", "golang" ]
title = "Golang で書いた何かを RPM パッケージにして yum install したい"
+++

表題のようなことをやる必要が出てきたので、その検証をした。

<br />
## 事前調査

「rpm package つくり方」でググるとこから始めるレベルの知識しかなく、先行きがだいぶ怪しまれたので、きちんと動いてる目指すべき状態を明確にしたいなと思った。で、身近に同じようなことやってるプロダクトはないだろうかと考えたとき、[mackerel-agent](https://github.com/mackerelio/mackerel-agent) が真っ先に頭に浮かんだので、迷う度に参考にさせていただきました。ありがとうございます🙏

<br />
## 下準備

クロスコンパイルできる環境が整っていない場合、整えておく。例えば macbook で homebrew を使っているような場合は、以下のようにオプションをつけて go をインストールすればいける。

``` sh
$ brew install go --cross-compile-common
```

一旦はさくっと検証だけしたかったので `--cross-compile-common` にしたが、全部入りの `--cross-compile-all` もあるので、それはお好きな方でどうぞ。

<br />
## rpmbuild するための準備

### SOURCES

Golang で [こんな感じのコード](https://github.com/m0t0k1ch1/go-rpm-sample/blob/master/main.go) を描いた。

最終的にはデーモンっぽい状態で起動する予定なので、最低限の挙動だけを実装してみた。ざっくりとは

- 基本的には sleep 入れつつ for でループしながらただただログを垂れ流す
- SIGHUP・SIGTERM・SIGQUIT はハンドリングしてあげる
- 起動時に自分の pid をファイルに吐いて、defer で削除する

という感じ。後ろ2つは、後々 RPM パッケージとしてインストールして、例えば `service poyo stop` とかされたときの挙動を想定していたりする。

実装が終わったらクロスコンパイルする。今回は Linux - amd64 用。

``` sh
$ GOOS=linux GOARCH=amd64 go build
```

できたバイナリは [ここ](https://github.com/m0t0k1ch1/go-rpm-sample/tree/master/package/SOURCES) に `go-rpm-sample-0.1.0` として置いておいた。ホントは tar とかで固めといた方がよい気がする。

また、start・stop・status・restart くらいはできるよう、[init スクリプト](https://github.com/m0t0k1ch1/go-rpm-sample/blob/master/package/SOURCES/go-rpm-sample-0.1.0.initd) もがんばって準備した。エラーハンドリングが甘いけど一旦はこれで。

### SPECS

rpmbuild に必要な spec ファイルを準備する。ググったり mackerel-agent の spec ファイルを参考にしたりしながら辿り着いたのが以下。結構シンプルで、記法を知っていなくてもなんとなく読めるレベルなんじゃないかなと思う。

``` txt
%define _binaries_in_noarch_packages_terminate_build 0

Summary: a simple sample application
Name:    go-rpm-sample
Version: 0.1.0
Release: 1
License: MIT
Group:   Applications/System
URL:     https://github.com/m0t0k1ch1/go-rpm-sample

Source0:   %{name}-%{version}
Source1:   %{name}-%{version}.initd
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}-root

%description
%{summary}

%prep

%build

%install
%{__rm} -rf %{buildroot}
%{__install} -Dp -m0755 %{SOURCE0} %{buildroot}/usr/local/bin/%{name}
%{__install} -Dp -m0755 %{SOURCE1} %{buildroot}/%{_initrddir}/%{name}

%clean
%{__rm} -rf %{buildroot}

%post
/sbin/chkconfig --add %{name}

%files
%defattr(-,root,root)
/usr/local/bin/%{name}
%{_initrddir}/%{name}
```

### ここまででできたもの

[ここ](https://github.com/m0t0k1ch1/go-rpm-sample/tree/master/package) に置いておいた。構造は以下。

``` txt
package
├── SOURCES
│   ├── go-rpm-sample-0.1.0
│   └── go-rpm-sample-0.1.0.initd
└── SPECS
    └── go-rpm-sample.spec
```

<br />
## CentOS 6 で rpmbuild する

まずは必要なものを yum で入れる。

``` sh
$ yum install rpmdevtools yum-utils
```

で、rpmbuild するための準備。

``` sh
$ rpmdev-setuptree
```

上記コマンドを実行すると、`$HOME/rpmbuild` というディレクトリができて、その中が以下のようにセットアップされる。

``` txt
rpmbuild/
├── BUILD
├── RPMS
├── SOURCES
├── SPECS
└── SRPMS
```

で、先ほど準備した諸々を SOURCES と SPECS の中にそのまま配置して、以下を実行。

``` sh
$ rpmbuild -bb $HOME/rpmbuild/SPECS/go-rpm-sample.spec
```

無事成功すると、`$HOME/rpmbuild/RPMS/x86_64` 以下に `go-rpm-sample-0.1.0-1.x86_64.rpm` という RPM パッケージができているはず。めでたい。

試しに直接インストールしてみる。

``` sh
$ rpm -ivh $HOME/rpmbuild/RPMS/x86_64/go-rpm-sample-0.1.0-1.x86_64.rpm
```

インストールが終わったら念のため確認。

``` sh
$ rpm -qa | grep go-rpm-sample
```

``` txt
go-rpm-sample-0.1.0-1.x86_64
```

正常にインストールできた模様なので、一通りの動作を確認してみる。

``` sh
$ service go-rpm-sample start
Starting go-rpm-sample:                                    [  OK  ]
$ service go-rpm-sample status
go-rpm-sample (pid  5560) is running...
$ service go-rpm-sample restart
Stopping go-rpm-sample:                                    [  OK  ]
Starting go-rpm-sample:                                    [  OK  ]
$ service go-rpm-sample stop
Stopping go-rpm-sample:                                    [  OK  ]
```

大丈夫そう。

<br />
## yum install できるようにする

yum 用の repo をつくって、先ほどこしらえた RPM パッケージを yum install できるようにするところまでやってみる。

まずは必要なものを yum で入れる。

``` sh
$ yum install createrepo
```

適当に mkdir する。

``` sh
$ mkdir repo
```

repo 以下に先ほどこしらえた `go-rpm-sample-0.1.0-1.x86_64.rpm` を配置して、以下を実行する。

``` sh
$ createrepo repo
```

これでおしまいらしい。思ってたよりかなりお手軽である。

で、今回はこの repo を GitHub に置いて、gh-pages ブランチを切っておいた。gh-pages ブランチを切っておくと、https://m0t0k1ch1.github.io/rpm/centos/latest/x86_64/go-rpm-sample-0.1.0-1.x86_64.rpm という URL で wget したりできる。

<div class="github-card" data-user="m0t0k1ch1" data-repo="yum-repo"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

で、例えばまっさらな別 CentOS 6 に `/etc/yum.repos.d/m0t0k1ch1.repo` を設置する。内容は以下。

``` txt
[m0t0k1ch1]
name=m0t0k1ch1
baseurl=https://m0t0k1ch1.github.io/yum-repo/centos/latest/$basearch
gpgcheck=0
```

これで yum install できる。初めて自分のつくったものを yum install する記念すべき瞬間。

``` sh
$ yum install --enablerepo=m0t0k1ch1 go-rpm-sample
```

``` txt
読み込んだプラグイン:fastestmirror
インストール処理の設定をしています
Loading mirror speeds from cached hostfile
 * base: ftp.iij.ad.jp
 * epel: ftp.iij.ad.jp
 * extras: ftp.iij.ad.jp
 * updates: ftp.iij.ad.jp
m0t0k1ch1                                                                         | 2.9 kB     00:00
m0t0k1ch1/primary_db                                                              | 1.8 kB     00:00
依存性の解決をしています
--> トランザクションの確認を実行しています。
---> Package go-rpm-sample.x86_64 0:0.1.0-1 will be インストール
--> 依存性解決を終了しました。

依存性を解決しました

=========================================================================================================
 パッケージ                  アーキテクチャ       バージョン               リポジトリー             容量
=========================================================================================================
インストールしています:
 go-rpm-sample               x86_64               0.1.0-1                  m0t0k1ch1               451 k

トランザクションの要約
=========================================================================================================
インストール         1 パッケージ

総ダウンロード容量: 451 k
インストール済み容量: 1.4 M
これでいいですか? [y/N]y
パッケージをダウンロードしています:
go-rpm-sample-0.1.0-1.x86_64.rpm                                                  | 451 kB     00:00
rpm_check_debug を実行しています
トランザクションのテストを実行しています
トランザクションのテストを成功しました
トランザクションを実行しています
警告: RPMDB は yum 以外で変更されました。
  インストールしています  : go-rpm-sample-0.1.0-1.x86_64                                             1/1
  Verifying               : go-rpm-sample-0.1.0-1.x86_64                                             1/1

インストール:
  go-rpm-sample.x86_64 0:0.1.0-1

完了しました!
```

完了しました！

先ほどと同様、挙動を確認してみる。

``` sh
$ service go-rpm-sample start
Starting go-rpm-sample:                                    [  OK  ]
$ service go-rpm-sample status
go-rpm-sample (pid  5560) is running...
$ service go-rpm-sample restart
Stopping go-rpm-sample:                                    [  OK  ]
Starting go-rpm-sample:                                    [  OK  ]
$ service go-rpm-sample stop
Stopping go-rpm-sample:                                    [  OK  ]
```

大丈夫そう。

とりあえず、これでやりたかった検証は終わり。目的とはちょっとズレるけど、自分用の repo もっとくと便利なことありそうだなあと思った。
