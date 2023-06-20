package main

import (
	"sort"
	"strconv"
	"strings"
)

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	t := 0
	last := len(needle) - 1
	for t < len(haystack)-last {
		p := 0
		for p <= last && haystack[t+p] == needle[p] {
			p++
		}
		if p == len(needle) {
			return t
		}
		t++
	}
	return -1
}

// https://leetcode.com/problems/merge-strings-alternately/
// 1768. Merge Strings Alternately
func mergeAlternately(word1 string, word2 string) string {
	var res string
	l := min(len(word1), len(word2))

	for i := 0; i < l; i++ {
		res += string(word1[i])
		res += string(word2[i])

	}

	res += word1[l:]
	res += word2[l:]

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://leetcode.com/problems/find-the-difference/
// 389. Find the Difference
func findTheDifference(s string, t string) byte {
	chars := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		chars[t[i]]++
		chars[s[i]]--
	}
	chars[t[len(t)-1]]++

	for char, value := range chars {
		if value != 0 {
			return char
		}
	}
	return 0
}

func findTheDifferenceVer2(s string, t string) byte {
	var res byte

	for i := 0; i < len(s); i++ {
		res += t[i]
		res -= s[i]
	}
	res += t[len(t)-1]
	return res
}

// 242. Valid Anagram
// https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]int)

	var counter int = 0

	for i := 0; i < len(s); i++ {
		if m[s[i]]++; m[s[i]] == 0 {
			counter++
		}

		if m[t[i]]--; m[t[i]] == 0 {
			counter++
		}
	}

	return counter == len(s)
}

// 459. Repeated Substring Pattern
// https://leetcode.com/problems/repeated-substring-pattern/
func repeatedSubstringPattern(s string) bool {

	var sub string = string(s[0])

	for i := 0; i < len(s)/2; i++ {

		flag := true

		for j := i + 1; j < len(s); j += len(sub) {
			if len(sub) > len(s[j:]) {
				flag = false
				break
			}

			if s[j:j+len(sub)] != sub {
				flag = false
				break
			}
		}

		if flag {
			return true
		}

		sub += string(s[i+1])
	}
	return false

}

// 1. Two Sum - Golang edition with using map
// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) (arr []int) {

	var d = make(map[int]int)
	arr = []int{0, 0}
	for i, v := range nums {
		var r = target - v

		if _, ok := d[r]; ok {
			arr = []int{d[r], i}
			return
		}

		d[v] = i
	}
	return
}

// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/
func moveZeroes(nums []int) {

	var c int16 = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			c++
			i--
		}
	}
	nums = append(nums, make([]int, c)...)
}

// 66. Plus One
// https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append([]int{1}, digits...)
	return digits

}

// 1822. Sign of the Product of an Array
// https://leetcode.com/problems/sign-of-the-product-of-an-array/
func arraySign(nums []int) int {

	flag := true
	for _, num := range nums {
		if num < 0 {
			flag = !flag

		} else if num == 0 {
			return 0
		}
	}
	if flag {
		return 1
	} else {
		return -1
	}
}

// 1502. Can Make Arithmetic Progression From Sequence
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
func canMakeArithmeticProgression(arr []int) bool {

	sum, max, min := 0, arr[0], arr[0]

	for _, value := range arr {
		sum += value

		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	var d int = max - min

	if d%(len(arr)-1) != 0 {
		return false
	}

	d /= len(arr) - 1
	if d == 0 {
		return true
	}

	m := make(map[int]int)
	for i := min; i <= max; i += d {
		m[i]++
	}

	for _, val := range arr {
		m[val]--
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}

	return true
}

// 896. Monotonic Array
// https://leetcode.com/problems/monotonic-array/
func isMonotonic(nums []int) bool {
	d := 0

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			if d == 0 {
				d = 1
			} else if d == -1 {
				return false
			}
		} else if nums[i-1] > nums[i] {
			if d == 0 {
				d = -1
			} else if d == 1 {
				return false
			}
		}
	}
	return true
}

// 13. Roman to Integer
// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {

	dict := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0
	for i := 0; i < len(s)-1; i++ {
		if dict[s[i]] < dict[s[i+1]] {
			res -= dict[s[i]]
		} else {
			res += dict[s[i]]
		}
	}
	return res + dict[s[len(s)-1]]
}

// 58. Length of Last Word
// https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {

	var count int = 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if s[i] == ' ' && count != 0 {
			return count
		}
	}
	return count
}

// 709. To Lower Case
// https://leetcode.com/problems/to-lower-case/
func toLowerCase(s string) string {

	for i := 0; i < len(s); i++ {
		if s[i] >= 41 && s[i] <= 90 {
			if i != len(s)-1 {
				s = s[:i] + string(s[i]|32) + s[i+1:]
			} else {
				s = s[:i] + string(s[i]|32)
			}

		}
	}

	return s
}

func toLowerCaseV2(s string) string {
	return strings.ToLower(s)
}

// 682. Baseball Game
// https://leetcode.com/problems/baseball-game/
func calPoints(operations []string) int {
	var scores []int
	for _, operation := range operations {

		switch operation {
		case "D":
			if len(scores) != 0 {
				scores = append(scores, 2*scores[len(scores)-1])

			}
		case "C":
			scores = scores[:len(scores)-1]
		case "+":
			if len(scores) >= 2 {
				scores = append(scores, scores[len(scores)-1]+scores[len(scores)-2])
			}
		default:
			v, _ := strconv.Atoi(operation)

			scores = append(scores, v)
		}
	}

	var res int = 0
	for _, score := range scores {
		res += score
	}
	return res
}

// 657. Robot Return to Origin
// https://leetcode.com/problems/robot-return-to-origin/
func judgeCircle(moves string) bool {

	x, y := 0, 0

	for _, char := range moves {
		switch char {
		case 'L':
			x++
		case 'R':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}

	return x == 0 && y == 0
}

// 1275. Find Winner on a Tic Tac Toe Game
// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
func tictactoe(moves [][]int) string {

	// 0, 1, 2 - rows; 3,4,5 - columns; 6,7 diagonals
	var A, B []int = make([]int, 8), make([]int, 8)

	for i := 0; i < len(moves); i++ {
		var player []int

		if i%2 == 0 {
			player = A
		} else {
			player = B
		}

		player[moves[i][0]]++
		player[moves[i][1]+3]++

		if moves[i][0] == moves[i][1] {
			player[6]++
		}
		if moves[i][0]+moves[i][1] == 2 {
			player[7]++
		}
	}

	for i := 0; i < 8; i++ {
		if A[i] == 3 {
			return "A"
		}
		if B[i] == 3 {
			return "B"
		}
	}

	if len(moves) < 9 {
		return "Pending"
	}
	return "Draw"

}

// 1672. Richest Customer Wealth
// https://leetcode.com/problems/richest-customer-wealth/
func maximumWealth(accounts [][]int) int {
	max := 0

	accountSum := func(arr []int, ch chan int) {
		temp := 0
		for _, v := range arr {
			temp += v
		}
		ch <- temp
	}

	ch := make(chan int)

	for i := 0; i < len(accounts); i++ {
		go accountSum(accounts[i], ch)
	}

	for i := 0; i < len(accounts); i++ {
		temp := <-ch

		if temp > max {
			max = temp
		}
	}

	return max
}

// 1572. Matrix Diagonal Sum
// https://leetcode.com/problems/matrix-diagonal-sum/
func diagonalSum(mat [][]int) int {

	if len(mat) == 1 {
		return mat[0][0]
	}

	var sum int
	for i := 0; i < len(mat)/2; i++ {
		j := len(mat) - i - 1

		sum += mat[i][i]
		sum += mat[i][j]
		sum += mat[j][i]
		sum += mat[j][j]
	}

	if len(mat)%2 == 1 {
		sum += mat[len(mat)/2+1][len(mat)/2+1]
	}

	return sum
}

// 1491. Average Salary Excluding the Minimum and Maximum Salary
// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
func average(salary []int) float64 {
	min, max := salary[0], salary[0]
	sum := 0

	for _, v := range salary {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}

	sum -= min + max

	return float64(sum) / float64(len(salary)-2)
}

// 860. Lemonade Change
// https://leetcode.com/problems/lemonade-change/
func lemonadeChange(bills []int) bool {

	var count5, count10 int

	for _, bill := range bills {

		if bill == 5 {
			count5++
		} else if bill == 10 {
			if count5 < 1 {
				return false
			}

			count5--
			count10++

		} else if bill == 20 {
			if count10 >= 1 && count5 >= 1 {
				count10--
				count5--
			} else if count5 >= 3 {
				count5 -= 3
			} else {
				return false
			}
		}
	}

	return true
}

// 976. Largest Perimeter Triangle
// https://leetcode.com/problems/largest-perimeter-triangle/
func largestPerimeter(nums []int) int {

	sort.Sort(sort.IntSlice(nums)) // O(log(n)*n)

	// O(n)
	for i := len(nums) - 1; i >= 2; i-- {
		if !(nums[i] >= nums[i-1]+nums[i-2]) {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}

	return 0
}
