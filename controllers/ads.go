package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type AdsController struct {
	control.UIController
}

func NewAdsCtrl(ctrlMap *control.ControllerMap, settings mango.ThemeSetting) *AdsController {
	result := &AdsController{}
	result.SetTheme(settings)
	result.SetInstanceMap(ctrlMap)

	return result
}

func (c *AdsController) Get() {
	c.Setup("ads", "For Sale", false)

	result := []interface{}{}
	pagesize := c.Ctx.Input.Param(":pagesize")

	_, err := mango.DoGET(c.GetMyToken(), &result, c.GetInstanceID(), "Stock.API", "car", "all", pagesize)

	c.Serve(result, err)
}

func (c *AdsController) GetView() {
	c.Setup("adView", "View Ad", true)

}
