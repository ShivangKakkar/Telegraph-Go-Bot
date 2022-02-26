package main

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

func editPage(b *gotgbot.Bot, ctx *ext.Context) error {
	content, err := extractHTMLFromReply(b, ctx)
	if err != nil {
		return err
	}
	if content == "" {
		return nil
	}
	args := getArgs(ctx)
	if len(args) < 3 {
		_, err := ctx.EffectiveMessage.Reply(b, "Wrong format. You need to pass account number, page title and page path to edit a page. \n\nUse <code>/editpage account_number page_path page_title &lt;in reply to html message&gt;</code> instead.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send editPage message : %w", err)
		}
		return nil
	}
	accountNumber := args[0]
	pagePath := args[1]
	pageTitle := strings.Join(args[2:], " ")
	accountNumber, con, err := getToken(accountNumber, b, ctx)
	if err != nil {
		return fmt.Errorf("failed to send editPage tokens message : %w", err)
	}
	if con {
		return nil
	}
	pg, err := telegraph.EditPage(telegraph.EditPageOpts{Path: pagePath, AccessToken: accountNumber, Title: pageTitle, HTMLContent: content})
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, "failed to edit Page: "+err.Error(), &gotgbot.SendMessageOpts{ParseMode: "<code>"})
		if err != nil {
			return fmt.Errorf("failed to send editPage error message : %w", err)
		}
		return nil
	}
	abouteditPage := fmt.Sprintf("<b>Edited Page</b> \n\n<b>URL</b> : %v \n\n<b>Path</b> : %v \n\n<b>Title</b> : %v", pg.URL, pg.Path, pg.Title)
	_, err = ctx.EffectiveMessage.Reply(b, abouteditPage, &gotgbot.SendMessageOpts{ParseMode: "html"})

	if err != nil {
		return fmt.Errorf("failed to send editPage message: %w", err)
	}
	return nil
}
