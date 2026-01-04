package main

import (
	"errors"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	_, ok := users[name]
	userDelete := users[name].scheduledForDeletion
	err = errors.New("not found")
	
	if ok {
		if userDelete {
			delete(users, name)
			deleted = true
			err = nil
			return deleted, err
		} else {
			deleted = false
			err = nil
			return deleted, err
		}
	}

	return deleted, err
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

