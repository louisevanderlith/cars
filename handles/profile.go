package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func GetProfile(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "Profile", "./views/profile.html", nil))

		if err != nil {
			log.Println(err)
		}
	}
}
