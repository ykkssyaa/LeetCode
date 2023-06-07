package main

import "fmt"

func main() {

	word1, word2 := "a", "pqr"
	answer := "apbqcr"
	res := mergeAlternately(word1, word2)

	fmt.Println(res)
	fmt.Println(res == answer)

}

// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=programming-skills
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
