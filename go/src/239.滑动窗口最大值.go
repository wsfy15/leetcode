/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	n := len(nums)
	if n * k == 0 {
		return res
	}
	if k == 1 {
		return nums
	}

	// 存储元素下标的队列 如果比已有元素大，则将已有元素的下标从后边出队，否则从最后入队当前元素下标
	queue := []int{0,}
	for i := 1; i < k - 1; i++ {
		for len(queue) > 0 && nums[i] > nums[queue[len(queue) - 1]] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, i)
	}

	for i := k - 1; i < n; i++ {
		for len(queue) > 0 && nums[i] > nums[queue[len(queue) - 1]] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, i)
		res = append(res, nums[queue[0]])

		// 清除范围外元素
		j := 0
		for ; j < len(queue) - 1; j++ {
			if queue[j] > i - k + 1 {
				break
			}
		}
		queue = queue[j:]
	}

	return res
}
// @lc code=end

