/*
 * @lc app=leetcode.cn id=842 lang=golang
 *
 * [842] 将数组拆分成斐波那契序列
 */

// @lc code=start
func splitIntoFibonacci(S string) []int {
	var res []int
	if dfs(S, 0, 0, 0, 0, &res) {
		return res
	}
	return nil
}

func dfs(S string, index, pre1, pre2, deep int, cur *[]int) bool {
	if index >= len(S) {
		return deep >= 3
	}

	for i := 1; i <= 11; i++ {
		if index+i > len(S) || (i > 1 && S[index] == '0') {
			break
		}

		num, _ := strconv.Atoi(S[index : index+i])
		if num > math.MaxInt32 || (deep > 1 && num > pre1+pre2) {
			break
		}

		if deep <= 1 || pre1+pre2 == num {
			*cur = append(*cur, num)
			if dfs(S, index+i, pre2, num, deep+1, cur) {
				return true
			}
			*cur = (*cur)[:len(*cur)-1]
		}
	}

	return false
}

// @lc code=end

