package controller

import (
	"net/http"
	"strconv"

	"github.com/darienkentanu/Structuring-Project-with-MVC-Pattern/model"

	"github.com/labstack/echo"
)

type M map[string]interface{}

var DB = model.DB

type UserModel interface {
	GetAll() error
	Add() error
	GetOne(id int) error
	EditOne(id int) error
	DeleteOne(id int) error
}

type UserController struct {
	model UserModel
}

func NewUserController(m UserModel) UserController {
	return UserController{model: m}
}

// get all users
func (uc UserController) GetAll(c echo.Context) error {
	var users []model.User
	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, M{"message": "success get all users",
		"users": users})
}

// get user by id
func (uc UserController) GetOne(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user model.User
	if err := DB.First(&user, id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "id not found")
	}
	return c.JSON(http.StatusOK, user)
}

// create new user
func (uc UserController) Add(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, M{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func (uc UserController) DeleteOne(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}
	var user model.User
	if err := DB.First(&user, id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "id not found")
	}
	if err := DB.Delete(&user).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

func (uc UserController) EditOne(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, M{"message": "please input a valid id"})
	}
	var user model.User
	if err := DB.First(&user, id).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if user.ID == 0 {
		return c.String(http.StatusNotFound, "id not found")
	}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	if err := DB.Save(&user).Error; err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusOK, user)
}
