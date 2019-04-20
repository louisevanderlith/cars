package controllers

import (
	"github.com/louisevanderlith/mango/control"
)

type ProfileController struct {
	control.UIController
}

func NewProfileCtrl(ctrlMap *control.ControllerMap) *ProfileController {
	result := &ProfileController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *ProfileController) Get() {
	c.Setup("profile", "Profile", false)
}
