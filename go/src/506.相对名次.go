/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */

// @lc code=start
type rank struct {
	num int
	index int
}

type ranks struct {
	data []rank
}

func (r ranks) Len() int {
	return len(r.data)
}

func (r ranks) Less(i, j int) bool {
	return r.data[i].num > r.data[j].num
}

func (r ranks) Swap(i, j int) {
	r.data[i], r.data[j] = r.data[j], r.data[i]
}

func findRelativeRanks(nums []int) []string {
	r := ranks{}
	for i, num := range nums {
		r.data = append(r.data, rank{
			num:   num,
			index: i,
		})
	}

	sort.Sort(r)
	res := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		switch i {
		case 0:
			res[r.data[0].index] = "Gold Medal"
		case 1:
			res[r.data[1].index] = "Silver Medal"
		case 2:
			res[r.data[2].index] = "Bronze Medal"
		default:
			res[r.data[i].index] = strconv.Itoa(i + 1)
		}
	}

	return res
}
// @lc code=end

