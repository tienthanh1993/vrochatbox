package models

import (
	"github.com/astaxie/beego/orm"
)

type EventType int

type MsgEvent struct {
	Id int
	Type      EventType // JOIN, LEAVE, MESSAGE
	ClientId int
	UserId int
	TargetId int
	ConversationId int
	UserName string
	Avatar string
	Content   string
	Timestamp int // Unix timestamp (secs)
}

func (m *MsgEvent) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *MsgEvent) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MsgEvent) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MsgEvent) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func init(){
	orm.RegisterModel(new(MsgEvent))
}