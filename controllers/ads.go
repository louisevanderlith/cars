package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/do"
)

type Ads struct {
}

func (c *Ads) Get(c *gin.Context) {
	result := []interface{}{}
	pagesize := c.Param("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "car", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Ads) View(c *gin.Context) {
	//c.Setup("adView", "View Ad", true)

	return http.StatusOK, nil
}
