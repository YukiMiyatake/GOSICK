package lib

import (
	"log"

	"plugin"
)


func LoadPlugin(plug *map[string]plugin.Symbol, name string, path string) {
	log.Printf("loadPlugin:" + name + ": " + path)

	p, err := plugin.Open(path)
	if err != nil {
		log.Printf("fail to load plugin [%s]", path)
		return
	}

	init, e := p.Lookup("Init")
	if e != nil {
		log.Printf("fail to Lookup 'init'")
		return
	}
	init.(func())()

	pv, err := p.Lookup("OnMention")
	if err != nil {
		log.Printf("fail to Lookup 'OnMention'")
		return
	}
	(*plug)[name] = pv
}


