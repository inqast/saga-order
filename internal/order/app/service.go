package app

import (
	pb "github.com/inqast/saga-order/pkg/api/order"
)

type tserver struct {
	repo     Repository
	client   Client
	producer Producer
	pb.UnimplementedOrderServiceServer
}

func New(repo Repository, client Client, producer Producer) (*tserver, error) {
	return &tserver{repo: repo, client: client, producer: producer}, nil
}
