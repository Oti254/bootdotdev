package main

import "fmt"

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; ok {
			validUsers[user]++
		}
	}
	fmt.Println(validUsers)
}
