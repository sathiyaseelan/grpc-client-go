package main

import (
	"log"
	"os"
	"time"

	pb "github.com/moneysmartco/userdata-go/userdata_v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9003"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserDataClient(conn)

	// Contact the server and print out its response.
	var email string
	if len(os.Args) > 1 {
		email = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &pb.GetUserRequest{UniqueUserOneof: &pb.GetUserRequest_Email{Email: email}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}
