package controllers

import (
	"github.com/louisevanderlith/mango/control"
	"github.com/louisevanderlith/mango"
)

type ProfileController struct {
	control.UIController
}

func NewProfileCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *ProfileController {
	result := &ProfileController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ProfileController) Get() {
	c.Setup("profile", "Profile", false)
}
