package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("starting xulu ...")

	var (
		sampleSentences = []string{
			"abcd abcd aabbc ab a c ccd dede cccd cd",
			"dede abcd abd abddd ddada dac abcd de ed",
		}
	)

	for idx, s := range sampleSentences {
		fmt.Printf("sample %d: %s \n", idx+1, s)

		code, err := grammerCheck(s)
		if err != nil {
			fmt.Println("error:", err.Error())
			os.Exit(1)
		}

		fmt.Printf("number equivalent to the code: %d \n", code)
		fmt.Println("_____________________________________________________________________")
	}

	fmt.Println("xulu stopping")
	os.Exit(0)
}
