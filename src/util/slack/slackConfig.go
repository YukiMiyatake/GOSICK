package slack

import "util"

type ISlackConfig interface {
	GetPort() *string
	GetBotToken() *string
	GetVerificationToken() *string
	GetBotID() *string
	GetBotName() *[]string
	GetChannelID() *string
}

// SlackConfig ..
type SlackConfig struct {
	port              string   `json:"PORT"`
	botToken          string   `json:"BOT_TOKEN"`
	verificationToken string   `json:"VERIFICATION_TOKEN"`
	botID             string   `json:"BOT_ID"`
	botName           []string `json:"BOT_NAME"`
	channelID         string   `json:"CHANNEL_ID"`
}

// NewSlackConfig ..
func NewSlackConfig(file string) (ISlackConfig, error) {
	s := &SlackConfig{}

	return s, util.JsonFileLoader(file, s)
}

func (s *SlackConfig) GetPort() *string {
	return &s.port
}

func (s *SlackConfig) GetBotToken() *string {
	return &s.botToken
}
func (s *SlackConfig) GetVerificationToken() *string {
	return &s.verificationToken
}
func (s *SlackConfig) GetBotID() *string {
	return &s.botID
}
func (s *SlackConfig) GetBotName() *[]string {
	return &s.botName
}
func (s *SlackConfig) GetChannelID() *string {
	return &s.channelID
}
