package routes

import (
	"gonotes/controllers"
	"gonotes/middlewares"
	"io"
	"net/http"
	"reflect"

	//"github.com/CloudyKit/jet"
	"github.com/CloudyKit/jet/v6"
	"github.com/labstack/echo/v4"
)

type JetRenderer struct {
	views *jet.Set
}

func (r *JetRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	vars, ok := data.(jet.VarMap)
	if !ok {
		vars = jet.VarMap{}
	}
	tmpl, err := r.views.GetTemplate(name)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, vars, nil)
}

func InitRoutes(e *echo.Echo) {

	//Initialize Jet template engine
	renderer := &JetRenderer{
		views: jet.NewSet(
			jet.NewOSFileSystemLoader("./views"),
			//use inDevelopmentMode to auto referesh changes in templates
			jet.InDevelopmentMode(),
		),
	}
	e.Renderer = renderer

	e.GET("/home", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.jet", jet.VarMap{
			"title": reflect.ValueOf("Home"),
		})
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "../users/login.jet", jet.VarMap{
			"title": reflect.ValueOf("Login"),
		})
	})

	e.GET("/register", func(c echo.Context) error {
		return c.Render(http.StatusOK, "../users/register.jet", jet.VarMap{
			"title": reflect.ValueOf("Register"),
		})
	})

	e.GET("/profile", func(c echo.Context) error {
		return c.Render(http.StatusOK, "../users/profile.jet", jet.VarMap{
			"title": reflect.ValueOf("Profile"),
		})
	})

	e.GET("/userlist", func(c echo.Context) error {
		return c.Render(http.StatusOK, "../users/userlist.jet", jet.VarMap{
			"title": reflect.ValueOf("List of Users"),
		})
	})

	//user authentication routes
	e.POST("/register", controllers.RegisterUser)
	e.POST("/login", controllers.LoginUser)
	e.GET("/logout", controllers.LogoutUser)

	//post routes (protected)
	e.POST("/posts", controllers.CreatePost, middlewares.AuthRequired)
	e.GET("/posts", controllers.GetUserPosts, middlewares.AuthRequired)

	//admin routes
	e.GET("/users", controllers.GetUsers)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

}
