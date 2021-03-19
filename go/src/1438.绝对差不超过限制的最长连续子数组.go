/*
 * @lc app=leetcode.cn id=1438 lang=golang
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 */

// @lc code=start
func longestSubarray(nums []int, limit int) int {
	res, left, right, n := 0, 0, 0, len(nums)
	minQueue := []int{} // 存储最小值下标的单调队列, 队头元素最大
	maxQueue := []int{} // 存储最大值下标的单调队列, 队头元素最小

	for right < n {
		for len(maxQueue) > 0 && nums[right] > nums[maxQueue[len(maxQueue) - 1]] {
			maxQueue = maxQueue[:len(maxQueue) - 1]
		}
		maxQueue = append(maxQueue, right)

		for len(minQueue) > 0 && nums[right] < nums[minQueue[len(minQueue) - 1]] {
			minQueue = minQueue[:len(minQueue) - 1]
		}
		minQueue = append(minQueue, right)

		for len(maxQueue) > 0 && len(minQueue) > 0 && nums[maxQueue[0]] - nums[minQueue[0]] > limit {
			if maxQueue[0] <= left {
				maxQueue = maxQueue[1:]
			}
			if minQueue[0] <= left {
				minQueue = minQueue[1:]
			}
			left++
		}

		res = max(res, right - left + 1)
		right++
	}

	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
// @lc code=end

