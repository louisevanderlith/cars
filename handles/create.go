package handles

import (
	"github.com/louisevanderlith/cars/resources"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

func GetCreation(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Create", "./views/create.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		err := ctx.Serve(http.StatusOK, pge.Page(nil, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}

// /:vin
func GetStep2(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Step2", "./views/step2.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		vin := ctx.FindParam("vin")

		if len(vin) < 15 {
			http.Error(w, "vin is too short", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.LookupVIN(vin)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}

func GetStep3(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Step3", "./views/step3.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		vehicleKey, err := husk.ParseKey(ctx.FindParam("vehicleKey"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(vehicleKey, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}
