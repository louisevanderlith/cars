package handles

import (
	"github.com/louisevanderlith/cars/resources"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"html/template"
	"log"
	"net/http"
)

func GetAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/ads.html")

	return func(w http.ResponseWriter, r *http.Request) {

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockCars("A10")

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

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

func SearchAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/ads.html")
	return func(w http.ResponseWriter, r *http.Request) {
		pagesize := drx.FindParam(r, "pagesize")

		src := resources.APIResource(http.DefaultClient, r)

		result, err := src.FetchStockCars(pagesize)

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
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, r)
		result, err := src.FetchStockCar(key.String())

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
