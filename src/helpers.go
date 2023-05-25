package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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

func getCronDuration() (time.Duration, error) {
	timeoutHoursRaw := os.Getenv("ACTION_TIMEOUT_HOURS")
	if timeoutHoursRaw == "" {
		return 0, errors.New("action timeout is not set")
	}

	timeoutHours, err := strconv.ParseInt(timeoutHoursRaw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse action timeout: %w", err)
	}

	return time.Hour * time.Duration(timeoutHours), nil
}

func getDebugImedStart() (bool, error) {
	paramRaw := os.Getenv("DEBUG_IMMEDIATELY_START")
	if paramRaw == "" {
		return false, nil
	}

	paramVal, err := strconv.ParseBool(paramRaw)
	if err != nil {
		return false, fmt.Errorf("parse debug param: %w", err)
	}

	return paramVal, nil
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
