package main

import (
	"os"
	"path"

	"github.com/louisevanderlith/cars/routers"
	"github.com/louisevanderlith/droxolite"
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
	srv := droxolite.NewService(conf.Appname, pubPath, conf.HTTPPort, servicetype.APP)

	err = srv.Register()

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

	poxy := droxolite.NewColourEpoxy(srv, theme, "master.html")
	routers.Setup(poxy)

	err = poxy.Boot()

	if err != nil {
		panic(err)
	}
}
