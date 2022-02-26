package main

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

// UserID : Array of Links
var states = map[int64][]string{}

var uploadButtons = [][]gotgbot.InlineKeyboardButton{
	{{Text: "Upload as Separate Image Files", CallbackData: "direct"}},
	{{Text: "Upload as a Page", CallbackData: "page"}},
}

func directUploadCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	links := states[cb.From.Id]
	if links == nil {
		_, _, err := cb.Message.EditText(b, "You currently have no media in state. Please send again.", nil)
		if err != nil {
			return fmt.Errorf("failed to edit 'no media in state' directUploadCB: %w", err)
		}
		return nil
	}
	_, err := b.SendMessage(cb.From.Id, strings.Join(links, "\n\n"), nil)
	// Change the temp state
	states[cb.From.Id] = []string{}
	if err != nil {
		return fmt.Errorf("failed to send link in directUploadCB: %w", err)
	}
	cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{})
	return nil
}

func pageUploadCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	links := states[cb.From.Id]
	if links == nil {
		_, _, err := cb.Message.EditText(b, "You currently have no media in state. Please send again.", nil)
		if err != nil {
			return fmt.Errorf("failed to edit 'no media in state' directUploadCB: %w", err)
		}
		return nil
	}
	str := toImgTag(links)
	n := cb.From.FirstName
	if cb.From.Username != "" {
		n = cb.From.Username
	}
	page, err := telegraph.CreatePage(telegraph.CreatePageOpts{
		AccessToken: GetUser(cb.From.Id).Account,
		Title:       cb.From.FirstName,
		AuthorName:  n,
		HTMLContent: str,
	})
	if err != nil {
		_, _, err = cb.Message.EditText(b, "An Error Occured"+err.Error(), nil)
		if err != nil {
			return fmt.Errorf("failed to send pageUploadCB: %w", err)
		}
		return nil
	}
	_, err = b.SendMessage(cb.Message.Chat.Id, page.URL, &gotgbot.SendMessageOpts{ReplyToMessageId: 0})
	if err != nil {
		return fmt.Errorf("failed to send pageUploadCB: %w", err)
	}
	// Change the temp state
	states[cb.From.Id] = []string{}
	_, err = cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: "Done!"})
	if err != nil {
		return fmt.Errorf("failed to answer callback query: %w", err)
	}
	return nil
}

// Help, About Message
// Guide how to use

// Default Account
// New Account on message
