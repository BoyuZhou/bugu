package models

import (
	"time"
	"fmt"

	"github.com/astaxie/beego/orm"

	"bugu/redis"
	"bugu/utils"
)



type User struct {
	Id       int      		`orm:"column(id);auto"`
	Username string   		`orm:"column(username);size(64);null"`
	Password string   		`orm:"column(password);size(64);null"`
	Phone    string         `orm:"column(phone);size(11);null"`
	Createtime time.Time    `orm:"column(create_time);type(datetime);null"`
	Updatetime time.Time 	`orm:"column(update_time);type(datetime);null"`
}

func init() {
	orm.RegisterModel(new(User))
}


func GetUserById(id int) (v *User, err error) {
	idStr := fmt.Sprintf("%d", id)
	v = &User{Id: id}
	getErr := redis.GetCache(idStr, v)
	if getErr == nil {
		fmt.Println("key", v)
		return v, nil
	}
	o := orm.NewOrm()
		
	if err = o.Read(v); err == nil {
		redisErr := redis.SetCache(idStr, &v, 0)
		if redisErr != nil {
			return nil, redisErr
		}
		return v, nil
	}
	return nil, err
}

func Login(username string, password string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Username: username}
	if err = o.Read(v, "Username"); err == nil {
		if v.Password == password {
			return v, nil
		}
		return v, nil
	}
	return nil, err
}

func Register(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	if err != nil {
		utils.LogError(err)
		return 0, err
	}
	return id, err
}