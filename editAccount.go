package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func editShortName(b *gotgbot.Bot, ctx *ext.Context) error {
	err := editAccount("ShortName", b, ctx)
	return err
}

func editAuthorURL(b *gotgbot.Bot, ctx *ext.Context) error {
	err := editAccount("AuthorURL", b, ctx)
	return err
}

func editAuthorName(b *gotgbot.Bot, ctx *ext.Context) error {
	err := editAccount("AuthorName", b, ctx)
	return err
}
