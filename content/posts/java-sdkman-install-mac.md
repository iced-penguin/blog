---
title: MacにSDKMANでJavaをインストールする
date: 2020-09-11 22:31:47
---

Javaのインストール方法がいくつもあって迷いがちなのでメモ。今回はSDKMANでMacにJavaを入れます。

<!--more-->

## 環境

- MacOS 10.15.6
- SDKMAN 

# SDKMAN?

SDKMANはJavaやJava関連のツールのインストールやバージョン管理を行います。
Javaの他にKotlinなどのJVM言語やGradleなどのビルドツールを取り扱っています。

面倒なJavaのインストールやバージョンの切替を簡単にしてくれるものです。

# SDKMANでJavaをインストール　

## SDKMANをインストール

ターミナルを開いて以下のコマンドを実行：

```sh
curl -s "https://get.sdkman.io" | bash
source "$HOME/.sdkman/bin/sdkman-init.sh"
```

`sdk --version`でバージョン情報が表示されればOKです。

## SDKMANでJavaをインストール

インストール可能なJavaのバージョンを確認します。

```sh
sdk list java
```

こんな感じで表示されます。

```sh
================================================================================
Available Java Versions
================================================================================
 Vendor        | Use | Version      | Dist    | Status     | Identifier
--------------------------------------------------------------------------------
 AdoptOpenJDK  |     | 14.0.2.j9    | adpt    |            | 14.0.2.j9-adpt
               |     | 14.0.2.hs    | adpt    |            | 14.0.2.hs-adpt
               |     | 13.0.2.j9    | adpt    |            | 13.0.2.j9-adpt
               |     | 13.0.2.hs    | adpt    |            | 13.0.2.hs-adpt
               |     | 12.0.2.j9    | adpt    |            | 12.0.2.j9-adpt
               |     | 12.0.2.hs    | adpt    |            | 12.0.2.hs-adpt
               |     | 11.0.8.j9    | adpt    |            | 11.0.8.j9-adpt
               |     | 11.0.8.hs    | adpt    |            | 11.0.8.hs-adpt
               |     | 8.0.265.j9   | adpt    |            | 8.0.265.j9-adpt
               |     | 8.0.265.hs   | adpt    |            | 8.0.265.hs-adpt
 Amazon        |     | 11.0.8       | amzn    |            | 11.0.8-amzn
               |     | 8.0.265      | amzn    |            | 8.0.265-amzn
 Azul Zulu     |     | 14.0.2       | zulu    |            | 14.0.2-zulu
               |     | 13.0.4       | zulu    |            | 13.0.4-zulu
               |     | 13.0.4.fx    | zulu    |            | 13.0.4.fx-zulu

# 以下略
```

vendorにAdoptOpenJDKだのAmazonだのと書かれていますが、これはそのJavaの提供元を指します。

今回は`AdoptOpenJDK`のJavaをインストールします。バージョンは長期サポートが付いている`11`とします。
`hs`と`j9`はそれぞれHotSpot、OpenJ9の略で違いはJVMの実装ですが、`hs`の方が一般的らしいので、長いものには巻かれろということでこちらを使います。

インストールするJavaが決まったので、次のコマンドでインストールします。

```sh
sdk install java 11.0.8.hs-adpt
```

バージョン確認すると今入れたJavaがデフォルトになっています。

```sh
$ java -version
openjdk version "11.0.8" 2020-07-14
OpenJDK Runtime Environment AdoptOpenJDK (build 11.0.8+10)
OpenJDK 64-Bit Server VM AdoptOpenJDK (build 11.0.8+10, mixed mode)
```

他のバージョンを入れたいときは上と同様にインストールし、デフォルトのバージョン切替は

```sh
sdk default java 11.0.8.hs-adpt
```

のようにします。

# まとめ

SDKMANを使ってJavaのインストールやバージョン管理を簡単に行うことができます。

## 参考

- [Installation - SDKMAN! the Software Development Kit Manager]*https://sdkman.io/install)
- [Part 1: OpenJ9 versus HotSpot](https://www.royvanrijn.com/blog/2018/05/openj9-jvm-shootout/)
