package routers

import (
	"github.com/astaxie/beego"

	"vrochatbox/routers/auth"
	"vrochatbox/routers/tchat"
)

func init() {
	beego.Router("/", &tchat.ChatRouter{},"get:Index")
	beego.Router("/fetch", &tchat.ChatRouter{},"get:Fetch")
	beego.Router("/conversation", &tchat.ChatRouter{},"post:Conversation")
	beego.Router("/getconversation", &tchat.ChatRouter{},"get:GetConversation")
	beego.Router("/ws", &tchat.ChatRouter{})
	beego.Router("/ws/join", &tchat.ChatRouter{}, "get:Join")

	beego.Router("/members", &tchat.ChatRouter{},"get:Index")
	beego.Router("/help", &tchat.ChatRouter{},"get:Index")
	beego.Router("/login", &auth.UserRouter{}, "get:Login;post:LoginProcess")
	beego.Router("/logout", &auth.UserRouter{}, "get:Logout")
	beego.Router("/register", &auth.UserRouter{}, "get:Register;post:RegisterProcess")
	beego.Router("/changepassword", &auth.UserRouter{}, "get:ChangePassword;post:ChangePasswordProcess")
	beego.Router("/changedisplayname", &auth.UserRouter{}, "get:ChangeDisplayName;post:ChangeDisplayNameProcess")
}
