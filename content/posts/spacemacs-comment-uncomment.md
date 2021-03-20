---
date: 2020-5-16
title: spacemacsでコメント・アンコメント
categories: [Programming]
tags: [Spacemacs]
---

spacemacsでも（行頭に記号を挿入するのではなく）コマンドでコメント/アンコメントをしたい。
ということでそれを実現するレイヤーを導入しました。

<!--more-->

# Evil-commentaryレイヤー

コメント/アンコメントの機能を提供する`evil-commentary`を導入します。

## インストール

`.spacemacs`に追加してあげればOKです。

```
 dotspacemacs-configuration-layers
 '(
   evil-commentary  ;; 追加
   yaml
   markdown
```

`SPC f e R`で`.spacemacs`をリロードして完了。

## 基本的な使い方

`g c c`で現在行をコメント/アンコメント。
範囲選択している場合は選択範囲をコメント/アンコメント。

`SPC ;`+移動コマンドで現在位置から移動先までコメント/アンコメント。
例えば、`SPC ; 2 j`で現在行から２行下までまとめてコメントアウト/アンコメント。

## 参考

- [Evil-commentary layer](https://www.spacemacs.org/layers/+vim/evil-commentary/README.html)
- [How to comment/uncomment code in Spacemacs](https://simpletutorials.com/c/2878/How+to+comment-uncomment+code+in+Spacemacs)
