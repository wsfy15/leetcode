/*
 * @lc app=leetcode.cn id=480 lang=golang
 *
 * [480] 滑动窗口中位数
 */

// @lc code=start
func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	slide := make([]int, k)
	copy(slide, nums)
	sort.Ints(slide)

	res := make([]float64, n - k + 1)
	res[0] = getMedian(slide)
	for i := k; i < n; i++ {
		index := sort.SearchInts(slide, nums[i - k])
		slide[index] = nums[i]

		for index < k - 1 && slide[index] > slide[index + 1] {
			slide[index], slide[index + 1] = slide[index + 1], slide[index]
			index++
		}

		for index > 0 && slide[index] < slide[index - 1] {
			slide[index], slide[index - 1] = slide[index - 1], slide[index]
			index--
		}

		res[i - k + 1] = getMedian(slide)
	}

	return res
}

func getMedian(nums []int) float64 {
	return float64(nums[len(nums) / 2] + nums[(len(nums) - 1) / 2]) / 2.0
}
// @lc code=end

