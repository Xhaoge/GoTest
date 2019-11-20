package main

import (
	"log"
	"os"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "goGrpc/helloworld"
)

const (
	address     = "localhost:8089"
	defaultName = "xhaoge"
	defaultAge  = 24
	defaultWork = "coder"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreetingClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	rr, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name,Age:defaultAge,Work:defaultWork})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(rr)
	log.Println("Greeting: ", rr.Message,rr.Work,rr.Age)
}
