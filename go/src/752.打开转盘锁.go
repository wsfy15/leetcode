/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	t, _ := strconv.Atoi(target)
	if t == 0 {
		return 0
	}

	dead, m := map[int]struct{}{}, map[int]struct{}{}
	for _, d := range deadends {
		num, _ := strconv.Atoi(d)
		dead[num] = struct{}{}
	}

	if _, ok := dead[0]; ok {
		return -1
	}

	queue := []int{0}
	round := 0
	for round < 20 {
		round++
		newQueue := []int{}
		for _, num := range queue {
			for _, delta := range []int{1, 10, 100, 1000} {
				for _, add := range []int{-1, 1} {
					digit := num / delta % (delta * 10)
					digit += add
					if digit < 0 {
						digit += 10
					}
					digit %= 10

					newNum := num/(delta*10)*(delta*10) + digit*delta + num%delta
					if _, ok := m[newNum]; ok {
						continue
					}
					if _, ok := dead[newNum]; ok {
						continue
					}

					m[newNum] = struct{}{}
					if newNum == t {
						return round
					}
					newQueue = append(newQueue, newNum)
				}
			}
		}

		queue = newQueue
	}

	return -1
}

// @lc code=end

