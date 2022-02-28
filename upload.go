package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

func upload(b *gotgbot.Bot, ctx *ext.Context) error {
	str, s := uploadToTelegraph(b, ctx.EffectiveMessage)
	if str == "" {
		return nil
	}
	b.SendChatAction(ctx.Message.From.Id, "typing")
	if s {
		_, err := ctx.EffectiveMessage.Reply(b, str, &gotgbot.SendMessageOpts{ParseMode: "html"})
		if err != nil {
			return fmt.Errorf("failed to send upload str message: %w", err)
		}
		return nil
	}
	userID := ctx.EffectiveMessage.From.Id
	currentStates := states[userID]
	if currentStates == nil {
		states[userID] = []string{str}
	} else {
		states[userID] = append(currentStates, str)
	}
	_, err := ctx.EffectiveMessage.Reply(
		b,
		fmt.Sprintf("File is ready to Upload. Where do you want me to upload? \n\nIf you don't understand this, just tap on the first option. \n\nTotal Files : %v", len(states[userID])),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: uploadButtons,
			},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to send upload str message: %w", err)
	}
	return nil
}

func downloadFile(fileId string, b *gotgbot.Bot) ([]byte, error) {
	bs := make([]byte, 0)
	file, err := b.GetFile(fileId)
	if err != nil {
		return bs, fmt.Errorf("failed to getFile using FileID: %w", err)
	}
	r, err := http.Get("https://api.telegram.org/file/bot" + b.Token + "/" + file.FilePath)
	if err != nil {
		return bs, fmt.Errorf("failed to get file from servers: %w", err)
	}
	bs, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return bs, fmt.Errorf("failed to read response while getting file bytes: %w", err)
	}
	defer r.Body.Close()
	return bs, nil
}

// Returns link to telegraph uploaded media or error as string.
// String because, users should know the error. Weird logic anyway
func uploadToTelegraph(b *gotgbot.Bot, m *gotgbot.Message) (string, bool) {
	var f string
	mt := "photo"
	if m.Animation != nil {
		f = m.Animation.FileId
		mt = "video"
	} else if m.Audio != nil {
		return "", false
	} else if m.Document != nil {
		if m.Document.FileName[len(m.Document.FileName)-4:] == "html" {
			return "", false
		}
		f = m.Document.FileId
	} else if m.Video != nil {
		f = m.Video.FileId
		mt = "video"
	} else if m.Sticker != nil {
		return "", false
	} else if m.Voice != nil {
		return "", false
	} else if m.Photo != nil {
		// Usually 3 file ids are provided. (Maybe always)
		// Quality increases in ascending order.
		// Last one is usually the best one. (Maybe always)
		f = m.Photo[len(m.Photo)-1].FileId
	} else if m.Text != "" {
		return "", false
	} else if m.PinnedMessage != nil {
		return "", false
	} else {
		return "Please send a valid media", true
	}
	bs, err := downloadFile(f, b)
	if err != nil {
		return "An error occurred while downloading. \n\n" + err.Error(), true
	}
	res, err := telegraph.Upload(bytes.NewReader(bs), mt)
	if err != nil {
		return "An error occurred while uploading. \n\n" + err.Error(), true
	} else {
		return "https://telegra.ph/" + res, false
	}
}
