# Go Microservices Tutorial
A shipping container management platform using Go for Microservices using the following tutorial: https://ewanvalentine.io/microservices-in-golang-part-1/?fbclid=IwAR1jS_32UIDe10cRRVO28J7oB4Meh1KKeB2bgskWMfL6aEqlwOWwpkoD1dQ

*Part #1 - Introducing protobuf and gRPC*
https://ewanvalentine.io/microservices-in-golang-part-1/?fbclid=IwAR1jS_32UIDe10cRRVO28J7oB4Meh1KKeB2bgskWMfL6aEqlwOWwpkoD1dQ

## Tech and Libraries
go-micro: Go microservice library.  

### Protobuffs
Generate: `protoc -I. --go_out=plugins=grpc:. proto/consignment/consignment.proto`
Managing packages: https://developers.google.com/protocol-buffers/docs/reference/go-generated#package  

## Required libraries
`go get -u google.golang.org/grpc`  
`go get github.com/golang/protobuf/protoc-gen-go@v1.3`   

## Futher reading
gRPC: https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/  


