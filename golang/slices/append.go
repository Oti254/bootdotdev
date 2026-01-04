package main

type cost struct {
	day int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCost := []float64{}
	days := []int{}
	for i := 0; i < len(costs); i++ {
		price := costs[i]
		days = append(days, price.day)

		/** Appends only the specified day **/
		if price.day == day {
			dayCost = append(dayCost, price.value)
		}
	}
	return dayCost
}
