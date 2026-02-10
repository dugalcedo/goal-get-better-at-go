package wog

import (
	"strings"
	"time"

	"github.com/dugalcedo/egoji"
	"github.com/fatih/color"
)

type Wogger struct {
	Name   string
	Head   string
	Emojis map[string]string
}

var logTypeColors = map[string]*color.Color{
	"default":     color.RGB(255, 255, 255),
	"fatal":       color.RGB(128, 0, 0),
	"error":       color.RGB(255, 0, 0),
	"urgent":      color.RGB(255, 165, 0),
	"warning":     color.RGB(255, 255, 0),
	"unintended":  color.RGB(128, 128, 0),
	"intended":    color.RGB(0, 255, 255),
	"success":     color.RGB(0, 255, 0),
	"clientError": color.RGB(255, 141, 185),
	"spam":        color.RGB(183, 170, 133),
}

type W struct {
	Msg   string
	Emoji string
	Type  string
}

func (wogger *Wogger) Wog(w W, args ...any) {
	// log type color
	logTypeColor, ok := logTypeColors[w.Type]
	if !ok {
		logTypeColor = logTypeColors["default"]
	}

	// emoji
	emoji, ok := wogger.Emojis[w.Emoji]
	if !ok {
		emoji, ok = wogger.Emojis["default"]
	}
	if !ok {
		emoji = egoji.RedQuestionMark
	}

	// message text
	msgText := w.Msg
	if msgText == "" {
		msgText = "!-NO-MESSAGE-!"
	}

	// message
	msg := logTypeColor.Sprintf(msgText, args...)

	// timestamp
	now := time.Now()
	timestamp := strings.ToUpper(now.Format("Mon 02 Jan 2006 | 15:04:05"))

	// final print
	logTypeColor.Printf(
		"%s  [%s - %s] %s\n",
		emoji,
		wogger.Head,
		timestamp,
		msg,
	)
}
