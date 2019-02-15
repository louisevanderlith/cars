package main

import (
	"log"

	_ "github.com/louisevanderlith/mango/core/auto"
	"github.com/louisevanderlith/mango/enums"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/auto/routers"
)

func main() {
	// Register with router
	mode := beego.BConfig.RunMode
	name := beego.BConfig.AppName

	srv := mango.NewService(mode, name, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
