package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Najwan160/go-grpc-1/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {

	log.Printf("Straming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
