package models

type Orders interface {
}

type Order struct {
	Id                int
	ProductsWarehouse string
	Products          map[string]Products
}
