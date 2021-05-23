module github.com/buzzology/go-microservices-tutorial/shippy-service-consignment

go 1.16

replace github.com/buzzology/go-microservices-tutorial/shippy-service-consignment => ../shippy-service-consignment

replace google.golang.org/grpc v1.38.0 => google.golang.org/grpc v1.26.0

require (
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d // indirect
	github.com/golang/protobuf v1.5.2
	github.com/kr/pretty v0.2.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1-0.20200723075038-fbdf1f2c1c4c
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	golang.org/x/net v0.0.0-20210521195947-fe42d452be8f // indirect
	golang.org/x/sys v0.0.0-20210521203332-0cec03c779c1 // indirect
	google.golang.org/genproto v0.0.0-20210521181308-5ccab8a35a9a // indirect
	google.golang.org/grpc v1.38.0 // indirect; Dropped from version 1.38 to fix naming issue
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
