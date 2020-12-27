/*
 * @lc app=leetcode.cn id=321 lang=golang
 *
 * [321] 拼接最大数
 */

// @lc code=start
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	// 先得到单个数组保留k个数的最大值，然后合并
	n, m := len(nums1), len(nums2)
	res := make([]int, k)
	for i := 0; i <= k; i++ {
		if n < i || k-i > m {
			continue
		}

		part1 := pickMax(nums1, i)
		part2 := pickMax(nums2, k-i)

		val := merge(part1, part2)
		if sliceBigger(val, res) {
			res = val
		}
	}

	return res
}

func pickMax(nums []int, k int) []int {
	stack := []int{}
	drop := len(nums) - k
	for i := 0; i < len(nums); i++ {
		for drop > 0 && len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			drop--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}

	return stack[:k]
}

func merge(part1, part2 []int) []int {
	i, j := 0, 0
	m, n := len(part1), len(part2)
	res := []int{}
	for i < m && j < n {
		if sliceBigger(part1[i:], part2[j:]) {
			res = append(res, part1[i])
			i++
		} else {
			res = append(res, part2[j])
			j++
		}
	}

	if i < m {
		res = append(res, part1[i:]...)
	} else if j < n {
		res = append(res, part2[j:]...)
	}

	return res
}

func sliceBigger(a, b []int) bool {
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return len(a) > len(b)
}

func min(a ...int) int {
	ans := a[0]
	for _, v := range a {
		if v < ans {
			ans = v
		}
	}
	return ans
}

// @lc code=end

