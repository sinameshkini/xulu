package models

import "xulu/utils"

// verbs of Xulu lang and map to operation func
var Verbs = map[string]func([]int) int{
	"abcd": utils.Addition,
	"bcde": utils.Subtraction,
	"dede": utils.Multiplication,
}

// check word is verb
func IsVerb(word string) bool {
	for verb := range Verbs {
		if word == verb {
			return true
		}
	}

	return false
}
