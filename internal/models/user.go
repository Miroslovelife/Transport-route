package models

type Users interface {
	GetValue()
}

type User struct {
	Id       int
	Name     string
	Password string
}
