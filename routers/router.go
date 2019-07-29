package routers

import (
	"github.com/louisevanderlith/cars/controllers"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e *droxolite.Epoxy) {
	//Home
	homeCtrl := &controllers.HomeController{}
	homeGroup := droxolite.NewRouteGroup("", homeCtrl)
	homeGroup.AddRoute("/", "GET", roletype.Unknown, homeCtrl.Get)
	e.AddGroup(homeGroup)
	/*ctrlmap := EnableFilter(s)

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
	beego.Router("/create/step3/:vehicleKey", controllers.NewStep3Ctrl(ctrlmap, theme), "get:Get")
	//beego.Router("/create/:step", createCtrl, "get:GetStep")
	*/
}

/*
func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	ctrlmap.Add("/", emptyMap)
	ctrlmap.Add("/profile", emptyMap)

	ownMap := make(secure.ActionMap)
	ownMap["GET"] = roletype.Owner
	ctrlmap.Add("/create", ownMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
*/
