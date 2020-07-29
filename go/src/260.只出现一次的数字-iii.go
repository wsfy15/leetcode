/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	// 两个不同元素异或，肯定有某一位是1，即在这一位上，一个数是1，另一个数是0
	// 找到最低位的1
	flag := xor & (-xor)

	a, b := 0, 0
	for _, num := range nums {
		if num & flag > 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
// @lc code=end

