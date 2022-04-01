# gRPCサーバー起動確認

```
$ go run cmd/user/main.go

$ grpc_cli ls localhost:50051 user.UserService
Create
Get
```

# 動作確認

```
$ grpc_cli call localhost:50051 user.UserService.Create 'name: "name", email: "user@example.com", password: "password"'
```