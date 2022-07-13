package main

import (
	"github.com/joealisson/odin/api/grpc/ptypes"
	"github.com/joealisson/odin/ingestion"
	"google.golang.org/grpc"
	"log"
	"net"

	odin "github.com/joealisson/odin/api/grpc"
	"github.com/joealisson/odin/env"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	env.Load()

	ls, err := net.Listen("tcp", env.BindAddress)
	if err != nil {
		log.Fatalf("Error listening on %s: %v", env.BindAddress, err)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	ingestor := &ingestion.InMemoryIngestor{}
	ptypes.RegisterOdinServer(server, odin.NewServer(ingestor))

	log.Println("Odin gRPC Listening on ", env.BindAddress)
	err = server.Serve(ls)
	if err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
