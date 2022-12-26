package main

import (
	"log"
	"net"

	"github.com/MichaelGit-beep/grpcc/chat"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	log.Println("listen on :9000/tcp")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server gRPC server on port 9000: %v", err)
	}
}
