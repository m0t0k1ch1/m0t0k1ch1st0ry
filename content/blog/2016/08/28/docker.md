+++
date = "2016-08-28T02:08:15+09:00"
tags = [ "docker" ]
title = "Docker Toolbox をアンインストールして Docker for Mac を試してみる"
+++

そろそろ Docker の使いドコロだなというタイミングが来そうなので、表題の通り、以前にインストールしていた [Docker Toolbox](https://www.docker.com/products/docker-toolbox) をアンインストールして [Docker for Mac](https://www.docker.com/products/docker) を試してみることにした。

<!--more-->

<br />
## Docker Toolbox をアンインストールする

まず、以前にインストールした Docker Toolbox とさようならする。こういうとこをきちんとやっとかないと、後々「日頃の行い」起因でよからぬことが発生する。

[こちらのエントリ](http://qiita.com/minamijoyo/items/ec5b35382797ac08e067) で紹介されているように、[OSX 用のアンインストールスクリプト](https://github.com/docker/toolbox/blob/master/osx/uninstall.sh) が公式の repo にあるので、それを落としてきて実行する。

``` sh
$ git clone git@github.com:docker/toolbox.git
$ cd toolbox
$ sudo osx/uninstall.sh
```

## Docker for Mac をインストールする

[Getting Started with Docker for Mac](https://docs.docker.com/docker-for-mac) の手順に従って、stable の方をインストール。簡単。

インストールできたらバージョンを確認。

``` sh
$ docker --version
```

<pre>
Docker version 1.12.0, build 8eab29e
</pre>

``` sh
$ docker-compose --version
```

<pre>
docker-compose version 1.8.0, build f3628c7
</pre>

``` sh
$ docker-machine --version
```

<pre>
docker-machine version 0.8.0, build b85aac1
</pre>

念を入れてきちんとインストールできているか確認。

``` sh
$ docker run hello-world
```

<pre>
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
c04b14da8d14: Pull complete
Digest: sha256:0256e8a36e2070f7bf2d0b0763dbabdd67798512411de4cdcf9431a1feb60fd9
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker Hub account:
 https://hub.docker.com

For more examples and ideas, visit:
 https://docs.docker.com/engine/userguide/
</pre>

> Hello from Docker!
> This message shows that your installation appears to be working correctly.

とのこと。大丈夫そう。

<br />
## 試しに既存のコンテナを起動してみる

``` sh
$ docker run -d -p 80:80 --name webserver nginx
```

<pre>
Unable to find image 'nginx:latest' locally
latest: Pulling from library/nginx
357ea8c3d80b: Pull complete
0fc04568277e: Pull complete
0bed9719ddcb: Pull complete
Digest: sha256:d33834dd25d330da75dccd8add3ae2c9d7bb97f502b421b02cecb6cb7b34a1b6
Status: Downloaded newer image for nginx:latest
9be7159c9c160130032598dd52b71442226be957e766c82b94d645b1703db168
</pre>

http://localhost に行くと、お馴染みの nginx 画面が表示される。

## docker/whalesay を題材にして基本的な挙動を確認する

とりあえず、run してみる。

``` sh
$ docker run docker/whalesay cowsay poyo!
```

<pre>
Unable to find image 'docker/whalesay:latest' locally
latest: Pulling from docker/whalesay
e190868d63f8: Pull complete
909cd34c6fd7: Pull complete
0b9bfabab7c1: Pull complete
a3ed95caeb02: Pull complete
00bf65475aba: Pull complete
c57b6bcc83e3: Pull complete
8978f6879e2f: Pull complete
8eed3712d2cf: Pull complete
Digest: sha256:178598e51a26abbc958b8a2e48825c90bc22e641de3d31e18aaf55f3258ba93b
Status: Downloaded newer image for docker/whalesay:latest
 _______
< poyo! >
 -------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
</pre>

poyo!

続いて、くじらくんがありがたい言葉を発せるようにするために image を改良する。

``` sh
$ mkdir mydockerbuild
$ cd mydockerbuild
$ touch Dockerfile
```

チュートリアルには以下のように記載されている。

> Open the Dockerfile in a visual text editor like Atom or Sublime, or a text based editor like vi, or nano (https://www.nano-editor.org/).

emacs が候補にないことに若干の悲しみを感じながら、emacs で以下の内容を Dockerfile に描く。

<pre>
FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD /usr/games/fortune -a | cowsay
</pre>

`docker/whalesay` をベースに、fortunes を追加でインストールしてから cowsay している。

Dockerfile が準備できたので、build する。

``` sh
$ docker build -t docker-whale .
```

<pre>
Sending build context to Docker daemon 2.048 kB
...
Removing intermediate container 1c1754725007
Successfully built d649997c162a
</pre>

新しく `docker-whale` ができていることを確認する。

``` sh
$ docker images
```

<pre>
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
docker-whale        latest              d649997c162a        36 seconds ago      274.9 MB
nginx               latest              4efb2fcdb1ab        3 days ago          183.4 MB
hello-world         latest              c54a2cc56cbb        8 weeks ago         1.848 kB
docker/whalesay     latest              6b362a9f73eb        15 months ago       247 MB
</pre>

できているようなので、run してみる。

``` sh
$ docker run docker-whale
```

<pre>
 ________________________________________
/ The Golden Rule is of no use to you    \
| whatever unless you realize it is your |
| move.                                  |
|                                        |
\ -- Frank Crane                         /
 ----------------------------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
</pre>

「どんな黄金律も行動に移さなかったら何の意味もないで」って感じかな。

<br />
## Docker Hub を使う

[Docker Hub](https://hub.docker.com) に m0t0k1ch1 というアカウントを登録済み。

先ほど build した `docker-whale` に `m0t0k1ch1/docker-whale` という tag をつける。バージョンはとりあえず latest。

``` sh
$ docker tag d649997c162a m0t0k1ch1/docker-whale:latest
```

確認する。

``` sh
$ docker images
```

<pre>
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
docker-whale                latest              d649997c162a        54 minutes ago      274.9 MB
m0t0k1ch1/docker-whale      latest              d649997c162a        54 minutes ago      274.9 MB
nginx                       latest              4efb2fcdb1ab        3 days ago          183.4 MB
hello-world                 latest              c54a2cc56cbb        8 weeks ago         1.848 kB
docker/whalesay             latest              6b362a9f73eb        15 months ago       247 MB
</pre>

続いて Docker Hub にログイン。

``` sh
$ docker login --username=m0t0k1ch1
```

で、先ほど tag をつけた image を Docker Hub に push する。

``` sh
$ docker push m0t0k1ch1/docker-whale
```

Docker Hub の Repositories に `m0t0k1ch1/docker-whale` が存在することが確認できたら、一旦、ローカルにある image を削除する。

``` sh
$ docker rmi -f d649997c162a
```

消えてるか確認する。

``` sh
$ docker images
```

<pre>
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
nginx               latest              4efb2fcdb1ab        3 days ago          183.4 MB
hello-world         latest              c54a2cc56cbb        8 weeks ago         1.848 kB
docker/whalesay     latest              6b362a9f73eb        15 months ago       247 MB
</pre>

消えてる。

Docker Hub にログインした状態で、ローカルに image がなくても run できるか確認する。

``` sh
$ docker run m0t0k1ch1/docker-whale
```

<pre>
Unable to find image 'm0t0k1ch1/docker-whale:latest' locally
latest: Pulling from m0t0k1ch1/docker-whale
e190868d63f8: Already exists
909cd34c6fd7: Already exists
0b9bfabab7c1: Already exists
a3ed95caeb02: Already exists
00bf65475aba: Already exists
c57b6bcc83e3: Already exists
8978f6879e2f: Already exists
8eed3712d2cf: Already exists
31259c62dd40: Already exists
Digest: sha256:91f6c308fe72436f45a848df389f0c32a45385bbf7bac003bc64276e38adc3b7
Status: Downloaded newer image for m0t0k1ch1/docker-whale:latest
 _____________________________________
/ <WildTHing> ok guys .. so whens the \
| next commit :PP <taniwha> when they |
\ come to get me                      /
 -------------------------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
</pre>

無事 run できた。おしまい。
