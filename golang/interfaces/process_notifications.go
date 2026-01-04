package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dM directMessage) importance() int {
	if dM.isUrgent {
		return 50
	}
	return dM.priorityLevel
}

func (gM groupMessage) importance() int {
	return gM.priorityLevel
}

func (sA systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch pN := n.(type){
	case directMessage:
		return pN.senderUsername, n.importance()

	case groupMessage:
		return pN.groupName, n.importance()

	case systemAlert:
		return pN.alertCode, n.importance()      

	default:
		return "", 0
	}
}


