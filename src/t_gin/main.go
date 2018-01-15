package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"t_gin/g_views"
	"t_gin/models"
	"t_gin/views"
)

func init() {
	// Beego ORM
	//var w io.Writer
	//orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",
		"host=127.0.0.1 port=5432 user=postgres password=postgres dbname=blog sslmode=disable")
	orm.Debug = true
	//orm.NewLog(w)

	// gorm
	//db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=blog sslmode=disable")
	//if err != nil {
	//	print(err)
	//}
	//db.AutoMigrate(&models.User{})
	//defer db.Close()
	//var users []models.User
	//db.Find(&users)
	//fmt.Print(users)
}

func main() {
	r := gin.Default()

	// Disable Console Color, you don't need console color when writing the logs to file.
	//gin.DisableConsoleColor()

	// Logging to a file.
	//f, _ := os.Create("ceshi.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// gorm
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=blog sslmode=disable")
	if err != nil {
		print(err)
	}
	db.AutoMigrate(&models.User{}, &models.Post{})
	defer db.Close()

	// beego orm
	user := r.Group("/user")
	user.GET("/info", views.GetInfo)

	r.GET("/ping", views.Ping)
	r.GET("/header", views.Header)
	r.GET("/check", views.CheckLogIn)
	r.GET("/login", views.LogIn)

	// gorm
	r.GET("/gorm_info", g_views.Info)
	r.GET("/gorm_post", g_views.Post)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
