package auth

import (
"vrochatbox/routers/base"
"vrochatbox/modules/auth"
"vrochatbox/modules/models"
	"github.com/astaxie/beego"
)

type UserRouter struct {
	routers.BaseRouter
}

func (this*UserRouter) Register() {
	// no need login
	if this.CheckLoginRedirect(false) {
		return
	}

	this.Data["IsRegister"] = true
	this.TplName = "auth/register.html"

	form := auth.RegisterForm{}
	this.Data["form"] = form

}

func (this*UserRouter) RegisterProcess() {
	this.Data["IsRegister"] = true
	this.TplName = "auth/register.html"

	// no need login
	if this.CheckLoginRedirect(false) {
		return
	}

	form := auth.RegisterForm{}
	if err := this.ParseForm(&form); err == nil {
		// Create new user.
		user := new(models.User)
		if form.Password != form.PasswordRe {
			this.Data["Form"] = form;
			this.Data["Errors"] = "Password not match"
			return;
		}

		if auth.GetUser(user,form.UserName) {
			this.Data["Form"] = form;
			this.Data["Errors"] = "Username has been used"
			return;
		}

		if err := auth.Register(user, &form); err == nil {

			this.LoginUser(user, false)
			this.Redirect("/",302)

			return

		} else {
			beego.Error("Register: Failed ", err)
		}
	} else {
		beego.Error("Register: Failed ", err)
	}

}


func (this*UserRouter) Login() {
	if this.CheckLoginRedirect(false) {
		return
	}

	this.TplName = "auth/login.html"

	form := auth.LoginForm{}
	this.Data["form"] = form
}


func (this*UserRouter) LoginProcess() {
	if this.CheckLoginRedirect(false) {
		return
	}

	user := models.User{}
	form := auth.LoginForm{}
	var err error
	if err = this.ParseForm(&form); err == nil {
		if auth.VerifyUser(&user, form.Username, form.Password) {
			auth.LoginUser(&user, this.Ctx, form.Remember)
			this.Redirect("/", 302)
		} else {
			this.Data["Form"] = form
			this.Data["Errors"] = "Wrong username or password"
		}
	} else {
		this.Data["Form"] = form
		this.Data["Errors"] = "Login failed"
	}

	this.TplName = "auth/login.html"

}

func (this* UserRouter) Logout() {
	auth.LogoutUser(this.Ctx)
	this.Redirect("/login", 302)
}

func (this*UserRouter) ChangePassword() {
	if this.IsLogin == false  {
		this.Redirect("/login", 302)
		return
	}

	this.TplName = "auth/changepassword.html"

	form := auth.ChangePasswordForm{}
	this.Data["form"] = form
}

func (this*UserRouter) ChangePasswordProcess() {
	if this.IsLogin == false  {
		this.Redirect("/login", 302)
		return
	}
	form := auth.ChangePasswordForm{}
	if err:= this.ParseForm(&form);err == nil {
		if auth.VerifyPassword(form.PasswordOld, this.User.Password, this.User.Salt) {
			if err := auth.ChangePassword(&this.User, form.Password); err == nil {
				this.Redirect("/", 302)
			} else {
				this.Data["Error"] = "Can't change password"
			}
		} else {
			this.Data["Error"] = "Your old password is incorrect"
		}
	} else {
		this.Data["Error"] = "Unknown error"
	}
	this.TplName = "auth/changepassword.html"
}


func (this*UserRouter) ChangeDisplayName() {
	if this.IsLogin == false  {
		this.Redirect("/login", 302)
		return
	}

	this.TplName = "auth/changedisplayname.html"

	form := auth.ChangeDisplayNameForm{}
	this.Data["form"] = form
}

func (this*UserRouter) ChangeDisplayNameProcess() {
	form := auth.ChangeDisplayNameForm{}
	if err := this.ParseForm(&form); err == nil  {
		if err:= auth.ChangeDisplayName(&this.User,form.DisplayName,form.Avatar); err == nil {
			this.Redirect("/",302)
		} else {
			this.Data["Error"] = "Unknown error"
		}
	} else {
		this.Data["Error"] = "Missing form"
	}
	this.TplName = "auth/changedisplayname.html"
}