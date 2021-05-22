package main

import (
	"context"
	"log"
	"net"
	"sync"

	// This is the generated protobuf code
	pb "github.com/buzzology/go-microservices-tutorial/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - a dummy repository to simulate the user of a datastore. We can replace with a real implementation later.
type Repository struct {
	mu           sync.RWMutex // Reader/writer mutual exclusion lock
	consignments []*pb.Consignment
}

// Create a new consignment - method on repository
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()

	// TODO: Is the lock automatically released if an error occurs here???
	updated := append(repo.consignments, consignment)
	repo.consignments = updated

	repo.mu.Unlock()
	return consignment, nil
}

// Get consignments - method on repository
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

// Create consignment - method on the service
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	// Save consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return matches the structure of the Response we created in our protobuf definition
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

// Retrieve consignments - method on the service
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {

	// Retrieve consignments
	consignments := s.repo.GetAll()

	// Return matches the structure of the Response we created in our protobuf definition
	return &pb.Response{Consignments: consignments}, nil
}

func main() {

	repo := &Repository{}

	// Setup our gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Register our service with the gRPC server. This will tie our implementation to our proto def.
	pb.RegisterShippingServiceServer(s, &service{repo})

	// Register reflection service on gRPC server
	// TODO: What is the reflection service???
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
