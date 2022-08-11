## db migrate
# NOTE: DBコンテナ起動から3秒待ってから実行
db-migrate:
	sleep 3
	cd ./infra/migrations
	goose mysql "${DSN}" up

generate:
	go generate ./...
