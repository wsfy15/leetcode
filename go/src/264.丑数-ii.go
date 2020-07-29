/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	p2, p3, p5 := 0, 0, 0
	nums := []int{1}
	for i := 1; i < n; i++ {
		num := min(2 * nums[p2], 3 * nums[p3], 5 * nums[p5])
		if num != nums[i - 1] {
			nums = append(nums, num)
		} else {
			i--
		}

		if num == 2 * nums[p2] {
			p2++
		} else if num == 3 * nums[p3] {
			p3++
		} else {
			p5++
		}
	}
	return nums[n - 1]
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

