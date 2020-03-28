/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
// 本题dp 比 dfs+贪心 慢

// dfs + 贪心：优先选择大面值的，不行再减少大面值的
var ans int
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	// 降序排序
	sort.Slice(coins, func(i, j int) bool {
    return coins[i] > coins[j]
	})

	ans = math.MaxInt32
	util(amount, 0, 0, coins)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func util(amount, count, index int, coins []int) {
	if amount == 0 {
		if count < ans {
			ans = count
		}
		return
	}

	if index == len(coins) {
		return
	}

	for i := amount / coins[index]; i >= 0 && i + count < ans; i-- {
		util(amount - i * coins[index], count + i, index + 1, coins)
	}
}


func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Ints(coins)
	states := make([]int, amount + 1)
	ans := dp(amount, coins, states) 
	return ans
}

// amount: 剩余待找零钱数
func dp(amount int, coins, states []int) int {
	if states[amount] != 0 {
		return states[amount]
	}
	if amount == 0 {
		return 0
	}

	min := math.MaxInt32
	for i := len(coins) - 1; i >= 0 ; i-- {
		if amount - coins[i] >= 0 {
			 ans := dp(amount - coins[i], coins, states) 
			 if ans >= 0 && ans < min {
				 min = ans + 1 // ans个硬币构成ammount-coins[i]，1个硬币构成coins[i]
			 }
		}
	}
	if min == math.MaxInt32 {
		states[amount] = -1
	} else {
		states[amount] = min
	}
	return states[amount]
}

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	states := make([]bool, amount + 1)
	states[0] = true
	i := 0
	for {
		i++
		change := false
		for j := amount; j >= 0; j-- {
			if states[j] {
				for k := range coins {
					if j + coins[k] <= amount && !states[j + coins[k]] {
						change = true
						states[j + coins[k]] = true
					}
				}
			}
		}
		if states[amount] {
			return i
		} 
		if !change {
			return -1
		}
	}
}
// @lc code=end

