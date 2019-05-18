+++
date = "2013-04-30"
tags = [ "emacs", "perl" ]
title = "Emacs で Perl 描きたい"
+++

Vim と Emacs どっちも触れないのにエンジニアなんですか？へっ？とか言われるのは心にぐさっとくるので、Emacs でなんとか Perl を書けるようにしてみました。

<!--more-->

ちなみに当方、Emacs についても Perl についても初心者です。ｳｪｰｲ。

あ、Vim じゃなくて Emacs を選んだ理由は特にないです。気分です。

## 方針

### 余分な設定書きたくない

- よくわかんない他人の設定はコピペしたくない
- とりあえず最低限でシンプルな感じが良い

### ELPA 使ってみたい

- ELPA にあるものは ELPA でインストール
- ELPA にないものは auto-install でインストール

### Perl と Markdown は書きたい

- 仕事柄、Perl は必須
- Markdown はこのブログ執筆用

## インストールした Elisp

- init-loader

### auto-install でインストール

- perl-completion

### ELPA でインストール

- anything
- auto-complete
- auto-install
- flymake
- markdown-mode

## できあがった設定

<div class="github-card" data-user="m0t0k1ch1" data-repo="dotfiles"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

### 設定概要

- 設定ファイルは分割して init-loader で読み込む
- cperl-mode の hook として、auto-complete + perl-completion + flymake
- .md と .markdown は markdown-mode で開く

``` lisp
(defalias 'perl-mode 'cperl-mode)

(add-to-list 'auto-mode-alist '("\\.t$" . cperl-mode))
(add-to-list 'auto-mode-alist '("\\.psgi$" . cperl-mode))

(add-hook 'cperl-mode-hook
          (lambda()
            (setq cperl-indent-level 4
                  cperl-close-paren-offset -4
                  cperl-continued-statement-offset 4
                  cperl-indent-parens-as-block t
                  cperl-tab-always-indent t)))

;; perl-completion
(require 'perl-completion)
(add-hook 'cperl-mode-hook
          (lambda()
            (perl-completion-mode t)))

;; auto-complete
(add-hook 'cperl-mode-hook
          (lambda ()
            (when (require 'auto-complete nil t) ; no error whatever auto-complete.el is not installed.
              (auto-complete-mode t)
              (make-variable-buffer-local 'ac-sources)
              (setq ac-sources '(ac-source-perl-completion)))))

;; flymake-mode
(add-hook 'cperl-mode-hook
          (lambda ()
            (flymake-mode t)))
```

``` lisp
(add-to-list 'auto-mode-alist '("\\.md$" . markdown-mode))
(add-to-list 'auto-mode-alist '("\\.markdown$" . markdown-mode))
```

## まとめ

- 目的通り、最低限のものがにシンプルにまとまった気がする
- 次はキーバインドちゃんと覚えます←
- 開発環境は仕事の効率に大きく影響するし、月1くらいで見直していきたい

## 参考

- [Perl Hacks on Emacs](http://typester.stfuawsc.com/slides/perlcasual2/start.html)
  - インストールする Elisp を決める際に参考になりました
  - 基本的な設定の部分も [typesterさんのもの](https://github.com/typester/emacs-config/blob/master/conf/init.el) を参考にさせていただきました
- [perl-completion をインストールした](http://d.hatena.ne.jp/a666666/20100524/1274634774)
  - auto-complete と perl-completion がうまく動かなかったときに現れた救世主
