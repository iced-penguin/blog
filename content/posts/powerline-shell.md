---
date: 2020-6-14
updated: 2020-8-28
title: Powerline-Shellを使ってみる
categories: [Programming]
tags: [iTerm2, Powerline-Shell]
---

時々見かけるなんだかおしゃれな鉛筆みたいなプロンプトを取り入れたかった+カレントディレクトリの短縮表示をしたかったので、Powerline-Shellを導入してみました。

<!--more-->

# 環境

- Mac OS 10.15.5
- iterm2 3.3.1
- zsh 5.7.1

僕はiterm2を使っていますが、Mac標準のターミナルでも同じように導入できるんじゃないかと思います。

# 導入手順

## Powerline Fonts を導入

Powerline fonts が必要なのでそれをインストールします。

```
# リポジトリをクローン
git clone https://github.com/powerline/fonts.git --depth=1
# インストール
cd fonts
./install.sh
```

iterm2でフォントを変更します。

`Command ,`で`Preferences`を開き、`Profiles > Text > Font`でフォントを`○○ powerline`のような名前のものに変更します。powerline用のフォントはいろいろ出てきますが、今回は`source code pro for powerline`にしました。

## Powerline Shell を導入

```
git clone https://github.com/b-ryan/powerline-shell
cd powerline-shell
python setup.py install
```

どうもpython3を使わないとうまく行かないみたい。

次に`~/.zshrc`に以下を追加：

```sh
function powerline_precmd() {
    PS1="$(powerline-shell --shell zsh $?)"
}

function install_powerline_precmd() {
  for s in "${precmd_functions[@]}"; do
    if [ "$s" = "powerline_precmd" ]; then
      return
    fi
  done
  precmd_functions+=(powerline_precmd)
}

if [ "$TERM" != "linux" ]; then
    install_powerline_precmd
fi
```

反映：

```
source ~/.zshrc
```

これで例の鉛筆みたいなプロンプトになっているはず。

# カスタマイズ

今回は表示項目と色を変更し、プロンプトに改行を入れてみます。

修正後はこんな感じになりました。

![](https://drive.google.com/uc?export=view&id=1f2J_0j50ThaJCk_m_0Sjg1tqISgJ7SP-)

## 設定ファイルを作成

まず設定ファイルを適切な場所に作成します。

```
mkdir -p ~/.config/powerline-shell
powerline-shell --generate-config > ~/.config/powerline-shell/config.json
```

これでデフォルトの設定ファイルが生成されました。
Powerline Shell をカスタマイズするには主にこれを編集します。

## 表示項目の変更

先ほど生成された`config.json`には表示項目（segments）が書かれています。
ユーザ名（username）、ホスト名（hostname）、右端の`%`（root）を表示項目から外します。

```
{
  "segments": [
    "virtual_env",
    "ssh",
    "cwd",
    "git",
    "hg",
    "jobs"
  ],
}
```

## 改行を入れる

`.zshrc`を編集します。
Powerline Shell インストール時に書き込んだ内容を修正します。改行を入れた箇所でそのまま改行がなされます。

```sh
function powerline_precmd() {
# customized prompt
    PS1="
$(powerline-shell --shell zsh $?)
$ "
}
```

## 色を調整

クローンしたリポジトリ以下に`powerline-shell/powerline_shell/themes/`のようなディレクトリがあり、そこにカラーテーマが入っています。その中に`default.py`があり、基本的にはこれを適当な場所にコピーしていじることになります。

```
cp path/to/default.py path/to/my_theme.py
```

テーマファイルのパスを`config.json`に追記します。

```
{
  "segments": [
    "virtual_env",
    "ssh",
    "cwd",
    "git",
    "hg",
    "jobs"
  ],
  "theme": "path/to/my_theme.py"
}
```

例えばホームディレクトリの部分の色を変えたいとすると、

```
HOME_BG = 116  # nordic cyan
HOME_FG = 0  # white
```

のように編集します。
色の指定には8ビットカラーが使えます。

[256色一覧](https://jonasjacek.github.io/colors/)

現在の僕のカラーテーマのファイルは[こちら](https://github.com/swimpenguin/dotfiles/blob/master/powerline-shell/themes/nord_like.py)。

iterm2にNordというテーマを入れているので、それに合わせた配色にしています。

# まとめ

- Powerline Shell を導入することでおしゃれかつ機能的なプロンプトを比較的簡単に作れる
- Powerline Shell をカスタマイズする方法を紹介

## 参考

- [Powerline-Shell - コマンドラインを美しく便利にパワーアップ！](https://linuxfan.info/powerline-shell)
- [powerline/fonts: Patched fonts for Powerline users. - GitHub](https://github.com/powerline/fonts)
- [b-ryan/powerline-shell: A beautiful and useful prompt for your shell](https://github.com/b-ryan/powerline-shell)
