package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/vehicle/api"
	"golang.org/x/oauth2"
	"html/template"
	"log"
	"net/http"
)

func GetAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/ads.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		result, err := api.FetchAllVehicles(clnt, Endpoints["vehicle"], "A10")

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
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		pagesize := drx.FindParam(r, "pagesize")

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		result, err := api.FetchAllVehicles(clnt, Endpoints["vehicle"], pagesize)

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
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		result, err := api.FetchVehicleInfo(clnt, Endpoints["vehicle"], key)

		if err != nil {
			log.Println("Vehicle Info Error", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
