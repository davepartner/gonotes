package controllers

import (
	"gonotes/database"
	"gonotes/models"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	//fetch user detail from session
	sess, _ := session.Get("session", c)
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	
	post.UserID = userID
	database.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)

}

func GetUserPosts(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userID, ok := sess.Values["user_id"].(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}
	
	var posts []models.Post
	database.DB.Where("user_id = ?", userID).Find(&posts)
	return c.JSON(http.StatusOK, posts)
}
