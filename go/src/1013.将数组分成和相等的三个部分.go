/*
 * @lc app=leetcode.cn id=1013 lang=golang
 *
 * [1013] 将数组分成和相等的三个部分
 */

// @lc code=start
func canThreePartsEqualSum(A []int) bool {
	total := 0
	for _, v := range A {
		total += v
	}
	if total % 3 != 0 {
		return false
	}

	avg := total / 3
	cur := 0
	count := 0
	for i, v := range A {
		cur += v
		if cur == avg {
			cur = 0
			count++
			if count == 2 && i < len(A) - 1 { // 后续还有元素能组成avg
				return true
			}
		}
	}
	return false
}
// @lc code=end

