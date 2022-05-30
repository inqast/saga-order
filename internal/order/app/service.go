package app

import (
	pb "github.com/inqast/saga-order/pkg/api/order"
)

type tserver struct {
	repo Repository
	pb.UnimplementedOrderServiceServer
}

func New(repo Repository) *tserver {
	return &tserver{repo: repo}
}
