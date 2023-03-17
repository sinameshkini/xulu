package models

import (
	"errors"
	"strings"
)

// Definition of errors
var (
	ErrNotBeginWithVerb = errors.New("the sentence does not begin with a verb")
	ErrNoNamesAfterVerb = errors.New("no words were found after the verb")
	ErrInvalidCharacter = errors.New("this character is not an alphabetic character")
)

// sentence data model
type Sentence struct {
	Verb         string      // the verb of the sentence contains one of the permitted verbs
	Names        []string    // words of the sentence
	SubSentences []*Sentence // parts of separate subsentences
	following    string      // the unparsed part of the sentence
}

// Function to compute the equivalent number of a Xulu sentence
func (s *Sentence) CalculateNumber() (number int) {
	var numbers []int
	if len(s.Names) != 0 {
		// this means that the sentence normally has one verb and several names
		for _, name := range s.Names {
			numbers = append(numbers, NameToNumber(name))
		}

	} else if len(s.SubSentences) != 0 {
		// this means that it is a compound sentence and has a verb and several sub-sentences
		for _, subSentence := range s.SubSentences {
			numbers = append(numbers, subSentence.CalculateNumber())
		}
	}

	// perform the specified operation of the verb on the sentence components
	number = Verbs[s.Verb](numbers)

	return
}

// ParseSentence checks if a Xulu sentence follows the grammar rules and map string into model
func ParseSentence(str string) (sentence *Sentence, err error) {
	var (
		subSentence *Sentence
	)

	// find first, return error when not started with verb
	if sentence, err = ParseSentenceWithFollowing(str); err != nil {
		return
	}

	// this means no words were found after the verb
	if len(sentence.Names) == 0 && sentence.following == "" {
		return nil, ErrNoNamesAfterVerb
	}

	if len(sentence.Names) == 0 {
		// this means sentence has no names and has other sentences in following
		subSentence = sentence
		// TODO refactor this func with concurrent methods
		for {
			if subSentence, err = ParseSentenceWithFollowing(subSentence.following); err != nil {
				return
			}

			// append sub-sentence and clear following
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

	// clear following
	sentence.following = ""

	return
}

// ParseSentenceWithFollowing Parse a sentence and put continue in following filed for next steps
func ParseSentenceWithFollowing(str string) (sentence *Sentence, err error) {
	// split words into slice
	words := strings.Split(str, " ")

	// sentence must begins with verb
	if !IsVerb(words[0]) {
		return nil, ErrNotBeginWithVerb
	}

	// set verb
	sentence = &Sentence{
		Verb: words[0],
	}

	for idx, word := range words[1:] {
		if IsVerb(word) {
			// this means this sentence is over
			// put other words to following to next steps proccessing
			sentence.following = strings.Join(words[idx+1:], " ")
			return
		}

		// check word characters validation
		for _, char := range word {
			if !IsAlphabet(byte(char)) {
				return nil, ErrInvalidCharacter
			}
		}

		sentence.Names = append(sentence.Names, word)
	}

	return
}
