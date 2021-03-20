---
date: 2020-6-9
updated: 2020-8-30
title: iterm2にNordテーマを適用してみた
categories: [Programming]
tags: [iTerm2, Nord]
---

item2がすごいらしいとの噂だったので、ターミナルエミュレータをMac標準のものからiterm2に乗り換え、テーマをNordに変更しました。

<!--more-->

# 環境

- MacOS 10.15.5
- iterm2 3.3.11

# 手順

[リポジトリ](https://github.com/arcticicestudio/nord-iterm2)のREADMEにきっちり書いてありますが、細かいところを補いながらインストールして適用するまでの手順をまとめます。

1. 適当なデイレクトリにnord-iterm2のリポジトリをクローン

```
git clone https://github.com/arcticicestudio/nord-iterm2.git 
```

2. `Command + ,`で`Preferences`を開く

3. 上部の`Profiles`をクリック

4. `Colors`タブをクリック

5. 右下の`Color Presets`から`import`を選び、先ほどクローンしたリポジトリ中の`Nord.itermcolors`を選択

`nord-iterm2/src/xml/Nord.itermcolors`にありました。

6. `Color Presets`に`Nord`が出現しているのでそれを選択

これでiterm2にNordテーマが適用されます。

# まとめ

iterm2にNordを入れてみました。
クールで見やすいおすすめのテーマです。

## 参考

- [nord-iterm2](https://github.com/arcticicestudio/nord-iterm2)
