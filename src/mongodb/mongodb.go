/*
  用户用户权限设置以及用户登录
*/
package mongodb

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// 新建一个用户，和一个数据库，并保存信息
type User struct {
	Name     string
	Password string
	Ip       string
	DB       string
}

func (user *User) GetDB() {
	session, err := mgo.Dial(user.Ip)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return session.DB(user.DB)
}

func (user *User) Save() err {
	user.GetDB().AddUser(user.Name, user.Password)
}
