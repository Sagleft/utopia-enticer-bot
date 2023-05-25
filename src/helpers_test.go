package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetChats(t *testing.T) {
	// given
	fromID := "E36921F241EA9A3ABE03308BF41E0C37"
	secondID := "EF978A706EEE71AB24A29D069CCCAA7D"
	chatIDsRaw := "E36921F241EA9A3ABE03308BF41E0C37:7059882759A545496050909C910D8AA0,EB3D3060924FE529271B2511DF51071F;EF978A706EEE71AB24A29D069CCCAA7D:7059882759A545496050909C910D8AA0"

	// when
	chatsCfg, err := getChats(chatIDsRaw)

	// then
	require.NoError(t, err)
	assert.Equal(t, 2, len(chatsCfg))
	assert.Equal(t, 2, len(chatsCfg[fromID]))
	assert.Equal(t, 1, len(chatsCfg[secondID]))
}
