---
title: MacへのSDKMANのインストールでエラー
date: 2020-09-11 23:17:49
categories: [Programming]
tags: [SDKMAN]
---

`No such file or directory`みたいなエラーが吐かれたが、原因や解決法が同じと思われる人が見当たらなかったので記録を残しておきます。

<!--more-->

## 環境

- Mac OS 10.15.6

# エラー内容

SDKMANをインストールしようとして、公式にしたがって以下のコマンドを実行したところうまくいっていない、、`~/.sdkman`もできていない、、

```sh
curl -s "https://get.sdkman.io" | bash
```

メッセージをよく見たところ、

```sh
mkdir: /Users/hoge/.sdkman/bin: Permission denied
```

などといったエラーが出ていました（hogeは実際の値ではありません）。


# 解決方法

初め`Permission denied`の方に気を取られていたのですが、よく考えたら私のユーザ名は hoge ではない。そんなユーザは存在しない。つまり誤ったホームディレクトリを指定してしまっている。

SDKMANのインストール先は`SDKMAN_DIR`によって指定されているようなので、確認したところ、

```sh
echo $SDKMAN_DIR  # /Users/hoge
```

やはり間違ったパスになっていました。

そこで正しい値を環境変数にセットしてから再度インストールを試みることにします。　

```sh
export SDKMAN_DIR="/Users/penguin"
curl -s "https://get.sdkman.io" | bash
```

今度はうまく行きました。

今回の問題とは直接関係ありませんが、インストールを完遂するのに次のコマンドを実行しました。

```sh
source "$HOME/.sdkman/bin/sdkman-init.sh"
```

# まとめ

SDKMANのインストールに失敗するときはユーザ名が誤っているかもしれないので確認して見るといいかもしれません。
