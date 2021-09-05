# Go製ブログサイト

## 環境構築

### ツール群インストール

1.  `command + shift + P` でコマンドパレットを表示
2. コマンドパレットに go tool と入力して `Go: Install/Update Tools` を選択し、インストール

## 言語、ミドルウェア

```sh
$ go version
go version go1.16.1 darwin/amd64
$ mysql --version
mysql  Ver 14.14 Distrib 5.6.51, for osx10.16 (x86_64) using  EditLine wrapper
```

## パッケージ管理


### 初期化

```sh
go mod init $(git config user.name)/$(basename `pwd`)
```

### 依存管理

```sh
go mod tidy
```

- ソースコードを検査して、どのような外部パッケージを利用しているかを判定する
- ソースコード内で利用されている外部パッケージは go.mod と go.sum というファイルに書き出される
- 直接的に利用しているパッケージは go.mod に、間接的に利用しているパッケージは go.sum に記載される

### ダウンロード

```sh
go mod download
```

- ダウンロードされた外部パッケージのソースコードは `$HOME/go/pkg/mod/` に配置される

## 起動

### MySQL Server 起動

```sh
# 起動
$ mysql.server start
Starting MySQL
.. SUCCESS!

# ステータス確認
$ mysql.server status
 SUCCESS! MySQL running (XXXXX)

#  停止
$ mysql.server stop
Shutting down MySQL
.. SUCCESS!

# ログイン
$ mysql -u root -p
Enter password: # 何も入力せずEnter
mysql> #プロンプト出現したら成功

# ログアウト
mysql> \q
```
