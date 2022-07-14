package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/baksman/grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) mustEmbedUnimplementedGreetServiceServer() {}
func (s *server) Greet(context.Context, *(greetpb.GreetRequest)) (*(greetpb.GreetResponse), error)

func main() {
	fmt.Println("hellow ord")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("failed to lsiten %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}

}
