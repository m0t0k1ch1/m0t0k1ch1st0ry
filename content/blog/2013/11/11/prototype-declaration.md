+++
date = "2013-11-11"
tags = [ "perl" ]
title = "サブルーチンの引数としてサブルーチンリファレンスを指定する - Perl"
+++

Perl の例外処理の勉強がてら [Try::Tiny](http://search.cpan.org/~doy/Try-Tiny-0.18/lib/Try/Tiny.pm) の実装を眺めていたところ、今更ながらサブルーチンの引数としてサブルーチンリファレンスを指定できることを知った若輩者でございます。。こういうテクニシャンぽいことはすぐ忘れるのでメモ。

<!--more-->

<br />
## 例えば

サブルーチンリファレンスを引数にとると、`grep { ... } @array` や `map { ... } @array` みたいな、いかにも構文チックな関数を定義できる。

試しに、`grep` と逆のことをする eliminate なるものをつくってみた。

``` perl
use strict;
use warnings;

use feature 'say';

sub eliminate (&@) {
    my ($block, @array) = @_;

    my @result;
    for (@array) {
        unless ($block->()) {
            push @result, $_;
        }
    }

    @result;
}

my @array = (3,'m',-2,0,'t',-8,2,0,5,'k',1,9,'c',-4,-6,'h',7,1);
say eliminate { $_ =~ /^-?\d+$/ && ($_ > 1 || $_ < 0) } @array;
```

実行結果は以下。

<pre>
m0t0k1ch1
</pre>

<br />
## 注意点

* サブルーチンリファレンスを指定可能なのは第一引数だけで、`&` で指定する
