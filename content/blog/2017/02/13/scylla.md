+++
date = "2017-02-13T02:35:40+09:00"
tags = [ "scylla", "azure" ]
title = "WORLD FASTEST NoSQL DATABASE な Scylla のクラスタを組んでみる on Azure"
+++

最近、[Scylla](http://www.scylladb.com) という C++ で書かれた Cassandra コンパチな NoSQL データベースを触っている。マスコットのおばけがかわいい。gopher くんと遊ばせたい。という話はさておき、Scylla 公式サイトのトップには「WORLD FASTEST NoSQL DATABASE」という一文。この星の一等賞とのこと。なんだけど、日本語の情報はまだ全然落ちてないので、今回はとりあえず Azure でクラスタを組んで、動作確認くらいのつもりで cassandra-stress を使って軽くベンチマーキングするところまでやってみましたメモを残してみようと思う。日本でも知見がもっと増えるといいな。

<!--more-->

<br />
## Scylla についてもう少し

公式サイトを読めば書いてあることだけど、気になった文言をいくつか。

> Fully compatible with Apache Cassandra at 10x the throughput and jaw dropping low latency

Cassandra コンパチでスループットは 10 倍、さらに jaw dropping なくらいレイテンシが低いとのこと。jaw dropping。

> Apache Cassandra compatible column store, with the low latency of Redis

レイテンシの低さは Redis 級とのこと。

<br />
## サクッと動かすだけ動かしてみる

カジュアルに動かしたい場合は、Docker で試してみるのが簡単そう。自分は [Docker Hub の公式 repo](https://hub.docker.com/r/scylladb/scylla) の通りにやったら特に問題なく動いた。今回の目的は Azure でクラスタを組んでみることなので、詳しくは書かない。

<br />
## Azure でクラスタを組む

<br />
### VM の準備

今回は以下のような VM 3 台でクラスタを組むことにした。最近は Ubuntu しか触ってないので Ubuntu で。

- OS：Ubuntu 16.04 LTS
- サイズ：Standard DS3 v2（4 コア・14 GB メモリ）
- OS ディスク：30 GB Premium_LRS（Managed Disk）
- データディスク：1,023 GB Premium_LRS（Managed Disk）

Scylla が使用するポートに関しては [公式ドキュメント](http://www.scylladb.com/admin) に記載されているので、それを参考にして適切にネットワークセキュリティグループを設定してあげる。

ちなみに、VM のディスクについては、つい先日リリースされた [Managed Disk](https://azure.microsoft.com/en-us/blog/announcing-general-availability-of-managed-disks-and-larger-scale-sets) を利用している。便利。

<br />
### Scylla をインストールする

基本的には [公式ドキュメント](http://docs.scylladb.com/getting-started/ubuntu-16-04) を参考に進めていけばよいが、[Scylla production recommendation](https://github.com/scylladb/scylla/wiki/Scylla-production-recommendation) に以下の記載があることに注意。

> Out of all of the above, setting the file system to XFS is the most important and mandatory for production. Scylla will significantly slow down without it.

なので、まずはデータディスクの設定から。今回は、Azure ポータルから 1,023 GB の Premium_LRS なデータディスクを追加しているので、それをまず確認する。

``` sh
$ lsblk
```

``` txt
NAME   MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
fd0      2:0    1    4K  0 disk
sda      8:0    0 29.3G  0 disk
└─sda1   8:1    0 29.3G  0 part /
sdb      8:16   0   28G  0 disk
└─sdb1   8:17   0   28G  0 part /mnt
sdc      8:32   0 1023G  0 disk
sr0     11:0    1  1.1M  0 rom
```

sdc が追加したデータディスクなので、パーティションを切ってあげる。

``` sh
$ fdisk /dev/sdc
```

``` txt
Welcome to fdisk (util-linux 2.27.1).
Changes will remain in memory only, until you decide to write them.
Be careful before using the write command.

Device does not contain a recognized partition table.
Created a new DOS disklabel with disk identifier 0x4307e611.

Command (m for help): n
Partition type
   p   primary (0 primary, 0 extended, 4 free)
   e   extended (container for logical partitions)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-2145386495, default 2048):
Last sector, +sectors or +size{K,M,G,T,P} (2048-2145386495, default 2145386495):

Created a new partition 1 of type 'Linux' and of size 1023 GiB.

Command (m for help): p
Disk /dev/sdc: 1023 GiB, 1098437885952 bytes, 2145386496 sectors
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 4096 bytes
I/O size (minimum/optimal): 4096 bytes / 4096 bytes
Disklabel type: dos
Disk identifier: 0x4307e611

Device     Boot Start        End    Sectors  Size Id Type
/dev/sdc1        2048 2145386495 2145384448 1023G 83 Linux

Command (m for help): w
The partition table has been altered.
Calling ioctl() to re-read partition table.
Syncing disks.
```

パーティションが切れてることを確認する。

``` sh
$ lsblk
```

``` txt
NAME   MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
fd0      2:0    1    4K  0 disk
sda      8:0    0 29.3G  0 disk
└─sda1   8:1    0 29.3G  0 part /
sdb      8:16   0   28G  0 disk
└─sdb1   8:17   0   28G  0 part /mnt
sdc      8:32   0 1023G  0 disk
└─sdc1   8:33   0 1023G  0 part
sr0     11:0    1  1.1M  0 rom
```

XFS でフォーマットする。

``` sh
$ mkfs.xfs /dev/sdc1
```

``` txt
meta-data=/dev/sdc1              isize=512    agcount=4, agsize=67043264 blks
         =                       sectsz=4096  attr=2, projid32bit=1
         =                       crc=1        finobt=1, sparse=0
data     =                       bsize=4096   blocks=268173056, imaxpct=25
         =                       sunit=0      swidth=0 blks
naming   =version 2              bsize=4096   ascii-ci=0 ftype=1
log      =internal log           bsize=4096   blocks=130943, version=2
         =                       sectsz=4096  sunit=1 blks, lazy-count=1
realtime =none                   extsz=4096   blocks=0, rtextents=0
```

フォーマットを確認する。

``` sh
$ file -s /dev/sdc1
```

``` txt
/dev/sdc1: SGI XFS filesystem data (blksz 4096, inosz 512, v2 dirs)
```

`/var/lib/scylla` にマウントする。

``` sh
$ mkdir /var/lib/scylla
$ mount /dev/sdc1 /var/lib/scylla
```

 マウントされているか確認する。

``` sh
$ lsblk
```

``` txt
NAME   MAJ:MIN RM  SIZE RO TYPE MOUNTPOINT
fd0      2:0    1    4K  0 disk
sda      8:0    0 29.3G  0 disk
└─sda1   8:1    0 29.3G  0 part /
sdb      8:16   0   28G  0 disk
└─sdb1   8:17   0   28G  0 part /mnt
sdc      8:32   0 1023G  0 disk
└─sdc1   8:33   0 1023G  0 part /var/lib/scylla
sr0     11:0    1  1.1M  0 rom
```

ディスクの準備はこれで ok。再起動時にもディスクがマウントされるよう、`/etc/fstab` の設定もしておいた。

ここからは公式ドキュメント通りにインストールを進める。

``` sh
$ wget -O /etc/apt/sources.list.d/scylla.list http://downloads.scylladb.com/deb/ubuntu/scylla-1.6-xenial.list
$ apt-get update
$ apt-get install scylla
```

今回インストールされたバージョンは以下。

``` sh
$ scylla --version
```

``` txt
1.6.0-20170202.7e1b245
```

<br />
### Scylla をセットアップをする

まず、`/etc/scylla/scylla.yaml` の調整。これも基本的には公式ドキュメント参照。今回は以下の項目について設定を変更した。

- seed_provider.parameters.seeds：最初に Scylla を起動する VM のプライベート IP を指定
- listen_address：それぞれの VM のプライベート IP を指定
- rpc_address：それぞれの VM のプライベート IP を指定
- endpoint_snitch：GossipingPropertyFileSnitch

GossipingPropertyFileSnitch にしたので、`/etc/scylla/cassandra-rackdc.properties` にも設定を追記。

``` txt
dc=dc1
rack=rack1
```

次にセットアップ用のスクリプトを走らせる。

``` sh
$ scylla_setup
```

どんどん質問されるので、yes / no で適切に答えていく。

セットアップ用のスクリプトの中で、ベンチマークをかけて io 関連の設定を自動でやってくれるステップがあるので、これは 3 台ともやっておくとよいかなと思う。3 台それぞれで設定された数字を参考に、最終的には `/etc/scylla.d/io.conf` の内容を以下で揃えた。

``` txt
SEASTAR_IO="--max-io-requests=20"
```

<br />
### Scylla を起動してクラスタを組む

`seed_provider.parameters.seeds` に設定した VM から順番に起動していく。

``` sh
$ systemctl start scylla-server
```

3 台とも起動できたら、正常にクラスタが組めているか確認する。

``` sh
$ nodetool status
```

``` txt
Datacenter: dc1
===============
Status=Up/Down
|/ State=Normal/Leaving/Joining/Moving
--  Address    Load       Tokens  Owns (effective)  Host ID                               Rack
UN  10.1.1.23  276.56 KB  256     69.3%             7a208e6d-96cb-4b5e-84bb-89b64106702d  rack1
UN  10.1.1.22  311.97 KB  256     63.8%             212c7ddc-4b63-4ace-8dc4-89eda382d5c3  rack1
UN  10.1.1.24  187.08 KB  256     67.0%             0cd1bd68-8f00-4abd-878f-be3f8b4b96cb  rack1
```

無事クラスタが組めた模様。

<br />
## cassandra-stress をかけてみる

cassandra-stress を実行する VM は以下。

- OS：Ubuntu 16.04 LTS
- サイズ：Standard F8s（8 コア・16 GB メモリ）

scylla-tools をインストールしておく。

``` sh
$ wget -O /etc/apt/sources.list.d/scylla.list http://downloads.scylladb.com/deb/ubuntu/scylla-1.6-xenial.list
$ apt-get update
$ apt-get install scylla-tools
```

<br />
### 手順

まずは以下を実行して keyspace を作成する。

``` sh
$ cassandra-stress write n=1 cl=ALL -schema 'replication(strategy=NetworkTopologyStrategy,dc1=3)' -mode native cql3 -node 10.1.1.22,10.1.1.23,10.1.1.24
```

次に、以下を実行して QUORUM で write 負荷を与える。

``` sh
$ cassandra-stress write cl=QUORUM n=1000000 -mode native cql3 -node 10.1.1.22,10.1.1.23,10.1.1.24
```

続けて以下を実行して QUORUM で read 負荷を与える。

``` sh
$ cassandra-stress read cl=QUORUM n=1000000 -mode native cql3 -rate threads=512 -node 10.1.1.22,10.1.1.23,10.1.1.24
```

再測定する場合は keyspace を drop してから最初に戻る。

<br />
### 結果

op rate と latency mean だけ抜き出して表にしてみる。

op rate (/sec)|write|read
:---:|:---:|:---:
1 回目|63,490|97,718
2 回目|64,625|99,711
3 回目|65,479|102,407

latency mean (msec)|write|read
:---:|:---:|:---:
1 回目|3.1|5.2
2 回目|3.1|5.1
3 回目|3.0|4.9

Scylla 公式サイトのトップのグラフでは、3-node cluster で write も read も 2,000 k transaction/sec くらいいってるので、今回の条件ではまだまだ Scylla の力を引き出せていない模様。ベンチマークの条件は公式サイトにきっちり載ってるので、VM のスペックを上げたり、複数のデータディスクで RAID0 を組んだりなどして条件を近づけて、同じような数字が出るのかは今後確かめていきたい。

チューニングに関しては、JVM 関連のパラメータも調整しないといけない Cassandra と比較すると、考えることが少ないのでラクだなと思う一方、細かい調整ができないデメリットもありそうかなと勝手に思っている。Cassandra コンパチとはいえ、一部のパラメータは `Not currently supported, reserved for future use` となっていたりもする。まあ、細かいこと考えないでも安定して高パフォーマンスが出てくれるならありがたい話ではある。きちんとアーキテクチャレベルからその性質を理解して、活用していきたい。大規模なクラスタ構成、かつ運用まで考えると、今見えてないことがもっと見えてくるとは思う。

最近は色々あって Azure にお世話になっているので、とりあえずは Azure 上でどこまでパフォーマンス引き出せるのかはやってみようと思う。
