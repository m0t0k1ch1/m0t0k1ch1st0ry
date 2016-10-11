+++
date = "2016-07-30T02:23:32+09:00"
tags = [ "python", "mecab", "cabocha" ]
title = "CentOS で Python から MeCab と CaboCha を使えるようにする"
+++

手順が無駄なく綺麗にまとまってるとこをなかなか見つけられなくてごにゃごにゃしたので、最終的に自分が辿り着いた手順を後々のためにメモ。

<!--more-->

<br />
## 環境

``` sh
$ cat /etc/redhat-release
CentOS release 6.8 (Final)
$ arch
x86_64
```

<br />
## 下準備

``` sh
$ yum install bzip2 bzip2-devel gcc gcc-c++ git make openssl-devel readline-devel sqlite sqlite-devel zlib-devel
```

<br />
## pyenv

``` sh
$ git clone https://github.com/yyuu/pyenv.git ~/.pyenv
```

`~/.bashrc` に以下を追記。

``` sh
export PYENV_ROOT="${HOME}/.pyenv"
if [ -d "${PYENV_ROOT}" ]; then
  export PATH=${PYENV_ROOT}/bin:$PATH
  eval "$(pyenv init -)"
fi
```

で、読み込みなおす。

``` sh
$ source ~/.bashrc
```

Python 3.5.2 をインストール。

``` sh
$ pyenv --version
pyenv 20160726
$ pyenv install 3.5.2
```

インストールが終わったら確認して切り替える。

``` sh
$ pyenv versions
* system (set by /home/app/.pyenv/version)
  3.5.2
$ pyenv global 3.5.2
$ pyenv rehash
$ pyenv versions
  system
* 3.5.2 (set by /home/app/.pyenv/version)
```

一応 pip を最新版にしておく。

``` sh
$ pip install --upgrade pip
$ pip list | grep pip
pip (8.1.2)
```

<br />
## MeCab

まず本体。

``` sh
$ wget https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7cENtOXlicTFaRUE -O mecab-0.996.tar.gz
$ tar zxvf mecab-0.996.tar.gz
$ cd mecab-0.996
$ ./configure --with-charset=utf8 --enable-utf8-only
$ make
$ make install
$ mecab --version
mecab of 0.996
```

次に辞書。

``` sh
$ wget https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7MWVlSDBCSXZMTXM -O mecab-ipadic-2.7.0-20070801.tar.gz
$ tar zxvf mecab-ipadic-2.7.0-20070801.tar.gz
$ ./configure --with-charset=utf8
$ make
$ make install
```

で、Python から使えるようにしておく。

``` sh
$ pip install mecab-python3
$ pip list | grep mecab-python3
mecab-python3 (0.7)
```

<br />
## CaboCha

まず、CRF++ をインストールする。

``` sh
$ wget https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7QVR6VXJ5dWExSTQ -O CRF++-0.58.tar.gz
$ tar zxvf CRF++-0.58.tar.gz
$ cd CRF++-0.58
$ ./configure
$ make
$ make install
```

続いて CaboCha 本体のインストール。

``` sh
$ wget https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7SDd1Q1dUQkZQaUU -O cabocha-0.69.tar.bz2
$ tar jxvf cabocha-0.69.tar.bz2
$ cd cabocha-0.69
$ ./configure --with-charset=utf8 --enable-utf8-only
$ make
$ make install
$ cabocha --version
cabocha of 0.69
```

で、Python から使えるようにしておく。

``` sh
$ cd cabocha-0.69/python
$ python setup.py install
$ pip list | grep cabocha-python
cabocha-python (0.69)
```

おしまい。
