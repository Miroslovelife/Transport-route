package models

type Providers interface{}

type Provider struct {
	Id          int
	Name        string
	ContactInfo string
	Products    map[string]Products
}
