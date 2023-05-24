package main

import (
	"log"

	"github.com/Sagleft/uchatbot-engine"
)

const APIToken = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
const chatIDsSeparator = ","

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatalln(err)
	}

	chats, err := getChats()
	if err != nil {
		log.Fatalln(err)
	}

	_, err = uchatbot.NewChatBot(uchatbot.ChatBotData{
		Config: cfg,
		Chats:  chats,
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
}
