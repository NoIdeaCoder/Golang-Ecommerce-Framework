package models

type Product struct {
	Id            uint `gorm:"primary_key"`
	Name          string
	Description   string
	NumberOfStock int
	Price         int
	Image         string
}

type Order struct {
	Id       uint
	UFName   string
	ULname   string
	UEmail   string
	UPhone   string
	Location string
	UserId   uint
	Order    string
}
