---
title: "cutで切り出した列を行に変換"
date: 2021-03-28T19:52:57+09:00
categories: [Programming]
tags: [シェルスクリプト]
toc: false
draft: false
---

# 概要

cutなどで取得した１列の縦のデータを行に変換する。

<!--more-->

# 方法

例えばcsvファイルから１列を切り出して

```
cut -d ',' -f 1 sample.csv

# 1
# 2
# 3
```

これをカンマ区切りの行に変換したいとき

```
cut -d ',' -f 1 sample.csv | tr '\n' ',' | sed 's/,$/\n/g'

# 1,2,3
```

`tr`は文字を置換するコマンド。ここでは改行をカンマに変換している。

`sed`で行末のカンマを改行に変換。

`cut`以外の方法で列データを取り出す場合やカンマ以外の区切り文字を用いる場合も同様の方法で対処可能。
