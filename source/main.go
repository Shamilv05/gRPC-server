package main

import (
	"./api"
	"./sender"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func run(addr string) {
	s := grpc.NewServer()
	serv := &sender.GRPCServer{}
	api.RegisterUUIDSenderServer(s, serv)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func main() {
	for i := 0; i < 4; i++ {
		go run("localhost:404" + strconv.Itoa(i))
	}
	fmt.Scanf("%d")
}