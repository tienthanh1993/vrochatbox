package tchat


import (
	"vrochatbox/routers/base"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"vrochatbox/modules/models"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type ChatRouter struct {
	routers.BaseRouter
}
type ReceiveMsg struct {
	Message string `json:"message"`
	Conversation int `json:"conversation"`
	TargetId int `json:"targetid"`
}

func (this*ChatRouter) Index() {
	o := orm.NewOrm()
	qs := o.QueryTable("conversation")

	cond := orm.NewCondition().Or("UserId",this.User.Id).Or("TargetId",this.User.Id).Or("UserId",0)

	var conversation []models.Conversation
	qs.SetCond(cond).Limit(50).OrderBy("id").All(&conversation)
	this.Data["conversation"] = conversation

	this.Data["users"] = Clients
	this.TplName ="chat/index_new.html"

}
func (this*ChatRouter) Fetch() {

	id,_ := this.GetInt("id")
	page,_ := this.GetInt("page")

	o := orm.NewOrm()
	qs := o.QueryTable("msg_event")
	cond := orm.NewCondition()
	if (id != 1) {
		cond = cond.AndCond(orm.NewCondition().And("ConversationId",id)).AndCond(orm.NewCondition().Or("TargetId",this.User.Id).Or("UserId",this.User.Id))
	} else {
		cond = cond.And("ConversationId",id)
	}

	var msg_event []models.MsgEvent
	//n,_ := qs.SetCond(cond).Count()
	var offset int64
	//if n <= 30 {
	//	offset = 0
	//} else {
	//	offset= 30*int64(page)
	//}
	offset= 30*int64(page)
	qs.SetCond(cond).Limit(30,offset).OrderBy("-id").All(&msg_event)
	this.Data["json"] = msg_event
	this.ServeJSON();

}

func (this* ChatRouter) GetConversation()  {
	id,_ := this.GetInt("id")
	conv := models.Conversation{}
	conv.Id = id;
	orm := orm.NewOrm()
	if err := orm.Read(&conv); err == nil {
		this.Data["json"] = conv;
		this.ServeJSON();
	}
	this.ServeJSON();
}
func (this*ChatRouter) Conversation() {

	id,_ := this.GetInt("id")

	if this.User.Id == 0 || id == this.User.Id {
		this.ServeJSON();
	}
	o := orm.NewOrm()
	target := models.User{}
	target.Id = id
	err := o.Read(&target)
	if err == nil {
		qs := o.QueryTable("conversation")
		cond := orm.NewCondition()
		cond = cond.OrCond(orm.NewCondition().And("UserId",this.User.Id).And("TargetId",id)).OrCond(orm.NewCondition().And("TargetId",this.User.Id).And("UserId",id))



		if qs.SetCond(cond).Exist() == false {
			var conversation models.Conversation
			conversation.UserId = this.User.Id
			conversation.Name = this.User.DisplayName + "|" + target.DisplayName
			conversation.TargetId = id
			o.Insert(&conversation)
			this.Data["json"] = conversation
		}

	}



	this.ServeJSON();

}


// Join method handles WebSocket requests for WebSocketController.
func (this *ChatRouter) Join() {
	this.TplName ="chat/index.html"
	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	ClientId := Join(this.User.Id,this.User.DisplayName,this.User.Avatar, ws)


	defer Leave(ClientId)

	// Message receive loop.
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			break
		}
		o := orm.NewOrm()

		var xx ReceiveMsg
		err = json.Unmarshal(p,&xx)
		var event models.MsgEvent
		event.UserId= this.User.Id
 		event.ClientId = ClientId
		event.TargetId = xx.TargetId
		event.UserName = this.User.DisplayName
		event.Avatar = this.User.Avatar
		event.Content = models.ReplaceSmile(xx.Message)
		event.ConversationId = xx.Conversation
		event.Timestamp = int(time.Now().Unix())
		event.Type =  EVENT_MESSAGE
		id, err := o.Insert(&event)
		if err == nil {
			//fmt.Println(id)
		}

		publish <- Event{Id:id, Type:EVENT_MESSAGE,UserId : this.User.Id,
			ClientId : ClientId,
			UserName :this.User.DisplayName,
			Avatar : this.User.Avatar,
			Timestamp:int(time.Now().Unix()),
			ConversationId:xx.Conversation,
			Content:models.ReplaceSmile(xx.Message),
			TargetId : event.TargetId,

		}
	}


}

