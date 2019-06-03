package controllers

import (
	"errors"
	"log"

	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

//Step2Controllers is used to view and confirm VIN details.
type Step2Controller struct {
	control.UIController
}

func NewStep2Ctrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *Step2Controller {
	result := &Step2Controller{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

// /:vin
func (req *Step2Controller) Get() {
	req.Setup("step2", "Validate VIN", true)
	req.Data["StepNo"] = 2

	vin := req.Ctx.Input.Param(":vin")

	if len(vin) < 15 {
		req.Serve(nil, errors.New("vin is too short man"))
		return
	}

	result := make(map[string]interface{})
	_, err := mango.DoGET(req.GetMyToken(), &result, req.GetInstanceID(), "VIN.API", "lookup", vin)

	if err != nil {
		log.Println(err)
		req.Serve(nil, err)
		return
	}

	req.Serve(result, nil)
}
