package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testFromID = "E36921F241EA9A3ABE03308BF41E0C37"
const testSecondID = "EF978A706EEE71AB24A29D069CCCAA7D"
const dest1 = "7059882759A545496050909C910D8AA0"
const dest2 = "EB3D3060924FE529271B2511DF51071F"
const testChatIDsRaw = testFromID + ":" + dest1 + "," + dest2 + ";" + testSecondID + ":" + dest1

func TestGetChats(t *testing.T) {
	// when
	chatsCfg, err := getChats(testChatIDsRaw)

	// then
	require.NoError(t, err)
	assert.Equal(t, 2, len(chatsCfg))
	assert.Equal(t, 2, len(chatsCfg[testFromID]))
	assert.Equal(t, 1, len(chatsCfg[testSecondID]))
}

func TestReverseChatCfg(t *testing.T) {
	// given
	chatsCfg, _ := getChats(testChatIDsRaw)

	// when
	rev := reverseChatCfg(chatsCfg)

	// then
	assert.Equal(t, 2, len(rev))
	assert.Equal(t, 2, len(rev[dest1]))
	assert.Equal(t, 1, len(rev[dest2]))
}
