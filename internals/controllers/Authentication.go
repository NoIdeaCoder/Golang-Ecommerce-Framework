package controllers

import (
	"fmt"
	"github.com/NoIdeaCoder/Saniama/internals/database"
	"github.com/NoIdeaCoder/Saniama/internals/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const SECRET string = "SUPERSECRETSHIT" // TODO: user enviroment variables to store and use this
func Signup(c *fiber.Ctx) error {
	var body models.User
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error Hashing Password",
		})

	}
	newuser := models.User{
		Fname:    body.Fname,
		Lname:    body.Lname,
		Email:    body.Email,
		Password: string(encryptedPassword),
	}
	if result := database.DB_USER.Create(&newuser); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error Creating User",
		})
	}
	cart := models.Cart{
		UserId: newuser.Id,
	}
	if result := database.DB_USER.Create(&cart); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error Creating Cart",
		})
	}

	return c.Redirect("/login")
}

func Login(c *fiber.Ctx) error {
	var loginRequest models.LoginRequest
	var user models.User
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}
	result := database.DB_USER.Where("email = ?", loginRequest.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Email or Password ", // TODO :instead of sending an error can i send html snippet that sends over the error on a div thats right above the password and email field?
		})

	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Email or Password ",
		})

	}
	token, err := generatejwt(strconv.FormatUint(uint64(user.Id), 10))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})

	}
	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	c.Cookie(cookie)
	fmt.Println("Successful login")
	return c.Redirect("/")
}

func generatejwt(userId string) (string, error) {
	claims := &jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
func AuthMiddleware(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return c.Redirect("/login")
	}

	claims := token.Claims.(*jwt.StandardClaims)
	if time.Now().Unix() > claims.ExpiresAt {
		// Token has expired, redirect to the login page
		return c.Redirect("/login")
	}

	return c.Next()
}
