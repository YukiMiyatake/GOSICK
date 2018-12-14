package util

import (
	"plugin"
	"io/ioutil"
	"log"
	"encoding/json"
)

type PluginData struct {
	Name string `json:"NAME"`
	Path string `json:"PATH"`
}
type PluginManager struct {
	Promiscuous  map[string]plugin.Symbol
	Mention map[string]plugin.Symbol

}


func NewPluginManager()(*PluginManager){
	s := PluginManager{}
	s.Promiscuous = map[string]plugin.Symbol{}
	s.Mention = map[string]plugin.Symbol{}

	return &s
}

func (s *PluginManager) LoadPlugins(file string)(error){

	pluginJson, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("[Error] %s", err)
		return err
	}

	var pd []PluginData
	err = json.Unmarshal(pluginJson, &pd)
	if err != nil {
		log.Printf("[Error] %s", err)
		return err
	}

	for _, p := range pd {
		LoadPlugin(&s.Mention, p.Name, p.Path)
	}
	return nil
}


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
