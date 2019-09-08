package controllers

import (
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/do"
)

type Ads struct {
}

func (c *Ads) Get(ctx context.Requester) (int, interface{}) {
	result := []interface{}{}
	pagesize := ctx.FindParam("pagesize")

	code, err := do.GET(ctx.GetMyToken(), &result, ctx.GetInstanceID(), "Stock.API", "car", pagesize)

	if err != nil {
		log.Println(err)
		return code, err
	}

	return http.StatusOK, result
}

func (c *Ads) View(ctx context.Requester) (int, interface{}) {
	//c.Setup("adView", "View Ad", true)

	return http.StatusOK, nil
}
