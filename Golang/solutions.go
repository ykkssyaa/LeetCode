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
