package main

import (
	"os"
	"strings"
)

func validName(name string) bool {
	return len(strings.TrimSpace(name)) >= 3
}

func validDict() bool {
	_, err := os.Stat(DictPath)
	if err != nil {
		return false
	}
	return true
}
