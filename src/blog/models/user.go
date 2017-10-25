package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id   int
	Age  int16
	Name string
}

func init() {
	orm.RegisterModel(new(User))
}