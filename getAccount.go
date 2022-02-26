package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

func getAccount(b *gotgbot.Bot, ctx *ext.Context) error {
	args := getArgs(ctx)
	if len(args) < 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "Wrong format. You need to pass account number to get account info for an account.", &gotgbot.SendMessageOpts{ParseMode: "html", DisableWebPagePreview: true})
		if err != nil {
			return fmt.Errorf("failed to send getAccount message : %w", err)
		}
		return nil
	}
	accountNumber := args[0]
	accountNumber, con, err := getToken(accountNumber, b, ctx)
	if err != nil {
		return fmt.Errorf("failed to send getAccount tokens message : %w", err)
	}
	if con {
		return nil
	}
	acc, err := telegraph.GetAccountInfo(telegraph.GetAccountInfoOpts{AccessToken: accountNumber, Fields: []string{"short_name", "page_count", "author_name", "author_url", "auth_url"}})
	if err != nil {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("An Error Occured \n\n%v", err), &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send getAccount message : %w", err)
		}
		return nil
	}
	if acc.AuthorName == "" {
		acc.AuthorName = "Not Set"
	}
	if acc.AuthorURL == "" {
		acc.AuthorURL = "Not Set"
	}
	infoStr := "<b>Account Info</b>"
	infoStr += fmt.Sprintf("\n\n<b>Access Token</b> : <code>%v</code>", accountNumber)
	infoStr += fmt.Sprintf("\n\n<b>Short Name</b> : %v", acc.ShortName)
	infoStr += fmt.Sprintf("\n\n<b>Author Name</b> : %v", acc.AuthorName)
	infoStr += fmt.Sprintf("\n\n<b>Author URL</b> : %v", acc.AuthorURL)
	infoStr += fmt.Sprintf("\n\n<b>Total Pages</b> : %v", acc.PageCount)
	infoStr += fmt.Sprintf("\n\n<b>Authorization URL</b> : %v", acc.AuthURL)
	_, err = ctx.EffectiveMessage.Reply(b, infoStr, &gotgbot.SendMessageOpts{ParseMode: "html", DisableWebPagePreview: true})
	if err != nil {
		return fmt.Errorf("failed to send getAccount status message: %w", err)
	}
	return nil
}
