package rout

import (
	"encoding/json"
	"net/http"

	"github.com/dugalcedo/goal-get-better-at-go/wog"
)

type Context struct {
	W      http.ResponseWriter
	R      *http.Request
	Wogger wog.Wogger
	Data   map[string]any
}

func (c *Context) Reject(status int, msg string) {
	// wog
	c.Wogger.Wog(
		wog.W{
			Emoji: "clientError",
			Type:  "clientError",
			Msg:   "Responding with error %d. Message: '%s'",
		},
		status,
		msg,
	)

	// set headers
	c.W.WriteHeader(status)
	c.W.Header().Set("Content-Type", "application/json")

	// create data obj and send
	if c.Data == nil {
		c.Data = map[string]any{}
	}
	c.Data["msg"] = msg
	c.Data["error"] = true
	json.NewEncoder(c.W).Encode(c.Data)
}

func (c *Context) Respond(status int) {
	// set headers
	c.W.WriteHeader(status)
	c.W.Header().Set("Content-Type", "application/json")
	if c.Data == nil {
		c.Data = map[string]any{}
	}
	json.NewEncoder(c.W).Encode(c.Data)
}
