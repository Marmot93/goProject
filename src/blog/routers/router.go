package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/form/", &controllers.TestController{})
    beego.Router("/json_t/", &controllers.Json_tController{})
}
