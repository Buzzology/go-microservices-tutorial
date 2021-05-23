module github.com/buzzology/go-microservices-tutorial/shippy-cli-assignment

go 1.16

//replace github.com/buzzology/go-microservices-tutorial/shippy-service-consignment => ../shippy-service-consignment

replace google.golang.org/grpc v1.38.0 => google.golang.org/grpc v1.26.0

require (
	github.com/buzzology/go-microservices-tutorial/shippy-service-consignment v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro/v2 v2.9.1-0.20200723075038-fbdf1f2c1c4c
	google.golang.org/grpc v1.38.0
)
