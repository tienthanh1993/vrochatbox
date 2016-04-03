package models

import (
	"github.com/astaxie/beego/orm"
)

type Conversation struct {
	Id int
	UserId int
	TargetId int
	Name string
}

func (m *Conversation) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Conversation) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Conversation) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Conversation) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func init(){
	orm.RegisterModel(new(Conversation))
}