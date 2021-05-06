/*
 * @lc app=leetcode.cn id=1169 lang=golang
 *
 * [1169] 查询无效交易
 */

// @lc code=start
func invalidTransactions(transactions []string) []string {
	n := len(transactions)
	items := make([][]string, n)
	var res []string
	times := make([]int, n)
	amounts := make([]int, n)

	for i := 0; i < n; i++ {
		items[i] = strings.Split(transactions[i], ",")
		amount, _ := strconv.Atoi(items[i][2])
		time, _ := strconv.Atoi(items[i][1])
		times[i] = time
		amounts[i] = amount
	}

	for i := 0; i < n; i++ {
		if amounts[i] > 1000 {
			res = append(res, transactions[i])
			continue
		}

		for j := 0; j < n; j++ {
			if j == i {
				continue
			}

			if items[j][0] == items[i][0] && items[i][3] != items[j][3] && abs(times[i]-times[j]) <= 60 {
				res = append(res, transactions[i])
				break
			}
		}
	}

	return res
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}

// @lc code=end

