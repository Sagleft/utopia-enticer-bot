package main

import (
	"fmt"
	"os"
	"strings"

	utopiago "github.com/Sagleft/utopialib-go/v2"
	"github.com/Sagleft/utopialib-go/v2/pkg/structs"
)

const (
	deactivateBotHashTag = "no-bots"
)

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

func sendMessages(client utopiago.Client, chatsCfg chatsConfig) error {
	revCfg := reverseChatCfg(chatsCfg)

	for toChannelID, adChannels := range revCfg {
		if err := handleAdToChannels(client, adChannels, toChannelID); err != nil {
			return fmt.Errorf("handle channel: %w", err)
		}
	}
	return nil
}

func isBotDeactivatedInChannel(channelHashTags string) bool {
	return strings.Contains(channelHashTags, deactivateBotHashTag)
}

func handleAdToChannels(
	client utopiago.Client,
	adChannels fromChannelsCfg,
	toChannelID string,
) error {
	channelData, err := client.GetChannelInfo(toChannelID)
	if err != nil {
		return fmt.Errorf("get channel info: %w", err)
	}

	if isBotDeactivatedInChannel(channelData.HashTags) {
		return nil // skip channel
	}

	// TODO: filter: how many messages have been written in the chat
	// since the last mention of these channels

	adWelcome := os.Getenv("AD_MESSAGE1") // TODO

	messages, err := getAdMessages(client, adChannels, adWelcome)
	if err != nil {
		return fmt.Errorf("get ad messages: %w", err)
	}

	for _, msg := range messages {
		if _, err := client.SendChannelMessage(toChannelID, msg); err != nil {
			return fmt.Errorf("send channel message: %w", err)
		}
	}
	return nil
}

func getAdMessages(
	client utopiago.Client,
	adChannels fromChannelsCfg,
	adWeclome string,
) ([]string, error) {
	if len(adChannels) == 1 {
		return getAdMessagesForSingleChannel(adChannels, adWeclome), nil
	}

	return getAdMessagesForChannels(client, adChannels, adWeclome)
}

func getAdMessagesForSingleChannel(adChannels fromChannelsCfg, adWeclome string) []string {
	var lastChannelID string
	for adChannelID := range adChannels {
		lastChannelID = adChannelID
	}

	return []string{
		adWeclome,
		lastChannelID,
	}
}

func getAdMessagesForChannels(
	client utopiago.Client,
	adChannels fromChannelsCfg,
	adWeclome string,
) ([]string, error) {
	var lines []string
	for adChannelID := range adChannels {
		channelData, err := client.GetChannelInfo(adChannelID)
		if err != nil {
			return nil, fmt.Errorf("get channel info: %w", err)
		}

		lines = append(lines, getChannelLine(channelData, adChannelID))
	}

	return []string{
		fmt.Sprintf(
			"%s\n\n%s",
			adWeclome,
			strings.Join(lines, "\n"),
		),
	}, nil
}

func getChannelLine(channelData structs.ChannelData, ID string) string {
	return fmt.Sprintf(
		"%s %s",
		channelData.Title,
		ID,
	)
}
