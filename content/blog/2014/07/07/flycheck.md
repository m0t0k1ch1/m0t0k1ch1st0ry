+++
date = "2014-07-07"
tags = [ "emacs", "perl" ]
title = "flycheck しながら Perl を描くときに"
+++

Emacs、「モダン」という言葉に誘われて flymake から flycheck に乗り換えたときに Perl の @INC 周りの設定で少しがんばったのでメモ。

<!--more-->

## 解決したいこと

`my-project/scripts/poyo.pl` とかを編集しているとき、`my-project/lib` や `my-project/local/lib/perl5`（Carton でインストールしたモジュールが入ってる）は @INC に含まれていないので、例えば自分のプロジェクト内のモジュールを `use MyApp;` という具合で読み込もうとすると flycheck さんに怒られてしまう。これを解決したい。

## 解決策

解決策を探してみても、flymake の話が少々出てくるくらい。flycheck の話はなかなか出てこなかった。それでもめげずに探したところ、以下のリンクに辿り着いた。

- [良い感じで flycheck する](https://gist.github.com/co-me/7363369)
- [flycheck で C/C++ のエラーチェック定義を追加する](http://qiita.com/akisute3@github/items/6fb94c30f92dae2a24ee)

どうやら、`flycheck-define-checker` というのを使うことで、エラーチェックの定義を追加することができる模様。見よう見まねで、できるだけシンプルに以下のような設定を書いてみた。

``` lisp
(flycheck-define-checker perl-project-libs
  "A perl syntax checker."
  :command ("perl"
            "-MProject::Libs lib_dirs => [qw(local/lib/perl5)]"
            "-wc"
            source-inplace)
  :error-patterns ((error line-start
                          (minimal-match (message))
                          " at " (file-name) " line " line
                          (or "." (and ", " (zero-or-more not-newline)))
                          line-end))
  :modes (cperl-mode))
```

Project::Libs、便利である。で、これを `cperl-mode-hook` として登録する。

``` lisp
(add-hook 'cperl-mode-hook
          (lambda ()
            (flycheck-mode t)
            (setq flycheck-checker 'perl-project-libs)))
```

すると、きちんと解決できていた。めでたし。
