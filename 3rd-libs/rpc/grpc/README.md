## grpc setup

```shell
$ brew tap grpc/grpc
$ brew install protobuf
$ go get -u github.com/golang/protobuf/proto
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ go install github.com/golang/protobuf/proto 
$ go install github.com/golang/protobuf/protoc-gen-go
export PATH=$PATH:$(go env GOPATH)/bin
```

## generation

```shell
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  helloworld.proto
```

## verification

通过对客户端和服务器端的通信抓包发现：

* grpc 底层采用http2.0 作为传输底层

```
0000   1e 00 00 00 60 0b 0a 00 00 38 06 40 00 00 00 00   ....`....8.@....
0010   00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 00   ................
0020   00 00 00 00 00 00 00 00 00 00 00 01 e6 89 c3 83   ................
0030   b1 0a 9b 9a bf b5 f6 e8 80 18 18 e3 00 40 00 00   .............@..
0040   01 01 08 0a ee 61 74 6b b8 6e b5 4a 50 52 49 20   .....atk.n.JPRI 
0050   2a 20 48 54 54 50 2f 32 2e 30 0d 0a 0d 0a 53 4d   * HTTP/2.0....SM
0060   0d 0a 0d 0a                                       ....

```

## .pb.go

* 脚手架生成了针对server和client的调用逻辑代码，在server端的载体是`GreeterServer`
* 它被注册到grpcserver中，维持在services 的map结构中。 
* 对SayHello逻辑的调用是通过accept connection -> 检索services -> 通过 servicename + methodName 来触发。