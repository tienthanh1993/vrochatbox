package auth

import (
	"vrochatbox/modules/models"
	"vrochatbox/utils"
	"github.com/astaxie/beego/orm"
"strings"
"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
"github.com/astaxie/beego/session"
)

func Register(user * models.User,registerform* RegisterForm) error {
	user.UserName = registerform.UserName
	user.Salt = utils.Random(5)
	user.DisplayName = registerform.DisplayName
	user.Password = utils.GetMD5Hash(registerform.Password+user.Salt)

	user.Avatar = "/static/img/avatar.png"
	return user.Insert()
}

func UserExists(username string, email string) (bool,bool,error) {
	cond := orm.NewCondition()
	cond = cond.Or("UserName", username).Or("Email", email)

	var maps []orm.Params
	o := orm.NewOrm()
	n, err := o.QueryTable("user").SetCond(cond).Values(&maps, "UserName", "Email")
	if err != nil {
		return false, false, err
	}

	e1 := true
	e2 := true

	if n > 0 {
		for _, m := range maps {
			if e1 && orm.ToStr(m["UserName"]) == username {
				e1 = false
			}
			if e2 && orm.ToStr(m["Email"]) == email {
				e2 = false
			}
		}
	}

	return e1, e2, nil
}

func GetUser(user* models.User, username string) bool {
	var err error
	qs := orm.NewOrm()
	if strings.IndexRune(username, '@') == -1 {
		user.UserName = username
		err = qs.Read(user, "UserName")
	} else {
		user.Email = username
		err = qs.Read(user, "Email")
	}
	if err == nil {
		return true
	}
	return false
}

func ChangePassword(user* models.User, password string) error {
	salt := utils.Random(5)
	user.Salt = salt
	user.Password = utils.GetMD5Hash(password+salt)
	return user.Update("Password","Salt","Updated")
}
func ChangeDisplayName(user* models.User, displayname string,avatar string ) error {
	user.DisplayName = displayname
	if avatar != "" {
		user.Avatar = avatar
	}
	return user.Update("DisplayName","Avatar","Updated")
}

func LoginUser(user* models.User, ctx *context.Context, remember bool)  {
	ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
	ctx.Input.CruSession = beego.GlobalSessions.SessionRegenerateID(ctx.ResponseWriter, ctx.Request)
	ctx.Input.CruSession.Set("user_id", user.Id)

	if remember {
		WriteRememberCookie(user, ctx)
	}
}

func WriteRememberCookie(user *models.User, ctx *context.Context) {
	secret := utils.GetMD5Hash(user.Password+user.Salt)
	days := 86400 * 7
	ctx.SetCookie("username", user.UserName, days)
	ctx.SetSecureCookie(secret, "remember", user.UserName, days)
}

func DeleteRememberCookie(ctx *context.Context) {
	ctx.SetCookie("username", "", -1)
	ctx.SetCookie("remember", "", -1)
}

func LoginUserFromRememberCookie(user *models.User, ctx *context.Context) (success bool) {
	userName := ctx.GetCookie("username")
	if len(userName) == 0 {
		return false
	}

	defer func() {
		if !success {
			DeleteRememberCookie(ctx)
		}
	}()

	user.UserName = userName
	if err := user.Read("UserName"); err != nil {
		return false
	}

	secret := utils.GetMD5Hash(user.Password+user.Salt)
	value, _ := ctx.GetSecureCookie(secret, "remember")
	if value != userName {
		return false
	}

	LoginUser(user, ctx, true)

	return true
}

func VerifyUser(user *models.User, username, password string) (success bool) {
	if GetUser(user, username) == false {
		return
	}

	if VerifyPassword(password, user.Password, user.Salt) {
		success = true
	}
	return
}

func VerifyPassword(rawPwd, encodedPwd, salt string) bool {
	return utils.GetMD5Hash(rawPwd+salt) == encodedPwd
}

// logout user
func LogoutUser(ctx *context.Context) {
	DeleteRememberCookie(ctx)
	ctx.Input.CruSession.Delete("user_id")
	ctx.Input.CruSession.Flush()
	beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
}

func GetUserIdFromSession(sess session.Store) int {
	if id, ok := sess.Get("user_id").(int); ok && id > 0 {
		return id
	}
	return 0
}

// get user if key exist in session
func GetUserFromSession(user *models.User, sess session.Store) bool {
	id := GetUserIdFromSession(sess)
	if id > 0 {
		u := models.User{Id: id}
		if u.Read() == nil {
			*user = u
			return true
		}
	}

	return false
}
