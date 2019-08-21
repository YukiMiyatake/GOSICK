package main

// Refactor
import (
	"bufio"
	"log"
	"os"
	"strings"

	"util"
	slackUtil "util/slack"
)

func init() {
	log.Printf("[Info] Start CommandLine driver ")
}

func Echo(sc slackUtil.SlackConfig, pm *util.PluginManager) {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		log.Print(text)
		if strings.ToLower(text) == "quit" {
			break
		}

		msgs := strings.Fields(text)
		// TODO: load from Env or JSON
		md := util.NewMessageDispatcher(&sc.BotName, &sc.BotID)
		if md.GetMessageType(msgs[0]) == util.Mention {
			for key, value := range pm.Mention {
				if msgs[1] == key {
					log.Print(value.(func([]string) string)(msgs[2:]))
				}
			}
		}
	}
}
