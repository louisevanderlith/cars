package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type CreateController struct {
}

func (req *CreateController) Get(ctx context.Contexer) (int, interface{}) {
	//req.Setup("step1", "Create Car", true)
	//req.Data["StepNo"] = 1

	return http.StatusOK, nil
}

/*
func (c *CreateController) GetStep() {
	step := c.Ctx.Input.Param(":step")
	c.Setup(step, step, true)

	c.Data["StepNo"] = step
}*/

// POST must start ad upload and verification
func (c *CreateController) Post(ctx context.Contexer) (int, interface{}) {

	// Verify VIN
	// Upload Photos
	// Confirm Vehicle Match
	// Do something with tag...
	// Save Object

	return http.StatusOK, nil
}
