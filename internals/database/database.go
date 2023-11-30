package database

import (
	"log"

	"github.com/NoIdeaCoder/Saniama/internals/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB_USER *gorm.DB
var DB_PRODUCT *gorm.DB
var DB_ORDERS *gorm.DB

func InitUserDatabase() {
	db, err := gorm.Open(sqlite.Open("internals/database/users.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB_USER = db
	err = db.AutoMigrate(&models.User{}, &models.Cart{}, &models.CartProduct{})
	if err != nil {
		log.Fatal(err)
	}
}

func InitProductDatabase() {
	db, err := gorm.Open(sqlite.Open("internals/database/products.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB_PRODUCT = db
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal(err)
	}
}
func InitOrderDatabase(){
	db, err := gorm.Open(sqlite.Open("internals/database/orders.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB_ORDERS = db
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal(err)
	}
}
