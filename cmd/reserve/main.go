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
	"github.com/inqast/saga-order/internal/reserve/app"
	reserveHandler "github.com/inqast/saga-order/internal/reserve/handler"
	"github.com/inqast/saga-order/internal/reserve/repository"
	pb "github.com/inqast/saga-order/pkg/api/reserve"
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfigFromFile()

	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true

	db.Migrate(&cfg.Reserve.Database)
	adp, err := db.New(ctx, &cfg.Reserve.Database)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.New(adp)

	producer, err := sarama.NewSyncProducer([]string{cfg.Kafka.Brokers}, saramaCfg)
	if err != nil {
		log.Fatalf("%v", err)
	}
	handler, err := reserveHandler.New(saramaCfg, repo, producer)
	handler.StartConsuming(ctx, cfg.Kafka)

	lis, err := net.Listen("tcp", cfg.Reserve.Grpc.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	newServer := app.New(repo)
	grpcServer := grpc.NewServer()
	pb.RegisterReserveServiceServer(grpcServer, newServer)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second)
	}
}
