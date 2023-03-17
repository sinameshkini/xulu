package models

import (
	"errors"
	"strings"
)

var (
	ErrNotBeginWithVerb = errors.New("The sentence does not begin with a verb")
)

type Sentence struct {
	Verb         string
	Names        []string
	SubSentences []*Sentence
}

func ParseSentence(str string) (sentence *Sentence, err error) {
	var (
		words []string
	)

	words = strings.Split(str, " ")

	if !IsVerb(words[0]) {
		return nil, ErrNotBeginWithVerb
	}

	sentence = &Sentence{
		Verb:  words[0],
		Names: words[1:],
	}
	// IsVerb(words[1])

	return
}

func findVerb(sentence string) (verb string) {
	for v, _ := range Verbs {
		if strings.HasPrefix(sentence, v) {
			return v
		}
	}

	return
}
