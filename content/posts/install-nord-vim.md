---
title: VimにNordテーマを適用
date: 2020-06-21 14:20:05
categories: [Programming]
tags: [Vim, Nord]
---

VimにもNordをインストールしてみたので、その手順のメモです。

<!--more-->

## 環境

- macOS 10.15.6
- iterm2 3.3.12

僕はitemr2上でvimを起動しているので、先にiterm2にNordを適用しておきます。具体的な方法は[こちら](https://penguin-code.com/iterm2-install-nord/)。

# Nordとは

寒色系のクールなダークテーマ。

![](https://raw.githubusercontent.com/arcticicestudio/nord-docs/develop/assets/images/ports/vim/overview-go-nerdtree.png)

# インストール手順

iTerm2にNordをインストールしていない場合は、[こちら](https://penguin-code.com/iterm2-install-nord/)。を参考にインストールしておきます。

## プラグインマネージャをインストール

すでにvimのプラグインマネージャを使っている場合はそれを使えば良いかと思います（自分はvim-plugでのみ確認しています）。

もし使っていない場合はインストールします。vim-plugが推奨されているのでそれにします。

ターミナルを開いて以下のコマンドを実行：

```sh
curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```

以上。

## Nordをインストール

vim-plugを使ってnordをインストールします。
`.vimrc`を開き、`arcticicestudio/nord-vim`をプラグインに追加します。
さらに`colorscheme nord`を追加するか、`:colorscheme nord`を実行してテーマを反映します。

```vim
call plug#begin('~/.vim/plugged')
Plug 'arcticicestudio/nord-vim'
call plug#end()

colorscheme nord
```

# まとめ

- vimにNordテーマを適用できる

プラグインマネージャを使用するだけでした。
紹介したのはvim-plugを使う方法ですが、他のプラグインマネージャでも入手できるはずです。

## 参考

- [nord-vim](https://github.com/arcticicestudio/nord-vim)
