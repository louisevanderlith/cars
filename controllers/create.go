package controllers

import "github.com/louisevanderlith/droxolite/xontrols"

type CreateController struct {
	xontrols.UICtrl
}

func (req *CreateController) Get() {
	req.Setup("step1", "Create Car", true)
	req.Data["StepNo"] = 1
}

/*
func (c *CreateController) GetStep() {
	step := c.Ctx.Input.Param(":step")
	c.Setup(step, step, true)

	c.Data["StepNo"] = step
}*/

// POST must start ad upload and verification
func (c *CreateController) Post() {

	// Verify VIN
	// Upload Photos
	// Confirm Vehicle Match
	// Do something with tag...
	// Save Object
}
