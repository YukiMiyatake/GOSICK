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
	"github.com/YukiMiyatake/GOSICK/util"

	"github.com/nlopes/slack"
)



func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	sc := util.SlackConfig{}
	err := sc.LoadSlackConfig("./slack.json")

	if(err != nil){
		log.Printf("[Error] %s", err)
		return 1
	}

	log.Printf("[Info] Start slack event listening ")
	client := slack.New(sc.BotToken)


	pm := util.NewPluginManager()
	pm.LoadPlugins("./plugin.json")
	if err != nil {
		log.Printf("[Error] %s", err)
		return 1
	}

	slackListener := &SlackListener{
		client:    client,
		botID:     &sc.BotID,
		channelID: &sc.ChannelID,
		allmsg:    &pm.Promiscuous,
		mention:   &pm.Mention,
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

