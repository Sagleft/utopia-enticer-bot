package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Sagleft/uchatbot-engine"
	utopiago "github.com/Sagleft/utopialib-go/v2"
	"github.com/fatih/color"
)

func getConfig() (utopiago.Config, error) {
	cfg := utopiago.Config{}

	cfg.Host = os.Getenv("UTOPIA_HOST")
	if cfg.Host == "" {
		log.Fatalln("host is not set")
	}

	apiPortRaw := os.Getenv("UTOPIA_PORT")
	if apiPortRaw == "" {
		log.Fatalln("port is not set")
	}
	apiPort, err := strconv.ParseInt(apiPortRaw, 10, 64)
	if err != nil {
		return cfg, fmt.Errorf("parse port: %w", err)
	}
	cfg.Port = int(apiPort)

	apiWsPortRaw := os.Getenv("UTOPIA_WS_PORT")
	if apiWsPortRaw == "" {
		log.Fatalln("port is not set")
	}
	apiWsPort, err := strconv.ParseInt(apiWsPortRaw, 10, 64)
	if err != nil {
		log.Fatalf("parse port: %s\n", err.Error())
	}
	cfg.WsPort = int(apiWsPort)

	return cfg, nil
}

func getChats() ([]uchatbot.Chat, error) {
	chatIDsRaw := os.Getenv("CHAT_IDS")
	if chatIDsRaw == "" {
		return nil, fmt.Errorf(
			"chat ids is not set. " +
				"it is necessary to specify the channel ID " +
				"in the corresponding parameter, separated by commas",
		)
	}

	chats := []uchatbot.Chat{}
	chatIDs := strings.Split(chatIDsRaw, chatIDsSeparator)
	for i, chatID := range chatIDs {
		chats[i] = uchatbot.Chat{ID: chatID}
	}

	return chats, nil
}

func onError(err error) {
	color.Red(err.Error())
}
