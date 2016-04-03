package auth


type RegisterForm struct {
	DisplayName string `form:"displayname"`
	UserName string `form:"username"`
	Password string `form:"password"`
	PasswordRe string `form:"passwordre"`
}
type ChangePasswordForm struct {
	PasswordOld string `form:"passwordold"`
	Password string `form:"password"`
	PasswordRe string `form:"passwordre"`
}
type ChangeDisplayNameForm struct {
	DisplayName string `form:"displayname"`
	Avatar string `form:"avatar"`
}
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Remember bool `form:remember`
}
