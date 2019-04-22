package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/cars/controllers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	createCtrl := controllers.NewCreateCtrl(ctrlmap)

	beego.Router("/", controllers.NewHomeCtrl(ctrlmap))
	beego.Router("/profile", controllers.NewProfileCtrl(ctrlmap))
	beego.Router("/create", createCtrl, "get:Get")
	beego.Router("/create/:step", createCtrl, "get:GetStep")
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)

	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/profile", emptyMap)
	ctrlmap.Add("/create", emptyMap)


	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
