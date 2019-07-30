package controllers

import (
	"github.com/louisevanderlith/cars/logic"
	"github.com/louisevanderlith/droxolite/xontrols"
)

type HomeController struct {
	xontrols.UICtrl
}

func (c *HomeController) Get() {
	c.Setup("home", "Cars", true)
	c.CreateTopMenu(logic.GetMenu())
}
