package rout

import (
	"fmt"
	"net/http"

	"github.com/dugalcedo/egoji"
	"github.com/dugalcedo/goal-get-better-at-go/wog"
)

type Router struct {
	Mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		Mux: http.NewServeMux(),
	}
}

func (r *Router) Handle(pattern string, handler func(ctx Context)) {
	wogger := wog.Wogger{
		Head: "New request",
		Emojis: map[string]string{
			"default":     egoji.Bird,
			"GET":         egoji.Giraffe,
			"POST":        egoji.PolarBear,
			"PUT":         egoji.Parrot,
			"PATCH":       egoji.Phoenix,
			"DELETE":      egoji.Deer,
			"clientError": egoji.Cat,
		},
	}

	r.Mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		wogger.Wog(
			wog.W{
				Emoji: r.Method,
				Msg:   "%s %s",
			},
			r.Method,
			r.URL,
		)

		ctx := Context{
			W:      w,
			R:      r,
			Wogger: wogger,
		}

		handler(ctx)
	})
}

func (r *Router) Listen(port string) {
	http.ListenAndServe(fmt.Sprintf(":%s", port), r.Mux)
}
