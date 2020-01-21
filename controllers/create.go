package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Create struct {
}

func (req *Create) Get(c *gin.Context) {
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

func (x *Create) Search(c *gin.Context) {
	return http.StatusMethodNotAllowed, nil
}

// POST must start ad upload and verification
func (c *Create) Create(c *gin.Context) {

	// Verify VIN
	// Upload Photos
	// Confirm Vehicle Match
	// Do something with tag...
	// Save Object

	return http.StatusOK, nil
}
