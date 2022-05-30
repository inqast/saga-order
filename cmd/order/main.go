package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	"github.com/inqast/saga-order/internal/config"
	"github.com/inqast/saga-order/internal/db"
	"github.com/inqast/saga-order/internal/order/app"
	"github.com/inqast/saga-order/internal/order/repository"
	pb "github.com/inqast/saga-order/pkg/api/order"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfigFromFile()

	db.Migrate(&cfg.Order.Database)

	adp, err := db.New(ctx, &cfg.Order.Database)
	if err != nil {
		log.Fatal(err)
	}

	newServer := app.New(repository.New(adp))
	lis, err := net.Listen("tcp", cfg.Order.Grpc.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
