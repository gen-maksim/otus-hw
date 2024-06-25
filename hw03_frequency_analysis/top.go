package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func Top10(s string) []string {
	words := map[string]int{}

	// count words
	for _, word := range strings.Fields(s) {
		words[word]++
	}

	// move from map to slice
	wordSlice := make([]wordCount, 0, len(words))
	for word, value := range words {
		if word == "" {
			continue
		}
		wordSlice = append(wordSlice, wordCount{word: word, count: value})
	}

	// sort slice by count and lex
	sort.Slice(wordSlice, func(i, j int) bool {
		if wordSlice[i].count == wordSlice[j].count {
			return (wordSlice[i].word) < (wordSlice[j].word)
		}
		return wordSlice[i].count > wordSlice[j].count
	})

	// leave only first ten
	res := []string{}
	i := 0
	for _, wc := range wordSlice {
		res = append(res, wc.word)
		i++
		if i == 10 {
			break
		}
	}

	return res
}
