package main

func countConnections(groupsize int) int {
	connections := 0
	for c := 1; c < groupsize; c++ {
		connections += c
	}
	return connections
}
