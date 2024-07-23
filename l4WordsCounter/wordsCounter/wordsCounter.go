package wordsCounter

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func WordsCounter(str string) map[string]int {
	storage := make(map[string]int)
	temp := ""
	runeCount := utf8.RuneCountInString(str)
	for i, char := range str {
		if unicode.IsLetter(char) {
			temp += string(char)
		}
		if (i+1 == runeCount || !unicode.IsLetter(char)) && len(temp) > 0 {
			temp = strings.ToLower(temp)
			if _, ok := storage[temp]; ok == false {
				storage[temp] = 1
			} else {
				storage[temp] += 1
			}
			temp = ""
		}
	}

	keys := make([]string, 0, len(storage))
	for k := range storage {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return storage[keys[i]] > storage[keys[j]]
	})
	length := 10
	if len(keys) < 10 {
		length = len(keys)
	}
	subSlice := keys[0:length]
	result := make(map[string]int)
	for _, key := range subSlice {
		result[key] = storage[key]
	}
	return result
}
