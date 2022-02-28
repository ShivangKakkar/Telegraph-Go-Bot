package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func stats(b *gotgbot.Bot, ctx *ext.Context) error {
	if os.Getenv("OWNER_ID") == "" {
		fmt.Println("Set OWNER_ID to use this command.")
		return nil
	}
	o, err := strconv.ParseInt(os.Getenv("OWNER_ID"), 10, 64)
	if err != nil {
		fmt.Printf(parseIntFail, "OWNER_ID")
		return nil
	}
	if ctx.EffectiveMessage.From.Id != o {
		return nil
	}
	c := UsersCount()
	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Total Users : %v", c), &gotgbot.SendMessageOpts{})
	if err != nil {
		return fmt.Errorf("failed to send removeAccount message: %w", err)
	}
	return nil
}
