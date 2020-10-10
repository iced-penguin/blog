---
date: 2020-6-24
updated: 2020-9-6
title: .ideavimrcが反映されない問題の解決
---

Intellij IDEA のVimエミュレータの設定ファイルが反映されない問題の解決策のメモです。

<!--more-->

# 状況

インサートモードを抜ける時のキーバインドをjjにする以下の設定を`~/.ideavimrc`に加えたが、反映されない。設定を書いた場合即時反映されていたような気がするが...。

```
inoremap jj <ESC>
```

# 解決方法

Intellij上で以下を実行したら反映された。

```
:source path/to/.ideavimrc
```

# まとめ

設定が反映されない時は上の方法を試してみると良いと思います。

## 参考

- [ここ3年ほど使ってみて得られたIdeaVimの知見](https://yukidarake.hateblo.jp/entry/2019/11/18/220737)
