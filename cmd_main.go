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

	"github.com/kelseyhightower/envconfig"
)

type cmdConfig struct {
	BotName             []string `envconfig:"BOT_NAME" default: {"gosick", "regina"}  required: "false"`
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	log.Printf("[Info] Start CommandLine driver ")

/*
	va := []int  {1, 2, 3, 4}
	log.Printf( strconv.FormatBool( Contains( va, "1")) )
	log.Printf(strconv.FormatBool(Contains( va, 100)))

	vs := []string  {"1", "2", "3", "4"}
	log.Printf( strconv.FormatBool( Contains( vs, "1")) )
*/


	var env cmdConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Print("error")
	}
	//client.SetDebug(true)
//	var allmsg = map[string]plugin.Symbol{}
	var mention = map[string]plugin.Symbol{}

	// TODO: yuki load from Json
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
		// TODO: load from Env or JSON
		if (Contains( []string{"regina","gosick"}, msgs[0])) {
//		if (msgs[0] == "regina") {
			for key, value := range mention {
				if msgs[1] == key {
					log.Print(value.(func([]string) string)(msgs[2:]))
				}
			}
		}
	}


		//*/


	return 0
}

