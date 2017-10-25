package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"blog/models"
)

type MainController struct {
	beego.Controller
}

func query() {
	var j, i int64

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	var users [] models.User
	num, err := o.Raw("SELECT age, id, name FROM public.user").QueryRows(&users)
	fmt.Println("user", users)

	if err == nil {
		fmt.Println("user nums: ", num)
	}

	i = num
	for j = 0; j < i; j++ {
		fmt.Printf("Element[%d] : name:%s, age: %d Id: %d\n", j, users[j].Name, users[j].Age, users[j].Id)
	}
}

func (self *MainController) Get() {
	query()
	self.Ctx.WriteString("ok")
}
