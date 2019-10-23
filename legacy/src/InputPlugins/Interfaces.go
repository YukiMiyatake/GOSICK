package InputPlugins

type ISlackConfig interface {
	GetPort() *string
	GetBotToken() *string
	GetVerificationToken() *string
	GetBotID() *string
	GetBotName() *[]string
	GetChannelID() *string
}

type IChatListener interface {
	Listen() error
}
