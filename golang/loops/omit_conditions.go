package main


func maxMessages(thresh int) int {
	var cost int
	var messageCost int
	var total int
	
	for i := 0;; i++ {
		messageCost = 100 + cost
		if thresh < messageCost {
			break
		}
		thresh -= messageCost
		cost ++
		total ++
	}
	return total
}

