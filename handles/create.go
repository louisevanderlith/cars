package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/vin/api"
	"html/template"
	"log"
	"net/http"
)

type Step struct {
	No   int
	Info interface{}
}

func GetCreation(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step 1", tmpl, "./views/create.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		err := mix.Write(w, pge.Create(r, Step{No: 1}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

// /:vin
func GetStep2(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step 2", tmpl, "./views/create.html")
	return func(w http.ResponseWriter, r *http.Request) {
		vin := drx.FindParam(r, "vin")

		if len(vin) < 15 {
			http.Error(w, "vin is too short", http.StatusBadRequest)
			return
		}

		clnt := CredConfig.Client(r.Context())
		result, err := api.LookupVIN(clnt, Endpoints["vin"], vin)

		if err != nil {
			log.Println("VIN Lookup Error", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		err = mix.Write(w, pge.Create(r, Step{No: 2, Info: result}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep3(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step 3", tmpl, "./views/create.html")
	return func(w http.ResponseWriter, r *http.Request) {
		vehicleKey, err := keys.ParseKey(drx.FindParam(r, "vehicleKey"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		pge.ChangeTitle("Sell your Vehicle: Step 3")
		err = mix.Write(w, pge.Create(r, Step{No: 3, Info: vehicleKey}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep4(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step 4", tmpl, "./views/create.html")
	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Sell your Vehicle: Step 4")
		err := mix.Write(w, pge.Create(r, Step{No: 4}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep5(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step 5", tmpl, "./views/create.html")
	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Sell your Vehicle: Step 5")
		err := mix.Write(w, pge.Create(r, Step{No: 5}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep6(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step 6", tmpl, "./views/create.html")
	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Sell your Vehicle: Step 6")
		err := mix.Write(w, pge.Create(r, Step{No: 6}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep7(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step 7", tmpl, "./views/create.html")
	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Sell your Vehicle: Step 7")
		err := mix.Write(w, pge.Create(r, Step{No: 7}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep8(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Step 8", tmpl, "./views/create.html")
	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Sell your Vehicle: Step 8")
		err := mix.Write(w, pge.Create(r, Step{No: 8}))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
