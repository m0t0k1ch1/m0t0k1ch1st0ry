+++
date = "2013-05-01"
tags = [ "chef", "perl", "ruby" ]
title = "vs Chef 第2回戦"
+++

Chef と戯れ始めてはや1週間。  
先輩方の手厚いサポートもあり、少しずつ敵の姿をとらえつつあります m0t0k1ch1 です。

<!--more-->

ちなみに、この記事は [第1回戦]({{< ref "blog/2013/04/25/vs-chef.md" >}}) の続編的なノリでお届けさせていただいております。

## 今回の課題

- [perl-build](https://github.com/tokuhirom/Perl-Build) で任意のバージョンの Perl をインストールしてくれる cookbook をつくってみよう

## 下準備

- いつも通り repository の中に入って cookbook の雛形をつくる、ただそれだけ

``` sh
$ knife cookbook create perl -o site-cookbooks
```

## 完成した cookbook の内容

- recipe
  - Perl のバージョンは指定できるように
  - 同じバージョンの Perl はインストールしない

``` ruby
perl_user    = node['perl']['user']
perl_group   = node['perl']['group']
perl_version = node['perl']['version']

binary_dir = "/home/#{perl_user}/bin"
perl5_dir  = "/home/#{perl_user}/perl5"

directory "#{binary_dir}" do
  owner perl_user
  group perl_group
  mode  0755
  action :create
end

remote_file "#{binary_dir}/perl-build"  do
  source node['perl']['perl-build']['binary-url']
  owner  perl_user
  group  perl_group
  mode   0755
  action :create_if_missing
end

bash 'install perl' do
  user  perl_user
  group perl_group
  not_if { File.exists?("#{perl5_dir}/perl-#{perl_version}") }
  code <<-"..."
    #{binary_dir}/perl-build #{perl_version} #{perl5_dir}/perl-#{perl_version}
  ...
end
```

- attribute

``` ruby
default['perl']['user']    = 'user-name'
default['perl']['group']   = 'user-group'
default['perl']['version'] = '5.16.3'

default['perl']['perl-build']['binary-url'] = 'https://raw.github.com/tokuhirom/Perl-Build/master/perl-build'
```

- JSON
  - `run_list` の中に Perl を含める

## まとめ

- Chef の雰囲気、なんとなくつかめてきた
- `owner` や `user` は、無指定だと極論なんでもいいってことなので、ちゃんと指定する
- `remote_file` とか `bash` 使いだすとなんでもできてしまう感があって乱用してしまいがちだが、なるべく「これでしかムリだわ！」ってときしか使わないように心がけた方が良い（今回は使っちゃいましたが…）
- `remote_file` するときに `:create_if_missing` は便利
  - 先輩に教えていただくまで気づかなかった…
  - ドキュメント読む癖をつけれ俺
