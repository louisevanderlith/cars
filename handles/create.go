package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/vin/api"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

type Step struct {
	No   int
	Info interface{}
}

func GetCreation(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := Step{No: 1}
		err := mix.Write(w, fact.Create(r, "Step 1", "./views/create.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

// /:vin
func GetStep2(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vin := drx.FindParam(r, "vin")

		if len(vin) < 15 {
			http.Error(w, "vin is too short", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		info, err := api.LookupVIN(clnt, Endpoints["vin"], vin)

		if err != nil {
			log.Println("VIN Lookup Error", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		data := Step{No: 2, Info: info}
		err = mix.Write(w, fact.Create(r, "Step 2", "./views/create.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep3(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vehicleKey, err := keys.ParseKey(drx.FindParam(r, "vehicleKey"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		//pge.ChangeTitle("Sell your Vehicle: Step 3")
		data := Step{No: 3, Info: vehicleKey}
		err = mix.Write(w, fact.Create(r, "Step 3", "./views/create.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep4(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pge.ChangeTitle("Sell your Vehicle: Step 4")
		data := Step{No: 4}
		err := mix.Write(w, fact.Create(r, "Step 4", "./views/create.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep5(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pge.ChangeTitle("Sell your Vehicle: Step 5")
		data := Step{No: 5}
		err := mix.Write(w, fact.Create(r, "Step 5", "./views/create.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep6(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pge.ChangeTitle("Sell your Vehicle: Step 6")
		data := Step{No: 6}
		err := mix.Write(w, fact.Create(r, "Step 6", "./views/create.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep7(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pge.ChangeTitle("Sell your Vehicle: Step 7")
		data := Step{No: 7}
		err := mix.Write(w, fact.Create(r, "Step 7", "./views/create.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func GetStep8(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pge.ChangeTitle("Sell your Vehicle: Step 8")
		data := Step{No: 8}
		err := mix.Write(w, fact.Create(r, "Step 8", "./views/create.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
