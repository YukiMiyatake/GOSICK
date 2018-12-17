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
	//	"plugins/echo"
	"plugin"
	//  "echo"

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


func _main(args []string) int {
	sc,err := LoadSlackSettings("./slack.json")

	if(err != nil){
		log.Printf("[Error] %s", err)
		return 1
	}
		
	log.Printf("[Info] Start slack event listening ")
	client := slack.New(sc.BotToken)

	//client.SetDebug(true)
	var allmsg = map[string]plugin.Symbol{}
	var mention = map[string]plugin.Symbol{}

	pluginJson, err := ioutil.ReadFile("./plugin.json")
	if err != nil {
		log.Printf("[Error] %s", err)
		return 1
	}

	type PluginData struct {
		Name string `json:"NAME"`
		Path string `json:"PATH"`
	}

	var pd []PluginData
	err = json.Unmarshal(pluginJson, &pd)
	if err != nil {
		log.Printf("[Error] %s", err)
		return 1
	}

	for _, p := range pd {
		loadPlugin(&mention, p.Name, p.Path)
	}

	slackListener := &SlackListener{
		client:    client,
		botID:     sc.BotID,
		channelID: sc.ChannelID,
		allmsg:    allmsg,
		mention:   mention,
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

