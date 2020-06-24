/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	outDegree := make([][]int, numCourses)
	for _, prerequisity := range prerequisites {
		inDegree[prerequisity[0]]++
		outDegree[prerequisity[1]] = append(outDegree[prerequisity[1]], prerequisity[0])
	}

	res, queue := []int{}, []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			inDegree[i]--
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)

		for _, v := range outDegree[cur] {
			if inDegree[v]--; inDegree[v] == 0 {
				inDegree[v] = -1
				queue = append(queue, v)
			}
		}
	}
	
	if len(res) != numCourses {
		return nil
	}
	return res
}
// @lc code=end

