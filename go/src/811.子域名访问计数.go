/*
 * @lc app=leetcode.cn id=811 lang=golang
 *
 * [811] 子域名访问计数
 */

// @lc code=start
func subdomainVisits(cpdomains []string) []string {
	m := map[string]int{}
	for _, domain := range cpdomains {
		parts := strings.Split(domain, " ")
		count, _ := strconv.Atoi(parts[0])
		domain = parts[1]
		for {
			m[domain] += count
			parts := strings.SplitN(domain, ".", 2)
			if len(parts) != 2 {
				break
			}
			domain = parts[1]
		}
	}

	res := []string{}
	for k, v := range m {
		res = append(res, strconv.Itoa(v) + " " + k)
	}
	return res
}
// @lc code=end

