package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/cars/controllers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
	"github.com/louisevanderlith/secure/core/roletype"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	siteName := beego.AppConfig.String("defaultsite")
	theme, err := mango.GetDefaultTheme(ctrlmap.GetInstanceID(), siteName)

	if err != nil {
		panic(err)
	}

	beego.Router("/", controllers.NewHomeCtrl(ctrlmap, theme))
	beego.Router("/profile", controllers.NewProfileCtrl(ctrlmap, theme))

	createCtrl := controllers.NewCreateCtrl(ctrlmap, theme)
	beego.Router("/create", createCtrl, "get:Get")

	beego.Router("/create/step2/:vin", controllers.NewStep2Ctrl(ctrlmap, theme), "get:Get")
	//beego.Router("/create/:step", createCtrl, "get:GetStep")
}

func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/profile", emptyMap)

	emptyMap["GET"] = roletype.Owner
	ctrlmap.Add("/create", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
