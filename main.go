/*
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
*/
package main

import (
	"log"
	"net/http"
	"os"
	//	"plugins/echo"
	"plugin"
	//  "echo"

	"github.com/kelseyhightower/envconfig"
	"github.com/nlopes/slack"
)

type envConfig struct {
	Port              string `envconfig:"PORT" default: "3000"`
	BotToken          string `envconfig:"BOT_TOKEN" required: "true"`
	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required: "true"`
	BotID             string `envconfig:"BOT_ID" require: "true"`
	ChannelID         string `envconfig:"CHANNEL_ID" require: "true"`
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[Error] Failed to process env var: %s", err)
		return 1
	}

	log.Printf("[Info] Start slack event listening ")
	client := slack.New(env.BotToken)

	//client.SetDebug(true)
	var allmsg = map[string]plugin.Symbol{}
	var mention = map[string]plugin.Symbol{}

	loadPlugin(&mention, "memo", "plugins/memo/memo.so")
	loadPlugin(&mention, "echo", "plugins/echo/echo.so")
	//	loadPlugin(&mention, "aws", "plugins/aws/aws.so")
	//	loadPlugin(&mention, "sqs", "plugins/sqs/sqs.so")
	//loadPlugin(&mention, "ecr", "plugins/ecr/ecr.so")
	//loadPlugin(&mention, "ecs", "plugins/ecs/ecs.so")

	slackListener := &SlackListener{
		client:    client,
		botID:     env.BotID,
		channelID: env.ChannelID,
		allmsg:    allmsg,
		mention:   mention,
	}

	go slackListener.ListenAndResponse()

	http.Handle("/interaction", interactionHandler{
		verificationToken: env.VerificationToken,
	})

	log.Printf("[Info] Server listening on: %s", env.Port)
	if err := http.ListenAndServe(":"+env.Port, nil); err != nil {
		log.Printf("[Error] %s", err)
		return 1
	}

	return 0
}

func loadPlugin(plug *map[string]plugin.Symbol, name string, path string) {
	log.Printf("loadPlugin:" + name + ": " + path)

	p, err := plugin.Open(path)
	if err != nil {
		log.Printf("fail to load plugin [%s]", path)
		return
	}

	init, e := p.Lookup("Init")
	if e != nil {
		log.Printf("fail to Lookup 'init'")
		return
	}
	init.(func())()

	pv, err := p.Lookup("OnMention")
	if err != nil {
		log.Printf("fail to Lookup 'OnMention'")
		return
	}
	(*plug)[name] = pv
}
