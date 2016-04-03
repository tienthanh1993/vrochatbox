package tchat

import (
	"github.com/gorilla/websocket"
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
)
type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
	EVENT_DELETE_MESSAGE
	EVENT_BAN_USER
	EVENT_UNBAN_USER
	EVENT_CURRENT_USER
)
type Event struct {
	Id int64
	Type      EventType // JOIN, LEAVE, MESSAGE
	ClientId int
	UserId int
	TargetId int
	ConversationId int

	UserName string
	Avatar string
	Timestamp int // Unix timestamp (secs)
	Content   string
}

type ClientConn struct {
	websocket *websocket.Conn
	ClientId int
	UserId int
	UserName string
	Avatar string
}

var (
	Clients = make(map[ClientConn]int) // map containing clients
	subscribe = make(chan ClientConn, 10)
	unsubscribe = make(chan int, 10)
	publish = make(chan Event, 10)
	ClientId = 0
)

func Join(UserId int,UserName string,Avatar string, ws *websocket.Conn) int {
	ClientId++
	if UserName == "" {
		UserName = fmt.Sprintf("Guest %v",ClientId)
		Avatar = "/static/img/avatar.png";
	}
	subscribe <- ClientConn{ClientId: ClientId, UserId :UserId, UserName :UserName,Avatar :Avatar, websocket: ws}
	return ClientId
}


func Leave(ClientId int)  {
	unsubscribe <- ClientId
}
func MainLoop() {
	for {
		select {
		case sub := <-subscribe:
			sockCli := ClientConn{ClientId: sub.ClientId, UserId :sub.UserId, UserName :sub.UserName, Avatar :sub.Avatar, websocket: sub.websocket}
			Clients[sockCli] = 0
			{
				//var id int64
				//if sub.UserId != 0 {
				//	o := orm.NewOrm()
				//	var event models.MsgEvent
				//	event.UserId= sub.UserId
				//	event.ClientId = sub.ClientId
				//	event.UserName = sub.UserName
				//	event.Avatar = sub.Avatar
				//	event.Content = sub.UserName+" join chat room"
				//	event.ConversationId = 1
				//	event.Timestamp = int(time.Now().Unix())
				//	event.Type =  EVENT_JOIN
				//
				//	id, _ = o.Insert(&event)
				//}


				publish <- Event{Id:0,Type:EVENT_JOIN,UserId :sub.UserId,
					UserName :sub.UserName,
					Avatar :sub.Avatar,
					Timestamp:int(time.Now().Unix()),
					Content:sub.UserName+" join chat room",
					ClientId: sub.ClientId,
					TargetId:0,
					ConversationId: 1,

				}
			}
		case event := <-publish:

			data, err := json.Marshal(event)
			if err != nil {
				beego.Error("Fail to marshal event:", err)
				return
			}


			for cs, _ := range Clients {

				if event.ConversationId != 1 {
					if event.TargetId == cs.UserId || event.UserId == cs.UserId {
						sock := cs.websocket;
						if sock != nil {
							if sock.WriteMessage(websocket.TextMessage, data) != nil {
								unsubscribe <- event.ClientId
							}

						} else {
							sock.Close()
							beego.Error("WebSocket closed 1:", cs.ClientId)
						}
					}
				} else {

					sock := cs.websocket;
					if sock != nil {
						if sock.WriteMessage(websocket.TextMessage, data) != nil {
							unsubscribe <- event.ClientId
						}

					} else {
						sock.Close()
						beego.Error("WebSocket closed 2:", cs.ClientId)
					}
				}
			}

		case unsub := <-unsubscribe:

			for cs, _ := range Clients {
				if unsub == cs.ClientId {
					sock := cs.websocket
					if sock != nil {
						sock.Close()
						fmt.Printf("Close socket %v\n",cs.ClientId)
						beego.Error("WebSocket closed:", cs.ClientId)
					}

					if cs.UserId != 0 {
						//o := orm.NewOrm()
						//var event models.MsgEvent
						//event.UserId= cs.UserId
						//event.ClientId = cs.ClientId
						//event.UserName = cs.UserName
						//event.Avatar = cs.Avatar
						//event.Content = cs.UserName+" leave chat room"
						//event.ConversationId = 1
						//event.Timestamp = int(time.Now().Unix())
						//event.Type =  EVENT_LEAVE
						//id, _ := o.Insert(&event)
						publish <- Event{Id:0,Type:EVENT_LEAVE, UserId :cs.UserId,
							UserName :cs.UserName,
							Avatar :cs.Avatar,
							Timestamp:int(time.Now().Unix()),
							Content:cs.UserName + " Leave Chat Room",
							ClientId: cs.ClientId,
							ConversationId: 1,

						}
					} else {
						publish <- Event{Id:0,Type:EVENT_LEAVE, UserId :cs.UserId,
							UserName :cs.UserName,
							Avatar :cs.Avatar,
							Timestamp:int(time.Now().Unix()),
							Content:cs.UserName + " Leave Chat Room",
							ClientId: cs.ClientId,
							ConversationId: 1,

						}
					}

					delete(Clients,cs)
					break
				}
			}
		}
	}
}

func init()  {
	go MainLoop()
}