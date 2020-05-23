package main

import (
	"context"
	"log"

	"arni.com/golang-grpc/chat"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	if resp, err := c.SayHello(
		context.Background(),
		&chat.Message{Body: "Hello from Client!"}); err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	} else {
		log.Printf("Response from server: %s", resp.Body)
	}

	if resp, err := c.SayBye(
		context.Background(),
		&chat.Message{Body: "Bye from Client!"}); err != nil {
		log.Fatalf("Error when calling SayBye: %s", err)
	} else {
		log.Printf("Response from server: %s", resp.Body)
	}

}
