package models

type Wharehouses interface {
}

type Warehouse struct {
	Id        int
	Location  string
	Transport map[string]Transport
}
