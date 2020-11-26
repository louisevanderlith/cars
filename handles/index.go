package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/records"
	"github.com/louisevanderlith/vehicle/api"
	"github.com/louisevanderlith/vehicle/core"
	"golang.org/x/oauth2"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Index", tmpl, "./views/index.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Cars")

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		result, err := api.FetchAllVehicles(clnt, Endpoints["vehicle"], "A10")

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
	if items.GetRecords() == nil {
		return StockStats{}
	}

	itor := items.GetEnumerator()

	min := int64(999999)
	max := int64(0)
	stats := make(map[string]int)
	for itor.MoveNext() {
		rec := itor.Current().(hsk.Record)
		obj := rec.GetValue().(*core.Vehicle)
		if obj.EstValue < min {
			min = obj.EstValue
		}

		if obj.EstValue > max {
			max = obj.EstValue
		}

		tags := []string{
			strconv.Itoa(obj.Year),
			obj.Series.Manufacturer,
			obj.Series.Model,
			obj.Info,
		}
		tags = append(tags, obj.Extra...)

		for _, tag := range tags {
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
