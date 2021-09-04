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
