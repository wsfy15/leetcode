/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return nil
	}

	res := []string{}
	for i := 0; i < n; i++ {
		j := i + 1
		for j < n && nums[j] == nums[j - 1] + 1 {
			j++
		}
		if j > i + 1 {
			res = append(res, strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j - 1]))
		} else {
			res = append(res, strconv.Itoa(nums[i]))
		}
		i = j - 1
	}

	return res
}
// @lc code=end

