package main

import (
	"context"
	"log"

	// This is the generated protobuf code
	pb "github.com/buzzology/go-microservices-tutorial/shippy-service-consignment/proto/consignment"
	"github.com/micro/go-micro/v2"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - a dummy repository to simulate the user of a datastore. We can replace with a real implementation later.
type Repository struct {
	consignments []*pb.Consignment
}

// Create a new consignment - method on repository
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {

	// TODO: Is the lock automatically released if an error occurs here???
	updated := append(repo.consignments, consignment)
	repo.consignments = updated

	return consignment, nil
}

// Get consignments - method on repository
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type consignmentService struct {
	repo repository
}

// Create consignment - method on the service
func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Save consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matches the structure of the Response we created in our protobuf definition
	res.Created = true
	res.Consignment = consignment

	return nil
}

// Retrieve consignments - method on the service
func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	// Retrieve consignments
	res.Consignments = s.repo.GetAll()

	// Return matches the structure of the Response we created in our protobuf definition
	return nil
}

func main() {

	repo := &Repository{}

	// Creates the new service (can include options here)
	service := micro.NewService(

		// This name must match the package name given in protobuf definition
		micro.Name("shippy.service.consignment"),
	)

	service.Init()

	// Register service - instantiating err and checking in one line
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}