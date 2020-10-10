---
title: coc.nvimのインストール・設定
date: 2020-10-10 10:23:29
---

vimにcoc.nvimをインストールしてみたので、その方法をメモしておきます。
また、使い方に関してもまとめておき、随時更新していきます。

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

# 設定

使いたい言語の拡張を入れればそれだけで十分便利だとは思いますが、個人的に挙動を少し変更している部分を書いておきます。

設定は基本的に`:CocConfig`で開く`coc-settings.json`に書きます。

困ったら`:h coc-nvim`。

## ステータスラインにメッセージを表示

設定方法は`:h coc-status`を参考にしてください。

僕は`lightline`を使っているので、それに合わせて次のように設定します。

```vim
" Use Nord theme
" Show coc-vim status message
let g:lightline = {
      \ 'colorscheme': 'nord',
      \ 'active': {
      \   'left': [ [ 'mode', 'paste' ],
      \             [ 'cocstatus', 'readonly', 'filename', 'modified' ] ]
      \ },
      \ 'component_function': {
      \   'cocstatus': 'coc#status'
      \ },
      \ }
```

`colorscheme`はお使いのものに変更してください。

## 補完の挙動　

特に何もしなくても補完が効きます。が、挙動を他のIDEやエディタのようにしたいので少し設定をいじります。

具体的には、最初の候補がデフォルトで選択され、`return`で決定するようにします。

`:CocConfig`で`coc-settings.json`を開き、以下を書き込みます。

```
{
  "suggest.noselect": false
}
```

## coc-java

Java用の拡張です。インストール：

```
:CocInstall coc-java
```

# まとめ

coc.nvimをNeoVimにインストールしてみました。
デフォルトでいろいろな機能が揃っているのでとても便利です。

## 参考

- [coc.nvim](https://github.com/neoclide/coc.nvim)
- [MacにNode.jsをインストール](https://qiita.com/kyosuke5_20/items/c5f68fc9d89b84c0df09)
- [最小限なcoc.nvim導入手順](https://blog.sgry.jp/entry/2020/03/14/194130)
- [auto select of first item from the drop down menu #2221](https://github.com/neoclide/coc.nvim/issues/2221)
