---
date: 2020-08-18
title: Apache Commons IO をGradleプロジェクトに導入
categories: [Programming]
tags: [Gradle]
---

Apache Commons IO をGradleプロジェクトで使えるようにする

<!--more-->

# Apache Commons IO？

Apache Commons 含まれるJavaのIOライブラリ。ファイルの処理などに便利。

[Commons IO](https://commons.apache.org/proper/commons-io/)

# 方法

`build.gradle`の依存関係にApache Commons IO を追加します。

```
dependencies {
    implementation 'commons-io:commons-io:2.7'
}
```

次にGradleプロジェクトをリフレッシュすれば依存関係が解決され、Apache Commons IO が使用できるようになります。

## 参考

- [Dependency Information](https://commons.apache.org/proper/commons-io/dependency-info.html)
