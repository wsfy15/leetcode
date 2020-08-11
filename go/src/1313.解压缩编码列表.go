/*
 * @lc app=leetcode.cn id=1313 lang=golang
 *
 * [1313] 解压缩编码列表
 */

// @lc code=start
func decompressRLElist(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}

	return res
}

// @lc code=end

