package controllers

import (
)

type MainController struct {
	BaseController
}

func (c *MainController) Prepare() {
	c.BaseController.Prepare()
}

func (c *MainController) Finish() {
	c.Controller.Finish()
}

func (c *MainController) Get() {
	c.TplName = "user_login.html"
}

func (c *MainController) UserLogin() {
	c.TplName = "user_login.html"
}

func (c *MainController) UserRegist() {
	c.TplName = "user_regist.html"
}


