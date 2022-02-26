package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// ------------------- Basic Commands ------------------ //

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf(startMessage, ctx.EffectiveMessage.From.FirstName, b.User.FirstName), &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: mainButtons,
		},
		DisableWebPagePreview: true,
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

func help(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, helpMessage, &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: homeButtons,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to send help message: %w", err)
	}
	return nil
}

func about(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, aboutMessage, &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: homeButtons,
		},
		DisableWebPagePreview: true,
	})
	if err != nil {
		return fmt.Errorf("failed to send about message: %w", err)
	}
	return nil
}

func guide(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, guideLink, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})
	if err != nil {
		return fmt.Errorf("failed to send about message: %w", err)
	}
	return nil
}

func html(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf(htmlTutMessage, htmlTut), &gotgbot.SendMessageOpts{
		ParseMode:             "html",
		DisableWebPagePreview: true,
	})
	if err != nil {
		return fmt.Errorf("failed to send about message: %w", err)
	}
	return nil
}

// ---------------------------------------------- //

var guideLink = "https://telegra.ph/Telegraph-Bot-Usage-Guide-Stark-Bots-02-26"

var htmlTut = "https://geekyweb.tk/docs/html"

var startMessage = `
Hey %v 

Welcome to %s 

I can do anything related to Telegraph like create accounts, create pages, upload media and so on. Please see a full list in help message
`

var helpMessage = fmt.Sprintf(`
Please read <a href="%v">this guide</a> first. 

Still here? It's not optional. Go read it.

<b>Available Commands</b>

/guide - Usage Guide
/html - Available HTML Tags and Attributes
/create - Create a new account
/new - Alias for 'create'
/accounts - List all accounts
/editshortname - Edit account's short name
/editauthorname - Edit account's author name
/editauthorurl - Edit account's author url
/get - Get an account
/add - Add an existing account
/remove - Remove an account
/editpage - Edit a Page
/page - Create a new page
/newpage - Alias for 'page'
/pages - List pages for a particular account
/views - Get views of a particular page
/start - Check if bot is running
/help - Help Message
/about - About this bot
`, guideLink)

var aboutMessage = fmt.Sprintf(`
A telegraph bot by @StarkBots

<b>Language</b> - <a href="https://go.dev">Golang</a>

<b>Telegraph Library</b> - https://github.com/StarkBotsIndustries/telegraph

<b>Source Code</b> - <a href="https://github.com/StarkBotsIndustries/Telegraph-Go-Bot">GitHub Repository</a>

<b>Usage Guide</b> - %v

<b>Telegram Library</b> - <a href="https://github.com/PaulSonOfLars/gotgbot">gotgbot</a>

Developed with ❤️ by @StarkProgrammer
`, guideLink)

var htmlTutMessage = `
<b>HTML Content</b>

You can use HTML to create your Telegraph Pages. 
If you don't know HTML at all, please checkout this <a href="%v">HTML Tutorial</a> for a basic understanding.

<b>Available Tags</b>: a, aside, b, blockquote, br, code, em, figcaption, figure, h3, h4, hr, i, iframe, img, li, ol, p, pre, s, strong, u, ul, video

<b>Available Attributes</b>: href, src

<b>Providing HTML</b>
I will ask you for some HTML Content when you want to create or edit a page. You can provide it in two ways.
   1. Reply to a text message with HTML 
   2. Reply to an HTML document
`
var homeButtons = [][]gotgbot.InlineKeyboardButton{{
	{Text: "🏠 Return Home 🏠", CallbackData: "home"},
}}

var mainButtons = [][]gotgbot.InlineKeyboardButton{
	{
		{Text: "✨ Bot Status and More Bots ✨", Url: "https://t.me/StarkBots/7"},
	},
	{
		{Text: "How to Use ❔", CallbackData: "help"},
		{Text: "🎪 About 🎪", CallbackData: "about"},
	},
	{
		{Text: "♥ More Amazing bots ♥", Url: "https://t.me/StarkBots"},
	},
}

// ------------------------- Callbacks ------------------------ //

func homeCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	_, _, err := cb.Message.EditText(
		b,
		fmt.Sprintf(startMessage, cb.From.FirstName, b.User.FirstName),
		&gotgbot.EditMessageTextOpts{
			ParseMode: "html",
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: mainButtons,
			},
			DisableWebPagePreview: true,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to edit homecb text: %w", err)
	}
	return nil
}

func helpCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	_, _, err := cb.Message.EditText(
		b,
		helpMessage,
		&gotgbot.EditMessageTextOpts{
			ParseMode: "html",
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: homeButtons,
			},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to edit helpcb text: %w", err)
	}
	return nil
}

func aboutCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	_, _, err := cb.Message.EditText(
		b,
		aboutMessage,
		&gotgbot.EditMessageTextOpts{
			ParseMode: "html",
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: homeButtons,
			},
			DisableWebPagePreview: true,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to edit aboutcb text: %w", err)
	}
	return nil
}
