package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //数据库驱动
)

func init() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 登记驱动
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// 登记数据库
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("psqlInfo"))
	// 自动建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}