package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type User struct {
	Id   int
	Age  int
	Name string
}

type Post struct {
	gorm.Model
	Content string
}

func init() {
	DB, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=blog sslmode=disable")
	if err != nil {
		print(err)
	}
	DB.AutoMigrate(&User{}, &Post{})
	defer DB.Close()
}

func GetUsers() *User {
	var user User
	err := DB.Find(&user, 1).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("1111", user)
	return &user
}
