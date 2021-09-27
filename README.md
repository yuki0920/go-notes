# Go製ブログサイト

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

### ユーザー作成

- 記事の作成、編集、削除に必要となる
- コマンド引数にユーザー名、パスワードを渡す

```sh
$ docker-compose run --rm api go run db/seeds/main.go <user_name> <password>
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

---以下、記事用のまとめ---

## Herokuへのデプロイについて

```sh
$ heroku container:login # ログイン
$ heroku create -a app_name # herokuアプリの作成
$ heroku git:remote -a app_name # herokuリポジトリをgit登録
$ heroku addons:add cleardb:ignite # mysqlのアドオンを追加
$ heroku config # CLEARDB_DATABASE_URLが登録されていることを確認
$ heroku config:set DB_USER=b1ff2fbc65bd01 # その他、DB_PASSWORD,DB_PORT,DB_HOST,DB_NAMEも同様
$ heroku stack:set container # heroku.ymlを使う時に必要
$ git push heroku master # デプロイ
```

- `CLEARDB_DATABASE_URL` を DSN の型似合わせるため、ユーザー名などを環境変数として定義する
- デプロイ後に `db/migrations` に移動して、環境変数 `DSN` を利用しマイグレーションを実行する `goose mysql $DSN up`

## 開発環境構築準備

### VS Code用ツール群インストール

1.  `command + shift + P` でコマンドパレットを表示
2. コマンドパレットに go tool と入力して `Go: Install/Update Tools` を選択し、インストール

### 関数の定義ジャンプ

- アプリケーションのソースコードだけでなく標準ライブラリも定義ジャンプできる

### Repl

```sh
git clone https://github.com/x-motemen/gore.git
cd gore
docker build -t gore .
docker run -it --rm gore --autoimport
```

## 言語、ミドルウェア

Dockerfile で用意している

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



## ディレクトリ構成、設計

### 移行前

- db: DB用の設定ファイルやマイグレーションスクリプトを配置する
- model: データベースに保管されたデータをプログラム上においてどのようなデータ構造で利用するかを記述する
- repository: データストアからデータを取得するための処理を記述する。クエリ結果をモデルに(GoのStruct)にマッピングし、モデルを返す
- handler: ハンドラ関数(リクエストを受け取りレスポンスを返す関数)を記述する
- server: リクエストとハンドラのマッピングをする(ルーティングの役割)
- middleware: カスタムのミドルウェアを設定する
- util: ドメインモデルとは関連のないヘルパー関数を配置する

### 移行後

- domain/model: ドメインモデル(構造体)、ドメインに即したメソッドを実装する
- domain/repository: モデルの永続化を行うリポジトリ。インターフェースのみを実装する
- infra/{article,user}: domain層のRepositoryで定義したinterfaceを満たすメソッドを実装する
- infra/sqlhandler: データベースとのコネクション

## リクエストからレスポンスの流れ

### 用語
- マルチプレクサ: `Echo` のインスタンスで、ミドルウェアを設定したりリクエストをハンドラに振り分ける
- ハンドラ: リクエストに対する具体的な処理をする

### 流れ

- サーバーがリクエストを受け取ったら、マルチプレクサがリクエストのURLに応じたハンドラに処理を振り分ける
- ハンドラは引数(`echo.Context`の型を持つ)からリクエストのパラメータやフォームの値をから読み取り、必要に応じてDBアクセスを伴いリソースの作成、更新、削除、返却を担う
  - リソース作成の場合は、
    - リクエストを構造体にバインドする
    - バリデーションを実行する
    - バインドした構造体をリポジトリに渡しDBを更新する
    - レスポンスを返す

## Go言語の文法

- 構造体(struct):
  - 複数の任意の型の値を1つにまとめたもの。typeと組み合わせて新しい型を定義することが多い
  - 値型のため、関数の引数に渡すとコピーが生成され元の構造体に影響を与えない。参照渡しにするためには、構造体へのポインタを渡す
- メソッド:
  - 構造体と手続きを結びつけるためのもの。任意の型に特化した関数を定義する
  - オブジェクト指向においてクラスやインスタンスの手続きとしてのメソッドとは異なる
- 型(type):
  - すでに定義されている型をもとに、新しい型を定義するための機能
- インターフェース:
  - 任意の型がどのようなメソッドを実装するべきかを規定するための枠組み
  - インターフェースに定義するメソッドは外部から参照されることが多い。そのため、大文字始まりのメソッド名となる
  - `interface{}`型は実装すべきメソッドが1つも定義されていないインターフェースのこと。TypeScript的には`any`型
  - 型アサーション: `interface{}`型によって隠蔽されたもとの型を復元する仕組み

```go
// Go組み込みのerror型インターフェフェース(予約後)
type error interface {
  Error() string
}
```

```go
// 構造体
type MyError struct {
  Message string
  ErrCode int
}
// メソッド
func (e * MyError) Error() string {
  return e.Message
}
// インスタンスをstructを返す関数
func RaiseError() error {
  // &でMyErrorのstructのポインタを表す
  return &MyError{Message: "エラーがー発生しました", ErrCode: 500}
}

err := RaiseError()
// 返り値はerror型
err.Error // == "エラーが発生しました"

// 型アサーションによって本来の方を取り出す
e, ok := err.(*MyError)
if ok {
  e.ErrCode // 500
}
```

## JWTについて

### JWTの構成

- ヘッダ、ペイロード、署名からなり、 `ヘッダ.ペイロード.署名` となる
- ヘッダ: 暗号アルゴリズムとトークンタイプ
- ペイロード: JWT発行時刻、識別子
- 署名: 署名なしトークン(`ヘッダ.ペイロード`)を暗号アルゴリズムと秘密鍵によって生成

### 検証

- JWTを署名なしトークンと署名に分離する
- 署名なしトークンを復号化した値と署名を比較する

### 認証での利用

- ユーザー登録時
  - `bcrypt`を利用してパスワードはハッシュ化してDBに保存する
  - APIはユーザーIDと秘密鍵により、トークンを生成し返却する
- ユーザー情報取得時:
  - ブラウザはクッキーにJWTを入れAPIへのリクエスト時に送信する(`Authorization`ヘッダーにJWTを設定し、APIにリクエストする場合もある)
  - APIは秘密鍵を用いてトークンを検証する
  - トークンからユーザーIDを取得し、DBに問い合わせ、返却する
- ログイン時:
  - `bcrypt`を利用して入力値の平文のパスワードとDBに保存されたハッシュ化済みのパスワードを比較し、検証をする
  - ユーザーIDと秘密鍵により、トークンを生成し返却する

## SPAにおけるクッキーの設定
### バックエンド

- Access-Control-Allow-Origin ヘッダーを指定する。ワイルドカードを指定すると失敗する
- Access-Control-Allow-Credentials ヘッダーを true を指定する
- Access-Control-Allow-Methods ヘッダーには許可するHTTPメソッドを指定する
- SetCookie ヘッダーの属性に、
  - SameSite に `none` を指定する
  - Path に `/` を指定する
  - Secure に `true` を指定する
  - HttpOnly に `true` を指定する

### フロントエンド

- withCredentials でリクエスト時にクッキーをサーバーに送信する
