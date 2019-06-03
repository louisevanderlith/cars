package controllers

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

//Step3Controllers is used to add images of the vehicle
type Step3Controller struct {
	control.UIController
}

func NewStep3Ctrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *Step3Controller {
	result := &Step3Controller{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

// /:vehicleKey
func (req *Step3Controller) Get() {
	req.Setup("step3", "Upload Images", true)
	req.Data["StepNo"] = 3

	vehicleKey, err := husk.ParseKey(req.Ctx.Input.Param(":vehicleKey"))

	if err != nil {
		req.Serve(nil, err)
		return
	}

	req.Serve(vehicleKey, nil)
}
