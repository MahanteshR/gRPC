package main

import (
	protos "gRPC/currency/protos/currency"
	"gRPC/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	log2 "log"
	"net"
	"os"
)

func main() {
	log := log2.Default()

	gs := grpc.NewServer()
	c := server.NewCurrency(log)

	reflection.Register(gs)
	protos.RegisterCurrencyServer(gs, c)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalf("unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
