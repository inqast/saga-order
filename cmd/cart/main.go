package main

import (
	"context"
	"github.com/Shopify/sarama"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	"github.com/inqast/saga-order/internal/cart/app"
	cleanHandler "github.com/inqast/saga-order/internal/cart/handler"
	"github.com/inqast/saga-order/internal/cart/repository"
	"github.com/inqast/saga-order/internal/config"
	"github.com/inqast/saga-order/internal/db"
	pb "github.com/inqast/saga-order/pkg/api/cart"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfigFromFile()

	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true

	db.Migrate(&cfg.Cart.Database)
	adp, err := db.New(ctx, &cfg.Cart.Database)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.New(adp)

	handler, err := cleanHandler.New(saramaCfg, repo)
	handler.StartConsuming(ctx, cfg.Kafka)

	newServer := app.New(repo)
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
