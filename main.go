package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/cars/routers"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/bodies"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/servicetype"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	host := os.Getenv("HOST")
	profile := os.Getenv("PROFILE")
	pubPath := path.Join(keyPath, pubName)

	conf, err := droxolite.LoadConfig()

	if err != nil {
		panic(err)
	}

	// Register with router
	srv := bodies.NewService(conf.Appname, pubPath, conf.HTTPPort, servicetype.APP)

	routr, err := droxolite.GetServiceURL("", "Router.API", false)

	if err != nil {
		panic(err)
	}

	err = srv.Register(routr)

	if err != nil {
		panic(err)
	}

	err = droxolite.UpdateTheme(srv.ID)

	if err != nil {
		panic(err)
	}

	theme, err := droxolite.GetDefaultTheme(host, srv.ID, profile)

	if err != nil {
		panic(err)
	}

	secur, err := droxolite.GetServiceURL(srv.ID, "Auth.APP", true)

	if err != nil {
		panic(err)
	}

	poxy := resins.NewColourEpoxy(srv, theme, "master.html", secur)
	routers.Setup(poxy)

	err = droxolite.Boot(poxy)

	if err != nil {
		panic(err)
	}
}
