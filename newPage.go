package main

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

func newPage(b *gotgbot.Bot, ctx *ext.Context) error {
	content, err := extractHTMLFromReply(b, ctx)
	if err != nil {
		return err
	}
	if content == "" {
		return nil
	}
	args := getArgs(ctx)
	if len(args) < 2 {
		_, err := ctx.EffectiveMessage.Reply(b, "Wrong format. You need to pass account number and page title to create a new page. \n\nUse <code>/page account_number page_title &lt;in reply to text message&gt;</code> instead.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send newPage message : %w", err)
		}
		return nil
	}
	accountNumber := args[0]
	pageTitle := strings.Join(args[1:], " ")
	accountNumber, con, err := getToken(accountNumber, b, ctx)
	if err != nil {
		return fmt.Errorf("failed to send listPages tokens message : %w", err)
	}
	if con {
		return nil
	}
	opts := telegraph.CreatePageOpts{
		AccessToken: accountNumber,
		Title:       pageTitle,
		Content:     telegraph.HTMLToNode(content),
	}
	// pg, err := telegraph.CreatePage(opts)
	pg, err := telegraph.Post("createPage", opts)
	// pg, err := telegraph.Get("createPage", opts)
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, "An error occurred while creating page: "+err.Error(), &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send newPage error message : %w", err)
		}
		return nil
	}
	aboutNewPage := fmt.Sprintf("<b>Created New Page</b> \n\n<b>URL</b> : %v \n\n<b>Path</b> : <code>%v</code>  \n\n<b>Title</b> : %v", pg.URL, pg.Path, pg.Title)
	_, err = ctx.EffectiveMessage.Reply(b, aboutNewPage, &gotgbot.SendMessageOpts{ParseMode: "html"})

	if err != nil {
		return fmt.Errorf("failed to send newPage message: %w", err)
	}
	return nil
}
