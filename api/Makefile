## db migrate
# NOTE: DBコンテナ起動から3秒待ってから実行
db-migrate:
	sleep 3
	cd ./infra/migrations
	goose mysql "${DSN}" up

generate:
	go generate ./...

coverage:
	go test ./... -race -coverprofile=tmp_coverage.out
	cat tmp_coverage.out | grep -v -e "_gen.go" -v -e "mock" > coverage.out
	rm tmp_coverage.out
