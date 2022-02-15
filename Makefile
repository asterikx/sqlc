.PHONY: build build-endtoend test test-ci test-examples test-endtoend regen start psql mysqlsh

build:
	go build ./...

test:
	go test ./...

test-examples:
	go test --tags=examples ./...

build-endtoend:
	cd ./internal/endtoend/testdata && go build ./...

test-endtoend:
	cd ./internal/endtoend/testdata && go test ./...

test-ci: test-examples build-endtoend test-endtoend

regen: sqlc-dev
	go run ./scripts/regenerate/

sqlc-dev:
	go build -o ~/bin/sqlc-dev ./cmd/sqlc/

sqlc-pg-gen:
	go build -o ~/bin/sqlc-pg-gen ./internal/tools/sqlc-pg-gen

start:
	docker-compose up -d

fmt:
	go fmt ./...

psql:
	PGPASSWORD=mysecretpassword psql --host=127.0.0.1 --port=5432 --username=postgres dinotest

mysqlsh:
	mysqlsh --sql --user root --password mysecretpassword --database dinotest 127.0.0.1:3306

# $ protoc --version
# libprotoc 3.17.3
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
proto:
	protoc -I ./protos --go_out=. --go_opt=module=github.com/asterikx/sqlc ./protos/**/*.proto
