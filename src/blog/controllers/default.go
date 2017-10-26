package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"blog/models"
	//"encoding/json"
	"encoding/json"
)

// /
type MainController struct {
	beego.Controller
}

func query(id string) ([]models.User) {
	//var j, i int64

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	var user [] models.User
	num, err := o.Raw("SELECT * FROM public.user WHERE public.user.id=?", id).QueryRows(&user)
	//fmt.Println("user", users)
	//
	if err == nil {
		fmt.Println("num", num)
	}
	//
	//i = num
	//for j = 0; j < i; j++ {
	//	fmt.Printf("Element[%d] : name:%s, age: %d Id: %d\n", j, users[j].Name, users[j].Age, users[j].Id)
	//}
	return user
}

func (self *MainController) Get() {
	id := "1"
	user := query(id)
	self.Ctx.WriteString(user[0].Name)
}



func (self *MainController) Post() {
	self.Ctx.WriteString("ok")
}


// /json_t
type Json_tController struct {
	beego.Controller
}

func (this *Json_tController) Post() {
	u := models.User{}
	//fmt.Println(this.Ctx.Input.RequestBody)
	err := this.ParseForm(&u)
	if err != nil{
		panic(err)
	}
	fmt.Println("user", u.Id)
	// json 解析和对象绑定
	//json.Unmarshal(u, &u)
	objectid := models.User(u)
	this.Data["json"] = objectid
	this.ServeJSON()
	return
}