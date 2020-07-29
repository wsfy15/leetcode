/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果交换
 */

// @lc code=start
func fairCandySwap(A []int, B []int) []int {
	counterA, counterB := map[int]bool{}, map[int]bool{}
	sumA, sumB := 0, 0
	for _, num := range A {
		counterA[num] = true
		sumA += num
	}

	for _, num := range B {
		counterB[num] = true
		sumB += num
	}

	tmp := (sumB - sumA) / 2
	for k, _ := range counterA {
		if _, ok := counterB[k + tmp]; ok {
			return []int{k, k + tmp}
		}
	}
	return nil
}
// @lc code=end

