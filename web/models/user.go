package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	Name       string
	Pass       string
	Email      string
	Rights     string
	Created_at string
}

func init() {
	orm.RegisterModel(new(User))
}

func FindAllUser() ([]User, error) {
	o := orm.NewOrm()
	var user []User
	_, err := o.QueryTable("user").All(&user)
	return user, err
}

func FindUserByName(name string) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("name", name).One(&user)
	return user, err
}

func FindUserById(id int) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("id", id).One(&user)
	return user, err
}

func AddUser(name, pass, email, right string) (int64, error) {
	o := orm.NewOrm()
	sql := "insert into user (name, pass,email,rights) values( ?, ?, ?, ?)"
	res, err := o.Raw(sql, name, pass, email, right).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func EditUser(name, email, right string, id int) (int64, error) {
	o := orm.NewOrm()
	sql := "update user  set name=?,email=?,rights=? where id=?"
	res, err := o.Raw(sql, name, email, right, id).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}

}

func DeleteUser(id int) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&User{Id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}