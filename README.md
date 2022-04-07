# gRPCサーバー起動確認

```
$ docker-compose up

$ grpc_cli ls localhost:50051 user.UserService
Create
Get
```

# 動作確認

```
$ grpc_cli call localhost:50051 user.UserService.Create 'name: "name", email: "user@example.com", password: "password"'
```