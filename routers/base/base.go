package routers

import (
	"github.com/astaxie/beego"
	"vrochatbox/modules/models"
	"fmt"
	"net/url"
"strings"
"vrochatbox/modules/auth"
)

type BaseRouter struct {
	beego.Controller
	User models.User
	IsLogin	bool
}

func (this*BaseRouter) Prepare() {

	this.StartSession()

	switch {
	// save logined user if exist in session
	case auth.GetUserFromSession(&this.User, this.CruSession):
		this.IsLogin = true
	// save logined user if exist in remember cookie
	case auth.LoginUserFromRememberCookie(&this.User, this.Ctx):
		this.IsLogin = true
	}

	if this.IsLogin {
		this.IsLogin = true
		this.Data["User"] = &this.User
		this.Data["IsLogin"] = this.IsLogin

		// if user forbided then do logout
		if this.User.IsForbid {
			auth.LogoutUser(this.Ctx)
			this.Redirect("/login", 302)
			return
		}
	}

}

func (this *BaseRouter) Finish() {

}

// check if not login then redirect
func (this *BaseRouter) CheckLoginRedirect(args ...interface{}) bool {
	var redirect_to string
	code := 302
	needLogin := true
	for _, arg := range args {
		switch v := arg.(type) {
		case bool:
			needLogin = v
		case string:
			// custom redirect url
			redirect_to = v
		case int:
			// custom redirect url
			code = v
		}
	}

	// if need login then redirect
	if needLogin && !this.IsLogin {
		if len(redirect_to) == 0 {
			req := this.Ctx.Request
			scheme := "http"
			if req.TLS != nil {
				scheme += "s"
			}
			redirect_to = fmt.Sprintf("%s://%s%s", scheme, req.Host, req.RequestURI)
		}
		redirect_to = "/login?to=" + url.QueryEscape(redirect_to)
		this.Redirect(redirect_to, code)
		return true
	}

	// if not need login then redirect
	if !needLogin && this.IsLogin {
		if len(redirect_to) == 0 {
			redirect_to = "/"
		}
		this.Redirect(redirect_to, code)
		return true
	}
	return false
}

func (this *BaseRouter) LoginUser(user *models.User, remember bool) string {
	loginRedirect := strings.TrimSpace(this.Ctx.GetCookie("login_to"))
	//if utils.IsMatchHost(loginRedirect) == false {
	//	loginRedirect = "/"
	//} else {
	//	this.Ctx.SetCookie("login_to", "", -1, "/")
	//}

	// login user
	auth.LoginUser(user, this.Ctx, remember)

	return loginRedirect
}
