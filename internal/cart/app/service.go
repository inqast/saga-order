package app

import (
	pb "github.com/inqast/saga-order/pkg/api/cart"
)

type tserver struct {
	repo Repository
	pb.UnimplementedCartServer
}

func New(repo Repository) *tserver {
	return &tserver{repo: repo}
}
