---
date: 2020-5-16
title: "[Kotlin] コレクション"
categories: [Programming]
tags: [Kotlin]
---

Kotlinのコレクションについてのメモ。

<!--more-->

# コレクション

次のインターフェースがある。

|        | 読み取り専用 | 書き込み可  |
| --- | --- | --- |
| リスト | List         | MutableList |
| セット | Set          | MutableSet  |
| マップ | Map          | MutableMap  |

`List`, `Set`, `Map`は生成した後に変更しない場合に用いるが、厳密にはイミュータブルではないらしい。
これについては最後にまとめる。

以下のコードは全てREPLで実行した。

## リスト

要素を順序つきで並べる。

### List

読み取り専用のリストを提供するインターフェース

```kotlin
>>> val x: List<Int> = listOf(1, 2, 3)
>>> x
res2: kotlin.collections.List<kotlin.Int> = [1, 2, 3]
>>> x.size
res3: kotlin.Int = 3
>>> x[0]
res4: kotlin.Int = 1
```

### MutableList

書き込み可のリストを提供するインターフェース

```kotlin
>>> val x: MutableList<Int> = mutableListOf(1, 2, 3)
>>> x.add(4)
res2: kotlin.Boolean = true
>>> x
res3: kotlin.collections.MutableList<kotlin.Int> = [1, 2, 3, 4]
>>> x.remove(4)
res4: kotlin.Boolean = true
>>> x
res5: kotlin.collections.MutableList<kotlin.Int> = [1, 2, 3]
>>> x.removeAt(2)
res6: kotlin.Int = 3
```

## セット

重複しない要素を順序なしで保持する。数学の集合のイメージ。

順序を持たないのでリストのようにインデックスを用いてアクセスすることはできない。

### Set

読み取り専用のセットを提供するインターフェース

```kotlin
>>> val x: Set<Int> = setOf(1, 1, 2, 3)
>>> x
res1: kotlin.collections.Set<kotlin.Int> = [1, 2, 3]
>>> x.size
res2: kotlin.Int = 3
```

### MutableSet

書き込み専用のセットを提供するインターフェース

```kotlin
>>> val x: MutableSet<Int> = mutableSetOf(1, 1, 2, 3)
>>> x
res1: kotlin.collections.MutableSet<kotlin.Int> = [1, 2, 3]
>>> x.add(4)
res2: kotlin.Boolean = true
>>> x
res3: kotlin.collections.MutableSet<kotlin.Int> = [1, 2, 3, 4]
>>> x.remove(1)
res4: kotlin.Boolean = true
>>> x
res5: kotlin.collections.MutableSet<kotlin.Int> = [2, 3, 4]
```

## マップ

キーと値のペアを要素として保持する。キーと値は`Pair`型で表現する。
Pairオブジェクトは`key to value`の形で得られる。

### Map

読み取り専用のマップを提供するインターフェース

```kotlin
>>> val x: Map<String, String> = mapOf("id" to "1", "name" to "John Doe")
>>> x
res1: kotlin.collections.Map<kotlin.String, kotlin.String> = {id=1, name=John Doe}
>>> x["name"]
res2: kotlin.String? = John Doe
```

### MutableMap

書き込み専用のマップを提供するインターフェース

```kotlin
>>> val x: MutableMap<String, String> = mutableMapOf("id" to "1", "name" to "John Doe")
>>> x
res1: kotlin.collections.MutableMap<kotlin.String, kotlin.String> = {id=1, name=John Doe}
>>> x["age"] = "24"
>>> x
res2: kotlin.collections.MutableMap<kotlin.String, kotlin.String> = {id=1, name=John Doe, age=24}
```

## イミュータブルと読み取り専用の違い

イミュータブルとは変更不可を意味するらしい。
リストを例にとると、Listは実はイミュータブルではない。これはMutableListがListを継承していることによりに次のような操作ができてしまうため。

```kotlin
>>> val mutableList: MutableList<Int> = mutableListOf(1, 2, 3)
>>> val list: List<Int> = mutableList
>>> mutableList.add(4)
res1: kotlin.Boolean = true
>>> list
res2: kotlin.collections.List<kotlin.Int> = [1, 2, 3, 4]
```

つまり、MutableListオブジェクトを代入することでListオブジェクトを作成することができ、MutableListオブジェクトを変更することでListオブジェクトも変更されてしまう。

# まとめ

- Kotlinのコレクションにはリスト、セット、マップがあり、それぞれ読み取り専用と書き込み可のインターフェースが用意されている
- 読み取り専用はイミュータブルとは異なり、変更されうる

## 参考

- 長澤太郎. Kotlinスタートブック - 新しいAndroidプログラミング. 株式会社リックテレコム. 2018.

