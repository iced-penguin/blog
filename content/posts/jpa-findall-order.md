---
title: JpaRepository#findAllを特定のフィールドでソートしたい
date: 2020-05-20 21:47:59
---

SpringBootでJPAを使ってデータベースアクセスを実現することが多い。
今回はJpaRepositoryの全件取得のソートを、メソッド自動生成機能を用いて実現する。

<!--more-->


JpaRepositoryではリポジトリに適切な名前のメソッド名を追加することで、そのメソッドの処理を書かなくても自動的に生成してくれる。

`@Repository`を付与したインターフェースをメソッドを追加する。

# 方法

結論としては、次のようなメソッド名になる。

```
findAllByOrderBy○○();
```

○○にはソートのキーとなるフィールド名を入れる。例えば`findAllByOrderById()`など。

なお、Allを入れなくても次のメソッドでも動作するようだ。

```
findByOrderBy○○();
```

## 降順にする

降順にしたい場合は`Desc`を付ける。

```
findAllByOrderBy○○Desc();
```

## 参考

- [Spring Data JPA のfind+OrderByで、No property desc foundエラー時の対処法](https://ti-tomo-knowledge.hatenablog.com/entry/2018/10/18/095121)
