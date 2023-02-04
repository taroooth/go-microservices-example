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

$ curl http://localhost:8080/books
```

# MySQL

```
$ docker exec -it db /bin/sh

# mysql

> use book_app;
```
