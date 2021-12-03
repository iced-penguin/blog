---
title: "csvヘッダーの列名がそれぞれ何列目かを出力するワンライナー"
date: 2021-12-03T18:13:41+09:00
categories: [Programming]
tags: [シェルスクリプト]
toc: false
draft: false
---

csvのどの列が何列目かを知りたいが、たくさん列があり数えるのが大変な時に。

<!--more-->

以下のダミーデータを使ってみる。

`sample.csv`
```
firstname,lastname,age,birthday,gender,email
百合,野中,45,1976/7/6,女,EVd0wN4@test.org
宏美,越川,35,1986/8/22,女,N1OdZ@example.com
銀蔵,田宮,30,1991/7/29,男,RC4MJZU@test.org
```

これのヘッダーだけを取り出し、なおかつそれぞれが何列目かを表示する。

```
$ head -1 sample.csv | tr ',' '\n' | nl
     1	firstname
     2	lastname
     3	age
     4	birthday
     5	gender
     6	email
```

`tr` でカンマを改行に変換し、`nl`で行数を表示している。
