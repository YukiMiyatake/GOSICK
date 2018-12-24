// +build cmd_driver

/*
	Slackに接続を行わず　コマンドラインインタフェースで行う
*/
package main

import (
	"log"
	"os"
	"github.com/YukiMiyatake/GOSICK/util"

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

	pm := util.NewPluginManager()
	pm.LoadPlugins("./plugin.json")

	if err != nil {
		log.Printf("[Error] %s", err)
		return 1
	}


//*
//	コマンドライン入力
	echo := util.LoadOutputProcess("plugins/cmd/cmd.so")
	echo.(func(util.SlackConfig, *util.PluginManager))(sc, pm)




		//*/


	return 0
}

