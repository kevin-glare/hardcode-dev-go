package main

import (
	pb "github.com/kevin-glare/hardcode-dev-go/hw14/pkg/messenger_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"context"
)

func main()  {
	conn, err := grpc.Dial("localhost:12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewMessengerClient(conn)

	err = printAllMessages(client)
	if err != nil {
		log.Fatal(err)
	}

	messageList := pb.MessageList{Messages: []*pb.Message{}}
	messageList.Messages = append(messageList.Messages, &pb.Message{Text: "Hello"})
	messageList.Messages = append(messageList.Messages, &pb.Message{Text: "World"})
	client.Send(
		context.Background(),
		&messageList,
	)

	err = printAllMessages(client)
	if err != nil {
		log.Fatal(err)
	}
}

func printAllMessages(c pb.MessengerClient) error {
	list, err := c.Messages(context.Background(), &pb.Empty{})
	if err != nil {
		return err
	}

	log.Println(list.Messages)

	return nil
}