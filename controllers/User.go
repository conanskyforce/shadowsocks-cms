package controllers

import (
	"github.com/cheneylew/goutil/utils"
	"github.com/cheneylew/shadowsocks-cms/database"
	"github.com/cheneylew/shadowsocks-cms/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Prepare() {
	c.BaseController.Prepare()
}

func (c *UserController) Finish() {
	c.BaseController.Finish()
}

func (c *UserController) Get() {
	c.TplName = "user_login.html"
}

func (c *UserController) Login() {
	c.TplName = "user_login.html"
	email := c.GetString("email")
	password := c.GetString("password")
	utils.JJKPrintln(email, password)

	if len(email) > 0 && len(password) > 0 {
		users := database.DBQueryUserWithEmailOrMobile(email)

		var loginedUser models.User
		isLogin := false
		for _, user := range users {
			if user.Password == password {
				loginedUser = user
				isLogin = true
			}
		}

		if isLogin {
			c.SetLoginedUser(loginedUser)
			c.RedirectWithURL("/user/home")
		}
	}
}

func (c *UserController) Regist() {
	c.TplName = "user_regist.html"
}

func (c *UserController) Home() {
	c.TplName = "user_home.html"
	utils.JJKPrintln(c.Data["User"])
}

func (c *UserController) Logout() {
	c.SetUserLogout()
	c.RedirectWithURL("/user/login")
}


