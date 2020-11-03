package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/stock/api"
	"html/template"
	"log"
	"net/http"
)

func GetAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/ads.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllCars(clnt, Endpoints["stock"], "A10")

		if err != nil {
			log.Println("Fetch Cars Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/ads.html")
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	return func(w http.ResponseWriter, r *http.Request) {
		pagesize := drx.FindParam(r, "pagesize")

		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllCars(clnt, Endpoints["stock"], pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewAd(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/ads.html")
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchCar(clnt, Endpoints["stock"], key)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}
