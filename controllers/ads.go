package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/context"
)

type Ads struct {
}

func (c *Ads) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("ads", "For Sale", false)

	result := []interface{}{}
	pagesize := ctx.FindParam("pagesize")

	code, err := droxolite.DoGET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "car", "all", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Ads) View(ctx context.Contexer) (int, interface{}) {
	//c.Setup("adView", "View Ad", true)

	return http.StatusOK, nil
}
