package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/kong"
	"net/http"
)

func SetupRoutes(clnt, scrt, secureUrl, authUrl string) http.Handler {
	tmpl, err := droxolite.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	r.HandleFunc("/", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, Index(tmpl), "stock.cars.search")).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, SearchAds(tmpl), "stock.cars.search")).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, SearchAds(tmpl), "stock.cars.search")).Methods(http.MethodGet)
	r.HandleFunc("/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, ViewAd(tmpl), "stock.cars.view")).Methods(http.MethodGet)

	crt := r.PathPrefix("/create").Subrouter()
	crt.HandleFunc("", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, GetCreation(tmpl), "vin.validate")).Methods(http.MethodGet)
	crt.HandleFunc("/step2/{vin:[a-zA-Z0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, GetStep2(tmpl), "vin.validate")).Methods(http.MethodGet)
	crt.HandleFunc("/step3/{key:[0-9]+\\x60[0-9]+}", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, authUrl, GetStep3(tmpl), "vin.validate")).Methods(http.MethodGet)

	return r
}
