package lc

import (
	"strings"
)

func ReplaceWords(dictionary []string, sentence string) string {
	sentenceWords := strings.Split(sentence, " ")
	m := len(dictionary)
	n := len(sentenceWords)
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			if strings.HasPrefix(dictionary[j], dictionary[i]) {
				dictionary[j] = dictionary[i]
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if strings.HasPrefix(sentenceWords[j], dictionary[i]) {
				sentenceWords[j] = dictionary[i]
			}
		}
	}
	return strings.Join(sentenceWords, " ")
}
