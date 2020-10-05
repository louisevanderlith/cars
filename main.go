package main

import (
	"flag"
	"github.com/louisevanderlith/cars/handles"
	"github.com/louisevanderlith/droxolite/drx"
	"net/http"
	"time"
)

func main() {
	clientId := flag.String("client", "mango.cars", "Client ID which will be used to verify this instance")
	clientSecrt := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	authrty := flag.String("authority", "http://localhost:8094", "Authority Provider's URL")
	manager := flag.String("manager", "http://localhost:8097", "User Provider's URL")
	security := flag.String("security", "http://localhost:8086", "Security Provider's URL")

	flag.Parse()

	err := drx.UpdateTemplate(*clientId, *clientSecrt, *security)

	if err != nil {
		panic(err)
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8081",
		Handler:      handles.SetupRoutes(*clientId, *clientSecrt, *security, *manager, *authrty),
	}

	err = srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
