package main

import (
	"errors"
	"strings"
)

var (
	ErrNotBeginWithVerb = errors.New("the sentence does not begin with a verb")
	ErrNoNamesAfterVerb = errors.New("no words were found after the verb")
	ErrInvalidCharacter = errors.New("this character is not an alphabetic character")
)

type Sentence struct {
	Verb         string
	Names        []string
	SubSentences []*Sentence
	following    string
	// number       int
}

func (s *Sentence) CalculateNumber() (number int) {
	var numbers []int
	if len(s.Names) != 0 {
		for _, name := range s.Names {
			numbers = append(numbers, NameToNumber(name))
		}

	} else if len(s.SubSentences) != 0 {
		for _, subSentence := range s.SubSentences {
			numbers = append(numbers, subSentence.CalculateNumber())
		}
	}

	number = Verbs[s.Verb](numbers)

	return
}

func ParseSentence(str string) (sentence *Sentence, err error) {
	var (
		subSentence *Sentence
	)

	if sentence, err = ParseSentenceWithFollowing(str); err != nil {
		return
	}

	if len(sentence.Names) == 0 && sentence.following == "" {
		return nil, ErrNoNamesAfterVerb
	}

	if len(sentence.Names) == 0 {
		subSentence = sentence
		for {
			if subSentence, err = ParseSentenceWithFollowing(subSentence.following); err != nil {
				return
			}

			sentence.SubSentences = append(sentence.SubSentences, &Sentence{
				Verb:         subSentence.Verb,
				Names:        subSentence.Names,
				SubSentences: subSentence.SubSentences,
			})

			if subSentence.following == "" {
				break
			}
		}
	}

	sentence.following = ""
	// if subSentence

	return
}

func ParseSentenceWithFollowing(str string) (sentence *Sentence, err error) {
	words := strings.Split(str, " ")

	if !IsVerb(words[0]) {
		return nil, ErrNotBeginWithVerb
	}

	sentence = &Sentence{
		Verb: words[0],
	}

	for idx, word := range words[1:] {
		if IsVerb(word) {
			// sentence.Names = words[1:]
			sentence.following = strings.Join(words[idx+1:], " ")
			return
		}

		for _, char := range word {
			if !IsAlphabet(byte(char)) {
				return nil, ErrInvalidCharacter
			}
		}

		sentence.Names = append(sentence.Names, word)
	}

	return
}

// func findVerb(sentence string) (verb string) {
// 	for v, _ := range Verbs {
// 		if strings.HasPrefix(sentence, v) {
// 			return v
// 		}
// 	}

// 	return
// }
