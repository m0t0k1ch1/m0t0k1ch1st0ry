+++
title = 'Docker で private かつ standalone な Ethereum ノードを立てる'
tags = ['ethereum', 'blockchain', 'docker']
date = '2017-11-12'
+++

手元の環境で表題の通りのことがやりたいと思ったのだが、要望を満たすような最小限お手軽ツールが見当たらなかったので自分で書いてみた。

<!--more-->

{{< github "m0t0k1ch1" "ethereum-docker" >}}

README に書いてある通り、`docker-compose up -d` するだけで private かつ standalone な Ethereum ノードが立ち上がり、マイニングが始まる。なお、起動コマンドは以下。

```sh
$ geth --networkid 1234 --nodiscover --maxpeers 0 --datadir /root/.ethereum/privchain --mine --minerthreads 1 --etherbase 0x49d38ba99e0a1712388031345114d9da84110e9f --rpc --rpcaddr '0.0.0.0' --rpcport 8545 --rpccorsdomain '*' --rpcapi='admin,db,debug,eth,miner,net,personal,shh,txpool,web3' --unlock 0 --password /root/.ethereum/privchain/passwd
```

ベースイメージとして採用した [ethereum/go-client](https://hub.docker.com/r/ethereum/client-go) にて `ENTRYPOINT ["geth"]` となっているので、`docker-compose.yml` の command ではオプションだけを指定している。また、ノードには初期 etherbase 用のアカウントが 1 匹だけ登録してあり、これを起動時にアンロックしている（パスワードは空文字）。

正常にマイニングが開始された場合のログは以下のようになる。

```txt
INFO [11-11|19:43:48] Starting peer-to-peer node               instance=Geth/v1.7.2-stable/linux-amd64/go1.9.1
INFO [11-11|19:43:48] Allocated cache and file handles         database=/root/.ethereum/privchain/geth/chaindata cache=128 handles=1024
WARN [11-11|19:43:48] Upgrading database to use lookup entries
INFO [11-11|19:43:48] Initialised chain configuration          config="{ChainID: 1234 Homestead: 0 DAO: <nil> DAOSupport: false EIP150: <nil> EIP155: 0 EIP158: 0 Byzantium: <nil> Engine: unknown}"
INFO [11-11|19:43:48] Disk storage enabled for ethash caches   dir=/root/.ethereum/privchain/geth/ethash count=3
INFO [11-11|19:43:48] Disk storage enabled for ethash DAGs     dir=/root/.ethash                         count=2
INFO [11-11|19:43:48] Database deduplication successful        deduped=0
INFO [11-11|19:43:48] Initialising Ethereum protocol           versions="[63 62]" network=1234
INFO [11-11|19:43:48] Loaded most recent local header          number=0 hash=c29d61…3a3864 td=131072
INFO [11-11|19:43:48] Loaded most recent local full block      number=0 hash=c29d61…3a3864 td=131072
INFO [11-11|19:43:48] Loaded most recent local fast block      number=0 hash=c29d61…3a3864 td=131072
INFO [11-11|19:43:48] Regenerated local transaction journal    transactions=0 accounts=0
INFO [11-11|19:43:48] Starting P2P networking
INFO [11-11|19:43:48] RLPx listener up                         self="enode://d81d3f613dd1028919fe846e4fcb021cb193b5e7623205f9153e003a8c19adcd6ac0ed1b44689045058cb7a5bf1af484b55fd6a422839e751885efcbbb387ba3@[::]:30303?discport=0"
INFO [11-11|19:43:48] HTTP endpoint opened: http://0.0.0.0:8545
INFO [11-11|19:43:48] IPC endpoint opened: /root/.ethereum/privchain/geth.ipc
INFO [11-11|19:43:49] Unlocked account                         address=0x49d38BA99e0a1712388031345114d9DA84110E9f
INFO [11-11|19:43:49] Transaction pool price threshold updated price=18000000000
INFO [11-11|19:43:49] Starting mining operation
INFO [11-11|19:43:49] Commit new mining work                   number=1 txs=0 uncles=0 elapsed=447.294µs
INFO [11-11|19:43:57] Generating DAG in progress               epoch=0 percentage=0 elapsed=5.880s
INFO [11-11|19:44:00] Generating DAG in progress               epoch=0 percentage=1 elapsed=9.375s
INFO [11-11|19:44:06] Generating DAG in progress               epoch=0 percentage=2 elapsed=15.057s
INFO [11-11|19:44:11] Generating DAG in progress               epoch=0 percentage=3 elapsed=19.835s
INFO [11-11|19:44:16] Generating DAG in progress               epoch=0 percentage=4 elapsed=25.091s
INFO [11-11|19:44:19] Generating DAG in progress               epoch=0 percentage=5 elapsed=28.697s

（中略）

INFO [11-11|19:50:01] Generating DAG in progress               epoch=0 percentage=95 elapsed=6m9.997s
INFO [11-11|19:50:04] Generating DAG in progress               epoch=0 percentage=96 elapsed=6m13.453s
INFO [11-11|19:50:07] Generating DAG in progress               epoch=0 percentage=97 elapsed=6m16.638s
INFO [11-11|19:50:10] Generating DAG in progress               epoch=0 percentage=98 elapsed=6m19.652s
INFO [11-11|19:50:14] Generating DAG in progress               epoch=0 percentage=99 elapsed=6m23.110s
INFO [11-11|19:50:14] Generated ethash verification cache      epoch=0 elapsed=6m23.113s
INFO [11-11|19:50:21] Generating DAG in progress               epoch=1 percentage=0  elapsed=4.922s
INFO [11-11|19:50:27] Generating DAG in progress               epoch=1 percentage=1  elapsed=11.015s
INFO [11-11|19:50:33] Generating DAG in progress               epoch=1 percentage=2  elapsed=16.916s
INFO [11-11|19:50:36] Successfully sealed new block            number=1 hash=b6e63c…15283e
INFO [11-11|19:50:36] 🔨 mined potential block                  number=1 hash=b6e63c…15283e
INFO [11-11|19:50:36] Commit new mining work                   number=2 txs=0 uncles=0 elapsed=7.560ms
INFO [11-11|19:50:38] Generating DAG in progress               epoch=1 percentage=3  elapsed=22.292s
INFO [11-11|19:50:45] Generating DAG in progress               epoch=1 percentage=4  elapsed=28.557s
INFO [11-11|19:50:53] Generating DAG in progress               epoch=1 percentage=5  elapsed=36.857s
INFO [11-11|19:50:59] Successfully sealed new block            number=2 hash=a5ae1c…667b40
INFO [11-11|19:50:59] 🔨 mined potential block                  number=2 hash=a5ae1c…667b40
INFO [11-11|19:50:59] Commit new mining work                   number=3 txs=0 uncles=0 elapsed=216.156µs
INFO [11-11|19:51:00] Generating DAG in progress               epoch=1 percentage=6  elapsed=43.786s
INFO [11-11|19:51:06] Generating DAG in progress               epoch=1 percentage=7  elapsed=50.185s
INFO [11-11|19:51:11] Successfully sealed new block            number=3 hash=323f67…b0868f
INFO [11-11|19:51:11] 🔨 mined potential block                  number=3 hash=323f67…b0868f
```

ノードが起動したら以下を実行してノードに接続し、お馴染みの JavaScript コンソールさんでよしなに動作確認などを行う。

```sh
$ docker exec -it geth-standalone geth attach rpc:http://localhost:8545
```

```txt
Welcome to the Geth JavaScript console!

instance: Geth/v1.7.2-stable/linux-amd64/go1.9.1
coinbase: 0x49d38ba99e0a1712388031345114d9da84110e9f
at block: 13 (Sat, 11 Nov 2017 20:14:40 UTC)
 datadir: /root/.ethereum/privchain
 modules: admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

>
```

一旦今日はここまで。
