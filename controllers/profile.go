package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type ProfileController struct {
}

func (c *ProfileController) Get(ctx context.Requester) (int, interface{}) {
	//.Setup("profile", "Profile", false)
	return http.StatusOK, nil
}
