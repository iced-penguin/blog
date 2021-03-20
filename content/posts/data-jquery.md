---
title: "jQueryでデータ属性（data-*）を扱う"
date: 2020-11-14T22:15:41+09:00
categories: [Programming]
tags: [Javascript, jQuery]
---

データ属性をjQueryで扱う方法をメモしておく。

<!--more-->

# データ属性

`data-*`のような属性をデータ属性という。`*`には任意の文字列が入れることができる。

# データ属性をjQueryで取り扱う

## データをセット

 `data-hoge`に`baz`をセット

```html
<div data-hoge="fuga"></div>
```

```jsx
$('div').data('hoge', 'baz')
```

## データを取得

`data-hoge`の値を取得

```jsx
const hoge = $('div').data('hoge')
```

# 参考文献

[データ属性の使用](https://developer.mozilla.org/ja/docs/Learn/HTML/Howto/Use_data_attributes)

[.data()](https://api.jquery.com/data/)

