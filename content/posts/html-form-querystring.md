---
date: 2020-8-29
title: 【HTML】 formでGETメソッドを指定し、クエリ文字列付きのURLを構成する
categories: [Programming]
tags: [HTML]
---

formから`/search?age=20`みたいなURLにGETメソッドでリクエストを送る方法です。

<!--more-->

# やりたいこと

HTMLのformから、クエリ文字列を含んだURLに対してリクエストを送りたいと思います。

例えばこんな感じ。

| メソッド | URL            |
|----------|----------------|
| GET      | /search?age=20 |

その際、クエリパラメータの値はフォームに入力した値を渡します。

# 方法

formで`method`属性に`get`を指定した場合、`<input>`タグの`name`属性をキー、`value`属性を値として生成されたクエリ文字列がformの`action`で指定したURLの末尾にくっつけられます。

したがって例えば次のようにすればOKです。

```html
<form action="/search" method="get">
  <input type="text" name="age">
  <button class="btn">Send</button>
</form>
```

入力欄に入力した内容が`input`の`value`属性の値となります。

`input`を複数用意すると`/search?age=20&country=Japan`みたいになります。

# まとめ

実は最初、動的にURLを構成しようとしてJavascriptを使おうとしたのですが、HTMLだけでクエリパラメータ付きのURLにリクエストを送信できることがわかったのでメモしておきました。

formの挙動さえわかれば簡単だと思います。

## 参考

- [フォームデータの送信 - ウェブ開発を学ぶ | MDN](https://developer.mozilla.org/ja/docs/Learn/Forms/Sending_and_retrieving_form_data)
- [[HTML] GET メソッドの form で URL のクエリパラメータが消えてしまう](https://qiita.com/QUANON/items/d8528afb37e14e8866b9)

