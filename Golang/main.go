package main

import "fmt"

func main() {

	matrix := [][]int{{1, 0}, {2, 4}, {3, 6}, {1, 0}}
	setZeroes(matrix)

	fmt.Println(matrix)
}
