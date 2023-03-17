package models

import "xulu/utils/operations"

var Verbs = map[string]func([]int) int{
	"abcd": operations.Addition,
	"bcde": operations.Subtraction,
	"dede": operations.Multiplication,
}

func IsVerb(word string) bool {
	for verb := range Verbs {
		if word == verb {
			return true
		}
	}

	return false
}
