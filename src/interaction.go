package main

import (
	"fmt"
	"strings"

	"github.com/Sagleft/uchatbot-engine"
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

func sendMessages(client utopiago.Client, chats []uchatbot.Chat) error {
	for _, chat := range chats {
		if err := handleChannel(client, chat.ID); err != nil {
			return fmt.Errorf("handle channel: %w", err)
		}
	}
	return nil
}

func isBotDeactivatedInChannel(channelHashTags string) bool {
	return strings.Contains(channelHashTags, deactivateBotHashTag)
}

func handleChannel(client utopiago.Client, channelID string) error {
	channelData, err := client.GetChannelInfo(channelID)
	if err != nil {
		return fmt.Errorf("get channel info: %w", err)
	}

	if isBotDeactivatedInChannel(channelData.HashTags) {
		return nil // skip channel
	}

	return nil
}
