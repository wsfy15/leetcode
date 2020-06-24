/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	strs := []string{}
	count := 0
	for _, num := range nums {
		if num == 0 {
			count++
		}
		strs = append(strs, strconv.Itoa(num))
	}

	if count == len(nums) {
		return "0"
	}

	sort.Slice(strs, func(i, j int) bool {
		num1, _ := strconv.Atoi(strs[i] + strs[j])
		num2, _ := strconv.Atoi(strs[j] + strs[i])
		return num1 >= num2
	})
	
	var res string
	for _, str := range strs {
		res += str
	}
	return res
}
// @lc code=end

