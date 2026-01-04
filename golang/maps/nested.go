package main

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		firstLetter := runes[0]
		countsM, ok := counts[firstLetter]
		if !ok {
			countsM = make(map[string]int)
			counts[firstLetter] = countsM
		}
		countsM[name]++
	}
	return counts
}

