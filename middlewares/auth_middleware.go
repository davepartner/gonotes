package middlewares

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//check if user exists in session
		sess, _ := session.Get("session", c)
		if sess.Values["user_id"] == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized user"})
		}
		return next(c)
	}
	
}
