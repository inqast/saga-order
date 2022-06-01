package app

import (
	pb "github.com/inqast/saga-order/pkg/api/reserve"
)

type tserver struct {
	repo Repository
	pb.UnimplementedReserveServiceServer
}

func New(repo Repository) *tserver {
	return &tserver{repo: repo}
}
