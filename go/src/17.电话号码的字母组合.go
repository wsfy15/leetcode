/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	
	num2digits := map[byte][]byte{}
	var i, index byte = 0, '2'
	for i < 26 {
		num2digits[index] = []byte{i + 'a', i + 'a' + 1, i + 'a' + 2}
		if index == '7' || index == '9' {
			num2digits[index] = append(num2digits[index], i + 'a' + 3)
			i++
		}
		i += 3
		index++
	}

	return generate(num2digits, 0, len(digits) - 1, digits)
}

func generate(num2digits map[byte][]byte, index, end int, digits string) []string {
	if index > end {
		return []string{""}
	}

	res := []string{}
	partial := generate(num2digits, index + 1, end, digits)
	for _, v := range partial {
		for _, char := range num2digits[digits[index]] {
			res = append(res, string(char) + v)
		}
	}
	return res
}
// @lc code=end

