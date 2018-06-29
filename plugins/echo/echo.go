package main

import (
//	"fmt"
	"strings"
)

func Init() {
}

func OnMention(s []string) (string) {
        return strings.Join(s, " ");
}

