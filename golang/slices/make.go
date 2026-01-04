package main

func getMessageCosts(messages []string) []float64 {
	messageCost := make([]float64, len(messages))
	/**
	Setting the cap() to the len(messages)
	messageCost := make([]float64, 0,  len(messages))
	**/
	for i := 0; i < len(messages); i++ {
		/** messageCost = append(messageCost, float64(len(messages[i])) * 0.01) **/
		messageCost[i] = float64(len(messages[i])) * 0.01
	}
	return messageCost
}
