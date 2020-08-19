func countSubstrings(s string) int {
	n, res := len(s), 0
	for i := 0; i < n; i++ {
		count(s, i, i, &res)
		count(s, i, i+1, &res)
	}
	return res
}

func count(s string, midL, midR int, res *int) {
	for midL >= 0 && midR < len(s) && s[midL] == s[midR] {
		*res = *res + 1
		midL--
		midR++
	}
}