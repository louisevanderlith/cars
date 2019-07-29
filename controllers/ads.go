package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/xontrols"
)

type AdsController struct {
	xontrols.UICtrl
}

func (c *AdsController) Get() {
	c.Setup("ads", "For Sale", false)

	result := []interface{}{}
	pagesize := c.FindParam("pagesize")

	code, err := droxolite.DoGET(c.GetMyToken(), &result, c.Settings.InstanceID, "Stock.API", "car", "all", pagesize)

	if err != nil {
		log.Println(err)
		c.Serve(code, err, nil)
		return
	}

	c.Serve(http.StatusOK, nil, result)
}

func (c *AdsController) GetView() {
	c.Setup("adView", "View Ad", true)

	c.Serve(http.StatusOK, nil, nil)
}
