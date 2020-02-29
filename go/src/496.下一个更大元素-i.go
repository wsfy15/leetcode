/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var m = map[int]int{}
  res := make([]int, len(nums1))

	var stack []int
	for _, num := range nums2 {
		m[num] = -1
		for len(stack) > 0 && num > stack[len(stack) - 1] {
			m[stack[len(stack) - 1]] = num
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, num)
	}

	for i, num := range nums1 {
		res[i] = m[num]
	}
	return res
}
// @lc code=end

