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
	"os"
	slackUtil "util/slack"
)

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	sl, err := slackUtil.NewSlackListener("./slack.json")
	if err != nil {
		log.Print(err)
		return 1
	}

	err = sl.Listen()
	if err != nil {
		log.Print(err)
		return 1
	}

	return 0
}
