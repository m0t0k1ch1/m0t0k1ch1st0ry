+++
date = "2013-05-09"
tags = [ "perl" ]
title = "smart match の smart さに物申す - Perl"
+++

Perl の smart match ではまったのでめもめも。

<!--more-->

<br />
## 1. 比較的 smart な使い方

* smart か smart でないかと問われれば smart かなと思う
* `grep` とかでごりって感じでやるよりは smart かなと思う

``` perl
my $hoge = 3;
if ($hoge ~~ [1, 3, 5]) {
    warn '$hoge is 1 or 3 or 5';
}
```

* 当然、出力は以下

``` txt
$hoge is 1 or 3 or 5
```

<br />
## 2. smart さを発揮できない使い方

``` perl
if ('0.00' ~~ '0') {
    warn 'smart match!';
}

if ('0.00' == '0') {
    warn 'match!';
}
```

* なんと、出力は以下

``` txt
match!
```

* smart match も内部的に数値比較をする仕組みを含んでいるはずなのですが、この場合はそれよりも文字列比較が優先されてしまうよう…

<br />
## まとめ

* smart match は思ったより smart 感がない
* 1 みたいな使い方以外で乱用するのは控えた方が良さそう
