package main

import (
	//"gonotes/controllers"
	"gonotes/database"
	"gonotes/routes"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New() //instantiate echo framework
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)

	//enable CORS for all origins (you can limit it to specific origins if necessaries)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//adjust this if your front end is served from a different origin/domain
		AllowOrigins: []string{"http://localhost:8000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	//connect to the dabase
	database.ConnectDB()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("239239b23b2jb32jv3j23v2j3"))))
	//serve static files form "public" folder
	// e.Static("/", "public")
	e.Static("/css", "public/css")
	e.Static("/js", "public/js")

	//routing
	routes.InitRoutes(e)

	//start server
	e.Logger.Fatal(e.Start(":8000"))
}
