package controllers

import (
	"net/http"
)

type ProfileController struct {
}

func (c *ProfileController) Get(c *gin.Context) {
	//.Setup("profile", "Profile", false)
	return http.StatusOK, nil
}
