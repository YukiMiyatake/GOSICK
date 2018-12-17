// +build !cmd_driver

/*
  TODO: プラグイン interface作る
  TODO: ポインタまわり見直す
  TODO: 管理権限必要・・
  TODO: チャンネル名をチャンネルIDに変換
  TODO: チャンネルリスト対応
  TODO: BOT_NAMEから BotIDをひく
  TODO: VerificationToken調査
  TODO: 会話機能
  TODO: イメージリストを選択式に
  TODO: makefile
  TODO: ランチ機能
  TODO: リージョンなどを設定ファイルから読む
  TODO: 異リージョンの切り替え？
  TODO: AWSプラグインを同じフォルダに置きたいところ
  TODO: 設定ファイルの動的読み込み機能
  TODO: 全体ヘルプ
  TODO: リスナー形式に
  TODO: AWSプラグイン シングルトン設計検討
*/
package main

import (
	"log"
	"net/http"
	"os"
	"plugin"
	"./util"

	"github.com/nlopes/slack"
	"io/ioutil"
	"encoding/json"
)

type slackConfig struct {
	Port              string `json:"PORT"`
	BotToken          string `json:"BOT_TOKEN"`
	VerificationToken string `json:"VERIFICATION_TOKEN"`
	BotID             string `json:"BOT_ID"`
	ChannelID         string `json:"CHANNEL_ID"`
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func LoadSlackSettings(file string)(slackConfig, error){
	var sc slackConfig
	
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("[Error] %s", err)
		return sc, err
	}

	err = json.Unmarshal(jsonData, &sc)
	if err != nil {
		log.Printf("[Error] %s", err)
		return sc, err
	}
	
	return sc, nil
}


type PluginData struct {
	Name string `json:"NAME"`
	Path string `json:"PATH"`
}

type PluginManager struct {
	allmsg map[string]plugin.Symbol
	mention map[string]plugin.Symbol
}

func NewPluginManager()(*PluginManager){
	pm := PluginManager{}
	pm.allmsg = map[string]plugin.Symbol{}
	pm.mention = map[string]plugin.Symbol{}

	return &pm
}

func (s *PluginManager) LoadPlugins(file string)(error){
	pluginJson, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("[Error] %s", err)
		return err
	}

	var pd []PluginData
	err = json.Unmarshal(pluginJson, &pd)
	if err != nil {
		log.Printf("[Error] %s", err)
		return err
	}

	for _, p := range pd {
		util.LoadPlugin(&s.mention, p.Name, p.Path)
	}
	return nil
}

func _main(args []string) int {
	sc,err := LoadSlackSettings("./slack.json")

	if(err != nil){
		log.Printf("[Error] %s", err)
		return 1
	}
		
	log.Printf("[Info] Start slack event listening ")
	client := slack.New(sc.BotToken)

	//client.SetDebug(true)
	pm := NewPluginManager()
	pm.LoadPlugins("./plugin.json")
	if err != nil {
		log.Printf("[Error] %s", err)
		return 1
	}

	slackListener := &SlackListener{
		client:    client,
		botID:     sc.BotID,
		channelID: sc.ChannelID,
		allmsg:    pm.allmsg,
		mention:   pm.mention,
	}

	go slackListener.ListenAndResponse()

	http.Handle("/interaction", interactionHandler{
		verificationToken: sc.VerificationToken,
	})

	log.Printf("[Info] Server listening on: %s", sc.Port)
	if err := http.ListenAndServe(":"+sc.Port, nil); err != nil {
		log.Printf("[Error] %s", err)
		return 1
	}

	return 0
}

