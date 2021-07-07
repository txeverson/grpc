# gRPC
Demo gRPC Project with Golang

Server and Client communication.

**Golang com gRPC - Comunicação entre micro-serviços**
# Pontos positivos

- Estrutura RPC leve e de alto desempenho. Alem de usar HTTP/2 que é muito mais rápido que o HTTP/1.1 usado no REST padrão.

- Usa o protobuf para serializar dados, trafegando-os em binários.

- gRPC possui streaming de dados, e é bidirecional;

# Pontos Negativos

- Nao suporta comunicacao com Browsers (é mais indicado o uso entre servicos)

- Por trafegar os dados em binários, nao são legiveis aos humanos


# ProtoBuf model

```
syntax = "proto3";
option go_package = "proto/user";

service UserService {
  rpc GetUser(UserRequest) returns (UserResponse) {}
}

message UserRequest {
  string username = 1;
}

message UserResponse {
  int64 id = 1;
  string name = 2;
  string avatarurl = 4;
  Statistics statistics = 6;
  repeated string listURLs = 7;
}

message Statistics {
  int64 followers = 1; 
  int64 gists = 4;
}
```

# Running

Golang running on Port 8200.

```
//Server
make run

//client
make run_client
```

# Output

Golang

```
2021/07/08 12:28:36 Response from server: 
id: 3136203 
name: "Everson Teixeira" 
avatarurl: "https://avatars.githubusercontent.com/u/3136203?v=4" 
statistics: {
  followers:8  
  gists:1
}
listURLs: ["https://api.github.com/users/tx.everson"]
```

# Installation required:

ProtoBuff/Protoc
http://google.github.io/proto-lens/installing-protoc.html

# Links

https://developers.google.com/protocol-buffers