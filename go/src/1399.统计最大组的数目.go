/*
 * @lc app=leetcode.cn id=1399 lang=golang
 *
 * [1399] 统计最大组的数目
 */

// @lc code=start
func countLargestGroup(n int) int {
	counter := map[int]int{}
	for i := 1; i <= n; i++ {
		index := digits(i)
		counter[index]++
	}

	res, max := 0, 0
	for _, v := range counter {
		if v > max {
			res = 1
			max = v
		} else if v == max {
			res++
		}
	}

	return res
}

func digits(n int) int {
	res := 0
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}

// @lc code=end

