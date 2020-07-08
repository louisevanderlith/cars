package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Index", "./views/index.html")
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		err := ctx.Serve(http.StatusOK, pge.Page(nil, ctx.GetTokenInfo(), ctx.GetToken()))

		if err != nil {
			log.Println(err)
		}
	}
}
