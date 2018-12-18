package util

import (
	"plugin"
	"log"
	"github.com/nlopes/slack"
)


type SlackSendType struct {
	//	allmsg		map[string]func([]string)string
	//	mention		map[string]func([]string)string
	// TODO: オブジェクト化する
	// TODO: mentionと通常メッセージ分けるかも
	Promiscuous  map[string]plugin.Symbol
	Mention     map[string]plugin.Symbol
}

func NewSlackSendType()(*SlackSendType){
	s := SlackSendType{}

	pm := NewPluginManager()
	pm.LoadPlugins("./plugin.json")

	s.Promiscuous = pm.Promiscuous
	s.Mention = pm.Mention

	return &s
}

func (s *SlackSendType) SendSlackMessage(channelID string, msgs []string, rtm *slack.RTM)(error){

	log.Printf("*mention*")
	for key, value := range s.Mention {
		if msgs[1] == key {
			// TODO: 引数をjoinしておく？
			// TODO: 引数をポインタにする
			rtm.SendMessage(rtm.NewOutgoingMessage(value.(func([]string) string)(msgs[2:]), channelID))
			return nil
		}
	}
	return nil
}
