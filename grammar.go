package main

import "xulu/models"

func grammerCheck(input string) (code int, err error) {
	var (
		sentence *models.Sentence
	)

	if sentence, err = models.ParseSentence(input); err != nil {
		return
	}

	code = sentence.CalculateNumber()

	return
}
