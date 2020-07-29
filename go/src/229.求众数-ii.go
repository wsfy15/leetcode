/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */

// @lc code=start
// 摩尔投票法
// 抵消阶段 + 计数阶段
// 抵消阶段找到出现次数前二的元素，计数阶段判断这两个元素出现次数是否大于3/n
func majorityElement(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	x, y, xcount, ycount := nums[0], 0, nums[0], 0
	for _, num := range nums {
		if x == num {
			xcount++
			continue
		}
		if y == num {
			ycount++
			continue
		}

		if xcount == 0 {
			x = num
			xcount = 1
			continue
		}

		if ycount == 0 {
			y = num
			ycount = 1
			continue
		}

		xcount--
		ycount--
	}

	xcount, ycount = 0, 0
	for _, num := range nums {
		if num == x {
			xcount++
		} else if num == y {
			ycount++
		}
	}

	res := []int{}
	if xcount > n / 3 {
		res = append(res, x)
	}
	if ycount > n / 3 {
		res = append(res, y)
	}
	return res
}
// @lc code=end

