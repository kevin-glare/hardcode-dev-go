package main

import (
	"context"
	pb "github.com/kevin-glare/hardcode-dev-go/hw14/pkg/messenger_proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type MessengerServer struct {
	sync.Mutex
	MessageList pb.MessageList
	pb.MessengerServer
}

func NewMessengerServer() *MessengerServer {
	return &MessengerServer{
		MessageList: pb.MessageList{},
	}
}

func (m *MessengerServer) Send(_ context.Context, list *pb.MessageList) (*pb.Empty, error) {
	m.MessageList.Messages = append(m.MessageList.Messages, list.GetMessages()...)

	m.Lock()
	id := len(m.MessageList.Messages)
	for _, msg := range list.GetMessages() {
		msg.ID = int64(id)
		id++
	}
	m.Unlock()

	return nil, nil
}

func (m *MessengerServer) Messages(_ context.Context, _ *pb.Empty) (*pb.MessageList, error) {
	return &m.MessageList, nil
}

func main() {
	srv := NewMessengerServer()
	grpcServer := grpc.NewServer()

	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterMessengerServer(grpcServer, srv)
	grpcServer.Serve(lis)
}
