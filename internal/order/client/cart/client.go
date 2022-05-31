package cart

import (
	"context"
	"google.golang.org/grpc"

	pb "github.com/inqast/saga-order/pkg/api/cart"
)

type Client interface {
	GetCartItemsByUserId(context.Context, *pb.ID, ...grpc.CallOption) (*pb.GetCartItemsResponse, error)
}
