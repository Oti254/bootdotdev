package main

func sum(nums ...int) int {
	var total int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}
