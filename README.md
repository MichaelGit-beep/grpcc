1. To start from scratch
```
sudo apt install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"


# Generate go grcp code
protoc -I=. --go-grpc_out=. chat.proto
# Generate go protobuf
protoc -I=. --go_out=. chat.proto
```

2. Just to test
```
go run server.go &
go run client
```