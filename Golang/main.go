package main

import (
	"fmt"
)

func main() {

	word1, word2 := "a", "pqr"
	answer := "apbqcr"
	res := mergeAlternately(word1, word2)

	fmt.Println(res)
	fmt.Println(res == answer)

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
