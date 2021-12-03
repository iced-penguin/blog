---
title: "exportを実行するシェルスクリプトを別シェルスクリプト内で実行する"
date: 2021-08-12T20:01:20+09:00
categories: [Programming]
tags: [シェルスクリプト]
toc: false
draft: false
---

あまりにもニッチな内容であるが一応、備忘録を残しておく。

<!--more-->

以下のようなシェルスクリプトを考える。

`export.sh`
```
export HOGE=hoge
export FUGA=fuga
```

これを別のシェルスクリプト上で実行したいとする。
そのためには以下のようにする。

```
#!/bin/bash
source export.sh
# 何らかの処理...
```

ポイントは`source`で実行すること。

スクリプト内で一時的に環境変数をセットする場合に使うことがあるかも...？