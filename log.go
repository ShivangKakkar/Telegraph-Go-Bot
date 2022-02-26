package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

// Logs to Telegram Group and Prints to console
func logTg(b *gotgbot.Bot, s string) {
	var logGroup int64 = -1111111111111
	_, err := b.SendMessage(logGroup, s, &gotgbot.SendMessageOpts{})
	if err != nil {
		fmt.Println(err)
	}
}
