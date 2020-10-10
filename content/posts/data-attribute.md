---
title: data属性を利用してJavascriptにデータを渡す
date: 2020-08-22 21:07:14
---


HTML側からjavascript側にデータを渡す方法としてdata属性を使うことができます。
data属性の基本的な使い方の説明です。

<!--more-->

# data属性の使い方

HTMLのタグに`data-*`という名前の属性を与えます（`*`は任意）。

HTML:

```html
<div id="animal" data-cat="Cat" data-dog="Dog">
```

Javascriptからは次のようにアクセスします。

Javascript:

```javascript
const animal = document.getElementById('animal')

const cat = animal.dataset.cat
const dog = animal.dataset.dog

console.log(cat)  // Cat
console.log(dog)  // Dog
```

## Thymeleafでdata属性を使う

僕がよくSpringBootを使うのでThymleafでdata属性を扱う方法もメモしておきます。

HTML:

```html
<div id="animal" th:data-cat="Cat" th:data-dog="Dog">
```

Thymeleafなので、当然data属性に与える値には単なる文字列だけでなく様々な式が使えます。

## jQueryで使う

```javascript
const animal = $('#animal')

const cat = animal.data('cat')
const dog = animal.data('dog')

console.log(cat)  // Cat
console.log(dog)  // Dog
```

# まとめ

data属性を使うと簡単にHTMLからJSへ値を渡せました。
JSへ値を渡す方法はいくつかあるので場合によって使い分けると良いでしょう。

## 参考

- [data属性の使用 - MDN](https://developer.mozilla.org/ja/docs/Learn/HTML/Howto/Use_data_attributes)
- [jQueryのdata()で属性を取得・設定・変更する方法まとめ！](https://www.sejuku.net/blog/38263)
