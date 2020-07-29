/*
 * @lc app=leetcode.cn id=819 lang=golang
 *
 * [819] 最常见的单词
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	ban := map[string]bool{}
	for _, str := range banned {
		ban[str] = true
	}

	times := map[string]int{}
	paragraph = strings.ToLower(paragraph)
	for i := 0; i < len(paragraph); i++ {
		if !isAlphabet(paragraph[i]) {
			continue
		}

		j := i
		for i < len(paragraph) && isAlphabet(paragraph[i]) {
			i++
		}
		if _, ok := ban[paragraph[j:i]]; !ok {
			times[paragraph[j:i]]++
		}
		i--
	}

	res, count := "", 0
	for str, val := range times {
		if val > count {
			count = val
			res = str
		}
	}

	return res
}

func isAlphabet(element byte) bool {
	return ('a' <= element && element <= 'z') || ('A' <= element && element <= 'Z') 
}
// @lc code=end

