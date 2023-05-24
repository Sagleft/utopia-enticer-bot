package main

import (
	"fmt"

	"github.com/Sagleft/utopialib-go/v2/pkg/structs"
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
