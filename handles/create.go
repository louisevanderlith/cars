package handles

import (
	"github.com/louisevanderlith/cars/resources"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk"
	"html/template"
	"log"
	"net/http"
)

type Step struct {
	StepNo int
	Info   interface{}
}

func GetCreation(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Create", tmpl, "./views/create.html")
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, Step{StepNo: 1}))

		if err != nil {
			log.Println(err)
		}
	}
}

// /:vin
func GetStep2(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step2", tmpl, "./views/step2.html")
	return func(w http.ResponseWriter, r *http.Request) {
		vin := drx.FindParam(r, "vin")

		if len(vin) < 15 {
			http.Error(w, "vin is too short", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.LookupVIN(vin)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		err = mix.Write(w, pge.Create(r, Step{StepNo: 2, Info: result}))

		if err != nil {
			log.Println(err)
		}
	}
}

func GetStep3(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step3", tmpl, "./views/step3.html")
	return func(w http.ResponseWriter, r *http.Request) {
		vehicleKey, err := husk.ParseKey(drx.FindParam(r, "vehicleKey"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, Step{StepNo: 3, Info: vehicleKey}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
