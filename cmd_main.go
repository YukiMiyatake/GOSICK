// +build cmd_driver

/*
	Slackに接続を行わず　コマンドラインインタフェースで行う
*/
package main

import (
	"os"

	"util"
)

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	sc := util.SlackConfig{BotID: "id", BotName: {"name"}}

	pm, err := util.NewPluginManager("./plugin.json")
	//	pm.LoadPlugins()

	//*
	//	コマンドライン入力
	echo := util.LoadOutputProcess("plugins/cmd/cmd.so")
	echo.(func(util.SlackConfig, *util.PluginManager))(sc, pm)

	//*/

	return 0
}
