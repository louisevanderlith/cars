package routers

import (
	"github.com/louisevanderlith/droxolite/resins"
)

func Setup(e resins.Epoxi) {
	//Home
	//homeGroup := routing.NewInterfaceBundle("Home", roletype.Admin, &controllers.Home{})
	//e.AddBundle(homeGroup)
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
