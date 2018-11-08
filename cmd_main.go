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
	//	"plugins/echo"
	"plugin"
	//  "echo"

)

type envConfig struct {
//	Port              string `envconfig:"PORT" default: "3000"`
//	BotToken          string `envconfig:"BOT_TOKEN" required: "true"`
//	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required: "true"`
	BotID             string `envconfig:"BOT_ID" require: "true"`
//	ChannelID         string `envconfig:"CHANNEL_ID" require: "true"`
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {

	log.Printf("[Info] Start CommandLine driver ")

	//client.SetDebug(true)
//	var allmsg = map[string]plugin.Symbol{}
	var mention = map[string]plugin.Symbol{}

	loadPlugin(&mention, "memo", "plugins/memo/memo.so")
	loadPlugin(&mention, "echo", "plugins/echo/echo.so")
	//	loadPlugin(&mention, "aws", "plugins/aws/aws.so")
	//	loadPlugin(&mention, "sqs", "plugins/sqs/sqs.so")
	//loadPlugin(&mention, "ecr", "plugins/ecr/ecr.so")
	//loadPlugin(&mention, "ecs", "plugins/ecs/ecs.so")

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
		if (msgs[0] == "regina") {
			for key, value := range mention {
				if msgs[1] == key {
					log.Print(value.(func([]string) string)(msgs[2:]))
					return 0
				}
			}
		}
	}


//*/


	return 0
}

