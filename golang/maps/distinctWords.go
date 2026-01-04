package main

import "strings"

func countDistinctWords(messages []string) int {
	distinctWordsStruct := make(map[string]struct{})
	for _, word := range messages {
		lowerWords := strings.ToLower(word)
		separate := strings.Fields(lowerWords)
		for _, distinct := range separate {
			if _, ok  := distinctWordsStruct[distinct]; !ok {
				distinctWordsStruct[distinct] = struct{}{}
			}
		}
	}
	return len(distinctWordsStruct)
}
