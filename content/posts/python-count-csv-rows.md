---
title: "CSVファイルの行数をカウントするPythonスクリプトを書いた"
date: 2020-12-05T16:46:46+09:00
draft: false
---

普通は`wc -l`などとすると思うが、値の中に改行が含まれていたりしてunixコマンドでやるには面倒臭そうな行数カウントをPythonでやってみた。

<!--more-->

# スクリプト

## 使い方

コマンドライン引数としてCSVファイルのパスを与えると、その行数を表示する。

```python
python count_csv_rows.py sample.csv
```

## コード

`count_csv_rows.py`：

```python
import sys
import csv

def count_rows(file_path):
    with open(file_path) as f:
        reader = csv.reader(f)
        rows = [row for row in reader]
    return len(rows)

if __name__ == '__main__':
    args = sys.argv

    if (len(args) != 2):
        print('usage: python count_csv_rows.py path/to/your/file')
        sys.exit(1)

    file_path = args[1]
    print(count_rows(file_path))
```

処理としては非常に簡単で、2次元配列としてCSVを読み込み、行数を数えるだけ。

# まとめ

弱弱なのでシェルスクリプトを使いこなせず、Pythonに頼った。

`sed`や`awk`を駆使すればできるんだろうか。多分できるんだろうな。
