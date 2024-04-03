package models

type ProductStores interface {
}

type ProductStore struct {
	Id          int
	Warehouse   string
	Goods       string
	Quantity    int
	MinQuantity int
}
