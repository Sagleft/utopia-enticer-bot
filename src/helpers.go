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

func getChats(chatIDsRaw string) (chatsConfig, error) {
	if chatIDsRaw == "" {
		return chatsConfig{}, fmt.Errorf(
			"chat ids is not set. " +
				"it is necessary to specify the channel ID " +
				"in the corresponding parameter, separated by commas",
		)
	}

	r := chatsConfig{}
	var lastError error

	chatGroups := strings.Split(chatIDsRaw, chatIDsGroupSeparator)
	for _, chatGroup := range chatGroups {
		groupElements := strings.Split(chatGroup, groupElementsSeparator)
		if len(groupElements) != 2 {
			lastError = errors.New("invalid chat ids structure. see readme")
		}

		channelFromID := groupElements[0]
		chatIDs := strings.Split(groupElements[1], chatIDsSeparator)

		for _, chatID := range chatIDs {
			r[channelFromID] = append(r[channelFromID], uchatbot.Chat{ID: chatID})
		}
	}

	return r, lastError
}

func getChatsFromCfg(cfg chatsConfig) []uchatbot.Chat {
	c := []uchatbot.Chat{}
	for _, v := range cfg {
		c = append(c, v...)
	}
	return c
}

func onError(err error) {
	color.Red(err.Error())
}

func reverseChatCfg(cfg chatsConfig) chatsConfigReversed {
	r := chatsConfigReversed{}

	for fromChannelID, chats := range cfg {
		for _, channelData := range chats {
			if _, isExists := r[channelData.ID]; !isExists {
				r[channelData.ID] = make(fromChannelsCfg)
			}

			r[channelData.ID][fromChannelID] = struct{}{}
		}
	}

	return r
}
