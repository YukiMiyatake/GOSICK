package main

import (
	"plugin"
	"log"
	"reflect"
	"io/ioutil"
	"encoding/json"
)

//////// container lib algorithm
// TODO: package

func Copy(arr interface{}) []interface{} {
	dest := make([]interface{}, 0)

	switch val := arr.(type) {
	case []interface{}:
	case []int:
		for _, v := range val {
			dest = append(dest, v)
		}
	case []string:
		for _, v := range val {
			dest = append(dest, v)
		}
	case int:
		dest = append(dest, val)
	case string:
		dest = append(dest, val)
	}

	return dest
}


func Contains(i interface{}, val interface{}) bool{
	arr := Copy(i)

	if(reflect.TypeOf(arr[0]) != reflect.TypeOf(val)){
		log.Printf("missmatch Type in Contains [%s] and [%s]",reflect.TypeOf(arr[0]) ,reflect.TypeOf(val) )
		panic(0)
	}

	for _, v := range arr{
		if v == val{
			return true
		}
	}
	return false
}
/*
func AllOf(arr []Traits, val Traits) bool{
	for _, v := range arr{
		if v != val{
			return false
		}
	}
	return true
}

func NoneOf(arr []Traits, val Traits) bool{
	for _, v := range arr{
		if v == val{
			return false
		}
	}
	return true
}

func CountIf(arr []Traits, val Traits) int{
	n := 0
	for _, v := range arr{
		if v == val{
			n++
		}
	}
	return n
}

*/
/////////

func loadPlugin(plug *map[string]plugin.Symbol, name string) {

	jsonString, err := ioutil.ReadFile("./plugin.json")
	if err != nil {
		log.Printf("fail to load json")
		return
	}

	type PluginData struct {
		Name string `json:"NAME"`
		Path string `json:"PATH"`
	}

	var pd []PluginData
	err = json.Unmarshal(jsonString, &pd)
	if err != nil {
		log.Printf("fail to parse json")
		return
	}
	
	var pluginName,pluginPath string
	
	for _, p := range pd {
		if(p.Name == name){
			pluginName = p.Name
			pluginPath = p.Path
		}
	}
	
	log.Printf("loadPlugin:" + pluginName + ": " + pluginPath)

	p, err := plugin.Open(pluginPath)
	if err != nil {
		log.Printf("fail to load plugin [%s]", pluginPath)
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
	(*plug)[pluginName] = pv
}
