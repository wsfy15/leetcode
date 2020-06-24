/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start
func getHint(secret string, guess string) string {
	nums := make([]int, 10)
	n, bulls, cows := len(secret), 0, 0

	for i := 0; i < n; i++ {
		nums[secret[i] - '0']++
	}

	for i := 0; i < n; i++ {
		if guess[i] == secret[i] {
			bulls++
			nums[guess[i] - '0']--
		}
	}

	for i := 0; i < n; i++ {
		if guess[i] != secret[i] && nums[guess[i] - '0'] > 0 {
			cows++
			nums[guess[i] - '0']--
		}
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
// @lc code=end

