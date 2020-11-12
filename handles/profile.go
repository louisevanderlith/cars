package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func GetProfile(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Profile", tmpl, "./views/profile.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println(err)
		}
	}
}
