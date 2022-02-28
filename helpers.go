// Helpers Functions
// Upload and Download helpers are in 'upload.go'
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

var parseIntFail = "Could not parse %v var. Please check if it's a valid Integer and a valid Chat ID \n"

func getString(ctx *ext.Context) string {
	args := strings.SplitN(strings.TrimSpace(ctx.EffectiveMessage.Text), " ", 2)
	return args[1]
}

func getArgsCount(ctx *ext.Context) int {
	return len(getArgs(ctx))
}

func getArgs(ctx *ext.Context) []string {
	args := strings.Split(strings.TrimSpace(ctx.EffectiveMessage.Text), " ")
	_, args = args[0], args[1:]
	return args
}

func getToken(accountNumber string, b *gotgbot.Bot, ctx *ext.Context) (string, bool, error) {
	if accountNumberIndex, err := strconv.Atoi(accountNumber); err == nil {
		if accountNumberIndex == 0 {
			uDB := GetUser(ctx.EffectiveMessage.From.Id)
			accountNumber = uDB.Account
			return accountNumber, false, nil
		}
		// Get From Tokens
		tokens := GetAllAccounts(ctx.EffectiveMessage.From.Id)
		if len(tokens) < accountNumberIndex {
			_, err = ctx.EffectiveMessage.Reply(b, "You don't have that much accounts anyway. Give a valid account number.", &gotgbot.SendMessageOpts{ParseMode: "html"})
			return "", true, err
		}
		// Now of no use as implemented "Default Account"
		// if accountNumberIndex == 0 {
		// 	_, err = ctx.EffectiveMessage.Reply(b, "Oh really? You must be a programmer", &gotgbot.SendMessageOpts{ParseMode: "html"})
		// 	return "", true, err
		// }
		accountNumber = tokens[accountNumberIndex]
	}
	return accountNumber, false, nil
}

// Edit Account using different values
func editAccount(change string, b *gotgbot.Bot, ctx *ext.Context) error {
	status, err := ctx.EffectiveMessage.Reply(b, "Wait...", &gotgbot.SendMessageOpts{})
	if err != nil {
		return fmt.Errorf("failed to send editAccount message: %w", err)
	}
	args := getArgs(ctx)
	if len(args) < 2 {
		_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Wrong format. You need to pass account number and %v to edit an account.", change), &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send editAccount %v message : %w", change, err)
		}
		return nil
	}
	accountNumber := args[0]
	value := strings.Join(args[1:], " ")
	accountNumber, con, err := getToken(accountNumber, b, ctx)
	if err != nil {
		return fmt.Errorf("failed to send editAccount %v tokens message : %w", change, err)
	}
	if con {
		return nil
	}
	infoStr := fmt.Sprintf("<b>Edited Account [%v]</b>", change)
	var opts telegraph.EditAccountInfoOpts
	if change == "AuthorURL" {
		opts = telegraph.EditAccountInfoOpts{AccessToken: accountNumber, AuthorURL: value}
	} else if change == "AuthorName" {
		opts = telegraph.EditAccountInfoOpts{AccessToken: accountNumber, AuthorName: value}
	} else {
		opts = telegraph.EditAccountInfoOpts{AccessToken: accountNumber, ShortName: value}
	}
	acc, err := telegraph.EditAccountInfo(opts)
	if err != nil {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("An Error Occured \n\n%v", err), &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send editAccount %v message : %w", change, err)
		}
		return nil
	}
	infoStr += fmt.Sprintf("\n\n<b>Short Name</b>: %v \n\n<b>Access Token</b>: <code>%v</code>", acc.ShortName, accountNumber)
	if acc.AuthorName != "" {
		infoStr += fmt.Sprintf("\n\n<b>Author Name</b>: %v", acc.AuthorName)
	}
	if acc.AuthURL != "" {
		infoStr += fmt.Sprintf("\n\n<b>Auth URL</b>: %v", acc.AuthURL)
	}
	if acc.AuthorURL != "" {
		infoStr += fmt.Sprintf("\n\n<b>Author URL</b>: %v", acc.AuthorURL)
	}
	if acc.PageCount != 0 {
		infoStr += fmt.Sprintf("\n\n<b>Page Count</b>: %v", acc.PageCount)
	}
	_, _, err = status.EditText(b, infoStr, &gotgbot.EditMessageTextOpts{ChatId: status.Chat.Id, MessageId: status.MessageId, ParseMode: "html", DisableWebPagePreview: true})
	if err != nil {
		return fmt.Errorf("failed to send editAccount %v message : %w", change, err)
	}
	return nil
}

func checkIfLengthy(s string, b *gotgbot.Bot, ctx *ext.Context) bool {
	if len(s) > 4096 {
		r := []string{"<b>", "</b>", "<code>", "</code>", "<u>", "</u>", "<i>", "</i>"}
		for _, j := range r {
			s = strings.ReplaceAll(s, j, "")
		}
		file := fmt.Sprintf("%v_data.txt", ctx.EffectiveMessage.From.Id)
		f, _ := os.Create(file)
		f.WriteString(s)
		defer f.Close()
		f, _ = os.Open(file)
		_, err := b.SendDocument(ctx.Message.From.Id, gotgbot.NamedFile{File: f, FileName: "data.txt"}, &gotgbot.SendDocumentOpts{DisableNotification: true, ReplyToMessageId: ctx.Message.MessageId})
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}
		os.Remove(fmt.Sprintf("%v_data.txt", ctx.EffectiveMessage.From.Id))
		return true
	}
	return false
}

func toImgTag(a []Img) string {
	str := ""
	for _, l := range a { // SayNoToIJ
		str += fmt.Sprintf(`<img src="%v"><figcaption>%v</figcaption>`, l.Link, l.Caption)
	}
	return str
}

func extractHTMLFromReply(b *gotgbot.Bot, ctx *ext.Context) (string, error) {
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		_, err := ctx.EffectiveMessage.Reply(b, "Wrong format. You need to reply to some html content text message or document.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return "", fmt.Errorf("failed to send 'no html content' message : %w", err)
		}
		return "", nil
	}
	var content string
	if ctx.EffectiveMessage.ReplyToMessage.Document != nil {
		doc := ctx.EffectiveMessage.ReplyToMessage.Document
		if doc.FileName[len(doc.FileName)-4:] != "html" {
			_, err := ctx.EffectiveMessage.Reply(b, "File must be an HTML File. It must end with .html", &gotgbot.SendMessageOpts{ParseMode: "html"})
			if err != nil {
				return "", fmt.Errorf("failed to send 'no .html' message : %w", err)
			}
			return "", nil
		}
		bs, err := downloadFile(ctx.EffectiveMessage.ReplyToMessage.Document.FileId, b)
		if err != nil {
			return "", fmt.Errorf("failed to download html document : %w", err)
		}
		content = string(bs)
	} else {
		content = ctx.EffectiveMessage.ReplyToMessage.Text
	}
	return content, nil
}
