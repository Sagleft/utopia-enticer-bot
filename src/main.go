package main

import (
	"log"
	"os"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/Sagleft/uchatbot-engine"
	"github.com/fatih/color"
	simplecron "github.com/sagleft/simple-cron"
)

const APIToken = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
const chatIDsGroupSeparator = ";"
const groupElementsSeparator = ":"
const chatIDsSeparator = ","

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatalln(err)
	}

	actionTimeout, err := getCronDuration()
	if err != nil {
		log.Fatalln(err)
	}

	actionOnStart, err := getDebugImedStart()
	if err != nil {
		log.Fatalln(err)
	}

	chats, err := getChats(os.Getenv("CHAT_IDS"))
	if err != nil {
		log.Fatalln(err)
	}

	bot, err := uchatbot.NewChatBot(uchatbot.ChatBotData{
		Config: cfg,
		Chats:  getChatsFromCfg(chats),
		Callbacks: uchatbot.ChatBotCallbacks{
			OnContactMessage:        OnContactMessage,
			OnChannelMessage:        OnChannelMessage,
			OnPrivateChannelMessage: OnPrivateChannelMessage,
			WelcomeMessage:          OnWelcomeMessage,
		},
		UseErrorCallback: true,
		ErrorCallback:    onError,
	})
	if err != nil {
		log.Fatalln(err)
	}

	simplecron.NewCronHandler(func() {
		if err := sendMessages(bot.GetClient(), chats); err != nil {
			color.Red("send messages: %s", err.Error())
		}
	}, actionTimeout).Run(actionOnStart)

	swissknife.RunInBackground()
}
