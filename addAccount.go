package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

func addAccount(b *gotgbot.Bot, ctx *ext.Context) error {
	id := ctx.EffectiveMessage.From.Id
	if getArgsCount(ctx) < 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "Wrong format. You need to pass an access token to add an existing account to my database. \n\nUse <code>/add token</code> instead.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send addAccount message : %w", err)
		}
		return nil
	}
	token := getArgs(ctx)[0]
	_, err := telegraph.GetAccountInfo(telegraph.GetAccountInfoOpts{AccessToken: token})
	if err != nil {
		text := fmt.Sprintf("An error occurred while checking the account: \n\n%v", err)
		if err.Error() == "ACCESS_TOKEN_INVALID" {
			text = "This access token is invalid. Please pass a valid one."
		}
		_, err = ctx.EffectiveMessage.Reply(b, text, &gotgbot.
			SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send addAccount message: %w", err)
		}
		return nil
	}
	e := false
	tokens := GetAllAccounts(ctx.EffectiveMessage.From.Id)
	for _, j := range tokens {
		if j == token {
			e = true
			break
		}
	}
	if e {
		_, err = ctx.EffectiveMessage.Reply(b, "This account already exists in my database.", &gotgbot.
			SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send addAccount message: %w", err)
		}
		return nil
	}
	AddAccount(id, token)
	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Successfully added your account in my database. It's account number is %v", len(tokens)), &gotgbot.
		SendMessageOpts{ParseMode: "html"})
	if err != nil {
		return fmt.Errorf("failed to send addAccount message: %w", err)
	}
	return nil
}
