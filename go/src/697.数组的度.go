/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	counter := map[int]int{}
	max := 0
	for _, v := range nums {
		counter[v]++
		if counter[v] > max {
			max = counter[v]
		}
	}

	res := math.MaxInt32
	for k, v := range counter {
		if v == max {
			start, end := 0, len(nums) - 1
			for nums[start] != k {
				start++
			}
			for nums[end] != k {
				end--
			}
			res = min(end - start + 1, res)
		}
	}
	return res
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
// @lc code=end

