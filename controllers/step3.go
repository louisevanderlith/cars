package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/husk"
)

//Step3Controllers is used to add images of the vehicle
type Step3Controller struct {
}

// /:vehicleKey
func (req *Step3Controller) Get(ctx context.Contexer) (int, interface{}) {
	//req.Setup("step3", "Upload Images", true)
	//req.Data["StepNo"] = 3

	vehicleKey, err := husk.ParseKey(ctx.FindParam("vehicleKey"))

	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, vehicleKey
}
