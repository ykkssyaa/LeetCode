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
