/*
 * @lc app=leetcode.cn id=929 lang=golang
 *
 * [929] 独特的电子邮件地址
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	exists := map[string]bool{}
	for _, email := range emails {
		parts := strings.Split(email, "@")
		local, domain := parts[0], parts[1]

		sb := strings.Builder{}
		for i := 0; i < len(local); i++ {
			if local[i] ==  '+' {
				break
			}
			if local[i] == '.' {
				continue
			}
			sb.WriteByte(local[i])
		}

		sb.WriteByte('@')
		sb.WriteString(domain)
		exists[sb.String()] = true
	}

	return len(exists)
}
// @lc code=end

