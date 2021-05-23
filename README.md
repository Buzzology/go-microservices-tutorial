# Go Microservices Tutorial
A shipping container management platform using Go for Microservices using the following tutorial: https://ewanvalentine.io/microservices-in-golang-part-1/?fbclid=IwAR1jS_32UIDe10cRRVO28J7oB4Meh1KKeB2bgskWMfL6aEqlwOWwpkoD1dQ

*Part #1 - Introducing protobuf and gRPC*  
https://ewanvalentine.io/microservices-in-golang-part-1/?fbclid=IwAR1jS_32UIDe10cRRVO28J7oB4Meh1KKeB2bgskWMfL6aEqlwOWwpkoD1dQ

*Part #2 - Docker and go-micro*
https://ewanvalentine.io/microservices-in-golang-part-2/?fbclid=IwAR1jS_32UIDe10cRRVO28J7oB4Meh1KKeB2bgskWMfL6aEqlwOWwpkoD1dQ

## Tech and Libraries

### Protobuffs
Generate: `protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto`
Managing packages: https://developers.google.com/protocol-buffers/docs/reference/go-generated#package  

#### Required libraries
`go get -u google.golang.org/grpc`  
`go get github.com/golang/protobuf/protoc-gen-go@v1.3`   

### Docker

#### Consignment Service
Build docker image: `docker build -t shippy-service-consignment .`  
Run the docker image: `docker run -p 50051:50051 shippy-service-consignment` 

### Cli assignment
Build docker image (from root): `docker build -t shippy-cli-consignment -f ./shippy-cli-assignment/Dockerfile  .`

### Go-Micro
- Dependencies
  - `go get github.com/micro/micro/v2`  
  - `go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master`  
- Generating protos: `protoc --proto_path=. --go_out=. --micro_out=. proto/consignment/consignment.proto`
- Port should be set using the `MICRO_SERVER_ADDRESS` environment variable.

#### Futher reading
- Consul or etcd  

## Futher reading
gRPC: https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/  


