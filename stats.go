package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func stats(b *gotgbot.Bot, ctx *ext.Context) error {
	o, _ := strconv.ParseInt(os.Getenv("OWNER_ID"), 10, 64)
	if ctx.EffectiveMessage.From.Id != o {
		return nil
	}
	c := UsersCount()
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Total Users : %v", c), &gotgbot.SendMessageOpts{})
	if err != nil {
		return fmt.Errorf("failed to send removeAccount message: %w", err)
	}
	return nil
}

func users(b *gotgbot.Bot, ctx *ext.Context) error {
	o, _ := strconv.ParseInt(os.Getenv("OWNER_ID"), 10, 64)
	if ctx.EffectiveMessage.From.Id != o {
		return nil
	}
	users := GetAllUsers()
	list := fmt.Sprintf("<b>Users [%v]</b> \n\n", len(users)) // list is equal to string :)
	for i, user := range users {
		id := user.ID
		x, _ := b.GetChatMember(id, id)
		u := x.GetUser()
		list += fmt.Sprintf("%v) <a href='tg://user?id=%v'>%v</a> [<code>%v</code>]\n", i+1, u.Id, u.FirstName, u.Id)
	}
	_, err := ctx.EffectiveMessage.Reply(b, list, &gotgbot.SendMessageOpts{ParseMode: "html"})
	if err != nil {
		return fmt.Errorf("failed to send listUsers message: %w", err)
	}
	return nil
}
