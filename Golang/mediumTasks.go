package main

import "sync"

// 1041. Robot Bounded In Circle
// https://leetcode.com/problems/robot-bounded-in-circle/
func isRobotBounded(instructions string) bool {

	// 0, 1 - coordinates,
	// 2 - direction (0 - north, 1 - east, 2 - south, 3 - west)
	var position [3]int

	for _, instruction := range instructions {

		switch instruction {
		case 'G':
			switch position[2] {
			case 0:
				position[0]++
			case 1:
				position[1]++
			case 2:
				position[0]--
			case 3:
				position[1]--
			}

		case 'L':
			if position[2] < 3 {
				position[2]++
			} else {
				position[2] = 0
			}
		case 'R':
			if position[2] > 0 {
				position[2]--
			} else {
				position[2] = 3
			}
		}
	}

	return position[0] == 0 && position[1] == 0 || position[2] != 0
}

// 73. Set Matrix Zeroes
// https://leetcode.com/problems/set-matrix-zeroes/
func setZeroes(matrix [][]int) {

	columns := make([]int, len(matrix[0]))
	lines := make([]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				columns[j] = 1
				lines[i] = 1
			}
		}
	}

	var wg sync.WaitGroup

	for i := 0; i < len(columns); i++ {
		if columns[i] == 1 {
			wg.Add(1)
			go func(col int) {
				zeroColumn(matrix, col)
				wg.Done()
			}(i)
		}
	}
	for i := 0; i < len(lines); i++ {
		if lines[i] == 1 {
			wg.Add(1)
			go func(line int) {
				zeroLine(matrix, line)
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
}

func zeroColumn(matrix [][]int, column int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][column] = 0
	}
}
func zeroLine(matrix [][]int, line int) {
	for i := 0; i < len(matrix[line]); i++ {
		matrix[line][i] = 0
	}
}

// 54. Spiral Matrix
// https://leetcode.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {

	n, m := len(matrix), len(matrix[0])
	if n == 0 {
		return []int{}
	}

	arr := make([]int, 0, len(matrix)*len(matrix[0]))

	colBegin, colEnd := 0, m-1
	rowBegin, rowEnd := 0, n-1

	for colBegin <= colEnd && rowBegin <= rowEnd {

		for i := colBegin; i <= colEnd; i++ {
			arr = append(arr, matrix[rowBegin][i])
		}
		rowBegin++

		for j := rowBegin; j <= rowEnd; j++ {
			arr = append(arr, matrix[j][colEnd])
		}
		colEnd--

		if rowBegin <= rowEnd {
			for i := colEnd; i >= colBegin; i-- {
				arr = append(arr, matrix[rowEnd][i])
			}
			rowEnd--
		}

		if colBegin <= colEnd {
			for j := rowEnd; j >= rowBegin; j-- {
				arr = append(arr, matrix[j][colBegin])
			}
			colBegin++
		}
	}

	return arr
}

// 50. Pow(x, n)
// https://leetcode.com/problems/powx-n/
func myPow(x float64, n int) float64 {
	var res float64 = 1.0

	if x == 1 {
		return 1
	}
	if n < 0 {
		x = 1 / x
	}

	for n != 0 {
		if n%2 != 0 {
			res *= x
		}

		x *= x
		n /= 2
	}

	return res
}
