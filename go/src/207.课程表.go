/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	depends := make([][]int, numCourses)
	unfinish := map[int]int{}
	for i := 0; i < len(prerequisites); i++ {
		depends[prerequisites[i][1]] = append(depends[prerequisites[i][1]], prerequisites[i][0])
		unfinish[prerequisites[i][0]]++
	}

	if len(unfinish) == numCourses {
		return false
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if _, ok := unfinish[i]; !ok {
			queue = append(queue, i)
		}
	}

	finishNum := len(queue)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		for _, v := range depends[course] {
			if unfinish[v]--; unfinish[v] == 0 {
				queue = append(queue, v)
				finishNum++
			}
		}
		
		if finishNum == numCourses {
			return true
		}
	}

	return finishNum == numCourses
}
// @lc code=end

