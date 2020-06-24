/*
 * @lc app=leetcode.cn id=690 lang=golang
 *
 * [690] 员工的重要性
 */

// @lc code=start
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	m := map[int]*Employee{}
	for _, v := range employees {
		m[v.Id] = v
	}

	res := 0
	queue := []int{id}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		res += m[cur].Importance
		for _, v := range m[cur].Subordinates {
			queue = append(queue, v)
		}
	}

	return res
}
// @lc code=end

