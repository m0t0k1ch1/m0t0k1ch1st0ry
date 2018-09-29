+++
date = "2016-09-06T01:52:29+09:00"
tags = [ "phalcon", "php" ]
title = "PHP 7 で multi module な Phalcon 3.0.0 を動かしてみる"
+++

[Phalcon](https://phalconphp.com) は普段から使っているけれど、まだ 3.0.0 を触ってなかったので触ってみることにした。ついでに、view として HTML を返すエンドポイント群と js から叩かれる API エンドポイント群の切り分けを意識した multi module な構成の雛形をつくってみる。

<br />
## PHP と Phalcon をインストールする

もともと brew で php56 を入れていたので、とりあえず unlink しておく。もとに戻したいときの手順は後述。

``` sh
$ brew unlink php56
```

brew で php70 と php70-phalcon をインストール。

``` sh
$ brew install php70
$ brew install php70-phalcon
```

php のバージョンを確認。

``` sh
$ php --version
```

<pre>
PHP 7.0.10 (cli) (built: Aug 21 2016 19:14:33) ( NTS )
Copyright (c) 1997-2016 The PHP Group
Zend Engine v3.0.0, Copyright (c) 1998-2016 Zend Technologies
</pre>

Phalcon については、以下のように `/usr/local/etc/php/7.0/conf.d/ext-phalcon.ini` が読み込まれている。

``` sh
$ php -i | grep ext-phalcon.ini
```

<pre>
Additional .ini files parsed => /usr/local/etc/php/7.0/conf.d/ext-phalcon.ini
</pre>

中身を確認する。

``` sh
$ cat /usr/local/etc/php/7.0/conf.d/ext-phalcon.ini
```

<pre>
[phalcon]
extension="/usr/local/opt/php70-phalcon/phalcon.so"
</pre>

以下のように、`/usr/local/opt/php70-phalcon` は `../Cellar/php70-phalcon/3.0.0` への symlink なので、Phalcon のバージョンも 3.0.0 ということで大丈夫そう。

``` sh
$ ll /usr/local/opt | grep php70-phalcon
```

<pre>
lrwxr-xr-x   1 m0t0k1ch1  admin    29  9  1 10:23 php70-phalcon -> ../Cellar/php70-phalcon/3.0.0
</pre>

ちなみに、php56 に戻したい場合は以下のようにする。逆も然り。

``` sh
$ brew unlink php70
$ brew link php56
```

phpenv はなんかうまいこと動かなかったし、この 2 バージョン以外は触る予定ないし、そんなに頻繁に切り替えることはないので、これでいいかという感じ。

<br />
## phalcon-devtools をインストールする

自分は [ここ](https://github.com/phalcon/phalcon-devtools#installation-via-git) に書いてあるように git で落としてきてパスを通した。

こちらのバージョンも 3.0.0 であることを確認。

``` sh
$ phalcon
```

<pre>
Phalcon DevTools (3.0.0)

Available commands:
  commands         (alias of: list, enumerate)
  controller       (alias of: create-controller)
  module           (alias of: create-module)
  model            (alias of: create-model)
  all-models       (alias of: create-all-models)
  project          (alias of: create-project)
  scaffold         (alias of: create-scaffold)
  migration        (alias of: create-migration)
  webtools         (alias of: create-webtools)
</pre>

<br />
## プロジェクトの雛形をつくる

``` sh
$ phalcon create-project multi-module-phalcon
```

サーバーを立ち上げる。

``` sh
$ cd multi-module-phalcon
$ php -S 127.0.0.1:8000 -t public .htrouter.php
```

<pre>
PHP 7.0.10 Development Server started at Tue Sep  6 01:02:42 2016
Listening on http://127.0.0.1:8000
Document root is /Users/m0t0k1ch1/.ghq/src/github.com/m0t0k1ch1/multi-module-phalcon/public
Press Ctrl-C to quit.
</pre>

ブラウザで http://127.0.0.1:8000 にアクセス。

![phalcon.png](/img/entry/phalcon.png)

大丈夫そう。

<br />
## multi module な感じにする

公式の [ドキュメント](https://docs.phalconphp.com/ja/latest/reference/applications.html) や [GitHub repo](https://github.com/phalcon/mvc) を参考に、えいやっと最低限な状態をつくってみる。

できあがったのがこちら。

<div class="github-card" data-user="m0t0k1ch1" data-repo="multi-module-phalcon"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>

`apps` 以下はこんな感じ。

<pre>
apps
├── backend
│   ├── Module.php
│   ├── config
│   │   └── config.php
│   └── controllers
│       ├── ControllerBase.php
│       └── UserController.php
├── frontend
│   ├── Module.php
│   ├── config
│   │   └── config.php
│   ├── controllers
│   │   ├── ControllerBase.php
│   │   └── IndexController.php
│   └── views
│       ├── index
│       │   └── index.volt
│       ├── index.volt
│       └── layouts
└── routes.php
</pre>

`apps/routes.php` で以下のようにルーティングしている。

``` php
<?php
$router = new \Phalcon\Mvc\Router;
$router->setDefaultModule('frontend');

/*
 * View
 */
$router->add('/', [
    'module'     => 'frontend',
    'controller' => 'index',
    'action'     => 'index',
]);

/*
 * API
 */
$router->addGet('/user', [
    'module'     => 'backend',
    'controller' => 'user',
    'action'     => 'get',
]);

return $router;
```

これを先ほどと同じように立ち上げて http://127.0.0.1:8000 にアクセスすると、同じ画面が表示されるはず。内部的には frontend module に流れている。

で、http://127.0.0.1:8000/user を GET で叩くと、以下のようなレスポンスが返ってくる。

``` sh
$ curl -s http://127.0.0.1:8000/user | jq .
```

``` json
{
  "name": "m0t0k1ch1"
}
```

こちらは backend module に流れている。

ちなみに、`public/index.php` は以下のようになっている。`Phalcon\Mvc\Application` を使っていないのは、backend module（JSON 返すだけ）で view を DI コンテナに突っ込みたくなかったから。処理のフローが明示的になるので、エラーハンドリングもやりやすくなるといいなという想いもある。

``` php
<?php

use \Phalcon\Di\FactoryDefault;
use \Phalcon\Loader;
use \Phalcon\Mvc\Router;

error_reporting(E_ALL);

define('BASE_PATH', dirname(__DIR__));
define('APP_PATH', BASE_PATH . '/apps');

$di = new FactoryDefault;
$di->set('router', function() {
    return include APP_PATH . '/routes.php';
});

try {
    $router = $di['router'];
    $router->handle();

    $moduleName = $router->getModuleName();
    switch ($moduleName) {
    case 'frontend':
        require_once APP_PATH . '/frontend/Module.php';
        $module = new \Multi\Frontend\Module;
        break;
    case 'backend':
        require_once APP_PATH . '/backend/Module.php';
        $module = new \Multi\Backend\Module;
        break;
    default:
        throw new \RuntimeException('unknown module');
    }

    $module->registerAutoloaders($di);
    $module->registerServices($di);

    $dispatcher = $di['dispatcher'];
    $dispatcher->setModuleName($moduleName);
    $dispatcher->setControllerName($router->getControllerName());
    $dispatcher->setActionName($router->getActionName());
    $dispatcher->setParams($router->getParams());

    $dispatcher->dispatch();

    if ($moduleName == 'frontend') {
        $view = $di['view'];
        $view->start();
        $view->render(
            $dispatcher->getControllerName(),
            $dispatcher->getActionName(),
            $dispatcher->getParams()
        );
        $view->finish();

        $response = $di['response'];
        $response->setContent($view->getContent());
        $response->send();
    }
}
catch (\Exception $e) {
    echo $e;
}
```
