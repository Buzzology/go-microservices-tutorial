module github.com/buzzology/go-microservices-tutorial/shippy-cli-assignment

go 1.16

replace github.com/buzzology/go-microservices-tutorial/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/buzzology/go-microservices-tutorial/shippy-service-consignment v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.38.0
)
