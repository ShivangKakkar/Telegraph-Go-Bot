package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

func getViews(b *gotgbot.Bot, ctx *ext.Context) error {
	args := getArgs(ctx)
	if len(args) < 1 {
		_, err := ctx.EffectiveMessage.Reply(b, "Wrong format. You need to pass page path to get views. \n\nUse <code>/views page_path</code> instead.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send getViews message : %w", err)
		}
		return nil
	}
	pagePath := args[0]
	pg, err := telegraph.GetViews(telegraph.GetViewsOpts{Path: pagePath})
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b, "failed to edit Page: "+err.Error(), &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send getViews error message : %w", err)
		}
		return nil
	}
	_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Views : %v", pg.Views), &gotgbot.SendMessageOpts{ParseMode: "html"})

	if err != nil {
		return fmt.Errorf("failed to send getViews message: %w", err)
	}
	return nil
}
