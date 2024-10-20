package controllers

import (
	"gonotes/database"
	"gonotes/helpers"
	"gonotes/models"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type LoginInput struct {
	Email    string `json:"email" from:"email"`
	Password string `json:"password" from:"password"`
}

//register new user

func RegisterUser(c echo.Context) error {
	
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	//hash the password
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	user.Password = hashedPassword
	database.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func LoginUser(c echo.Context) error {
	input := new(LoginInput)
	//check for errors
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Invalid Input"})
	}

	var user models.User
	database.DB.Where("email = ?", input.Email).First(&user)
	if user.ID == 0 || !helpers.CheckPasswordHash(input.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"Error": "Invalid Login Credentials"})

	}

	//set session cookie
	sess, _ := session.Get("session", c)
	sess.Values["user_id"] = user.ID
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, map[string]string{"message": "Logged in successfully"})

}

func LogoutUser(c echo.Context) error {
	sess, _ := session.Get("session", c)
	delete(sess.Values, "user_id")
	sess.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
