package controllers

import (
	"fproject/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	abc := models.UserFindById()
	// abc := models.user{1, 2, "123"}
	// aaa := abc.FindBy()
	this.Data["Website"] = abc
	this.Data["Email"] = "406587678@qq.com"
	this.TplNames = "user.tpl"
}
