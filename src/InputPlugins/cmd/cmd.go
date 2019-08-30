package Cmd

// 使ってない
import (
	"InputPlugins"
	"bufio"

	"log"
	"os"
	"strings"

	"util"
)

func init() {
	log.Printf("[Info] Start CommandLine driver ")
}

func Echo(sc InputPlugins.ISlackConfig, pm *util.PluginManager) {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		log.Print(text)
		if strings.ToLower(text) == "quit" {
			break
		}

		msgs := strings.Fields(text)
		// TODO: load from Env or JSON
		md := util.NewMessageDispatcher(&sc.GetBotName(), &sc.GetBotID())
		if md.GetMessageType(msgs[0]) == util.Mention {
			for key, value := range pm.Mention {
				if msgs[1] == key {
					log.Print(value.(func([]string) string)(msgs[2:]))
				}
			}
		}
	}
}
