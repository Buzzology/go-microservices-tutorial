module github.com/buzzology/go-microservices-tutorial/shippy-cli-assignment

go 1.16

replace github.com/buzzology/go-microservices-tutorial/shippy-service-consignment => ../shippy-service-consignment

replace github.com/buzzology/go-microservices-tutorial/shippy-service-vessel => ../shippy-service-vessel

replace google.golang.org/grpc v1.38.0 => google.golang.org/grpc v1.26.0

require (
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b // indirect
	github.com/buzzology/go-microservices-tutorial/shippy-service-consignment v0.0.0-20210523041700-c6691415405e
	github.com/micro/go-micro/v2 v2.9.1
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf // indirect
)
