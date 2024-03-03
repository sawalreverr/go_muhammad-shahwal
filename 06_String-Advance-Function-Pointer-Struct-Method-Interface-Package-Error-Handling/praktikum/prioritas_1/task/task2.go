package task

import (
	"strings"
)

func SpinString(sentence string) string {
	n := len(sentence)
	mid := n / 2
	slcSentence := strings.Split(sentence, "")
	slcWord := make([]string, n)

	for i := 0; i < mid; i++ {
		slcWord[i*2] = slcSentence[i]
		slcWord[i*2+1] = slcSentence[n-i-1]
	}

	if n%2 != 0 {
		slcWord[n-1] = slcSentence[mid]
	}

	return strings.Join(slcWord, "")
}
