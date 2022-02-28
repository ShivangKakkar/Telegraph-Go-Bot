package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

// Logs to Telegram LOG_CHAT and prints to console
func logTg(b *gotgbot.Bot, s string) bool {
	fmt.Println(s)
	if os.Getenv("LOG_CHAT") == "" {
		return true // True because it's not an error and hence considered success
	}
	logGroup, err := strconv.ParseInt(os.Getenv("LOG_CHAT"), 10, 64)
	if err != nil {
		fmt.Printf(parseIntFail, "LOG_CHAT")
		return false
	}
	_, err = b.SendMessage(logGroup, s, &gotgbot.SendMessageOpts{})
	if err != nil {
		fmt.Println("Failed to send message to log chat. Please check if LOG_CHAT is a valid Chat ID")
		return false
	}
	return true
}
