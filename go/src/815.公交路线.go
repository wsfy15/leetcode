/*
 * @lc app=leetcode.cn id=815 lang=golang
 *
 * [815] 公交路线
 */

// @lc code=start
func numBusesToDestination(routes [][]int, source int, target int) int {
	if target == source {
		return 0
	}

	// 每个车站可以搭乘的公交车
	buses := map[int][]int{}
	for i, route := range routes {
		for _, st := range route {
			buses[st] = append(buses[st], i)
		}
	}

	if _, ok := buses[source]; !ok {
		return -1
	}
	if _, ok := buses[target]; !ok {
		return -1
	}

	step := 1
	used := map[int]bool{} // 指定公交车是否被乘坐过
	queue := buses[source] // step对应的公交车
	for len(queue) > 0 {
		newQueue := []int{}
		for _, bus := range queue {
			if used[bus] {
				continue
			}
			used[bus] = true

			for _, station := range routes[bus] {
				if station == target {
					return step
				}

				for _, anotherBus := range buses[station] {
					if used[anotherBus] {
						continue
					}

					newQueue = append(newQueue, anotherBus)
				}
			}
		}

		step++
		queue = newQueue
	}

	return -1
}

// @lc code=end

