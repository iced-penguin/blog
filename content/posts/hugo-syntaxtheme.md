---
title: "[Hugo] シンタックステーマを適用してカスタマイズする"
date: 2021-03-26T22:50:42+09:00
categories: [Programming]
tags: [Hugo]
toc: false
---

# 概要

HugoのテーマをNotepadiumにしてみたが、シンタックスハイライトだけ変えたかった。というわけでシンタックステーマを変更し、さらにそれをカスタマイズしてみた。

<!--more-->

# シンタックステーマを適用

- Hugoテーマ: Notepadium

HugoのテーマはNotepadiumを使用している。シンタックステーマの変更方法は基本的には公式の説明にあると思うので、それに従えば良い。今回はNotepadiumでの変更方法を述べる。
Nordというテーマが好きなので、それに対応しているhighlight.jsを使う。

`config.toml`を開き、以下を追記。

```
[params.syntax]
use = "hljs"  # 1. prismjs 2. hljs 3. none
theme = "nord"
darkTheme = "nord"
```

# シンタックステーマをカスタマイズ

上記の方法でシンタックステーマが適用される。
Nordは個人的にとても気に入っているのだが、コメントの色が暗すぎて見辛いため、そこだけ変更したい。

Nord用のCSSは`themes/notepadium/assets/css/hljs/nord.css`以下に書かれている。ここで、このファイルを`assets/css/hljs/nord.css`にコピーして編集する。`assets/css/hljs/nord.css`の方が優先して読み込まれるので、元のファイルを汚すことなくカスタマイズすることができる。ただし、`assets`以下のディレクトリ構造は同じである必要がある。

`assets/css/hljs/nord.css`を開いて編集する。変更したい箇所の色を自分の好きなように変更するのだが、今回はコメントの色を変える。

```
.hljs-comment {
  color: #4C566A;
}
```


# 参考

- [Hugo Themes - Notepadium](https://themes.gohugo.io/hugo-notepadium/)
