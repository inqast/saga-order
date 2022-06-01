package main

import (
	"context"
	"github.com/Shopify/sarama"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	"github.com/inqast/saga-order/internal/config"
	"github.com/inqast/saga-order/internal/db"
	"github.com/inqast/saga-order/internal/order/app"
	cartClient "github.com/inqast/saga-order/internal/order/client/cart"
	orderHandler "github.com/inqast/saga-order/internal/order/handler"
	"github.com/inqast/saga-order/internal/order/repository"
	pbcart "github.com/inqast/saga-order/pkg/api/cart"
	pb "github.com/inqast/saga-order/pkg/api/order"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfigFromFile()

	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true

	db.Migrate(&cfg.Order.Database)
	adp, err := db.New(ctx, &cfg.Order.Database)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.New(adp)

	conn, err := grpc.Dial(cfg.Cart.Grpc.Address(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := cartClient.New(pbcart.NewCartClient(conn))
	producer, err := sarama.NewSyncProducer([]string{cfg.Kafka.Brokers}, saramaCfg)
	if err != nil {
		log.Fatalf("%v", err)
	}

	handler, err := orderHandler.New(saramaCfg, repo)
	handler.StartConsuming(ctx, cfg.Kafka)

	lis, err := net.Listen("tcp", cfg.Order.Grpc.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	newServer, err := app.New(repo, client, producer)
	if err != nil {
		log.Fatalf("%v", err)
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
