// +build cmd_driver

/*
	Slackに接続を行わず　コマンドラインインタフェースで行う
*/
package main

import (
	"log"
	"bufio"
	"strings"
	"os"
	"github.com/YukiMiyatake/GOSICK/util"

)

type cmdConfig struct {
	BotName             []string `json:"BOT_NAME"`
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	log.Printf("[Info] Start CommandLine driver ")

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

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan(){
		text := stdin.Text()
		log.Print(text)
		if(  strings.ToLower(text) == "quit"){
			break
		}

		msgs := strings.Fields( text )
		// TODO: load from Env or JSON
		md := util.NewMessageDispatcher( &sc.BotName, &sc.BotID)
		if (md.GetMessageType(msgs[0]) == util.Mention) {
			for key, value := range pm.Mention {
				if msgs[1] == key {
					log.Print(value.(func([]string) string)(msgs[2:]))
				}
			}
		}
	}


		//*/


	return 0
}

