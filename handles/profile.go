package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func GetProfile(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Profile", tmpl, "./views/profile.html")
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println(err)
		}
	}
}
