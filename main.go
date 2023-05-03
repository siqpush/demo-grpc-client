package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"github.com/siqpush/demo-grpc-jg/hello"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new HelloService client.
	client := hello.NewHelloServiceClient(conn)

	// Call the SayHello method.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	

	ch := make(chan uint8, 2)
	for i := 0; i < 100; i++ {
		go func()  {
			ch <- 1
			response, err := client.SayHello(ctx, &hello.HelloRequest{})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			fmt.Println(response.GetCount())
		}()
		<-ch
	}

	// Print the response from the server.
	
}