package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)
const (
	port = ":8080"
)

func main()  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic("failed to listen")
	}
	s := grpc.NewServer()

	//discoveryservice.

	if err := s.Serve(lis); err != nil {
		panic("failed to serve")
	}
	fmt.Println("program ending......")
}
