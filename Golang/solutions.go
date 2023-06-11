package main

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
