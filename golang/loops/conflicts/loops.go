package main


func bulkSend(numMessages int) float64 {
	var cost float64
	var messageCost float64
	var totalCost float64

	for i := numMessages; i > 0; i-- {
		messageCost = 1 + cost
		totalCost += messageCost
		cost += 0.01
	}
	return totalCost
}
