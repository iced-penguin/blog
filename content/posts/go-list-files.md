---
title: "[Go] 指定ディレクトリのファイル一覧を取得する"
date: 2021-10-02T11:44:28+09:00
categories: [Programming]
tags: [Go]
toc: false
draft: false
---

指定したディレクトリ内の全ファイルを、サブディレクトリ内のものも含めて取得する。再起的にディレクトリを走査する関数を自前で実装する方法もあるが、今回は`filepath'パッケージを利用する。

<!--more-->

# ファイル一覧を取得

## 実装

`filepath.Walk`を使う。第2引数の関数に各ファイルに対して実行する処理を記述する。
今回はディレクトリは除いてリストアップすることにする。

`main.go`

```
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "."
	files, err := listFiles(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

func listFiles(root string) ([]string, error) {
	var files []string
	// root以下を走査
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ディレクトリは除く
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
```

上の実装ではカレントディレクトリ以下のファイル一覧を表示している。

## 実行してみる

試しに実行してみるとこんな感じ。

```text
$ tree
.
├── main.go
└── sample
    ├── sample.txt
    └── sub_sample
        └── sub_sample.txt
```

```text
$ go run main.go
main.go
sample/sample.txt
sample/sub_sample/sub_sample.txt
```

サブディレクトリのファイルも含めて、カレントディレクトリ以下の全ファイルを取得できている。