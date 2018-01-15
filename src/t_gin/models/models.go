package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/gorm"
)

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
	orm.RegisterModel(new(User))
}
