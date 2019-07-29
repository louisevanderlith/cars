package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
)

//Step2Controllers is used to view and confirm VIN details.
type Step2Controller struct {
	xontrols.UICtrl
}

// /:vin
func (req *Step2Controller) Get() {
	req.Setup("step2", "Validate VIN", true)
	req.Data["StepNo"] = 2

	vin := req.FindParam("vin")

	if len(vin) < 15 {
		req.Serve(http.StatusBadRequest, nil, errors.New("vin is too short man"))
		return
	}

	result := make(map[string]interface{})
	code, err := droxolite.DoGET(req.GetMyToken(), &result, req.Settings.InstanceID, "VIN.API", "lookup", vin)

	if err != nil {
		log.Println(err)
		req.Serve(code, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, result)
}
