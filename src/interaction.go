package main

import (
	"fmt"
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

	for adChannelID, _ := range adChannels {
		/*if err := handleAd(client, adChannelID, chat.ID); err != nil {
			return fmt.Errorf("handle ad: %w", err)
		}*/
	}

	return nil
}
