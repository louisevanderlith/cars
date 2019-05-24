package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type CreateController struct {
	control.UIController
}

func NewCreateCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *CreateController {
	result := &CreateController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *CreateController) Get() {
	c.Setup("step1", "Create Car", true)
	c.Data["StepNo"] = "step1"
}

func (c *CreateController) GetStep() {
	step := c.Ctx.Input.Param(":step")
	c.Setup(step, step, true)

	c.Data["StepNo"] = step
}

// POST must start ad upload and verification
func (c *CreateController) Post() {

	// Verify VIN
	// Upload Photos
	// Confirm Vehicle Match
	// Do something with tag...
	// Save Object
}
