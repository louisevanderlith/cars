package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/stock/api"
	"github.com/louisevanderlith/stock/core"
	"html/template"
	"log"
	"net/http"
)

func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Index", tmpl, "./views/index.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Cars")

		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllCars(clnt, Endpoints["stock"], "A10")

		if err != nil {
			log.Println("Fetch Cars Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		err = mix.Write(w, pge.Create(r, GenerateStockStats(result)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

type StockStats struct {
	Results  records.Page
	Min      int64
	Max      int64
	TagStats map[string]int
}

func GenerateStockStats(items records.Page) StockStats {
	itor := items.GetEnumerator()

	min := int64(999999)
	max := int64(0)
	stats := make(map[string]int)
	for itor.MoveNext() {
		rec := itor.Current().(core.Car)

		if rec.Price < min {
			min = rec.Price
		}

		if rec.Price > max {
			max = rec.Price
		}

		for _, tag := range rec.Tags {
			stat, ok := stats[tag]

			if !ok {
				stats[tag] = 0
			} else {
				stats[tag] = stat + 1
			}
		}
	}

	return StockStats{
		Results:  items,
		Min:      min,
		Max:      max,
		TagStats: stats,
	}
}
