package main

import (
	_ "vrochatbox/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"vrochatbox/modules/models"
)

func main() {

	//url := beego.AppConfig.String("dburl")

	//orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDriver("slite3", orm.DRSqlite)
	//orm.RegisterDataBase("default", "mysql", url)

	orm.RegisterDataBase("default", "sqlite3", "vrochatbox.db")
	force := false
	// Print log.
	verbose := true
	// Error.
	err := orm.RunSyncdb("default", force, verbose)
	if err != nil {
		panic(err)
	}
	orm.Debug = false
	o := orm.NewOrm()
	var conversation models.Conversation
	conversation.Id = 1
	err = o.Read(&conversation)
	if err != nil {
		conversation.UserId = 0
		conversation.Name = "Public"
		conversation.TargetId = 0
		o.Insert(&conversation)
	}


	beego.Run()
}

