# gRPC

## サーバー起動確認

```
$ cp .env.local .env
$ docker-compose build
$ docker-compose up

$ grpc_cli ls localhost:50052 book.BookService
GetBooks
```

## 自動生成
```
$ protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./services/book/proto/book.proto
```

# 動作確認

```
$ grpc_cli call localhost:50052 book.BookService.GetBooks
```

# MySQL

```
$ docker exec -it db /bin/sh

# mysql

> use book_app;
```
