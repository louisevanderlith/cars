package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/husk"
)

//Step3Controllers is used to add images of the vehicle
type Step3Controller struct {
	xontrols.UICtrl
}

// /:vehicleKey
func (req *Step3Controller) Get() {
	req.Setup("step3", "Upload Images", true)
	req.Data["StepNo"] = 3

	vehicleKey, err := husk.ParseKey(req.FindParam("vehicleKey"))

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, vehicleKey)
}
