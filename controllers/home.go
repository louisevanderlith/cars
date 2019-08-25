package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Home struct {
}

func (c *Home) Default(ctx context.Contexer) (int, interface{}) {
	//c.Setup("home", "Cars", true)
	//c.CreateTopMenu(logic.GetMenu())

	return http.StatusOK, nil
}
