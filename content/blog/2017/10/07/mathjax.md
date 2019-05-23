+++
title = "Hugo に MathJax を導入して数式を書けるようにする"
tags = [ "math" ]
date = "2017-10-07"
+++

数式を入れたエントリを書こうとしたけど、数式書けないやんってなったので [MathJax](https://www.mathjax.org) を導入することに。

<!--more-->

このブログは [自分でカスタマイズした ghostwriter](https://github.com/m0t0k1ch1/ghostwriter/tree/m0t0k1ch1st0ry) を theme として使っているので、素直に `layouts/partials/footer.html` に以下を追加した。

``` html
<script type="text/x-mathjax-config">
    MathJax.Hub.Config({
        tex2jax: {inlineMath: [['$','$'], ['\\(','\\)']]}
    });
</script>
<script type="text/javascript" async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.2/MathJax.js?config=TeX-MML-AM_CHTML"></script>
```

最低限、これだけで大丈夫そう。簡単。

実際に試してみる。

``` txt
$$
T = S((1 + \frac{E}{R})^{F} - 1)
$$
```

上記をそのまま markdown ファイルに書くと、以下のようになる。

$$
T = S((1 + \frac{E}{R})^{F} - 1)
$$

いい感じ。

インライン表示もできて、

``` txt
次回のエントリでは $T = S((1 + \frac{E}{R})^{F} - 1)$ について書こうと思います。
```

上記が以下のようになる。

次回のエントリでは \\(T = S((1 + \frac{E}{R})^{F} - 1)\\) について書こうと思います。
