package chat

import (
	context "context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recieved message body from client: %s\n", message.Body)
	return &Message{Body: "Hello From The server"}, nil
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {}
