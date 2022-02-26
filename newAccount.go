package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

func newAccount(b *gotgbot.Bot, ctx *ext.Context) error {
	id := ctx.EffectiveMessage.From.Id
	if getArgsCount(ctx) < 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "Wrong format. You need to pass a short name to create a new account. \n\nUse <code>/create short_name</code> or <code>/new short_name</code> instead.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send newAccount message : %w", err)
		}
		return nil
	}
	shortName := getString(ctx)
	opts := telegraph.CreateAccountOpts{ShortName: shortName}
	acc, err := telegraph.CreateAccount(opts)
	if err != nil {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("An error occurred while creating the account: \n\n%v", err), &gotgbot.
			SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send newAccount message: %w", err)
		}
		return nil
	}
	accInfoStr := "<b>Account Created</b>"
	accInfoStr += fmt.Sprintf("\n\n<b>Short Name</b> : %v", acc.ShortName)
	accInfoStr += fmt.Sprintf("\n\n<b>Access Token</b> : <code>%v</code>", acc.AccessToken)
	accInfoStr += fmt.Sprintf("\n\n<b>Auth URL</b> : %v", acc.AuthURL)
	_, err = ctx.EffectiveMessage.Reply(b, accInfoStr, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	AddAccount(id, acc.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to send newAccount message: %w", err)
	}
	return nil
}
