+++
date = "2017-11-18T21:00:45+09:00"
tags = [ "zabbix", "cassandra" ]
title = "Cassandra 監視用の Zabbix template"
+++

主要項目に絞ってつくったつもり。ご自由にお使いください。

<!--more-->

<br />
## 成果物

<div class="github-card" data-user="m0t0k1ch1" data-repo="zabbix-cassandra-template"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

細かい解説はしないが、用意したグラフ名を以下に列挙。名前からお察しください。

- Cassandra: compaction tasks
- Cassandra: GC count
- Cassandra: heap memory usage
- Cassandra: memory pool usage
- Cassandra: non-heap memory usage
- Cassandra: number of clients
- Cassandra: number of nodes
- Cassandra: number of unsuccessful requests
- Cassandra: read latency
- Cassandra: throughput
- Cassandra: write latency

<br />
## 参考にしたもの・利用したもの

JMX 経由で取得できる項目を [公式サイト](http://cassandra.apache.org/doc/latest/operating/metrics.html) で調べて、Zabbix に登録する前に [Command-line JMX Client](http://crawler.archive.org/cmdline-jmxclient) を利用して本当に値が取得できるか確かめた。

シュッと値が欲しいときに Command-line JMX Client は便利。

<br />
## ややこしかったこと

TotalLatency と Latency。

``` sh
$ java -jar cmdline-jmxclient-0.10.3.jar - 127.0.0.1:7199 org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=TotalLatency Count
```

``` txt
11/18/2017 11:10:53 +0000 org.archive.jmx.Client Count: 25386616763
```

^ これは書き込みレイテンシの合計。単位はマイクロ秒。

``` sh
$ java -jar cmdline-jmxclient-0.10.3.jar - 127.0.0.1:7199 org.apache.cassandra.metrics:type=ClientRequest,scope=Write,name=Latency Count
```

``` txt
11/18/2017 11:17:56 +0000 org.archive.jmx.Client Count: 105136578
```

^ これは捌いた書き込みリクエスト数の合計。スループットが欲しいときはこいつから計算する。

これはどっちがどっちかわかんなくなる人が多いのでは。。？上記は Write に関してだが、当然 Read でも同様。もう少し直感的にわかりやすい命名にできなかったのだろうか。

<br />
## やり残し

Table Metrics を discovery 使っていい感じに取得したい。
