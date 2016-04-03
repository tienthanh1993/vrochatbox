package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int
	DisplayName string
	Email string
	UserName string
	Password string
	Salt string
	Avatar string
	Extra string
	IsAdmin bool
	IsForbid bool
	IsForbidChat bool
	IsForbidTalk bool
	Description string
	Updated time.Time	`orm:"auto_now"`
	Created time.Time	`orm:"auto_now_add"`
}
func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
func init(){
	orm.RegisterModel(new(User))
}