/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	// 通过状态机的思路
	// 每一个位有三种状态，即0(00)、1(01)、2(10)、0的循环，每当一个数该位为1时，就进入下一个状态
	// 对于出现三次的元素，最后每个位还是0
	// 而对于只出现一次的元素，只有该数为1的位保持1
	// 既然每一位需要两位来表示,那就用两个int来表示
	a, b := 0, 0 // a表示低位,b表示高位
	for _, num := range nums {
		a = a ^ num & (^b)
		b = b ^ num & (^a)
	}

	return a
}
// @lc code=end

