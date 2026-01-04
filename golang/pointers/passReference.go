package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func analyzeMessage(analytics *Analytics, msg Message) {
	if msg.Success {
		analytics.MessagesTotal++
		analytics.MessagesSucceeded++
	} else {
		analytics.MessagesTotal++
		analytics.MessagesFailed++
	}
}
