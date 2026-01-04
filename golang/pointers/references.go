package main

import (
	"strings"
)

func removeProfanity(message *string) {
	bannedWords := []string{"fubb", "shiz", "witch"}
	for _, bad := range bannedWords {
		if strings.Contains(*message, bad) {
			replacement := strings.Repeat("*", len(bad))
			cleanString := strings.ReplaceAll(*message, bad, replacement)
			*message = cleanString
		}
	}
}

