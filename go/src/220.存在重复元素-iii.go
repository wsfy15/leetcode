/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// 维护一个大小为k+1的滑动窗口
	// 使用哈希表存储每个t范围内的数，[0, t], [t+1, 2t+1]
	if len(nums) < 2 || t < 0 {
		return false
	}

	w := t + 1 // 以免t为0，每个桶大小为t+1也可以满足桶内任意两元素差最大值为t
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if i > k {
			target := nums[i-k-1] / w
			delete(counter, target) // 移除窗口外元素
		}

		target := divide(nums[i], w)
		if _, ok := counter[target]; ok {
			return true
		} else {
			if v, ok := counter[target-1]; ok && nums[i]-v <= t {
				return true
			}
			if v, ok := counter[target+1]; ok && v-nums[i] <= t {
				return true
			}
			counter[target] = nums[i]
		}
	}

	return false
}

// -3 / 5 = 0， 3 / 5 = 0，导致[-4, -1] 与 [1, 4]这两个范围映射到同一个桶
// 因此让负数除完后-1
func divide(x, y int) int {
	if x < 0 {
		x = -x
		if x%y == 0 {
			return -(x / y)
		}
		return -1 - x/y
	}
	return x / y
}

// @lc code=end

