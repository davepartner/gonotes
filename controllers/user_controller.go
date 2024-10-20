package controllers

import (
	"fmt"
	"gonotes/database"
	"gonotes/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// get list of all users
func GetUsers(c echo.Context) error {
	var users []models.User
	//connect to the database and find users
	database.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

// create new user
func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		fmt.Printf("Error binding request: %v\n", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})

	}
	fmt.Printf("Received request: %+v\n", user)
	//otherwise connect to the databse
	database.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)

	//check for errors connecting to DB
	if err := c.Bind(user); err != nil {
		return err
	}
	//fetch the user using their ID
	database.DB.Model(&user).Where("id = ?", id).Updates(user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")                    // get user ID
	database.DB.Delete(&models.User{}, id) //delete the user
	return c.NoContent(http.StatusNoContent)
}
