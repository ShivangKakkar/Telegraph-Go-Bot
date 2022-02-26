package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func removeAccount(b *gotgbot.Bot, ctx *ext.Context) error {
	id := ctx.EffectiveMessage.From.Id
	args := getArgs(ctx)
	if len(args) < 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "Wrong format. You need to pass an account number to remove an account. \n\nUse <code>/remove account_number</code> instead.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send removeAccount message : %w", err)
		}
		return nil
	}
	accountNumber := args[0]
	accountNumber, con, err := getToken(accountNumber, b, ctx)
	if err != nil {
		return fmt.Errorf("failed to send removeAccount tokens message : %w", err)
	}
	if con {
		return nil
	}
	e := false
	tokens := GetAllAccounts(ctx.EffectiveMessage.From.Id)
	for _, j := range tokens {
		if j == accountNumber {
			e = true
			break
		}
	}
	if !e {
		_, err = ctx.EffectiveMessage.Reply(b, "This account doesn't exist in my database anyway.", &gotgbot.
			SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send rmeoveAccount message: %w", err)
		}
		return nil
	}
	RemoveAccount(id, accountNumber)
	_, err = ctx.EffectiveMessage.Reply(b, "Account removed successfully", &gotgbot.SendMessageOpts{})
	if err != nil {
		return fmt.Errorf("failed to send removeAccount message: %w", err)
	}
	return nil
}
