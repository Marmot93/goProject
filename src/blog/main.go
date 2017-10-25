package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //数据库驱动
	"fmt"
)

func init() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 登记驱动
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// 登记数据库
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		beego.AppConfig.String("host"), beego.AppConfig.String("port"),
			beego.AppConfig.String("user"), beego.AppConfig.String("password"),
				beego.AppConfig.String("dbname"))
	orm.RegisterDataBase("default", "postgres", psqlInfo)
	// 自动建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}