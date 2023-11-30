package helpers

import (
	"log"
	"net/url"

	"github.com/NoIdeaCoder/Saniama/internals/controllers"
	"github.com/NoIdeaCoder/Saniama/internals/database"
	"github.com/NoIdeaCoder/Saniama/internals/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func GetUserInfo(c *fiber.Ctx) (models.User, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {

		return []byte(controllers.SECRET), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return models.User{}, c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		c.Status(fiber.StatusUnauthorized)
		return models.User{}, c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	var user models.User
	if err := database.DB_USER.Preload("Cart").Preload("Cart.Products").Where("id = ?", claims.Id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetCurrentUserId(c *fiber.Ctx) (uint, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {

		return []byte(controllers.SECRET), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return 0, c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		c.Status(fiber.StatusUnauthorized)
		return 0, c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	var user models.User
	database.DB_USER.Where("id = ?", claims.Id).First(&user)
	return user.Id, nil
}

func GetProducts() []models.Product {
	var products []models.Product
	result := database.DB_PRODUCT.Find(&products)
	if result.Error != nil {
		panic(result.Error)
	}
	return products
}

func GetSpecificProduct(product_name string) models.Product {
	var product models.Product
	product_name, err := url.QueryUnescape(product_name)
	if err != nil {
		log.Fatal(err)
		return models.Product{}
	}
	err = database.DB_PRODUCT.Where("name = ?", product_name).First(&product).Error
	if err != nil {
		log.Fatal(err)
	}
	return product

}
