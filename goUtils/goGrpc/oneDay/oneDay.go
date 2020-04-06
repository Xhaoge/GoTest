package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

const PORT = "9091"
type SearchService struct {}

func (s *SearchService) Search(ctx context.Context){

}

func main(){
	fmt.Println("grpc main....")
	server := grpc.NewServer()
	fmt.Sprint(server)
}
