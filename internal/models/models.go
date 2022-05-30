package models

type CartItem struct {
	UserId    int
	ProductId int
	Count     int
}

type Product struct {
	Id    int
	Count int
}

type Order struct {
	Id       int
	UserId   int
	Status   int
	Products []*Product
}
