package cart

type Service struct {
	grpcClient  Client
	ErrNotFound error
}

func New(client Client) *Service {
	return &Service{
		grpcClient: client,
	}
}
