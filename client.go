package main

import (
	"context"
	"log"

	"github.com/MichaelGit-beep/grpcc/chat"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %s", err)
	}
	defer conn.Close()

	message := &chat.Message{
		Body: "Hello from the client",
	}

	c := chat.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), message)
	if err != nil {
		log.Fatalf("error when calling SayHello: %s", err)
	}
	log.Printf("Responce from Server: %s", response.Body)
}
