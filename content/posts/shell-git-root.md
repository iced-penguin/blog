---
title: "Gitリポジトリのルートに一気に移動するコマンドを書いた"
date: 2021-04-18T17:32:13+09:00
categories: [Programming]
tags: [Shell, Git]
toc: false
draft: false
---

# 概要

深い階層まで潜ったあと、プロジェクトルートまで戻りたいことがある。
ちまちまと`cd`を打ったり`cd ../../..`のように階層の数を指定したりするのが面倒になり、何も考えずに一気に戻りたくなったので、シェルスクリプトを書いてコマンド化した。

<!--more-->

# 実装

## ルートのパスを返すスクリプト

まず、Gitリポジトリのルートディレクトリのパスを標準出力に書き出すシェルスクリプトを書く。
ファイル名は`gitroot`とした。コマンド化することを考えて拡張子はつけていない。

`gitroot`

```
#!/bin/bash

if [ $# -ne 0 ]; then
  echo "gitroot: too many arguments"
  echo "usage: gitroot"
fi

# $1 パス
function is_root() {
  FILES=("$(ls -a $1)")
  FLAG=false
  for file in $FILES; do
    if [ $file = '.git' ]; then
      FLAG=true
    fi
  done
  echo $FLAG
}

DIR=$(pwd)
IS_ROOT=false

while [ $DIR != '/' -a $IS_ROOT = false ]
do
  IS_ROOT=$(is_root $DIR)
  if [ $IS_ROOT = false ]; then
    DIR=$(dirname $DIR)
  fi
done

if [ $IS_ROOT = false ]; then
  echo "gitroot: you are not in git repository"
  exit 1
fi

echo $DIR
```

ルートディレクトリであるかどうかは`.git`が存在しているかどうかによって判定する。
現在のディレクトリに`.git`が見つからない場合は親ディレクトリを探す、という処理を繰り返す。
`/`まで辿り着いてしまった場合にはエラーメッセージを出力する。

## ルートのパスを返すコマンドを作る

適当な自作コマンド用ディレクトリを用意し、パスを通す。
私の場合は`~/bin`に先ほど作成した`gitroot`を入れ、`.zshrc`に以下を追記した。
```
export PATH=$HOME/bin:$PATH
```

`source ~/.zshrc`などとすれば、`gitroot`コマンドが使えるようになるはず。

## ルート直下に移動できるようにする

先ほど作成したコマンドと'cat'を組み合わせて、ルートに移動できるようにする。
シェルスクリプトの中で`cat`を扱うのは面倒なので、このように機能を分割する方針とした。
全部`.zshrc`に書かなかったのは設定ファイルにごちゃごちゃと書きたくなかったため。

単純に
```
alias cdroot='cd $(gitroot)'
```
のようにしても良いと思うが、これだとGitリポジトリの外でこのコマンドを叩いた時にエラーが出てしまうので、`.zshrc`に次のような関数を定義した。

```
function cdroot() {
  if [ $(which gitroot) ]; then
    RESULT=$(gitroot)
    STATUS=$?
    if [ $STATUS -eq 0 ]; then
      cd $RESULT
    else
      echo $RESULT
    fi
  else
    echo 'cdroot: command gitroot does not exist'
  fi
}
```

`source ~/.zshrc`とすれば、`cdroot`でGitのルートに跳べるようになったはず。

# まとめ

- Gitリポジトリのルートのパスを出力するコマンドを実装した
- 上記のコマンドを`cd`と組み合わせることで、リポジトリのルート直下に一気に移動できるようにした。
