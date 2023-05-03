package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/siqpush/demo-grpc-jg/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)


func main() {

	log := log.Logger{}
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(f)


	creds, err := credentials.NewClientTLSFromFile("ca.crt", "")
	if err != nil {
		log.Fatalf("failed to load CA certificate: %v", err)
	}

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	
	// Create a new HelloService client.
	client := hello.NewHelloServiceClient(conn)

	
	// Call the SayHello method.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	response, err := client.SayHello(ctx, &hello.HelloRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(response.GetCount())

	// Print the response from the server.
	
}