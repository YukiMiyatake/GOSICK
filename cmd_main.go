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

	"io/ioutil"
	"encoding/json"
)

type cmdConfig struct {
	BotName             []string `json:"BOT_NAME"`
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

	jsonData, err := ioutil.ReadFile("./slack.json")
	if err != nil {
		log.Print("error")
	}

	var cc cmdConfig

	err = json.Unmarshal(jsonData, &cc)
	if err != nil {
		log.Print("error")
	}
	//client.SetDebug(true)
//	var allmsg = map[string]plugin.Symbol{}
	var mention = map[string]plugin.Symbol{}

	pluginJson, err := ioutil.ReadFile("./plugin.json")
	if err != nil {
		log.Printf("[Error] %s", err)
		return 1
	}
	type PluginData struct {
		Name string `json:"NAME"`
		Path string `json:"PATH"`
	}
	var pd []PluginData
	err = json.Unmarshal(pluginJson, &pd)
	if err != nil {
		log.Printf("[Error] %s", err)
		return 1
	}
	for _, p := range pd {
		loadPlugin(&mention, p.Name, p.Path)
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
		if (Contains( cc.BotName, msgs[0])) {
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

