package netsrv

import (
	"bufio"
	"fmt"
	"github.com/kevin-glare/hardcode-dev-go/hw11/pkg/crawler"
	"io"
	"log"
	"net"
)

const (
	network = "tcp4"
	address = "0.0.0.0:8000"
)

var queries = []string{
	"go",
	"golang",
	"language",
	"the",
}

//Client

type Client struct {
	conn net.Conn
}

func NewClient() (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func(c *Client) SendQuery(query string) (string, error) {
	_, err := c.conn.Write([]byte(query))
	if err != nil {
		return "", err
	}

	msg, err := io.ReadAll(c.conn)
	if err != nil {
		return "", err
	}

	return string(msg), nil
}

func(c *Client) Close() error {
	return c.conn.Close()
}

//Server

type Server struct {
	listener net.Listener
	handleFunc func(string) []crawler.Document
}

func NewServer(f func(string) []crawler.Document) (*Server, error){
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
		handleFunc: f,
	}, nil
}

func(s *Server) Run() error {
	log.Println("Server Run")
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}

		go s.handler(conn)
	}
}

func(s *Server) Close() error {
	err := s.listener.Close()
	if err != nil {
		log.Println("Error Server Stop")
		return err
	}

	log.Println("Server Stop")
	return nil
}

func(s *Server) handler(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}

		log.Printf("Requst: %s", string(msg))

		rest := fmt.Sprintf("%+v\n", s.handleFunc(string(msg)))

		_, err = conn.Write([]byte(rest))
		if err != nil {
			return
		}

		log.Printf("Response: %s", string(rest))
	}
}
