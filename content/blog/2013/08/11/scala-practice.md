+++
title = '第1回 Scala で遊ぼ☆ 「関数型のスタイル -基礎-」'
tags = ['scala']
date = '2013-08-11'
+++

最近なんとなく Scala を勉強しています（仕事柄、先に Perl もっと勉強しろという声が聴こえてきそうですが、勇気を振り絞って執筆しております）。

<!--more-->

学生の頃から ↓ こんな感じで興味はあったんです。

- オブジェクト指向と関数型プログラミングのハイブリッドって、なんかかっこいいやん！！！
- Twitter 社が使っているらしいし、なんかいい感じに違いない
- 単純に新しい雰囲気漂うもの好きなんです
- PHP からプログラミング始めた身、オブジェクト指向をちゃんと勉強する機会を求めてた

ホントに単純になんとなく興味があるだけで、欲求不満を解消するために勉強するようなものなので、Scala を使う利点とかそれを今学ぶ意義みたいなのはとりあえずあとづけでいいかなくらいのテンションであることはご了承ください。

ということで、「Scala で遊ぼ ☆」と題しまして、Scala ど素人がちょっとなんかつくれるようになるまでの過程を数回に分けてアウトプットしつつ、自分の備忘録的に活用できたらなあと考えております。  
生温かい目でお見守りいただければ幸いです。

## はじめに

かの [コップ本](http://www.amazon.co.jp/Scala%E3%82%B9%E3%82%B1%E3%83%BC%E3%83%A9%E3%83%96%E3%83%AB%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E7%AC%AC2%E7%89%88-Martin-Odersky/dp/4844330845/ref=pd_cp_b_0) の第 1 章に Scala の特徴はいろいろと書いてあるんですが、初めて読んだときは「全然わかんね！！！！！」って感じで、その中でも「そもそも関数型プログラミングってなんなの！！！？？？
」って感じだったので、実際にコードを書きつつ爽やかに触れていきます。

## 関数型プログラミング

とりあえず、[コップ本](http://www.amazon.co.jp/Scala%E3%82%B9%E3%82%B1%E3%83%BC%E3%83%A9%E3%83%96%E3%83%AB%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E7%AC%AC2%E7%89%88-Martin-Odersky/dp/4844330845/ref=pd_cp_b_0) から引用させていただきます。

> 関数型プログラミングは、主として 2 つの発想によって導かれている。1 つは、関数を一人前の値（first-class values）として扱うことである。関数型言語では、関数は整数や文字列などと同格の値である。
>
> （中略）
>
> 関数型プログラミングを導くもう 1 つの発想は、「プログラム内の操作は、データをその場で書き換え（変更）するのではなく、入力の値を出力の値にマップ（写像）すべきだ」というものである。
>
> （中略）
>
> 関数型プログラミングのこの第 2 の発想は、「メソッドはどんな副作用（side effect）も持ってはならない」と言い換えられる。メソッドは、引数をとり、結果値を返すという方法だけで環境と通信しなければならないのである。

…って言われても正直僕はよくわかりませんでした。  
ので、とりあえず、ここで言及されている 2 つの発想について、以下でもう少し詳しく見ていきます。

## 関数は一人前の値であり、どんな副作用も持ってはならない

- とりあえず関数は結果として値を出すべきや！それ以外は副作用！って感じの思想
- その思想は基本制御構造にも反映されてる

### 値を返す（副作用のない）関数

- 基本はこういう関数を書くように努める
- 関数本体の前の等号は「関数型の世界の視点では、関数とは値を結果として出す式を定義したものだよ」ということを示している
- だから、それ以外のはたらきは副作用ってことになる

```scala
def addPiyo(str: String): String = {
  str + " piyo!"
}

println(addPiyo("Hello world"))
```

- 実行すると、結果は当然の如く以下

```txt
Hello world piyo!
```

### 値を返さない（副作用のある）関数

- こういう関数は可能な限り定義しないように努める
- 実は `Unit` という結果型が返ってくる
- `Unit` 型は、関数が意味のある値を返してこないことを示す
- つまり、結果型が `Unit` の関数は、副作用のためにのみ実行されるということになる
- この場合の副作用は「標準出力への文字列出力」

```scala
def addPiyoAndPrint(str: String) = {
  println(str + " piyo!")
}

addPiyoAndPrint("Hello world")
```

- こちらも実行すると、結果は当然の如く以下

```txt
Hello world piyo!
```

### for も値を返せる

- `for ( ) yield { }` ってする
- Perl の `grep` に近いイメージ
- `for` に限らず、`if` とか `try` とか `match`（`switch` 的な）とかの基本制御構造も値を返せる

```scala
def searchPiyo() =
  for (
    arg <- args
    if arg.contains("piyo")
  ) yield arg

val result = searchPiyo
println(result.mkString(", "))
```

- ↑ これを `searchPiyo.scala` とかで保存する
- で、以下のように実行する

```sh
$ scala searchPiyo.scala puyopuyo no piyon piyo!
```

- と、結果は以下

```txt
piyon, piyo!
```

## まとめ

関数型プログラミングの基本思想は以下のような感じ。

- 基本、関数は値を返すもの
- 値を返さない = 副作用がある
- 関数は副作用を持つべきではない
- なので、`for` などの基本制御構造も値を返せちゃったりする

他にも、

- `var` じゃなくて `val` を使おう
- ミュータブルなオブジェクトじゃなくてイミュータブルなオブジェクトを使おう

っていう話もあったりするのですが、話が長くなるので今回は置いときました。  
次はこの辺に触れてもよいかもしれません。

あと、静的型付けっていう性質がこの辺とシナジー強い気がするので、その辺の関係性とかをちゃんと整理して言語思想あたりの話と絡めて説明できるようにならないとなっと思います。

また、Scala は「基本、関数型のスタイルでコード書いてね！でも、どうしてもってときのために命令型でも書けるようにしてるんだよ！ほら！」っていう感じで命令型のスタイルを選択肢として残してるあたり、そこまで関数型にふりきれてるわけではないようです（それはそれでどうなんだろうとは思ったりもしますが）。

正味、いろいろ自分で目的もってコード書いてみないと、メリット・デメリットが実感できるレベルまでは落ちてこないですね！きっと！

なので継続してがんばります。押忍。
