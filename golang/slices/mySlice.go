package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	switch {
	case plan == "pro":
		return messages[:], nil

	case plan == "free":
		return messages[:2], nil

	default:
		err := errors.New("unsupported plan")
		return nil, err
	}
}
