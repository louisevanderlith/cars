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

func GetAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Ads", "./views/ads.html")

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		src := resources.APIResource(http.DefaultClient, ctx)
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

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}

func SearchAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Ads", "./views/ads.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)
		pagesize := ctx.FindParam("pagesize")

		src := resources.APIResource(http.DefaultClient, ctx)

		result, err := src.FetchStockCars(pagesize)

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = ctx.Serve(http.StatusOK, pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}

func ViewAd(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Ads", "./views/ads.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		key, err := husk.ParseKey(ctx.FindParam("key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		src := resources.APIResource(http.DefaultClient, ctx)
		result, err := src.FetchStockCar(key.String())

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
