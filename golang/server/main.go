package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/danmrichards/learn-grpc/golang/protobuf"
)

const (
	port = 50051
)

// customerServer is used to implement customer.customerServer.
type customerServer struct {
	savedCustomers []*pb.CustomerRequest
}

// CreateCustomer - Creates a new Customer
//
// Params:
//     ctx context.Context - The context we are using.
//     in *pb.CustomerRequest - The request message for creating a new customer.
//
// Return:
//     *pb.CustomerResponse - The response message for creating a new customer.
//     error - An error if it occurred.
func (s *customerServer) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers - Get a filtered list of customers.
//
// Params:
//     filter *pb.CustomerFilter - Filter the customers list by this keyword.
//     stream pb.Customer_GetCustomersServer - The customers server.
//
// Return:
//     error - An error if it occurred.
func (s *customerServer) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
	// Filter the list of customers.
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}

		if err := stream.Send(customer); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, &customerServer{})
	s.Serve(lis)
}
