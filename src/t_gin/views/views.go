package views

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"net/http"
	"t_gin/models"
	"t_gin/mytool"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}

func CheckLogIn(c *gin.Context) {
	name := c.QueryArray("name")
	value := mytool.Get(name[0])
	fmt.Println(value)
	if value == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录已过期",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	}

}

func LogIn(c *gin.Context) {
	name := c.QueryArray("name")
	t := time.Now()
	fmt.Print(t)
	ok := mytool.Set(name[0], "321")
	c.JSON(http.StatusOK, gin.H{
		"msg": ok,
	})
}

func Header(c *gin.Context) {
	i := c.GetHeader("identity")
	c.JSON(http.StatusOK, gin.H{
		"msg": i,
	})
}

func GetInfo(c *gin.Context) {
	var users []*models.User
	o := orm.NewOrm()
	_, err := o.QueryTable("user").OrderBy("Id").All(&users)
	if err != nil {
		fmt.Print(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"info": users,
	})
}
