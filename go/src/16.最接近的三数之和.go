/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	cur := nums[0] + nums[1] + nums[2]

	for i := 0; i < n - 2; i++ {
		j, k := i + 1, n - 1
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			if abs(tmp - target) < abs(cur - target) {
				cur = tmp
			}

			if tmp > target {
				k--
			} else if tmp < target {
				j++
			} else {
				return target
			}
		}
	}

	return cur
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}
// @lc code=end

