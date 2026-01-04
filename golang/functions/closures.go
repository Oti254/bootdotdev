package main

func adder() func(int) int{
	var sum int
	return func(s int) int{
		sum += s
		return sum
	}
}
