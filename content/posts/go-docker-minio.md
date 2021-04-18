---
title: "AWS SDK for Go を利用してDockerコンテナからMinIoにアクセス"
date: 2021-04-16T19:57:29+09:00
categories: [Programming]
tags: [Go, Docker, AWS]
toc: false
draft: false
---

# 概要

AWS SDK for Go を利用してMinIOへのアクセスを試してみる。

環境構築はDockerで行う。Dockerで立てたコンテナ内のアプリケーションからMinIOに接続するときにハマったのでメモを残しておく。

<!--more-->

# MinIO?

Amazon S3 と互換性を持つオブジェクトストレージ。

ローカルにあるS3のような感覚で使える。

# 実装

Dockerコンテナ上でMinIOを動かす。

MinIOにアップロードされているファイル（オブジェクト）の情報を取得するAPIを作ってみる。

| HTTPメソッド | URL           | 説明         |
|--------------|---------------|--------------|
| GET          | /api/v1/files | ファイル取得 |

## 使用技術

- Go
- Echo（Goのフレームワーク）
- Docker
- MinIO
- AWS SDK for Go

## ファイル一覧

最終的に以下のような構成に。

```text
$ tree
.
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

わざわざ表示するまでもないかな...。

## 初期化・必要なライブラリをインストール

```text
$ mkdir go-minio-sample

$ cd go-minio-sample

# Modules 初期化
$ go mod init github.com/swimpenguin0504/go-minio-sample

# Echo
$ go get github.com/labstack/echo

# MinIO Go Client SDK
$ go get github.com/minio/minio-go

# AWS SDK for Go
$ go get github.com/aws/aws-sdk-go
```

## コード

`Dockerfile`

```
FROM golang:1.16

WORKDIR /app
```

`docker-compose.yml`

```
version: '3'
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
    environment:
      # minio で設定しているクレデンシャル情報と同じ値
      AWS_ACCESS_KEY_ID: hogehoge
      AWS_SECRET_ACCESS_KEY: fugafuga
    command: go run main.go
    tty: true
    volumes:
      - ".:/app"
    ports:
      # 他のアプリケーションと衝突しないように適当なポートを設定
      - "5000:80"
  minio:
    image: minio/minio
    environment:
      # クレデンシャル情報
      MINIO_ACCESS_KEY: hogehoge
      MINIO_SECRET_KEY: fugafuga
    # コンテナの /data 以下にデータが格納される
    command: server /data
    volumes:
      - "s3-data:/data"
    ports:
      # 他のアプリケーションと衝突しないように適当なポートを設定
      - "9090:9000"
volumes:
  s3-data:
```

`main.go`

```
package main

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo"
)

// レスポンス用
type Response struct {
	Objects []ObjInfo `json:"objects"`
}

// レスポンス用
type ObjInfo struct {
	Key          string    `json:"key"`
	LastModified time.Time `json:"lastModified"`
	Size         int64     `json:"size"`
}

func main() {
	e := echo.New()

	e.GET("/api/v1/files", getObjects)

	e.Logger.Fatal(e.Start(":80"))
}

// 特定バケット配下のオブジェクト一覧
func getObjects(c echo.Context) error {
	sess := createSession()
	svc := s3.New(sess)
	bucket := c.QueryParam("bucket")

	// オブジェクト取得
	res, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})
	// 本来はもっときちんとエラーハンドリングした方が良いが、簡単のため今回はこれで良しとする
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var objects []ObjInfo
	// オブジェクト情報をJSONにつめて返す
	for _, content := range res.Contents {
		objects = append(
			objects,
			ObjInfo{Key: *content.Key, LastModified: *content.LastModified, Size: *content.Size},
		)
	}
	return c.JSON(http.StatusOK, objects)
}

// セッションを返す
func createSession() *session.Session {
	// 特に設定しなくても環境変数にセットしたクレデンシャル情報を利用して接続してくれる
	cfg := aws.Config{
		Region:           aws.String("ap-northeast-1"),
		Endpoint:         aws.String("http://minio:9000"), // コンテナ内からアクセスする場合はホストをサービス名で指定
		S3ForcePathStyle: aws.Bool(true),                  // ローカルで動かす場合は必須
	}
	return session.Must(session.NewSession(&cfg))
}
```

ポイントは createSession() 内の処理で、

- `S3ForcePathStyle` を true にする
- エンドポイントの指定でホスト名を`minio`（`docker-compose.yml` のサービス名）とする

後者で少し詰まった...。

コンテナ間で通信をする際にサービス名を指定することで名前解決しているらしい。

# 動かしてみる

## MinIO上にバケットとファイルを用意

`http://localhost:9090` にアクセスするとMinIOの画面が開く。アクセスキーとシークレットキーを要求されるので、それぞれに `docker-compose.yml` で指定した `MINIO_ACCESS_KEY` 、`MINIO_SECRET_KEY` の値を入力すれば入れる。

![](https://drive.google.com/uc?export=view&id=1ofw5PgohS-E2X1dmdoSnUNaw7tGkZz-L)

右下の＋ボタンからバケットを作成できるので、予め作成しておく。また、そこに適当なファイルをアップロードしておく。

## ファイル一覧を取得

APIを叩く。

```text
$ curl http://localhost:5000/api/v1/files?bucket=sample
[{"key":"sample.txt","lastModified":"2021-04-16T08:49:41.941Z","size":15},{"key":"sample2.txt","lastModified":"2021-04-16T08:49:41.941Z","size":15}]
```

オブジェクト情報を取得できていることを確認。

# まとめ

MinIOに接続するときに気をつけること：

- `S3ForcePathStyle` を指定する
- コンテナから接続する場合はホスト名を `docker-compose.yml` で指定した MinIO のサービス名とする

# 参考

[AWS SDK for Goからminioへのアクセスでハマったこと | DevelopersIO](https://dev.classmethod.jp/articles/access-minio-using-aws-sdk-for-go/)
