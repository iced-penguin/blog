---
title: "[Go] Markdownのメタデータ（YAML記法）を取得"
date: 2021-08-22T19:00:12+09:00
categories: [Programming]
tags: [Go]
toc: false
draft: false
---

GoでMarkdownファイルを読み込み、メタデータ（YAML記法）を取得する方法

<!--more-->

# Markdownメタデータ

Markdownにメタデータを含める記法はいくつかあるようだが、ここではYAMLの場合を考える。

次のようなイメージ。セパレータ `---` で囲まれた部分にYAML記法でメタデータを書く。メタデータは文書の先頭に配置する。

`sample.md`

```text
---
title: Awesome title
date: 20210822
categories: [go, markdown]
---

# Sample

sample
```

# GoでMarkdownメタデータを読み込む

YAMLをパースするために[yaml.v2](https://github.com/go-yaml/yaml)を使う。以下でインストール。

```text
go get gopkg.in/yaml.v2
```

実際に `sample.md` からメタデータを取得して表示してみる。

`main.go`
```
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	separator = "---"
)

type Metadata struct {
	Title      string   `yaml:"title"`
	Date       string   `yaml:"date"`
	Categories []string `yaml:"categories"`
}

func main() {
	f, err := os.Open("sample.md")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", os.Args[0], err)
		os.Exit(-1)
	}
	defer f.Close()

	var frontMatter []string
	sepCount := 0

	// １つ目のセパレータから２つ目のセパレータまでを読み込む
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == separator {
			if sepCount > 0 {
				break
			}
			sepCount += 1
			continue
		}
		frontMatter = append(frontMatter, line)
	}

	// パース
	buf := []byte(strings.Join(frontMatter, "\n"))
	metadata := Metadata{}
	err = yaml.Unmarshal(buf, &metadata)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", os.Args[0], err)
		os.Exit(-1)
	}

	fmt.Println("title: ", metadata.Title)
	fmt.Println("date: ", metadata.Date)
	fmt.Println("categories: ", metadata.Categories)
}
```

実行結果

```text
$ go run main.go
title:  Awesome title
date:  20210822
categories:  [go markdown]
```

markdownファイルを１行づつ読み込み、１つ目のセパレータの次の行から２つ目のセパレータが現れるまでを文字列として取得し、あらかじめ定義しておいた構造体にパースするという流れ。

# 参考

- [Markdownのメタデータ記法について調べてみた](https://ytyaru.hatenablog.com/entry/2020/04/02/000000)
- [YAML support for the Go language](https://github.com/go-yaml/yaml)