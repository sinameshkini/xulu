package main

func Check(input string) (code int, err error) {
	var (
		sentence *Sentence
	)

	if sentence, err = ParseSentence(input); err != nil {
		return
	}

	code = sentence.CalculateNumber()

	return
}
