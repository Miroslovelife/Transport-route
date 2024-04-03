package models

type ProductsWarehouses interface{}

type ProductsWarehouse struct {
	Id        int
	Warehouse Wharehouses
	Product   Products
	Quantity  int
}
