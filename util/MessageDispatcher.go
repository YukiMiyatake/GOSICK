package util

import (
	// "log"
)

type MessageDispatcher struct {
	BotName  []string
	BotID      string
}

type MessageType string
const (
	Mention MessageType = "Mention"
	Promiscous MessageType = "Promiscous"
	Error MessageType = "Error"
)

func NewMessageDispatcher(botName []string, botID string)(*MessageDispatcher){
	s := MessageDispatcher{}
	s.BotName = botName
	s.BotID = botID
	return &s
}

func (s *MessageDispatcher) GetMessageType(botID string)(MessageType){
	if(Contains( s.BotName, botID) || Contains( s.BotID, botID)){
		return Mention
	}
	return Error
}
