package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

func listPages(b *gotgbot.Bot, ctx *ext.Context) error {
	args := getArgs(ctx)
	if len(args) < 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "Wrong format. You need to pass account number to list pages for an account.", &gotgbot.SendMessageOpts{ParseMode: "html", DisableWebPagePreview: true})
		if err != nil {
			return fmt.Errorf("failed to send listPages message : %w", err)
		}
		return nil
	}
	accountNumber := args[0]
	accountNumber, con, err := getToken(accountNumber, b, ctx)
	if err != nil {
		return fmt.Errorf("failed to send listPages tokens message : %w", err)
	}
	if con {
		return nil
	}
	pList, err := telegraph.GetPageList(telegraph.GetPageListOpts{AccessToken: accountNumber, Limit: 200})
	if err != nil {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("An Error Occured \n\n%v", err), &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send listPages message : %w", err)
		}
		return nil
	}
	infoStr := fmt.Sprintf("<b>Total Pages</b> : %v", pList.TotalCount)
	for i, page := range pList.Pages {
		infoStr += fmt.Sprintf("\n\n%v) <b>Title</b>: %v \n\n<b>Path</b>: <code>%v</code> \n\n<b>URL</b>: %v \n\n<b>Views</b>: %v", i+1, page.Title, page.Path, page.URL, page.Views)
	}
	l := checkIfLengthy(infoStr, b, ctx)
	if l {
		return nil
	}
	_, err = ctx.EffectiveMessage.Reply(b, infoStr, &gotgbot.SendMessageOpts{ParseMode: "html"})
	if err != nil {
		return fmt.Errorf("failed to send listPages status message: %w", err)
	}
	return nil
}
