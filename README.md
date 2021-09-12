# Go製ブログサイト

## 環境構築

### VS Code用ツール群インストール

1.  `command + shift + P` でコマンドパレットを表示
2. コマンドパレットに go tool と入力して `Go: Install/Update Tools` を選択し、インストール

## 言語、ミドルウェア

Dockerfile で用意している

```sh
$ go version
go version go1.16.1 darwin/amd64
$ mysql --version
mysql  Ver 14.14 Distrib 5.6.51, for osx10.16 (x86_64) using  EditLine wrapper
```

## マイグレーションツール

```
$ go get -u github.com/pressly/goose/v3/cmd/goose
$ which goose
/Users/yuki.watanabe/.goenv/shims/goose
```

## ホットリロードツール

```sh
$ go get -u github.com/cosmtrek/air
$ which air
/Users/yuki.watanabe/.goenv/shims/air
# 初期化
$ air init
```

## パッケージ管理


### 初期化

```sh
$  go mod init $(git config user.name)/$(basename `pwd`)
```

### 依存管理


参照しているパケージを整理する

```sh
$  go mod tidy
```

- ソースコードを検査して、どのような外部パッケージを利用しているかを判定する
- ソースコード内で利用されている外部パッケージは go.mod と go.sum というファイルに書き出される
- 直接的に利用しているパッケージは go.mod に、間接的に利用しているパッケージは go.sum に記載される
- indirect というコメントは、直接依存しているモジュールではないことを表現している

### ダウンロード

```sh
$ go mod download
```

- ダウンロードされた外部パッケージのソースコードは `$HOME/go/pkg/mod/` に配置される

## 起動

### マイグレーション

```sh
# コンテナに入る
$ docker-compose run --rm api bash

# 接続確認
$ goose mysql $DSN status

# 適用
cd migrations
$ goose mysql $DSN up

# ロールバック
$ cd migrations
$ goose mysql $DSN down
```

- マイグレーションを適用する際は、 `migrations` まで移動する必要がある
- 引数に`$DSN`(`"user:password@tcp(db:3306)/go_blog?charset=utf8&parseTime=true&loc=Asia%2FTokyo"`)が必要となる

### アプリケーション起動

```sh
# air コマンドでホットリロードで起動
$ docker-compose up
```

### テスト実行

```sh
$ docker-compose run --rm api go test ./handler
```

- テスト対象のディレクトリか関数を指定する

### OpenAPI Generatorによる型生成

```
$ make generate-api
```

## ディレクトリ構成

- db: DB用の設定ファイルやマイグレーションスクリプトを配置する
- model: データベースに保管されたデータをプログラム上においてどのようなデータ構造で利用するかを記述する
- repository: データストアからデータを取得するための処理を記述する
- handler: ハンドラ関数(リクエストを受け取りレスポンスを返す関数)を記述する


## リクエストからレスポンスの流れ

### 用語
- マルチプレクサ: `Echo` のインスタンスで、ミドルウェアを設定したりリクエストをハンドラに振り分ける
- ハンドラ: リクエストに対する具体的な処理をする

### 流れ

- サーバーがリクエストを受け取ったら、マルチプレクサがリクエストのURLに応じたハンドラに処理を振り分ける
- ハンドラは引数(`echo.Context`の型を持つ)からリクエストのパラメータやフォームの値をから読み取り、必要に応じてDBアクセスを伴いリソースの作成、更新、削除、返却を担う
