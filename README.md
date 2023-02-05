# 環境構築
## Docker

インストール

https://docs.docker.com/desktop/install/mac-install/

確認
```
$ docker version
$ docker-compose version
```

## gRPC

インストール

```
$ brew install protobuf
$ brew tap grpc/grpc
$ brew install grpc
```

### 自動生成
```
# 初回のみ
$ brew install protobuf
$ go get -u google.golang.org/grpc
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

# 2回目以降
$ protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./services/book/proto/book.proto
```

## env

.envファイルを作成

```
$ cp .env.local .env
```

# サーバ起動確認

```
$ docker-compose build
$ docker-compose up

$ grpc_cli ls localhost:50052 book.BookService
GetBooks

$ go mod tidy

# Get
$ curl http://localhost:8080/books

# Create
$ curl -X POST -H "Content-Type: application/json" -d '{"author_id":1, "title":"タイトル"}' localhost:8080/books

# Update
$ curl -X PUT -H "Content-Type: application/json" -d '{"author_id":2, "title":"New Title"}' localhost:8080/books/1

# Delete
$ curl -X DELETE http://localhost:8080/books/3
```

# MySQL

```
$ docker exec -it db /bin/sh

# mysql

> use book_app;
```

# リビルド

```
$ docker-compose up --build -d
```
