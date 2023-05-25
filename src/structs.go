package main

import "github.com/Sagleft/uchatbot-engine"

// from channel ID -> chats
type chatsConfig map[string][]uchatbot.Chat

// chat ID -> from channel ID -> empty struct
type chatsConfigReversed map[string]fromChannelsCfg

type fromChannelsCfg map[string]struct{}
