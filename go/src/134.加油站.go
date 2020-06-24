/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}

	diff := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		diff[i] = gas[i] - cost[i]
	}

	cur, index, minGas := 0, 0, 0
	for i := 0; i < len(diff); i++ { // 记录遍历过程中最低的汽油剩余
		cur += diff[i]
		if cur < minGas {
			minGas = cur
			index = i
		}
	}

	if cur < 0 { // 不能跑完全程
		return -1
	}
	if minGas == 0 { // 最低剩余汽油没变
		return 0
	}
	return (index + 1) % len(diff)
}

func canCompleteCircuit2(gas []int, cost []int) int {
	if sum(gas) < sum(cost) {
		return -1
	}

	diff := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		diff[i] = gas[i] - cost[i]
	}

	for i := 0; i < len(gas); i++ {
		if diff[i] >= 0 && canComplete(diff, i) {
			return i
		}
	}

	return -1
}

func canComplete(diff []int, start int) bool {
	total := diff[start]
	for i := start + 1; i < len(diff); i++ {
		total += diff[i]
		if total < 0 {
			return false
		}
	}

	for i := 0; i < start; i++ {
		total += diff[i]
		if total < 0 {
			return false
		}
	}

	return true
}

func sum(a []int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}
// @lc code=end

