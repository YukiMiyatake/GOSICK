package slack

import "util"

// SlackConfig ..
type SlackConfig struct {
	Port              string   `json:"PORT"`
	BotToken          string   `json:"BOT_TOKEN"`
	VerificationToken string   `json:"VERIFICATION_TOKEN"`
	BotID             string   `json:"BOT_ID"`
	BotName           []string `json:"BOT_NAME"`
	ChannelID         string   `json:"CHANNEL_ID"`
}

// NewSlackConfig ..
func NewSlackConfig(file string) (*SlackConfig, error) {
	s := &SlackConfig{}

	return s, util.JsonFileLoader(file, s)
}
