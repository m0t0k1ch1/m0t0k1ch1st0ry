---
date: '2017-03-31T01:24:20+09:00'
tags: ['azure']
title: 'Azure で Application Gateway 下の Healthy なホスト数を取得する'
---

表題の通り、パッと見簡単そうなことがやりたかったのだけど、シュッとやる方法をシュッと見つけられなかったのでメモ。

<!--more-->

[Azure CLI 2.0](https://docs.microsoft.com/en-us/cli/azure/overview) に `az network application-gateway show-backend-health` というコマンドがあり、これのレスポンスの中にバックエンドプールのホストの情報が含まれている。

ref. [Application Gateway - az network application-gateway](https://docs.microsoft.com/en-us/cli/azure/network/application-gateway)

`--query` オプションに対し [JMESPath](http://jmespath.org) 形式で取得したい情報を指定することができるので、これを駆使して Healthy なホストの数を抽出してみる。ちなみに、JMESPath は初体験。

例えば、以下のような感じ。

```sh
$ az network application-gateway show-backend-health \
  --resource-group <your resource group> \
  --name <your appilcation gateway name> \
  --query "backendAddressPools[0].backendHttpSettingsCollection[0].servers[?health=='Healthy'] | length(@)"
```

ただ、このコマンド、重い。20 〜 30 秒くらいはかかる。

Application Gateway 関連の操作は、所要時間が他のリソースと比べてだいぶ長い。祈りを捧ぐ 🙏
