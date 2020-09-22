/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 */

// @lc code=start
func findItinerary(tickets [][]string) []string {
	if len(tickets) == 0 {
		return nil
	}

	edges := map[string][]string{}
	for _, ticket := range tickets {
		start, end := ticket[0], ticket[1]
		edges[start] = append(edges[start], end)
	}

	for k, _ := range edges {
		sort.Strings(edges[k])
	}

	res := []string{}
	dfs(edges, "JFK", &res)

	return res
}

func dfs(edges map[string][]string, pos string, res *[]string) {
	for len(edges[pos]) > 0 {
		cur := edges[pos][0]
		edges[pos] = edges[pos][1:]
		dfs(edges, cur, res)
	}

	*res = append([]string{pos}, *res...)
}

// @lc code=end

