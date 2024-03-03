package task

import "strings"

func PickVocals(sentence string) string {
	var result string

	for _, v := range sentence {
		if strings.ContainsAny(string(v), "aiueo ") {
			result += string(v)
		}
	}

	return result
}
