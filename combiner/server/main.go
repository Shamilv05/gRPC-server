package main

import (
	"google.golang.org/grpc"
	"net"
	"log"
	"../api"
	"../protos"
)

func run(addr string) {
	s := grpc.NewServer()
	serv := &protos.CombinerServer{}
	api.RegisterCombinerServer(s, serv)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run(":4048")
}
