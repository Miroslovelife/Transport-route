package models

type Products interface{}

type Product struct {
	Id       int
	Name     string
	Provider string
	Price    float64
}
