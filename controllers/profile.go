package controllers

import "github.com/louisevanderlith/droxolite/xontrols"

type ProfileController struct {
	xontrols.UICtrl
}

func (c *ProfileController) Get() {
	c.Setup("profile", "Profile", false)
}
