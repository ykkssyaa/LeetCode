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
