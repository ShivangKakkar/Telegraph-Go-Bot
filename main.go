package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/joho/godotenv"
)

func main() {
	// Needed for Local Deploys
	godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }

	b, _ := gotgbot.NewBot(os.Getenv("TOKEN"), &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		ErrorLog: nil,
		DispatcherOpts: ext.DispatcherOpts{
			Error: func(b *gotgbot.Bot, _ *ext.Context, err error) ext.DispatcherAction {
				logTg(b, "An error occurred while handling update:"+err.Error())
				return ext.DispatcherActionNoop
			},
			Panic:       nil,
			ErrorLog:    nil,
			MaxRoutines: 0,
		},
	})
	// Dispatcher
	dispatcher := updater.Dispatcher
	// Basic Handlers
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("help", help))
	dispatcher.AddHandler(handlers.NewCommand("about", about))
	dispatcher.AddHandler(handlers.NewCommand("guide", guide))
	dispatcher.AddHandler(handlers.NewCommand("html", html))
	dispatcher.AddHandler(handlers.NewCommand("stats", stats)) // Owner Only
	dispatcher.AddHandler(handlers.NewCommand("users", users)) // Owner Only
	// Accounts Handlers
	dispatcher.AddHandler(handlers.NewCommand("create", newAccount))
	dispatcher.AddHandler(handlers.NewCommand("new", newAccount))
	dispatcher.AddHandler(handlers.NewCommand("accounts", account))
	dispatcher.AddHandler(handlers.NewCommand("editshortname", editShortName))
	dispatcher.AddHandler(handlers.NewCommand("editauthorname", editAuthorName))
	dispatcher.AddHandler(handlers.NewCommand("editauthorurl", editAuthorURL))
	dispatcher.AddHandler(handlers.NewCommand("remove", removeAccount))
	dispatcher.AddHandler(handlers.NewCommand("get", getAccount))
	dispatcher.AddHandler(handlers.NewCommand("add", addAccount))
	// Pages Handlers
	dispatcher.AddHandler(handlers.NewCommand("page", newPage))
	dispatcher.AddHandler(handlers.NewCommand("newpage", newPage))
	dispatcher.AddHandler(handlers.NewCommand("editpage", editPage))
	dispatcher.AddHandler(handlers.NewCommand("pages", listPages))
	dispatcher.AddHandler(handlers.NewCommand("views", getViews))
	// Callbacks Handlers
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("home"), homeCB))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("help"), helpCB))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("about"), aboutCB))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("direct"), directUploadCB))
	dispatcher.AddHandler(handlers.NewCallback(callbackquery.Equal("page"), pageUploadCB))
	// Message Handlers
	dispatcher.AddHandlerToGroup(handlers.NewMessage(message.Private, preWorks), -1)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(message.Private, upload), 1)
	dispatcher.AddHandlerToGroup(handlers.NewMessage(message.Private, checkForDefaultAccount), 2)
	// Polling
	err := updater.StartPolling(b, &ext.PollingOpts{
		Timeout:            11 * time.Second,
		GetUpdatesOpts:     gotgbot.GetUpdatesOpts{Timeout: 10},
		DropPendingUpdates: true,
	})
	if err != nil {
		panic("Failed to start polling: " + err.Error())
	}
	fmt.Printf("%s is now running...\n", b.User.Username)
	updater.Idle()
}
