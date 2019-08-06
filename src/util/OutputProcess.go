// TODO: Refractor
package util

import (
	"log"
	"plugin"
)

func LoadOutputProcess(path string) plugin.Symbol {
	log.Printf("loadPlugin:" + path)

	p, err := plugin.Open(path)
	if err != nil {
		log.Printf("fail to load plugin [%s]", path)
	}

	echo, err := p.Lookup("Echo")
	if err != nil {
		log.Printf("fail to Lookup 'Echo'")
	}

	return echo
}
