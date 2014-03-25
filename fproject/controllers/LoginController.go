package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.html"
}

func (this *LoginController) Post() {
	abc := this.Input()

	this.Ctx.WriteString(fmt.Sprint(abc.Get("name")))
	this.Ctx.WriteString(fmt.Sprint(abc))
}
