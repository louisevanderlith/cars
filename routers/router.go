package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/cars/controllers"
	"github.com/louisevanderlith/mango/control"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	beego.Router("/", controllers.NewHomeCtrl(ctrlmap))
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(control.ActionMap)

	ctrlmap.Add("/", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
