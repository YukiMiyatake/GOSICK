package main

import (
	"plugin"
	"log"
)

//////// container lib algorithm
type Traits struct {
	Val interface{}
}
/*
type IntVal struct{val int}
func (u IntVal) Val() Traits {
	return (u)
}

type StringVal struct{val string}
func (u StringVal) Val() Traits {
	return (u);
}
*/
func Contains(arr []interface{}, val interface{}) bool{
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

func loadPlugin(plug *map[string]plugin.Symbol, name string, path string) {
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


