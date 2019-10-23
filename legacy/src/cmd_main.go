package main

/*
	Slackに接続を行わず　コマンドラインインタフェースで行う
    作り直す
*/
/*

// +build cmd_driver
// Refactor

import (
	"os"

	"util"
)

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	sc := util.SlackConfig{BotID: "id", BotName: []string{"name"}}

	pm, err := util.NewPluginManager("./plugin.json")
	if err != nil {
		println(err)
		return 0
	}

	//	コマンドライン入力
	echo := util.LoadOutputProcess("plugins/cmd/cmd.so")
	echo.(func(util.SlackConfig, *util.PluginManager))(sc, pm)

	return 0
}
*/
