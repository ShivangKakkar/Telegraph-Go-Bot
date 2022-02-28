package main

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// UserID : Title
var titles = map[int64]string{}

func setTitle(b *gotgbot.Bot, ctx *ext.Context) error {
	id := ctx.EffectiveMessage.From.Id
	args := getArgs(ctx)
	var m string
	if len(args) == 0 {
		m = titles[id]
		if m == "" {
			m = "No Title is currently set. Set one using <code>/title your_title_here</code>."
		} else {
			m = fmt.Sprintf("Your Current Title : <code>%v</code>", m)
		}
	} else {
		m = strings.Join(args, " ")
		titles[id] = m
		m = "Title set successfully."
	}
	_, err := ctx.EffectiveMessage.Reply(b, m, &gotgbot.SendMessageOpts{ParseMode: "html"})
	if err != nil {
		return fmt.Errorf("failed to send setTitle message: %w", err)
	}
	return nil
}
