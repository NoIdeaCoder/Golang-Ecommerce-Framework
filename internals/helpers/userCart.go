package helpers

import (
	"net/url"

	"github.com/NoIdeaCoder/Saniama/internals/database"
	"github.com/NoIdeaCoder/Saniama/internals/models"
	"github.com/gofiber/fiber/v2"
)

func PlaceOrder(c *fiber.Ctx, productnames string) error {
	User, err := GetUserInfo(c)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "Cannot Get User's info",
		})
	}
	newOrder := models.Order{
		UFName:   User.Fname,
		ULname:   User.Lname,
		UPhone:   "",
		Location: "", //TODO: Add This Somehow from the Order Request Or Something  c.BodyParser
		// Infact Except For the Things We Already Got Have it Again
		UserId: User.Id,
		Order:  productnames,
		UEmail: User.Email,
	}

	if result := database.DB_ORDERS.Create(&newOrder); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error Placing Order",
		})

	}

	return c.JSON(fiber.Map{
		"message": "Sucessfully placed an order",
	})
}

func AddToCart(c *fiber.Ctx, productname string) error {
	userId, err := GetCurrentUserId(c)
	if err != nil {
		return err
	}

	var cart models.Cart
	if err := database.DB_USER.Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return err
	}
	var product models.Product
	productname, err = url.QueryUnescape(productname)
	if err != nil {
		return err
	}
	if err := database.DB_PRODUCT.Where("name = ?", productname).First(&product).Error; err != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product Not Found",
		})
	}

	testItem := models.CartProduct{
		CartId:      cart.Id,
		Name:        product.Name,
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    1,
	}
	cart.Products = append(cart.Products, testItem)
	if err := database.DB_USER.Save(&cart).Error; err != nil {
		return err
	}
	return nil

}
