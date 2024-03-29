[protobuf compiler](https://grpc.io/docs/protoc-installation/)

To install protobuf

```
$ brew install protobuf
$ protoc --version
```

[protoc](https://grpc.io/docs/languages/go/basics/)

1. install protoc
```
$ go install google.golang.org/grpc/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
2. set evn
```
$ sudo vim .zprofile
export PATH="$PATH:$(go env GOPATH)/bin"
export PROTOBUF=/opt/homebrew/bin/protoc
export PATH=$PROTOBUF/bin:$PATH
$ source .zprofile
```
3. cp protoc folders to /usr/local/bin/
```
$ which protoc-gen-go
$ cd <protoc-gen-go path>
$ sudo cp -r protoc-gen-go /usr/local/bin/
```

```
$ which protoc-gen-go-grpc
$ cd <protoc-gen-go-grpc path>
$ sudo cp -r protoc-gen-go-grpc /usr/local/bin/
```

4. Defining the service by a proto file
```
$ cd <project path>
$ go get -u github.com/golang/protobuf/proto
$ go get -u github.com/golang/protobuf/protoc-gen-go
```
5. Generating client and server code 
```
$ protoc --go-grpc_out=. ./programmer.proto \ --go-grpc_out=. ./programmer.proto
```