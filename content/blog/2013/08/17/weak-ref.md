+++
title = '弱い参照で循環参照を解決する - Perl'
tags = ['perl']
date = '2013-08-17'
+++

Mouse のプロパティの `weak_ref` の使いどころで「う…」ってなったのでメモ。

<!--more-->

先輩方に良記事を教えてもらいましたが、もともと参照とかそっち系の話にはてんで弱い系プログラマーなので、自分なりにまとめてみてもよいかと思いまして候う。

## 参照カウンタ

- 変数の値が参照されている箇所を数え上げているもの
- Perl では参照カウンタが 0 になるとその変数が解放される

簡単な例で見てみます。

``` perl
use strict;
use warnings;

use Devel::Peek;

my $piyo = 'piyo';
Dump $piyo;

my $piyo_ref = \$piyo;
Dump $piyo;
```

上記のプログラムを実行すると結果は以下のようになり、`REFCNT`（参照カウント）が増えているのがわかります。

``` txt
SV = PV(0x7fd4aa03ab90) at 0x7fd4aa08b408
  REFCNT = 1
  FLAGS = (PADMY,POK,pPOK)
  PV = 0x7fd4a9c21a90 "piyo"\0
  CUR = 4
  LEN = 16
SV = PV(0x7fd4aa03ab90) at 0x7fd4aa08b408
  REFCNT = 2
  FLAGS = (PADMY,POK,pPOK)
  PV = 0x7fd4a9c21a90 "piyo"\0
  CUR = 4
  LEN = 16
```

## 循環参照

- お互いに参照し合う状態のこと
- Perl は参照カウンタを用いてメモリ管理しているため、循環参照がメモリリークを引き起こす原因となる（プログラム終了時まで参照カウンタが 0 にならなくなっちゃう場合があるので）

例えば Mouse を用いて以下のようなオブジェクトを定義したとします。

``` perl
package Piyo;
use Mouse;

has poyo => (
    is => 'rw',
);

sub set_poyo {
    my ($self, $poyo) = @_;
    $self->poyo($poyo);
}

__PACKAGE__->meta->make_immutable;

no Mouse;

1;
```

``` perl
package Poyo;
use Mouse;

has piyo => (
    is => 'rw',
);

sub set_piyo {
    my ($self, $piyo) = @_;
    $self->piyo($piyo);
}

__PACKAGE__->meta->make_immutable;

no Mouse;

1;
```

で、Devel::Cycle を用いた以下のようなコードで循環参照になっちゃってるかどうかを調べてみます。

``` perl
use strict;
use warnings;

use Devel::Cycle;

use Piyo;
use Poyo;

my $piyo = Piyo->new();
my $poyo = Poyo->new();

$poyo->set_piyo($piyo);
$piyo->set_poyo($poyo);

find_cycle($piyo);
```

結果、以下のように循環参照が検出されます。

``` txt
Cycle (1):
                  $Piyo::A->{'poyo'} => \%Poyo::B
                  $Poyo::B->{'piyo'} => \%Piyo::A
```

## 循環参照を解決する

- その1：Scalar::Util::weaken を使って、参照カウンタを増やさないようにする

``` perl
package Piyo;
use Mouse;

has poyo => (
    is => 'rw',
);

sub set_poyo {
    my ($self, $poyo) = @_;
    $self->poyo($poyo);
    Scalar::Util::weaken($self->poyo);
}

__PACKAGE__->meta->make_immutable;

no Mouse;

1;
```

- その2：Moose とか Mouse のプロパティを弱い参照にして、スコープを抜けると参照が破棄されるようにする

``` perl
package Piyo;
use Mouse;

has poyo => (
    is       => 'rw',
    weak_ref => 1,
);

sub set_poyo {
    my ($self, $poyo) = @_;
    $self->poyo($poyo);
}

__PACKAGE__->meta->make_immutable;

no Mouse;

1;
```

## 参考

- [循環参照と弱い参照](http://memememomo.hatenablog.com/entry/20100528/1275005888)
- [リファレンスの循環参照によるメモリリークを Scalar::Util::weaken で解決する](http://d.hatena.ne.jp/naoya/20051012/1129115986)
- [循環参照のはなし](http://www.slideshare.net/hiratara/ss-10539893)
