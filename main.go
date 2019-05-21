package main

import (
	"log"
	"os"
	"path"

	_ "github.com/louisevanderlith/cars/core"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/cars/routers"
)

func main() {
	// Register with router
	mode := os.Getenv("RUNMODE")
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	pubPath := path.Join(keyPath, pubName)

	name := beego.BConfig.AppName

	srv := mango.NewService(mode, name, pubPath, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		err = mango.UpdateTheme(srv.ID)

		if err != nil {
			panic(err)
		}

		routers.Setup(srv)

		beego.SetStaticPath("/dist", "dist")
		beego.Run()
	}
}
