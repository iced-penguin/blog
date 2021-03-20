---
date: 2020-5-6
updated: 2020-8-28
title: Macにspacemacsを入れてみた
categories: [Programming]
tags: [Spacemacs]
---

spacemacsが何やらすごいらしいという噂を聞き、いてもたってもいられなくなりインストールしてしまいました。

<!--more-->

# 環境

- macOS 10.15.4
- Emacs 26.3
- Spacemacs

# spacemacsとは？

[ドキュメント](https://www.spacemacs.org/doc/DOCUMENTATION)

EmacsとVimを合体させたら最強、というコンセプトで作られたEmacs拡張。EmacsとVimのキーバインドを統合させたEvilなるパッケージ（？）を取り入れています。さらにデフォルトのキーバインドが人間に優しいらしい。

最近vimを使うことがありそのキーバインドも覚えたいと思っていたのですが、しばらくEmacsから離れており久々に触りたいと思っていた僕はこれは好機とばかりにこれに飛びつきました。

コンセプトといい、Evil（邪悪な）というネーミングといい、なんとなく中二病っぽいエディタ。
左手が疼く方はご検討を...。


# インストール

[リポジトリ](https://github.com/syl20bnr/spacemacs#documentation)に簡潔にまとまっています。

Emacsが入っていることを前提にします。

既存の`.emacs.d`をクローンしてきたリポジトリで置き換えるだけ。
ターミナルを開いて以下を実行：
```bash
cd ~
mv .emacs.d .emacs.d.bak  # バックアップ
mv .emacs .emacs.bak
```

これでEmacsを起動するとspacemacsが使えるようになっているはずです。

`master`ブランチは安定版です。開発途上のより先進的な機能を使いたい場合は`develop`ブランチを使用すれば良いようです。

# Dockのロゴを変更する

以下の内容はoptionalです。

僕はGUIでEmacsを使っているので、MacのDockにEmacsを追加しています。が、インストールしただけではDockのアイコンが元のEmacsのままです。これを変えたい。 

まず、[ここ](https://github.com/nashamri/spacemacs-logo)から.icnsファイルをダウンロードしてきます。

次に、Finderを開いて`アプリケーション`からEmacsを見つけ、右クリックから`情報を見る`をクリックします。

左上のアイコンにダウンロードしてきたアイコンファイルをドラッグアンドドロップします。

![](https://drive.google.com/uc?export=view&id=17DHBb36EDhNrGUbCqHHE0Z1wyDtuNJI0)

するとアイコンがspacemacsのものに変化します。

![](https://drive.google.com/uc?export=view&id=1b_u7bYXEtXFhQuMa3x5yodNW_YqbUjQD)

Emacsを再起動すると、Dockのアイコンがspacemacsのものになっています。

# まとめ

- spacemacsはemacs上でvimのUIが使える
- spacemacsをインストールした
- MacのDockのロゴをspacemacsのロゴに変えた

まだ使い始めたばかりで真髄を理解できていないのですが、デフォルトですでにかなり使えるようなので楽しみです。

# 参考

- [How to change a Mac app icon](https://www.idownloadblog.com/2014/07/16/how-to-change-app-icon-mac/)
- [spacemacs](https://github.com/syl20bnr/spacemacs#license)
