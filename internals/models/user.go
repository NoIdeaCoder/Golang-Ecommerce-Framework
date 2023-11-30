package models

type User struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Fname    string `json:"firstName"`
	Lname    string `json:"lastName"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Cart     Cart   `gorm:"foreignKey:UserId"`
}

type Cart struct {
	Id       uint `gorm:"primaryKey"`
	UserId   uint
	Products []CartProduct
}

type CartProduct struct {
	CartId      uint
	Name        string
	Image       string
	Description string
	Price       int
	Quantity    int
}

type LoginRequest struct {
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

