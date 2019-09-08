package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
)

//Step2Controllers is used to view and confirm VIN details.
type Step2Controller struct {
}

// /:vin
func (req *Step2Controller) Get(ctx context.Requester) (int, interface{}) {
	//req.Setup("step2", "Validate VIN", true)
	//req.Data["StepNo"] = 2

	vin := ctx.FindParam("vin")

	if len(vin) < 15 {
		return http.StatusBadRequest, errors.New("vin is too short man")
	}

	result := make(map[string]interface{})
	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "VIN.API", "lookup", vin)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}
