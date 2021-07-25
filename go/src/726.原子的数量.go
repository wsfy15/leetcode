/*
 * @lc app=leetcode.cn id=726 lang=golang
 *
 * [726] 原子的数量
 */

// @lc code=start
type node struct {
	name  string
	count int
	lp    int
}

func countOfAtoms(formula string) string {
	n := len(formula)
	lp := 0 // 左括号数量
	items := []node{}
	for i := 0; i < n; {
		if formula[i] >= 'A' && formula[i] <= 'Z' {
			name := []byte{formula[i]}
			i++
			for i < n && formula[i] >= 'a' && formula[i] <= 'z' {
				name = append(name, formula[i])
				i++
			}

			count := 0
			for i < n && isDigit(formula[i]) {
				count = count*10 + int(formula[i]-'0')
				i++
			}
			if count == 0 {
				count = 1
			}

			items = append(items, node{string(name), count, lp})
		} else if formula[i] == '(' {
			lp++
			i++
		} else if formula[i] == ')' {
			i++
			count := 0
			for i < n && isDigit(formula[i]) {
				count = count*10 + int(formula[i]-'0')
				i++
			}
			if count == 0 {
				count = 1
			}

			for j := len(items) - 1; j >= 0; j-- {
				if items[j].lp == lp {
					items[j].count *= count
					items[j].lp--
				} else {
					break
				}
			}
			lp--
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].name < items[j].name
	})

	sb := strings.Builder{}
	for i := 0; i < len(items); i++ {
		sb.WriteString(items[i].name)
		count := items[i].count
		for i+1 < len(items) {
			if items[i+1].name == items[i].name {
				count += items[i+1].count
				i++
			} else {
				break
			}
		}
		if count > 1 {
			sb.WriteString(strconv.Itoa(count))
		}
	}
	return sb.String()
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// @lc code=end

