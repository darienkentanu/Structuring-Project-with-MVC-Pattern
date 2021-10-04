package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	InitDB()
	InitialMigration()
}

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{"root", "password", "3306", "localhost", "crud_go"}
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username, config.DB_Password, config.DB_Host, config.DB_Port, config.DB_Name,
	)
	var err error
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserModel struct {
	data []User
}

func NewUserModel() *UserModel {
	return &UserModel{data: []User{}}
}

func (um *UserModel) GetAll() ([]User, error) {
	return um.data, nil
}

func (um *UserModel) Add(p User) (User, error) {
	um.data = append(um.data, p)
	return p, nil
}

func (um *UserModel) GetOne(id int) (User, error) {
	return um.data[id-1], nil
}

func (um *UserModel) EditOne(id int) (User, error) {
	return um.data[id-1], nil
}

func (um *UserModel) DeleteOne(id int) ([]User, error) {
	return um.data, nil
}
