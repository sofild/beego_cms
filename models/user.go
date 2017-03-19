package models

import (
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Addtime   int64
	Logintime int64
}

func AddUser(username string, password string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(User)
	user.Username = username
	user.Password = MD5(password)
	user.Addtime = time.Now().Unix()
	user.Logintime = time.Now().Unix()
	id, err := o.Insert(user)
	return id, err
}

func FindUser(username string, password string) User {
	o := orm.NewOrm()
	o.Using("default")
	var user User
	o.QueryTable("user").Filter("username", username).Filter("password", MD5(password)).One(&user)
	return user
}

func UpUser(userInfo User) bool {
	o := orm.NewOrm()
	o.Using("default")

	_, err := o.Update(&userInfo)
	if err != nil {
		return false
	}
	return true
}
