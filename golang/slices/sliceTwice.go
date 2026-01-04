package main

func createMatrix(rows, cols int) [][]int {
	mtx := [][]int{}

	for i := 0; i < rows; i++ {
		row := make([]int, cols)
		for j := 0; j < cols; j++ {
			row[j] = i * j
		}
		mtx = append(mtx, row)
	}
	return mtx
}
