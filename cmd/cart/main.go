package main

import (
	"context"
	"github.com/inqast/saga-order/internal/cart/app"
	"github.com/inqast/saga-order/internal/cart/repository"
	"github.com/inqast/saga-order/internal/config"
	"github.com/inqast/saga-order/internal/db"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	pb "github.com/inqast/saga-order/pkg/api/cart"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfigFromFile()

	db.Migrate(&cfg.Cart.Database)

	adp, err := db.New(ctx, &cfg.Cart.Database)
	if err != nil {
		log.Fatal(err)
	}

	newServer := app.New(repository.New(adp))
	lis, err := net.Listen("tcp", cfg.Cart.Grpc.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCartServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
