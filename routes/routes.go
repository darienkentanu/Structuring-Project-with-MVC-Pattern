package routes

import (
	c "github.com/darienkentanu/Structuring-Project-with-MVC-Pattern/controller"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// user routing
	e.GET("/users", c.GetUsersController)
	e.GET("/users/:id", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	e.DELETE("/users/:id", c.DeleteUserController)
	e.PUT("/users/:id", c.UpdateUserController)

	return e
}
