---
date: 2020-7-30
title: MySQLを停止できない：ERROR! MySQL server PID file could not be found!
categories: [DB]
tags: [MySQL]
---

MySQLを停止しようとしてPIDファイルが見つからないと怒られた時の症状と解決方法のメモ。

<!--more-->

## 環境

- macOS 10.15.8
- MySQL 5.7

# 症状

MySQLを停止したくて以下のコマンドを実行すると、

```
$ mysql.server stop
ERROR! MySQL server PID file could not be found!
```

こんな感じでPIDファイルが見つからないとエラーが出てしまいました。

```
$ mysql -u root -p
```

でログインはできるので起動はしている...。

# 解決方法

いったん止めてみることに。

```
ERROR! MySQL server PID file could not be found!
```

で見つかったプロセスを以下のコマンドでkillしてみる：

```
kill -9 プロセスID

```

すると、

```
$ mysql.server start
Starting MySQL
. SUCCESS!
```

で起動できました。

改めて停止してみます。

```
$ mysql.server stop
Shutting down MySQL
. SUCCESS!
```

今度はうまくいきました。

## まとめ

この手のエラーには時々出くわすので、解決策をメモしておきました。
参考になれば幸いです。
