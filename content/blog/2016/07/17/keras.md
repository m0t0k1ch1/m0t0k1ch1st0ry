+++
title = 'Keras のモデルと学習結果を保存して利用する'
tags = ['keras', 'python']
date = '2016-07-17T01:38:01+09:00'
+++

[Keras を使った簡単な Deep Learning](m0t0k1ch1st0ry.com/blog/2016/07/15/keras) はできたものの、そういえば学習結果は保存してなんぼなのでは、、、と思ったのでやってみた。

<!--more-->

## 準備

[公式の FAQ](http://keras.io/ja/getting-started/faq/#how-can-i-save-a-keras-model) に以下のような記載があるので、h5py を入れておく。

> モデルのweightパラメータを保存する場合，以下のようにHDF5を使います。  
> 注: HDF5とPythonライブラリの h5pyがインストールされている必要があります(Kerasには同梱されていません)。

``` sh
$ pip install h5py
```

## モデルと学習結果を保存する

[前回描いたコード](https://github.com/m0t0k1ch1/keras-sample/blob/master/mnist_mlp.py) に以下を追加して、モデルを `mnist_mlp_model.json` に、学習結果を `mnist_mlp_weights.h5` に保存する。

``` python
model_json_str = model.to_json()
open('mnist_mlp_model.json', 'w').write(model_json_str)
model.save_weights('mnist_mlp_weights.h5');
```

## 保存したモデルと学習結果を利用する

``` python
# -*- coding: utf-8 -*-

import numpy as np
np.random.seed(20160717)

from keras.datasets import mnist
from keras.models import model_from_json
from keras.utils import np_utils

import matplotlib.pyplot as plt

(X_train, y_train), (X_test, y_test) = mnist.load_data()

X_train = X_train.reshape(60000, 784).astype('float32') / 255
X_test  = X_test.reshape(10000, 784).astype('float32') / 255

Y_train = np_utils.to_categorical(y_train, 10)
Y_test  = np_utils.to_categorical(y_test, 10)

# モデルを読み込む
model = model_from_json(open('mnist_mlp_model.json').read())

# 学習結果を読み込む
model.load_weights('mnist_mlp_weights.h5')

model.summary();

model.compile(loss='categorical_crossentropy',
              optimizer='rmsprop',
              metrics=['accuracy'])

score = model.evaluate(X_test, Y_test, verbose=0)
print('Test loss :', score[0])
print('Test accuracy :', score[1])
```

結果は以下。きちんと学習後に近いスコアが出ている。

``` txt
Using TensorFlow backend.
____________________________________________________________________________________________________
Layer (type)                     Output Shape          Param #     Connected to
====================================================================================================
dense_1 (Dense)                  (None, 512)           401920      dense_input_1[0][0]
____________________________________________________________________________________________________
activation_1 (Activation)        (None, 512)           0           dense_1[0][0]
____________________________________________________________________________________________________
dropout_1 (Dropout)              (None, 512)           0           activation_1[0][0]
____________________________________________________________________________________________________
dense_2 (Dense)                  (None, 512)           262656      dropout_1[0][0]
____________________________________________________________________________________________________
activation_2 (Activation)        (None, 512)           0           dense_2[0][0]
____________________________________________________________________________________________________
dropout_2 (Dropout)              (None, 512)           0           activation_2[0][0]
____________________________________________________________________________________________________
dense_3 (Dense)                  (None, 10)            5130        dropout_2[0][0]
____________________________________________________________________________________________________
activation_3 (Activation)        (None, 10)            0           dense_3[0][0]
====================================================================================================
Total params: 669706
____________________________________________________________________________________________________
Test loss : 0.118880221067
Test accuracy : 0.9829
```
