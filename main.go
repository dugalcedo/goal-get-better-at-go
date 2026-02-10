package main

import (
	"fmt"

	"github.com/dugalcedo/goal-get-better-at-go/env"
	"github.com/dugalcedo/goal-get-better-at-go/rout"
)

// ===== MAIN =====
func main() {
	vars := env.Vars()

	router := rout.NewRouter()

	router.Handle("/", func(ctx rout.Context) {
		if ctx.R.Method == "GET" && ctx.R.URL.String() == "/" {
			ctx.Data = map[string]any{
				"message": "Hello",
			}
			ctx.Respond(200)
			return
		}

		// not found
		ctx.Reject(404, "Not found")
	})

	fmt.Printf("Now listening on port %s\n", vars.PORT)
	router.Listen(vars.PORT)
}
