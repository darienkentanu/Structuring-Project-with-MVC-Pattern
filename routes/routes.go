package routes

import (
	c "github.com/darienkentanu/Structuring-Project-with-MVC-Pattern/controller"
	m "github.com/darienkentanu/Structuring-Project-with-MVC-Pattern/model"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// user routing
	userModel := m.NewUserModel()
	uc := c.NewUserController(userModel)

	e.GET("/users", uc.GetAll)
	e.GET("/users/:id", uc.GetOne)
	e.POST("/users", uc.Add)
	e.DELETE("/users/:id", uc.DeleteOne)
	e.PUT("/users/:id", uc.EditOne)

	return e
}
