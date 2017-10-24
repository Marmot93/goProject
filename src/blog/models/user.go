package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type User struct {
	Id    int
	Phone string
	//UserProfile *UserProfile `orm:"rel(one)"`
	Password string
	Status   int
	Created  int64
	Changed  int64
}
type UserProfile struct {
	Id       int
	Realname string
	Sex      int
	Birth    string
	Email    string
	Phone    string
	Address  string
	Hobby    string
	Intro    string
	User     *User `orm:"reverse(one)"`
}

func (this *User) TableName() string {
	return "user"
}
func RegisterModel() {
	user := beego.AppConfig.String("user")
	passwd := beego.AppConfig.String("passwd")
	host := beego.AppConfig.String("host")
	port, err := beego.AppConfig.Int("port")
	dbname := beego.AppConfig.String("dbname")
	if nil != err {
		port = 5432
	}
	orm.Debug = true
	orm.RegisterModel(new(User)) //
	orm.RegisterDriver("postgres", orm.DRPostgres)
	//orm.RegisterDataBase("default", "mysql", "root:@/blog?charset=utf8", 30)
	orm.RegisterDataBase("default", "postgres", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))

}
