package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Index", tmpl, "./views/index.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Cars")
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
