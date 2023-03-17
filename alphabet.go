package main

import "math"

var Alphabets = map[byte]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
}

func IsAlphabet(char byte) bool {
	for a := range Alphabets {
		if char == a {
			return true
		}
	}

	return false
}

// Function to convert Xulu name to its numerical equivalent
func NameToNumber(name string) (number int) {
	var count = 1

	if len(name) == 0 {
		return 0
	}

	if len(name) == 1 {
		return int(math.Pow(float64((count*Alphabets[name[0]])%5), 2))
	}

	for idx, c := range name {
		if idx == len(name)-1 || c != rune(name[idx+1]) {
			number += int(math.Pow(float64((count*Alphabets[name[idx]])%5), 2))
			count = 1
		} else {
			count++
		}
	}

	return
}
