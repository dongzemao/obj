package models

import (
	"github.com/astaxie/beego/orm"
)

type user struct {
	Id  int
	Age int
	Sex string
}

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	orm.RegisterDataBase("default", "mysql", "root:@(127.0.0.1:3306)/blog?charset=utf8", 30)
}

func (ooo *user) FindBy() string {
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM user WHERE id = ?", 2).QueryRow(ooo)
	if err != nil {
		return "abc123"
	}
	return ooo.Sex
}

func UserFindById() string {
	u := user{}
	string := u.FindBy()
	return string
}
