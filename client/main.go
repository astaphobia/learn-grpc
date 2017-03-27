package main

import (
	"io"
	"log"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	pb "github.com/danmrichards/learn-grpc/customer"
)

const (
	address = "localhost:50051"
)

// createCustomer - Calls the RPC method CreateCustomer of CustomerServer.
//
// Params:
//     client pb.CustomerClient - The protocol buffer client.
//     customer *pb.CustomerRequest - The request message to create the customer.
func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	response, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}

	if response.Success {
		log.Printf("A new Customer has been added with id: %d", response.Id)
	}
}

// getCustomers - Calls the RPC method GetCustomers of CustomerServer.
//
// Params:
//     client pb.CustomerClient - The protocol buffer client.
//     filter *pb.CustomerFilter - Filter the customers list by this keyword.
func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
	// Call the streaming API.
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error on get customers: %v", err)
	}

	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}

		log.Printf("Customer: %v", customer)
	}
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)

	// Build a request message for a new customer and create.
	customer := &pb.CustomerRequest{
		Id:    101,
		Name:  "Shiju Varghese",
		Email: "shiju@xyz.com",
		Phone: "732-757-2923",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&pb.CustomerRequest_Address{
				Street:            "Greenfield",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}
	createCustomer(client, customer)

	// Build a second customer request and create.
	customer = &pb.CustomerRequest{
		Id:    102,
		Name:  "Irene Rose",
		Email: "irene@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}
	createCustomer(client, customer)

	// Filter with an empty Keyword
	// Example of filtering by name:
	//     filter := &pb.CustomerFilter{Keyword: "Rose"}
	filter := &pb.CustomerFilter{Keyword: ""}
	getCustomers(client, filter)
}
