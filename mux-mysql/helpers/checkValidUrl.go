// Package helpers contains the utility functions
package helpers

import (
	"log"
	"regexp"
	"strings"
)

func CheckValidURL(link string) bool {
	r, err := regexp.Compile("^(http|https)://")
	if err != nil {
		log.Print(err)
		return false
	}
	link = strings.TrimSpace(link)
	return r.MatchString(link)
}
