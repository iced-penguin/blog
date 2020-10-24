---
title: "coc-nvimの設定晒す"
date: 2020-10-24T21:52:35+09:00
draft: true
---

coc-nvimの個人的な設定のメモです。

<!--more-->

## 環境

- macOS 10.15.7
- NeoVim 0.4.4

# はじめに

NeoVimにcoc-vimを入れて色々と使ってみた結果、デフォルトでも便利ですが細かい挙動などを調整したくなったので、その設定方法をここにまとめておきます。
インストールした拡張についても書くかも。

# 設定

coc-nvimの設定は基本的に`:CocConfig`で開く`coc-settings.json`に書きます。

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

特に何もしなくてもデフォルトで補完が効きます。が、挙動を他のIDEやエディタのようにしたいので少し設定をいじります。

具体的には、最初の候補がデフォルトで選択され、`return`で決定するようにします。

`:CocConfig`で`coc-settings.json`を開き、以下を書き込みます。

```
{
  "suggest.noselect": false
}
```

# まとめ

## 参考資料

- [auto select of first item from the drop down menu #2221](https://github.com/neoclide/coc.nvim/issues/2221)
