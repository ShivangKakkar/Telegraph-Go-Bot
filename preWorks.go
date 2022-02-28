package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

// Change This As Per Convenience
func preWorks(b *gotgbot.Bot, ctx *ext.Context) error {
	id := ctx.EffectiveMessage.From.Id
	AddUser(id)
	if (os.Getenv("MUST_JOIN")) == "" {
		return nil
	}
	c, err := strconv.ParseInt(os.Getenv("MUST_JOIN"), 10, 64)
	if err != nil {
		fmt.Printf(parseIntFail, "MUST_JOIN")
		return nil
	}
	cm, err := b.GetChatMember(c, id)
	if err != nil {
		fmt.Println("Failed to send must join message. Please check if MUST_JOIN is a valid Chat ID")
		return nil
	}
	if cm.GetStatus() == "left" {
		ch, _ := b.GetChat(c)
		_, err = ctx.EffectiveMessage.Reply(
			b,
			fmt.Sprintf(`You must join <a href="%v">%v</a> to use this bot.`, ch.InviteLink, ch.Title),
			&gotgbot.SendMessageOpts{
				ParseMode: "html",
				// ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				// 	InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
				// 		{Text: "Join Channel", Url: ""},
				// 	}},
				// },
				DisableWebPagePreview: true,
			},
		)
		if err != nil {
			fmt.Println("Failed to send must join message")
		}
		return ext.EndGroups
	}
	return nil
}

func checkForDefaultAccount(b *gotgbot.Bot, ctx *ext.Context) error {
	user := ctx.EffectiveMessage.From
	uDB := GetUser(user.Id)
	if uDB.Account == "" {
		shortName := user.FirstName
		authorURL := ""
		authorName := shortName
		if user.Username != "" {
			shortName = user.Username
			authorURL = "https://t.me/" + user.Username
		}
		acc, _ := telegraph.CreateAccount(telegraph.CreateAccountOpts{ShortName: shortName, AuthorURL: authorURL, AuthorName: authorName})
		SetDefaultAccount(user.Id, acc.AccessToken)
		_, err := ctx.EffectiveMessage.Reply(b, defaultAccountMessage, &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send defaultAccount Message: %v", err)
		}
		return nil
	}
	return nil
}

// Important for new users
var defaultAccountMessage = `
You seem to be a brand new user. I've created a new account for you which is now the default one. You can see it's info in /accounts message. 

Whenever, the bot asks you for an account number to act on, use 0 to specify the default account. The number is same as seen in <code>/accounts</code> message.

But of course, you can create more accounts using the /create command. To use the newly created account while doing things, pass it's number as seen in <code>/accounts</code> message. 

So want to create a page? Use <code>/page 0 page_title &lt;in reply to html message&gt;</code>
`
