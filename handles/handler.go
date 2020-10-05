package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong/middle"
	"net/http"
)

func SetupRoutes(clnt, scrt, secureUrl, managerUrl, authUrl string) http.Handler {
	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	clntIns := middle.NewClientInspector(clnt, scrt, http.DefaultClient, secureUrl, managerUrl, authUrl)
	r.HandleFunc("/", clntIns.Middleware(Index(tmpl), map[string]bool{"stock.cars.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}", clntIns.Middleware(SearchAds(tmpl), map[string]bool{"stock.cars.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", clntIns.Middleware(SearchAds(tmpl), map[string]bool{"stock.cars.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/{key:[0-9]+\\x60[0-9]+}", clntIns.Middleware(ViewAd(tmpl), map[string]bool{"stock.cars.view": true})).Methods(http.MethodGet)

	crt := r.PathPrefix("/create").Subrouter()
	crt.HandleFunc("", clntIns.Middleware(GetCreation(tmpl), map[string]bool{"vin.validate": true})).Methods(http.MethodGet)
	crt.HandleFunc("/step2/{vin:[a-zA-Z0-9]+}", clntIns.Middleware(GetStep2(tmpl), map[string]bool{"vin.lookup": true})).Methods(http.MethodGet)
	crt.HandleFunc("/step3/{key:[0-9]+\\x60[0-9]+}", clntIns.Middleware(GetStep3(tmpl), map[string]bool{"vin.validate": true})).Methods(http.MethodGet)

	return r
}
