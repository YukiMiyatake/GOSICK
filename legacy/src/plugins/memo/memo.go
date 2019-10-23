// メモを保存呼び出しする
// xx memo hoge="aaaaaa""　で登録、 xx memo hoge で呼び出し  xx memo hoge="" で削除
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type Route struct {
	route *regexp.Regexp
	fnc   func([]string) string
	help  string
	//	arg   *[]string
}

// TODO: ダックタイプインタフェースにする
type SlackMsg struct {
	routes *[]Route
	header Header
}

var s SlackMsg

type Header struct {
	Version string `json:"version"`
	Data    []Data `json:"data"`
}

type Data struct {
	Keyword string `json:"keyword"`
	Value   string `json:"value"`
}

const (
	DATA_FILE_NAME = "memo.txt"
)

func Init() {

	routes := []Route{}
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^help`), fnc: cmdHelp, help: "'help'   ヘルプ表示"})

	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^(.*)=(.*)`), fnc: cmdRegist, help: "'keyword=value'  Keyword登録"})
	routes = append(routes, Route{route: regexp.MustCompile(`(?i)^(.*)`), fnc: cmdRead, help: "'Keyword'  Keyword表示"})

	s.routes = &routes

	type LatLng [2]float64

	raw, err := ioutil.ReadFile(DATA_FILE_NAME)
	if err != nil {
		fmt.Println(err.Error())
		//        os.Exit(1)
	}

	//var fc FeatureCollection

	var header []Header
	json.Unmarshal(raw, &header)

}

func OnMention(msgs []string) string {

	var response = ""

	ss := strings.TrimSpace(strings.Join(msgs[:], " "))

	for _, value := range *s.routes {
		match := value.route.FindStringSubmatch(ss)

		if match != nil {
			response = value.fnc(match[1:])
			break
		}
	}

	return response
}

func cmdHelp(msgs []string) string {
	var response string = "regist: keyword=value  read:Keyword"
	/*
		for _, value := range *s.routes {
			help := value.help

			if len(help) > 0 {
				response += help
				response += "\n"
			}
		}
	*/
	return (response)
}

func cmdRegist(msgs []string) string {

	keyword := ""
	value := ""
	msg := ""
	data := &s.header.Data

	if len(msgs) > 1 {
		// Regist or Delete
		keyword = msgs[0]
		value = msgs[1]

		if value == "" {
			b, n := contains(*data, keyword)
			if b {
				unset(*data, n)
				msg = "忘れたよ"
			} else {
				msg = "なぁに？"
			}
		} else {
			b, n := contains(*data, keyword)
			if b {
				(*data)[n].Value = value
			} else {

				*data = append(*data, Data{Keyword: keyword, Value: value})

			}
			msg = "覚えたよ"
		}
		//	} else if len(msgs) > 0 {
		// Read : cant reach
	}

	return (msg)
}

func cmdRead(msgs []string) string {

	keyword := ""
	//	value := ""
	msg := ""
	data := &s.header.Data

	if len(msgs) > 0 {
		keyword = msgs[0]
		b, n := contains(*data, keyword)
		if b {
			msg = (*data)[n].Value
		} else {
			msg = "なぁに？"
		}

	}

	return (msg)
}

func contains(h []Data, s string) (bool, int) {
	for n, v := range h {
		if v.Keyword == s {
			return true, n
		}
	}
	return false, 0
}

func unset(s []Data, i int) []Data {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
