+++
date = "2018-01-24T04:30:44+09:00"
tags = [ "bitcoin", "lightning network", "blockchain" ]
title = "lnd ＋ btcd を使って testnet で Lightning Network 体験"
+++

Bitcoin に限らず、今年はオフチェーンなテクノロジーが多々実用段階に突入しそうですねということで、把握必須であろう Lightning Network について手を動かしながら勉強していこうかと。Lightning Network ノードの実装はいくつかありますが、今回は [Lightning Labs](https://lightning.engineering) の [lnd](https://github.com/lightningnetwork/lnd) をチョイスしました。

<!--more-->

基本的には、

- https://github.com/lightningnetwork/lnd/blob/master/docs/INSTALL.md
- http://dev.lightning.community/tutorial/01-lncli/index.html

に従って進め、

__testnet にて、Lightning Network を利用した single hop payment と multi hop payment を成功させること__

をゴールとします。

双方向ペイメントチャネルや Lightning Network の原理に触れている記事などはたくさんありますので、仕組みの解説はそちらに任せ、今回は手を動かすことを主軸に据えようと思います。

なお、今回の検証は Ubuntu 16.04 上で行うこととします。

<br />
## インストール

<br />
### Golang のインストール

1.8 以上が必要とのことなので、今回は 1.9 をインストールします。

``` sh
$ sudo apt-get install golang-1.9-go
```

`~/.bashrc` などで、よしなに `GOPATH` と `PATH` を通しておきます。

``` sh
export GOPATH=~/go
export PATH=$PATH:/usr/lib/go-1.9/bin:$GOPATH/bin
```

バージョンを確認します。

``` sh
$ go version
```

``` txt
go version go1.9.2 linux/amd64
```

<br />
### Glide のインストール

Golang 用のパッケージ管理ツールである [Glide](https://glide.sh) をインストールします。lnd はこれを利用してパッケージ管理されています。

``` sh
$ go get -u github.com/Masterminds/glide
```

バージョンを確認します。

``` sh
$ glide --version
```

``` txt
glide version 0.13.2-dev
```

<br />
### lnd のインストール

lnd をインストールします。

``` sh
$ git clone https://github.com/lightningnetwork/lnd $GOPATH/src/github.com/lightningnetwork/lnd
...
$ cd $GOPATH/src/github.com/lightningnetwork/lnd
$ glide install
...
$ go install . ./cmd/...
```

バージョンを確認します。

``` sh
$ lnd --version
```

``` txt
lnd version 0.3.0-alpha
```

<br />
### btcd のインストール

今回は lnd のバックエンドとして [btcd](https://github.com/roasbeef/btcd) を利用します。

> lnd currently requires btcd with segwit support, which is not yet merged into the master branch. Instead, roasbeef maintains a fork with his segwit implementation applied.

とのことなので、Roasbeef さんが fork してメンテしているバージョンをインストールします。

``` sh
$ git clone https://github.com/roasbeef/btcd $GOPATH/src/github.com/roasbeef/btcd
...
$ cd $GOPATH/src/github.com/roasbeef/btcd
$ glide install
...
$ go install . ./cmd/...
```

バージョンを確認します。

``` sh
$ btcd --version
```

``` txt
btcd version 0.12.0-beta
```

<br />
## 起動

<br />
### btcd の起動

今回は検証目的なので、起動して雑にバックグランドに回しておきます。

``` sh
$ nohup btcd --testnet --txindex --rpcuser=btcdrpc --rpcpass=btcdrpc &
```

testnet とのデータの同期にはそれなりに時間がかかるので、完了するまでしばし放置します。

<br />
### lnd の起動

冒頭に記載した通り、multi hop payment も試したいので、lnd を 3 匹起動していきます。

まず、共通のオプションを `~/.lnd/lnd.conf` に設定しておきます。なお、検証目的なので、macaroons による認証は無効化しています。

``` txt
[Application Options]
datadir=data
logdir=log
debuglevel=info
debughtlc=true
no-macaroons=true

[Bitcoin]
bitcoin.active=1
bitcoin.testnet=1
bitcoin.rpcuser=btcdrpc
bitcoin.rpcpass=btcdrpc
```

3 匹それぞれのデータを保持するディレクトリを作成しておきます。Alice、Bob、Charlie の 3 人が lnd を起動する想定です。

``` sh
$ cd ~/.lnd
$ mkdir alice bob charlie
```

lnd を順番に起動していきます。ここも雑にバックグラウンドに回しておきます。

``` sh
$ cd ~/.lnd/alice
$ nohup lnd --rpcport=10001 --peerport=10011 --restport=8001 &
$ cd ~/.lnd/bob
$ nohup lnd --rpcport=10002 --peerport=10012 --restport=8002 &
$ cd ~/.lnd/charlie
$ nohup lnd --rpcport=10003 --peerport=10013 --restport=8003 &
```

<br />
## セットアップ

<br />
### ウォレットの生成

以下を実行してウォレットを生成します。パスワード（8 文字以上）の入力を求められますので、よしなに入力してください。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons create
...
$ lncli --rpcserver=localhost:10002 --no-macaroons create
...
$ lncli --rpcserver=localhost:10003 --no-macaroons create
...
```

成功するとバックエンド（btcd）との同期が始まり、各 `nohup.out` に以下のようなログが出力されるはずです。

``` txt
2018-01-15 23:46:45.582 [INF] LTND: Primary chain is set to: bitcoin
2018-01-15 23:46:45.582 [INF] LTND: Initializing btcd backed fee estimator
2018-01-15 23:46:51.336 [INF] LNWL: Opened wallet
2018-01-15 23:46:53.466 [INF] LNWL: The wallet has been unlocked without a time limit
2018-01-15 23:46:53.467 [INF] LNWL: Catching up block hashes to height 1258878, this will take a while...
2018-01-15 23:46:53.481 [INF] LTND: LightningWallet opened
2018-01-15 23:46:53.505 [INF] RPCS: RPC server listening on 127.0.0.1:10003
2018-01-15 23:46:53.505 [INF] RPCS: gRPC proxy started at localhost:8003
2018-01-15 23:46:53.513 [INF] LTND: Waiting for chain backend to finish sync, start_height=1258878
2018-01-15 23:46:58.266 [INF] LNWL: Caught up to height 10000
2018-01-15 23:47:03.309 [INF] LNWL: Caught up to height 20000
2018-01-15 23:47:08.280 [INF] LNWL: Caught up to height 30000

（中略）

2018-01-15 23:56:21.774 [INF] LNWL: Caught up to height 1230000
2018-01-15 23:56:24.160 [INF] LNWL: Caught up to height 1240000
2018-01-15 23:56:26.603 [INF] LNWL: Caught up to height 1250000
2018-01-15 23:56:28.907 [INF] LNWL: Done catching up block hashes
2018-01-15 23:56:28.953 [INF] LNWL: Started rescan from block 000000004d3c58ac87c987ceee34c350db3fb9cd25969877f4c56b9be022b2de (height 1258878) for 1 address
2018-01-15 23:56:28.956 [INF] LNWL: Catching up block hashes to height 1258880, this might take a while
2018-01-15 23:56:29.007 [INF] LNWL: Done catching up block hashes
2018-01-15 23:56:29.007 [INF] LNWL: Finished rescan for 1 address (synced to block 00000000000006bb9c24873200926c0a6446bad2b82887f4d05cfcdec8aa220e, height 1258880)
2018-01-15 23:56:29.236 [INF] LTND: Chain backend is fully synced (end_height=1258880)!
2018-01-15 23:56:29.282 [INF] HSWC: Starting HTLC Switch
2018-01-15 23:56:29.282 [INF] NTFN: New block epoch subscription
2018-01-15 23:56:29.283 [INF] DISC: Authenticated Gossiper is starting
2018-01-15 23:56:29.283 [INF] NTFN: New block epoch subscription
2018-01-15 23:56:29.283 [INF] BRAR: Starting contract observer with 0 active channels
2018-01-15 23:56:29.283 [INF] CRTR: FilteredChainView starting
2018-01-15 23:56:29.283 [ERR] DISC: unable to rebroadcast stale channels: error while retrieving outgoing channels: no graph edges exist
2018-01-15 23:56:29.750 [INF] CRTR: Filtering chain using 0 channels active
2018-01-15 23:56:29.751 [INF] CRTR: Prune tip for Channel Graph: height=1258880, hash=00000000000006bb9c24873200926c0a6446bad2b82887f4d05cfcdec8aa220e
2018-01-15 23:56:29.752 [INF] SRVR: Initializing peer network boostrappers!
2018-01-15 23:56:29.752 [INF] SRVR: Creating DNS peer boostrapper with seeds: [[nodes.lightning.directory soa.nodes.lightning.directory]]
2018-01-15 23:56:29.752 [INF] CMGR: Server listening on [::]:10013
2018-01-15 23:56:29.752 [INF] DISC: Attempting to bootstrap with: Authenticated Channel Graph
2018-01-15 23:56:29.754 [INF] DISC: Attempting to bootstrap with: BOLT-0010 DNS Seed: [[nodes.lightning.directory soa.nodes.lightning.directory]]
```

また、`getinfo` コマンドで基本的な情報を取得することもできます。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons getinfo
```

``` json
{
    "identity_pubkey": "03bc52cd50e93c3beb07aea040089c1eec7a1c702628cbee41c63a5101362ba8bf",
    "alias": "",
    "num_pending_channels": 0,
    "num_active_channels": 0,
    "num_peers": 0,
    "block_height": 1258880,
    "block_hash": "00000000000006bb9c24873200926c0a6446bad2b82887f4d05cfcdec8aa220e",
    "synced_to_chain": true,
    "testnet": true,
    "chains": [
        "bitcoin"
    ],
    "uris": [
    ]
}
```

<br />
### アドレスの生成

Alice、Bob、Charlie のオンチェーンアドレスを生成します。ここで指定している np2wkh というのは nested-pay-to-witness-key-hash の略で、[P2SH でネストされた P2WPKH](https://github.com/bitcoin/bips/blob/master/bip-0141.mediawiki#p2wpkh-nested-in-bip16-p2sh) のことです（ちょっと自信ないけど、たぶん合ってるはず）。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons newaddress np2wkh
```

``` json
{
    "address": "2MttnqvpvQyipkkbRBbjSu38W33qtMQ86yS"
}
```

``` sh
$ lncli --rpcserver=localhost:10002 --no-macaroons newaddress np2wkh
```

``` json
{
    "address": "2N5nawrANXwKo7aXWM3HiFkwZupUmT7RKCj"
}
```

``` sh
$ lncli --rpcserver=localhost:10003 --no-macaroons newaddress np2wkh
```

``` json
{
    "address": "2MviVP7BPcyYn2qA9dXWgvgbhadYApV9zCd"
}
```

<br />
### コインの付与

コインがなくては検証できないので、今回は Alice と Charlie にコインを付与します。適当な testnet の faucet をひねってきましょう。その後、残高を確認します。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons walletbalance
```

``` json
{
    "total_balance": "130000000",
    "confirmed_balance": "130000000",
    "unconfirmed_balance": "0"
}
```

``` sh
$ lncli --rpcserver=localhost:10003 --no-macaroons walletbalance
```

``` json
{
    "total_balance": "65000000",
    "confirmed_balance": "65000000",
    "unconfirmed_balance": "0"
}
```

<br />
### P2P ネットワークの構築

Alice、Bob、Charlie で P2P ネットワークを構築します。

まず、接続先となる Bob の公開鍵（`identity_pubkey`）を確認します。

``` sh
$ lncli --rpcserver=localhost:10002 --no-macaroons getinfo
```

``` json
{
    "identity_pubkey": "0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74",
    "alias": "",
    "num_pending_channels": 0,
    "num_active_channels": 0,
    "num_peers": 1,
    "block_height": 1259892,
    "block_hash": "00000000000003925458529edd7ebf929e37a4aa4630ac48d5b7a7d2066353d4",
    "synced_to_chain": true,
    "testnet": true,
    "chains": [
        "bitcoin"
    ],
    "uris": [
    ]
}
```

確認した Bob の公開鍵を指定して、Alice から Bob に接続します。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons connect 0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74@localhost:10012
```

同様に、Charlie から Bob に接続します。

``` sh
$ lncli --rpcserver=localhost:10003 --no-macaroons connect 0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74@localhost:10012
```

正常に接続されているか確認します。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons listpeers
```

``` json
{
    "peers": [
        {
            "pub_key": "0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74",
            "peer_id": 2,
            "address": "127.0.0.1:10012",
            "bytes_sent": "2512295",
            "bytes_recv": "2536047",
            "sat_sent": "0",
            "sat_recv": "0",
            "inbound": true,
            "ping_time": "0"
        }
    ]
}
```

``` sh
$ lncli --rpcserver=localhost:10002 --no-macaroons listpeers
```

``` json
{
    "peers": [
        {
            "pub_key": "03bc52cd50e93c3beb07aea040089c1eec7a1c702628cbee41c63a5101362ba8bf",
            "peer_id": 4,
            "address": "127.0.0.1:33292",
            "bytes_sent": "2536047",
            "bytes_recv": "2512295",
            "sat_sent": "0",
            "sat_recv": "0",
            "inbound": false,
            "ping_time": "0"
        },
        {
            "pub_key": "0288bedd304fccc8435db8073f1236f820162962a44fb3ef0068eccf287745c69b",
            "peer_id": 5,
            "address": "127.0.0.1:33308",
            "bytes_sent": "195235",
            "bytes_recv": "1080941",
            "sat_sent": "0",
            "sat_recv": "0",
            "inbound": false,
            "ping_time": "0"
        }
    ]
}
```

``` sh
$ lncli --rpcserver=localhost:10003 --no-macaroons listpeers
```

``` json
{
    "peers": [
        {
            "pub_key": "0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74",
            "peer_id": 2,
            "address": "127.0.0.1:10012",
            "bytes_sent": "2512295",
            "bytes_recv": "791559",
            "sat_sent": "0",
            "sat_recv": "0",
            "inbound": true,
            "ping_time": "0"
        }
    ]
}
```

<br />
## single hop payment

Alice から Bob への single hop payment を行ってみようと思います。

Alice と Bob の間にペイメントチャネルを開きます。今回は、Alice の 1,000,000 satoshi を利用することにします。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons openchannel --node_key=0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74 --local_amt=1000000
```

``` json
{
        "funding_txid": "8633a3c8e77e047c8c81c07b8777f6fd6eb0ff9fad2b44f37d798c9af12468c1"
}
```

[ペイメントチャネルを開くためのトランザクション](https://testnet.smartbit.com.au/tx/8633a3c8e77e047c8c81c07b8777f6fd6eb0ff9fad2b44f37d798c9af12468c1) がブロードキャストされました。このトランザクションが承認（デフォルトで 3 confirmation）されないとペイメントチャネルは開かないので、承認されるまで待ちます。

チャネルが開いたら確認します。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons listchannels
```

``` json
{
    "channels": [
        {
            "active": true,
            "remote_pubkey": "0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74",
            "channel_point": "8633a3c8e77e047c8c81c07b8777f6fd6eb0ff9fad2b44f37d798c9af12468c1:0",
            "chan_id": "1385268102766264320",
            "capacity": "1000000",
            "local_balance": "959456",
            "remote_balance": "0",
            "commit_fee": "40544",
            "commit_weight": "600",
            "fee_per_kw": "56000",
            "unsettled_balance": "0",
            "total_satoshis_sent": "0",
            "total_satoshis_received": "0",
            "num_updates": "0",
            "pending_htlcs": [
            ],
            "csv_delay": 144
        }
    ]
}
```

``` sh
$ lncli --rpcserver=localhost:10002 --no-macaroons listchannels
```

``` json
{
    "channels": [
        {
            "active": true,
            "remote_pubkey": "03bc52cd50e93c3beb07aea040089c1eec7a1c702628cbee41c63a5101362ba8bf",
            "channel_point": "8633a3c8e77e047c8c81c07b8777f6fd6eb0ff9fad2b44f37d798c9af12468c1:0",
            "chan_id": "1385268102766264320",
            "capacity": "1000000",
            "local_balance": "0",
            "remote_balance": "959456",
            "commit_fee": "40544",
            "commit_weight": "552",
            "fee_per_kw": "56000",
            "unsettled_balance": "0",
            "total_satoshis_sent": "0",
            "total_satoshis_received": "0",
            "num_updates": "0",
            "pending_htlcs": [
            ],
            "csv_delay": 144
        }
    ]
}
```

これを見ると、残高の内訳が以下のようになっていることが分かります。

- Alice：959,456 satoshi
- Bob：0 satoshi
- commit fee：40,544 satoshi

なお、40,544 satoshi はコミットメントトランザクションをブロードキャストするための手数料（commit fee）として確保されている（にしては多いような。。？）ようですので、Alice が準備した 1,000,000 satoshi から 40,544 satoshi を差し引いた値が残高の合計となっています。

無事ペイメントチャネルが開けましたので、Alice から Bob への single hop payment を行う準備が整いました。ということで、Alice から Bob に 10,000 satoshi の支払いを行ってみます。

まず、Bob が 10,000 satoshi 分の invoice を生成します。

``` sh
$ lncli --rpcserver=localhost:10002 --no-macaroons addinvoice --value=10000
```

``` json
{
        "r_hash": "354026c9a36be070f57371bbccdab636c90ee017bfd6a840fb3b49a908574c88",
        "pay_req": "lntb100u1pdxvf6hpp5x4qzdjdrd0s8patnwxauek4kxmysacqhhlt2ss8m8dy6jzzhfjyqdqqcqzysmj6fewu9fzyf6e3y793wkqq6s0dfsf7y4h5yxj4s8q30sk66yxs44qsh5q9lah8wlwugf9f9xgsged23sq4sq6ug0qjtg9trn4qe8fqqarz7dv"
}
```

生成された invoice に対して、Alice が支払いを行います。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons sendpayment --pay_req=lntb100u1pdxvf6hpp5x4qzdjdrd0s8patnwxauek4kxmysacqhhlt2ss8m8dy6jzzhfjyqdqqcqzysmj6fewu9fzyf6e3y793wkqq6s0dfsf7y4h5yxj4s8q30sk66yxs44qsh5q9lah8wlwugf9f9xgsged23sq4sq6ug0qjtg9trn4qe8fqqarz7dv
```

``` json
{
        "payment_error": "",
        "payment_preimage": "84880ffb3614a8d5939cb72406f3462cbd5fce9153caf20d29ac035bcaa15dd9",
        "payment_route": {
                "total_time_lock": 1260054,
                "total_amt": 10000,
                "hops": [
                        {
                                "chan_id": 1385268102766264320,
                                "chan_capacity": 1000000,
                                "amt_to_forward": 10000,
                                "expiry": 1260054
                        }
                ]
        }
}
```

再度ペイメントチャネルの状態を確認します。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons listchannels
```

``` json
{
    "channels": [
        {
            "active": true,
            "remote_pubkey": "0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74",
            "channel_point": "8633a3c8e77e047c8c81c07b8777f6fd6eb0ff9fad2b44f37d798c9af12468c1:0",
            "chan_id": "1385268102766264320",
            "capacity": "1000000",
            "local_balance": "949456",
            "remote_balance": "10000",
            "commit_fee": "40544",
            "commit_weight": "724",
            "fee_per_kw": "56000",
            "unsettled_balance": "0",
            "total_satoshis_sent": "10000",
            "total_satoshis_received": "0",
            "num_updates": "2",
            "pending_htlcs": [
            ],
            "csv_delay": 144
        }
    ]
}
```

``` sh
$ lncli --rpcserver=localhost:10002 --no-macaroons listchannels
```

``` json
{
    "channels": [
        {
            "active": true,
            "remote_pubkey": "03bc52cd50e93c3beb07aea040089c1eec7a1c702628cbee41c63a5101362ba8bf",
            "channel_point": "8633a3c8e77e047c8c81c07b8777f6fd6eb0ff9fad2b44f37d798c9af12468c1:0",
            "chan_id": "1385268102766264320",
            "capacity": "1000000",
            "local_balance": "10000",
            "remote_balance": "949456",
            "commit_fee": "40544",
            "commit_weight": "724",
            "fee_per_kw": "56000",
            "unsettled_balance": "0",
            "total_satoshis_sent": "0",
            "total_satoshis_received": "10000",
            "num_updates": "2",
            "pending_htlcs": [
            ],
            "csv_delay": 144
        }
    ]
}
```

これを見ると、残高の内訳が以下のようになっていることが分かります。

- Alice：949,456 satoshi
- Bob：10,000 satoshi
- commit fee：40,544 satoshi

無事支払いが行えたようです。

本来はこのペイメントチャネルを利用して Alice と Bob の間で複数回の支払いを行った後にペイメントチャネルを閉じ、最終的な残高をオンチェーンで確定させるべきですが、今回はここでペイメントチャネルを閉じてみます。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons closechannel --funding_txid=8633a3c8e77e047c8c81c07b8777f6fd6eb0ff9fad2b44f37d798c9af12468c1 --output_index=0
```

``` json
{
        "closing_txid": "849b08005e31891af7dc33e845d7835b78d4750cb59864efe4896009bda6a3dd"
}
```

[ペイメントチャネルを閉じるためのトランザクション](https://testnet.smartbit.com.au/tx/849b08005e31891af7dc33e845d7835b78d4750cb59864efe4896009bda6a3dd) がブロードキャストされました。トランザクションが承認されたら、Alice と Bob のオンチェーンアドレスの残高を確認します。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons walletbalance
```

``` json
{
    "total_balance": "129981438",
    "confirmed_balance": "129981438",
    "unconfirmed_balance": "0"
}
```

``` sh
$ lncli --rpcserver=localhost:10002 --no-macaroons walletbalance
```

``` json
{
    "total_balance": "10000",
    "confirmed_balance": "10000",
    "unconfirmed_balance": "0"
}
```

Alice は `openchannel` と `closechannel` のトランザクション手数料

- openchannel：4,218 satoshi
- closechannel：4,344 satoshi

を支払っているので、残高が 129,981,438 satoshi となっています。念のため、間違いがないか計算してみます。

$$
130000000 - 10000 - (4218 + 4344) = 129981438
$$

残高と計算結果が一致したので、問題なさそうです。

<br />
## multi hop payment

Alice から Charlie へ、Bob を経由して multi hop payment を行ってみようと思います。

Alice と Bob の間、Bob と Charlie の間にそれぞれペイメントチャネルを開きます。なお、Charlie がペイメントチャネルを開く際には `--push_amt` を指定し、`--local_amt` で指定した 800,000 satoshi のうち 200,000 satoshi を Bob の初期残高として割り当てておきます。これは、Bob が Alice の支払いを中継できるようにするためです。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons openchannel --node_key=0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74 --local_amt=1000000
```

``` json
{
        "funding_txid": "50a5d5dba4b293f6b932f33a1c50c20cfee7d778c9e8848b523ba71ae9399579"
}
```

``` sh
$ lncli --rpcserver=localhost:10003 --no-macaroons openchannel --node_key=0398d53fe171c4cf492122819ad3c3bc8c25ac9b285840b7279005ba5371a59a74 --local_amt=800000 --push_amt=200000
```

``` json
{
        "funding_txid": "8afa8b9685cbc23169cca7211d4bb97c373b411629f62e180757cc78f253980f"
}
```

`listchannels` コマンドで各ペイメントチャネルの残高の内訳を確認してみると、以下のようになっています。

- Alice - Bob
  - Alice：959,456 satoshi
  - Bob：0 satoshi
  - commit fee：40,544 satoshi
- Bob - Charlie
  - Bob：200,000 satoshi
  - Charlie：559,456 satoshi
  - commit fee：40,544 satoshi

Charlie が 10,000 satoshi 分の invoice を生成します。

``` sh
$ lncli --rpcserver=localhost:10003 --no-macaroons addinvoice --value=10000
```

``` json
{
        "r_hash": "bfa72c4fdefcdcd5a98684a2ff1af41ce1dd0f847e2e7f1cbd432fb1b5a49046",
        "pay_req": "lntb100u1pdxdxg5pp5h7njcn77lnwdt2vxsj307xh5rnsa6ruy0ch8789agvhmrddyjprqdqqcqzysp8uekllf8ay6wmmj29kyex0w97hwj92gyu9nf4keuuex4phle5u8jrnqfan9rzgcxj3a3r0cnzdlm0eeutv28dhlu89py25gqt0gytgp4hsz53"
}
```

生成された invoice に対して、Alice が支払いを行います。このとき、Alice と Charlie はペイメントチャネルを開いていないため、Bob を経由して支払いが行われることになります。


``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons sendpayment --pay_req=lntb100u1pdxdxg5pp5h7njcn77lnwdt2vxsj307xh5rnsa6ruy0ch8789agvhmrddyjprqdqqcqzysp8uekllf8ay6wmmj29kyex0w97hwj92gyu9nf4keuuex4phle5u8jrnqfan9rzgcxj3a3r0cnzdlm0eeutv28dhlu89py25gqt0gytgp4hsz53
```

``` json
{
        "payment_error": "",
        "payment_preimage": "828f01bd8f91bc3da3984edf2698579a5c82cf268b2bc93827e36473efefe84b",
        "payment_route": {
                "total_time_lock": 1260357,
                "total_fees": 1,
                "total_amt": 10001,
                "hops": [
                        {
                                "chan_id": 1385292292024696832,
                                "chan_capacity": 1000000,
                                "amt_to_forward": 10000,
                                "fee": 1,
                                "expiry": 1260213
                        },
                        {
                                "chan_id": 1385292292024827904,
                                "chan_capacity": 800000,
                                "amt_to_forward": 10000,
                                "expiry": 1260213
                        }
                ]
        }
}
```

`listchannels` コマンドで各ペイメントチャネルの残高の内訳を確認してみると、以下のようになっています。

- Alice - Bob
  - Alice：949,454 satoshi
  - Bob：10,001 satoshi
  - commit fee：40,545 satoshi
- Bob - Charlie
  - Bob：190,000 satoshi
  - Charlie：569,456 satoshi
  - commit fee：40,544 satoshi

中継を担った Bob は、Alice から 10,000 satoshi と手数料 1 satoshi を受け取っていますが、10,000 satoshi は Charlie に送っていますので、その残高は

$$
200,000 + (10,000 + 1) - 10,000 = 200,001
$$

となります。中継を行ったことで 1 satoshi の手数料を獲得しています。

無事支払いが行えたようなので、今回はここでペイメントチャネルを閉じます。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons closechannel --funding_txid=50a5d5dba4b293f6b932f33a1c50c20cfee7d778c9e8848b523ba71ae9399579 --output_index=0
```

``` json
{
        "closing_txid": "7b2bae6f4aef7c3990df784d082fb0e6b5ac38dcdfd0597061b87fd304b8401d"
}
```

``` sh
$ lncli --rpcserver=localhost:10003 --no-macaroons closechannel --funding_txid=8afa8b9685cbc23169cca7211d4bb97c373b411629f62e180757cc78f253980f --output_index=0
```

``` json
{
        "closing_txid": "49b4f98ecbbfa7a01bfc739a4c4039749b56aeeec7a5d22cd93b6e82be81a8be"
}
```

ペイメントチャネルを閉じるためのトランザクションが承認されたら、最終的なオンチェーンアドレスの残高を確認し、間違いがないか計算してみます。Alice と Charlie は、`openchannel` と `closechannel` のトランザクション手数料を支払っていることに注意します。

まず、Alice について。

``` sh
$ lncli --rpcserver=localhost:10001 --no-macaroons walletbalance
```

``` json
{
    "total_balance": "129961787",
    "confirmed_balance": "129961787",
    "unconfirmed_balance": "0"
}
```

$$
129981438 - (10000 + 1) - (5305 + 4345) = 129961787
$$

次に、Bob について。

``` sh
$ lncli --rpcserver=localhost:10002 --no-macaroons walletbalance
```

``` json
{
    "total_balance": "210001",
    "confirmed_balance": "210001",
    "unconfirmed_balance": "0"
}
```

$$
10000 + 200000 + (10000 + 1) - 10000 = 210001
$$

最後に、Charlie について。

``` sh
$ lncli --rpcserver=localhost:10003 --no-macaroons walletbalance
```

``` json
{
    "total_balance": "64801438",
    "confirmed_balance": "64801438",
    "unconfirmed_balance": "0"
}
```

$$
65000000 - 200000 + 10000 - (4218 + 4344) = 64801438
$$

残高と計算結果は全て一致したので、問題なさそうです。

<br />
## まとめ

- [Lightning Labs](https://lightning.engineering) が実装した Lightning Network ノードである [lnd](https://github.com/lightningnetwork/lnd) を用いて、プライベートな Lightning Network を構築しました
  - バックエンドには [btcd](https://github.com/roasbeef/btcd) を採用し、testnet と接続しました
- 構築した Lightning Network を利用して、single hop payment と multi hop payment の検証を行いました


<br />
## ちなみに

今回は lnd のバックエンドとして btcd を用いましたが、これは切り替えることが可能なつくりとなっています。

> lnd has several pluggable back-end chain services including btcd (a full-node) and neutrino (a new experimental light client).

上記紹介文にある通り、[neutrino](https://github.com/lightninglabs/neutrino) という新しいライトクライアントもバックエンドとして利用可能なようです。まだ experimental なようで、mainnet では利用しない方がよさそうですが、

- ライトクライアントベースの Lightning Network のセキュリティ
- そもそも neutrino 自体の仕組み

などは気になるところですので、今後の調査テーマとなりそうです。
