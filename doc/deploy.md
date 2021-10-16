# Herokuへのデプロイについて

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
