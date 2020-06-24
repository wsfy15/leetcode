/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	wordIndex := map[string]int{}
	res := []string{}
	for i, word := range list1 {
		wordIndex[word] = i
	}

	indexSum := math.MaxInt32
	for i, word := range list2 {
		if index, ok := wordIndex[word]; ok {
			if i + index < indexSum {
				res = []string{word}
				indexSum = i + index
			} else if i + index == indexSum {
				res = append(res, word)
			}
		}
	}
	return res
}
// @lc code=end

