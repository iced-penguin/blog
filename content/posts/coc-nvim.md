---
title: coc.nvimをインストールしてみた
date: 2020-10-10 10:23:29
categories: [Programming]
tags: [Vim]
---

vimにcoc.nvimをインストールしてみたので、その方法をメモしておきます。

<!--more-->

## 環境

- macOS 10.15.7
- NeoVim 0.4.4

# coc.nvim？

coc.nvimはvimのLSPをサポートするプラグインです。

ではそのLSPとはそもそも何かというと、これは Language Server Protocol の略で、開発を支援する諸機能（補完機能、定義ジャンプ、エラー解析などなど）をIDEやエディタから分離し、それらの仕様を定めたものです。

LSPについては以下の記事で詳しく解説されています。

[language server protocolについて (前編)](https://qiita.com/atsushieno/items/ce31df9bd88e98eec5c4)

vimにおいてもこれを利用してIDEライクな開発環境を作ることができます。そのためのプラグインの一つがcoc.nvimです（他にもいろいろなLSPプラグインがあります）。

nvimと付いていますが、vimでも動きます（一定以上のパージョンが必要かも？）。

# インストール

## Node.jsをインストール

もしお使いの環境にNode.jsが入っていなかったら以下の記事を参考にインストールしてください。

[MacにNode.jsをインストール](https://qiita.com/kyosuke5_20/items/c5f68fc9d89b84c0df09)

## coc.nvimをインストール

僕は`vim-plug`を使っているので、それを使ってインストールします。
`vim-plug`のプラグインに以下を追加して`:PlugInstall`を実行します。

```
Plug 'neoclide/coc.nvim', {'branch': 'release'}
```

他のプラグインマネージャーではそれぞれのやり方に従ってください。

# 使い方

デフォルトで使える機能の例：

- 入力補完
- リアルタイム文法チェック

## coc extensions

各言語の拡張を入れることによってその言語に関する様々な支援機能を得ることができます。

例えばJava用の拡張であれば`:CocInstall coc-java`で入れられます。

以下のリンクに拡張の詳細と一覧が載っています。

[Using coc extensions](https://github.com/neoclide/coc.nvim/wiki/Using-coc-extensions)


# まとめ

coc.nvimをNeoVimにインストールしてみました。
デフォルトでいろいろな機能が揃っているのでとても便利です。

## 参考

- [coc.nvim](https://github.com/neoclide/coc.nvim)
- [MacにNode.jsをインストール](https://qiita.com/kyosuke5_20/items/c5f68fc9d89b84c0df09)
- [最小限なcoc.nvim導入手順](https://blog.sgry.jp/entry/2020/03/14/194130)
