#!/bin/bash
# $1: file name, $2: category`

if [ $# -ne 2 ]; then
  echo "usage: $(basename $0) file_name category"
  exit 1
fi

# プロジェクトルートで実行
cd $(dirname $0); cd ..

# カテゴリー名が不正でなければファイル作成
if [ -n "$(cat ./scripts/categories.txt | grep ^$2$)" ]; then
  hugo new posts/$1.md
  sed -I '' "s/categories: \[\]/categories: \[$2\]/" ./content/posts/$1.md
else
  echo 'Invalid category name'
fi
