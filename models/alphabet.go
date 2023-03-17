package models

import "math"

// alphabets of Xulu lang and map to equivalent num
var Alphabets = map[byte]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
}

// check alphabet validation
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

	// empty string not acceptable
	if len(name) == 0 {
		return 0
	}

	// one char name, no need to proccess
	if len(name) == 1 {
		return xuluNumber(count, name[0])
	}

	for idx, c := range name {
		if idx == len(name)-1 || c != rune(name[idx+1]) {
			// this means repeat character ended or this is last character of name
			number += xuluNumber(count, name[idx])
			count = 1
		} else {
			count++
		}
	}

	return
}

// calculate xulu number of alphabet repetition
func xuluNumber(count int, alphabet byte) int {
	return int(math.Pow(float64((count*Alphabets[alphabet])%5), 2))
}
