package main

import (
	"fproject/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Run()
}
