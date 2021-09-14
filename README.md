# tic-tac-toe

Go が好きなので、Go で CLI の 〇✕ ゲームを作りました (ただの変態)。  
外部パッケージを使わず、標準パッケージのみでの実装です。

## インストール方法

### 直接ダウンロード

[Releases](https://github.com/logica0419/tic-tac-toe/releases) から直でDL・解凍・実行してください。

### Goを使ってインストール

前提: go

```sh
go install github.com/logica0419/tic-tac-toe@latest
```

してください。`Tic-Tac-Toe`コマンドでいつでも遊べます。

### リポジトリをクローン

環境を汚したくない人向け。

前提: go & git  
適当なディレクトリで

```sh
git clone https://github.com/logica0419/tic-tac-toe
cd tic-tac-toe
go run *.go
```

してください。  

## 遊び方

実行すると案内が出るので、それに従ってやれば大丈夫です。楽しんで。

## 敵 AI の裏話

難易度設定とそれぞれのアルゴリズムのお話です。
もっと強くするアイデア等あればご一報ください。

- Easy
  - 選べるセルの中から完全ランダムで選ぶ
- Normal
  - そのターンで勝てる手が存在する (2つ並びがある) 時、勝てる手を出す
  - 無ければ優先して真ん中を取りに行く、埋まっていたらランダム
- Hard
  - 次のターンで負けそうな (○ の 2つ並びがある) 時、それを阻止する手を出す
  - 無ければ Normal と同じ手順をたどる
