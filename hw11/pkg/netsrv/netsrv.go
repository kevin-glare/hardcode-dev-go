package netsrv

import (
	"bufio"
	"fmt"
	"github.com/kevin-glare/hardcode-dev-go/hw11/pkg/crawler"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	network = "tcp4"
	address = "0.0.0.0:8000"
)

var queries = []string{
	"go",
	"Golang",
	"get",
	"in",
}

//Client

type Client struct {
	conn net.Conn
	r    *rand.Rand
	l    int
}

func NewClient() (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		r: rand.New(rand.NewSource(time.Now().Unix())),
		l: len(queries) - 1,
	}, nil
}

func(c *Client) randomQuery() string {
	return queries[c.r.Intn(c.l)]
}

func(c *Client) Send() error {
	query := fmt.Sprintf("%s\n", c.randomQuery())

	_, err := c.conn.Write([]byte(query))
	if err != nil {
		return err
	}

	msg, err := io.ReadAll(c.conn)
	if err != nil {
		return err
	}

	fmt.Println(string(msg))

	return nil
}

func(c *Client) Close() error {
	return c.conn.Close()
}

//Server

type Server struct {
	listener net.Listener
	InCh chan string
	OutCh chan []crawler.Document
}

func NewServer(inCh chan string, outCh chan []crawler.Document) (*Server, error){
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener: listener,
		InCh: inCh,
		OutCh: outCh,
	}, nil
}

func(s *Server) Run() {
	go func() {
		log.Println("Server Run")
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				go s.handler(conn)
				go s.sendResponse(conn)
			}
		}
	}()
}

func(s *Server) Close() error {
	err := s.listener.Close()
	if err != nil {
		log.Println("Error Server Stop")
	}

	log.Println("Server Stop")
	return nil
}

func(s *Server) handler(conn net.Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Second * 10))

	r := bufio.NewReader(conn)

	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			fmt.Println(err.Error())
		}

		s.InCh <- string(msg)
	}
}

func(s *Server) sendResponse(conn net.Conn) {
	for result := range s.OutCh {
		_, err := conn.Write([]byte(fmt.Sprintf("%+v\n", result)))
		if err != nil {
			fmt.Println(err.Error())
		}

		conn.SetDeadline(time.Now().Add(time.Second * 10))
	}
}