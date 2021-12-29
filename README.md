# Go製学習ノート

## 言語・ツール

### 言語

Docker 環境を用意している

```sh
$ go version
go version go1.16.1 darwin/amd64
$ mysql --version
mysql  Ver 8.0.26 for Linux on x86_64 (MySQL Community Server - GPL)
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

## 起動

### マイグレーション

```sh
$ make db-migrate
```

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

### DB接続

```
docker exec -it <container_id> sh
mysql -u user -p
```

### ユーザー作成

- 記事の作成、編集、削除に必要となる
- コマンド引数にユーザー名、パスワードを渡す

```sh
$ docker-compose run --rm api go run infra/seeds/main.go <user_name> <password>
```

### テスト実行

```sh
# ./... で階層指定
$ docker-compose run --rm api go test ./... -v
```

- テスト対象のディレクトリか関数を指定する

### OpenAPI Generatorによる型生成

```
$ make generate-api
```

## ディレクトリ構成

クリーンアーキテクチャーをベースにした構成にしている

- handler
  - router: リクエストとハンドラのマッピングをする(ルーティングの役割)
  - handler: ハンドラ関数(リクエストを受け取りレスポンスを返す関数)を記述する
- infra
  - sqlhandler: DBとのコネクションを確立する
  - <model_name>: データストアからデータを取得するための処理を記述する。クエリ結果をモデルに(GoのStruct)にマッピングする。domain層のRepositoryで定義したinterfaceを満たすメソッドを持つ構造体を返す。
  - その他: マイグレーションスクリプトを配置する
- usecase: handlerから呼び出される。ドメイン層で定義されたメソッドを、リポジトリを使って呼び出すためのユースケース記述レベルの抽象度の高いコードを記載する。
- domain:
  - model: ドメインモデルを定義する。データベースに保管されたカラムをプログラム上においてどのようなデータ構造で利用するかを記述する。モデルに紐付いたビジネスロジックも対象となる。
  - repository: ドメインモデルのの永続化、再構築を担うためのインターフェースを定義する
- middleware: カスタムのミドルウェアを設定する
- util: ドメインモデルとは関連のないヘルパー関数を配置する
