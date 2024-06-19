package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	Count int
}

func Top10(s string) []string {
	words := map[string]wordCount{}

	// count words
	for _, word := range strings.Fields(s) {
		if _, ok := words[word]; ok {
			i := words[word]
			i.Count = i.Count + 1
			words[word] = i
		} else {
			words[word] = wordCount{word: word, Count: 1}
		}
	}

	// move from map to slice
	wordSlice := make([]wordCount, 0, len(words))
	for _, value := range words {
		if value.word == "" {
			continue
		}
		wordSlice = append(wordSlice, value)
	}

	// sort slice by count and lex
	sort.Slice(wordSlice, func(i, j int) bool {
		if wordSlice[i].Count == wordSlice[j].Count {
			return (wordSlice[i].word) < (wordSlice[j].word)
		}
		return wordSlice[i].Count > wordSlice[j].Count
	})

	// leave only first ten
	var res []string
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
