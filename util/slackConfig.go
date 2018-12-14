package util

import (
	"io/ioutil"
	"log"
	"encoding/json"
)

type SlackConfig struct {
	Port              string `json:"PORT"`
	BotToken          string `json:"BOT_TOKEN"`
	VerificationToken string `json:"VERIFICATION_TOKEN"`
	BotID             string `json:"BOT_ID"`
	BotName         []string `json:"BOT_NAME"`
	ChannelID         string `json:"CHANNEL_ID"`
}

func Hoge(){

}
func (s *SlackConfig) LoadSlackConfig(file string)(error){
//	var sc slackConfig

	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("[Error] %s", err)
		return err
	}
	err = json.Unmarshal(jsonData, s)
	if err != nil {
		log.Printf("[Error] %s", err)
		return err
	}

	return nil
}
