package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

func account(b *gotgbot.Bot, ctx *ext.Context) error {
	status, err := ctx.EffectiveMessage.Reply(b, "Sending...", &gotgbot.SendMessageOpts{})
	if err != nil {
		return fmt.Errorf("failed to send listAccount message: %w", err)
	}
	id := ctx.EffectiveMessage.From.Id
	tokens := GetAllAccounts(id)
	infoStr := fmt.Sprintf("Total Accounts : %v", len(tokens))
	for i, token := range tokens {
		acc, err := telegraph.GetAccountInfo(telegraph.GetAccountInfoOpts{AccessToken: token})
		if err != nil {
			infoStr = "An Error Occured : " + err.Error()
			break
		}
		if i == 0 {
			infoStr += "\n\n0) <u>Default Account</u> \n\n "
		} else {
			infoStr += fmt.Sprintf("\n\n%v)", i)
		}
		infoStr += fmt.Sprintf(" Short Name: %v \n\n   Access Token: <code>%v</code>", acc.ShortName, token)
		if acc.AuthorName != "" {
			infoStr += fmt.Sprintf("\n\n   Author Name: %v", acc.AuthorName)
		}
		if acc.AuthURL != "" {
			infoStr += fmt.Sprintf("\n\n   Auth URL: %v", acc.AuthURL)
		}
		if acc.AuthorURL != "" {
			infoStr += fmt.Sprintf("\n\n   Author URL: %v", acc.AuthorURL)
		}
		if acc.PageCount != 0 {
			infoStr += fmt.Sprintf("\n\n   Page Count: %v", acc.PageCount)
		}
	}
	l := checkIfLengthy(infoStr, b, ctx)
	if l {
		return nil
	}
	status.EditText(b, infoStr, &gotgbot.EditMessageTextOpts{
		ChatId:                status.Chat.Id,
		MessageId:             status.MessageId,
		ParseMode:             "html",
		DisableWebPagePreview: true,
	})
	return nil
}
