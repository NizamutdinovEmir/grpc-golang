package main

import (
	"context"
	pb "github.com/NizamutdinovEmir/grpc-golang/proto"
	"io"
	"log"
	"time"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesLists) {
	log.Printf("Calling HelloBidirectionalStream")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling SayHelloBidirectionalStreaming: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading stream: %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending message: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}
