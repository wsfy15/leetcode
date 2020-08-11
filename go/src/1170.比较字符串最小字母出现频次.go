/*
 * @lc app=leetcode.cn id=1170 lang=golang
 *
 * [1170] 比较字符串最小字母出现频次
 */

// @lc code=start
func numSmallerByFrequency(queries []string, words []string) []int {
	fw := make([]int, len(words))
	for index, word := range words {
		fw[index] = f(word)
	}

	res := make([]int, len(queries))
	for index, query := range queries {
		fq := f(query)
		for _, v := range fw {
			if v > fq {
				res[index]++
			}
		}
	}

	return res
}

func f(s string) int {
	counter := [26]int{}
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if counter[i] > 0 {
			return counter[i]
		}
	}
	return 0
}

// @lc code=end

