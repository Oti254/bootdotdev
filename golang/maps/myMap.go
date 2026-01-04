package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) == len(phoneNumbers) {
		for i, name := range names {
			userMap[name] = user{name, phoneNumbers[i]}
		}
		return userMap, nil
	}
	return userMap, errors.New("invalid sizes")
}

type user struct {
	name        string
	phoneNumber int
}

