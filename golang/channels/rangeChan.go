package main

func concurrentFib(n int) []int {
	/** creating a new channel **/
	chs := make(chan int)

	/** creating a slice to append values to **/
	numList := []int{}

	/** calling the fibonacci func concurrently **/
	go fibonacci(n, chs)
	
	/** ranging over the channel **/
	for num := range chs {
		numList = append(numList, num)
	}
	return numList
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

