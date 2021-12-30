package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/shamskhalil/gApp/gen/contactpb"
	"google.golang.org/grpc"
)

func main() {
	//connect to server
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to grpc server %v\n", err)
	}
	//create new client from service
	client := contactpb.NewContactServiceApiClient(conn)
	req := &contactpb.GetOneContactRequest{
		Index: 20,
	}
	//invoke

	stream, err := client.GetAll(context.Background(), req)
	if err != nil {
		log.Fatalf("Error reading server stream RPC  %v\n", err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server stream data finished!!!!!")
		}
		if err != nil {
			log.Fatalf("Error reading content of server stream %v\n", err)
		}
		fmt.Printf("Data Received >>> %+v\n", resp)
	}

}
