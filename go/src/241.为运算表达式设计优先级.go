/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
func diffWaysToCompute(input string) []int {
	n := len(input)
	if n == 0 {
		return nil
	}
	funcs := map[byte]func(a, b int)int{
		'+': func(a, b int) int { return a + b },
		'-': func(a, b int) int { return a - b },
		'*': func(a, b int) int { return a * b },
	}

	res := []int{}
	for i := 0; i < n; i++ {
		if fun, ok := funcs[input[i]]; ok {
			res1 := diffWaysToCompute(input[:i])
			res2 := diffWaysToCompute(input[i+1:])
			for _, num1 := range res1 {
				for _, num2 := range res2 {
					res = append(res, fun(num1, num2))
				}
			}
		}
	}

	if len(res) == 0 { // 说明是个数字
		num, _ := strconv.Atoi(input)
		return []int{num}
	} 
	return res
}
// @lc code=end

