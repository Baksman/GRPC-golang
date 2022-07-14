package main

import (
	"fmt"
	"log"

	"github.com/baksman/grpc/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("hello im a client")

	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error occured %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	_ = c
	fmt.Println("created client ")

}
