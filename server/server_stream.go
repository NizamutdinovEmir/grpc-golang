package main

import (
	pb "github.com/NizamutdinovEmir/grpc-golang/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesLists, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello, " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
