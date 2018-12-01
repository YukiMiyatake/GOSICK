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
	"strconv"
	//  "echo"

	"github.com/kelseyhightower/envconfig"
)

type cmdConfig struct {
	BotID             []string `envconfig:"BOT_NAME" default: {"gosick"}  required: "false"`
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	log.Printf("[Info] Start CommandLine driver ")


	va := []Traits  {IntVal{1},IntVal{2},IntVal{3},IntVal{4}}
	log.Printf( strconv.FormatBool( Contains( va, IntVal{1})) )
	log.Printf(strconv.FormatBool(Contains( va, IntVal{100})))

	vs := []Traits  {StringVal{"1"},StringVal{"2"},StringVal{"3"},StringVal{"4"}}
	log.Printf( strconv.FormatBool( Contains( vs, StringVal{"1"})) )



	var env cmdConfig
	if err := envconfig.Process("", &env); err != nil {
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
		if (msgs[0] == "regina") {
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

