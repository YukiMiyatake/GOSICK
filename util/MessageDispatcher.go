package util

import (
	"log"
)

type MessageDispatcher struct {
	BotName  []string
	BotID      string
}

func NewMessageDispatcher()(*MessageDispatcher){
	s := MessageDispatcher{}

	sc := SlackConfig{}
	err := sc.LoadSlackConfig("./slack.json")
	if(err != nil){
		log.Printf("[Error] %s", err)
		return &s
	}
	s.BotName = sc.BotName
	s.BotID = sc.BotID

	return &s
}

func (s *MessageDispatcher) ContainBot(botID string)(string){

	if(botID == "regina" || Contains( s.BotName, botID) || Contains( s.BotID, botID)){
		return "mention"
	}
	return "err"
}

