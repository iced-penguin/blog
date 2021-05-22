---
title: "[Go] XMLの全ての属性を取得"
date: 2021-05-22T17:00:06+09:00
categories: [Programming]
tags: [Go]
toc: false
draft: false
---


# 概要

Goで XML を読み込むときに全ての属性を取得する方法

<!--more-->

# 環境

- Go 1.16

# 方法

## Go で XML を扱う

Go で XML を取り扱うには `encoding/xml` パッケージを用いる。

基本的には自前で構造体を定義し、XMLをパースしてそれにマッピングする流れになる。

`encoding/xml` パッケージの全般的な使い方は本記事の主題から外れるので、良さそうな記事のリンクを貼っておくにとどめる。

基本↓

[GoでXMLをパースする - Qiita](https://qiita.com/ytkhs/items/948f516ec82c82eaa882)

もっと詳しい↓

[Goのencoding/xmlを使いこなす - Qiita](https://qiita.com/ono_matope/items/70080cc33b75152c5c2a)

全てが分かる↓

[xml - The Go Programming Language](https://golang.org/pkg/encoding/xml/)

# XML の全属性を取得する

必ずしも XML の構造がわかるわけではない（非構造化データ）場合などに、とりあえず全ての属性を取得しておきたいとする。

その場合は次のように構造体を定義する。

```
type XML {
	Attrs []xml.Attr `xml:",any,attr"`
}
```

できれば map で取りたいが、構造体のフィールドの型にmap は指定を指定するとパースできない。必要なら自前で変換するしかなさそう。

簡単な例を実装したので、以下にコードを貼り付けておく。

```text
$ tree
.
├── main.go
└── sample.xml
```

`sample.xml`

```
<sample attr1="attribute1" attr2="attribute2" attr3="attribute3">
  <tag1>Tag1</tag1>
  <tag2>Tag2</tag2>
</sample>
```

`main.go`

```
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Sample struct {
	XMLName xml.Name   `xml:"sample"`
	Attrs   []xml.Attr `xml:",any,attr"`  // 全ての属性
	Tag1    string     `xml:"tag1"`
	Tag2    string     `xml:"tag2"`
}

func main() {
	f, err := os.Open("sample.xml")
	if err != nil {
		fmt.Printf("cannot open file: %v", err)
		return
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("cannot read file: %v", err)
		return
	}

	sample := Sample{}

	// XMLをパース
	xml.Unmarshal(buf, &sample)

	fmt.Println(sample)
}
```

# 参考

[Golang - unmarshal extra XML attributes](https://stackoverflow.com/questions/25530780/golang-unmarshal-extra-xml-attributes/49950640)