package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Sagleft/uchatbot-engine"
	utopiago "github.com/Sagleft/utopialib-go/v2"
	"github.com/Sagleft/utopialib-go/v2/pkg/structs"
	"github.com/fatih/color"
)

const APIToken = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
const chatIDsSeparator = ","

func main() {
	chatIDsRaw := os.Getenv("CHAT_IDS")
	if chatIDsRaw == "" {
		log.Fatalln(
			"chat ids is not set. " +
				"it is necessary to specify the channel ID " +
				"in the corresponding parameter, separated by commas",
		)
	}

	apiHost := os.Getenv("UTOPIA_HOST")
	if apiHost == "" {
		log.Fatalln("host is not set")
	}

	apiPortRaw := os.Getenv("UTOPIA_PORT")
	if apiPortRaw == "" {
		log.Fatalln("port is not set")
	}
	apiPort, err := strconv.ParseInt(apiPortRaw, 10, 64)
	if err != nil {
		log.Fatalf("parse port: %s\n", err.Error())
	}

	apiWsPortRaw := os.Getenv("UTOPIA_WS_PORT")
	if apiWsPortRaw == "" {
		log.Fatalln("port is not set")
	}
	apiWsPort, err := strconv.ParseInt(apiWsPortRaw, 10, 64)
	if err != nil {
		log.Fatalf("parse port: %s\n", err.Error())
	}

	chats := []uchatbot.Chat{}
	chatIDs := strings.Split(chatIDsRaw, chatIDsSeparator)
	for i, chatID := range chatIDs {
		chats[i] = uchatbot.Chat{ID: chatID}
	}

	_, err = uchatbot.NewChatBot(uchatbot.ChatBotData{
		Config: utopiago.Config{
			Host:   apiHost,
			Token:  APIToken,
			Port:   int(apiPort),
			WsPort: int(apiWsPort),
		},
		Chats: chats,
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

func OnContactMessage(m structs.InstantMessage) {
	fmt.Printf("[CONTACT] %s: %s\n", m.Nick, m.Text)
}

func OnChannelMessage(m structs.WsChannelMessage) {
	fmt.Printf("[CHANNEL] %s: %s\n", m.Nick, m.Text)
}

func OnPrivateChannelMessage(m structs.WsChannelMessage) {
	fmt.Printf("[PRIVATE] %s: %s\n", m.Nick, m.Text)
}

func OnWelcomeMessage(userPubkey string) string {
	return fmt.Sprintf("Hello! Your pubkey is %s", userPubkey)
}

func onError(err error) {
	color.Red(err.Error())
}
