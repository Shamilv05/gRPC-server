package main

import (
	"google.golang.org/grpc"
	"./sender"
	"./api"
	"log"
	"net"
)


func main() {
	s := grpc.NewServer()
	serv := &sender.GRPCServer{}
	api.RegisterUUIDSenderServer(s, serv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}


