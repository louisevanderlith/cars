package controllers

import (
	"github.com/louisevanderlith/cars/logic"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type HomeController struct {
	control.UIController
}

func NewHomeCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *HomeController {
	result := &HomeController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *HomeController) Get() {
	c.Setup("home", "Cars", true)
	c.CreateTopMenu(c.Ctx, logic.GetMenu("/create"))
}
