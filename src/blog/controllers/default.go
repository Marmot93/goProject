package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (self *MainController) Get() {
	self.Ctx.WriteString("hello")
}