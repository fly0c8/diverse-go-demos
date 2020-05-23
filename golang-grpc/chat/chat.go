package chat

import (
	context "context"
)

// Server implements ChatServiceServer interface
type Server struct {
}

// SayHello say hello
func (s *Server) SayHello(context.Context, *Message) (*Message, error) {
	return &Message{Body: "Hello from the Server!"}, nil
}

// SayBye say bye
func (s *Server) SayBye(context.Context, *Message) (*Message, error) {
	return &Message{Body: "Good Bye from the Server!"}, nil
}
