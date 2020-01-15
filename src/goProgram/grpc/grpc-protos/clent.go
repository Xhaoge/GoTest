package main

import (
	"fmt"
	"log"
	// "os"
	pb "Pro_golang/Golang/grpc-protos/search"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)


const (
	address = "localhost:50001"
)

func main(){
	fmt.Println("this is grpc test client")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address,grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did noe connect :%v",err)
	}

	defer conn.Close()

	c := pb.NewGreetingClient(conn)
	// Contact the server and print out its response.
	// cts, cancel := context.WithTimeout(context.Background(),time.Second())

	r,err := c.SayHello(context.Background(),&pb.HelloRequest{Name:"xhaoge",Work:"程序员",Age:17})

	if err != nil {
		log.Fatalf("could not greet: %v",err)
	}

	log.Println("message:",r.Message)
	log.Println("work:",r.Work)
	log.Println("age:",r.Age)
}