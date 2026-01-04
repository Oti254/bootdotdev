package main
import (
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	/** Looping through each message **/
	for i, message := range messages{
		messages[i].tags = tagger(message)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	lowerCase := strings.ToLower(msg.content)
	isUrgent := false
	isPromo := false

	/** Check if the string is urgent **/
	if strings.Contains(lowerCase, "urgent") {
		isUrgent = true
	}

	/** Check if the string is promo **/
	if strings.Contains(lowerCase, "sale") {
		isPromo = true
	}

	/** Appending the necessary tags **/
	if isUrgent {
		tags = append(tags, "Urgent")
	}

	if isPromo {
		tags = append(tags, "Promo")
	}
	return tags
}

