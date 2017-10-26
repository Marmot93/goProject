package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
	"fmt"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Post() {
	u := models.User{}
	err := this.ParseForm(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u.Age)
	this.Ctx.WriteString("ok")
}