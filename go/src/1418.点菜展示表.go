/*
 * @lc app=leetcode.cn id=1418 lang=golang
 *
 * [1418] 点菜展示表
 */

// @lc code=start
type Table struct {
	no    int
	items map[string]int
}

func displayTable(orders [][]string) [][]string {
	food := map[string]bool{}
	tables := []Table{}
	no2table := map[int]int{}
	n := len(orders)
	for i := 0; i < n; i++ {
		foodName := orders[i][2]
		food[foodName] = true

		no, _ := strconv.Atoi(orders[i][1])
		if _, ok := no2table[no]; !ok {
			no2table[no] = len(tables)
			tables = append(tables, Table{
				no:    no,
				items: make(map[string]int),
			})
		}

		index := no2table[no]
		if _, ok := tables[index].items[foodName]; ok {
			tables[index].items[foodName] += 1
		} else {
			tables[index].items[foodName] = 1
		}
	}

	sort.Slice(tables, func(i, j int) bool {
		return tables[i].no < tables[j].no
	})

	foodNameSlice := []string{}
	for k, _ := range food {
		foodNameSlice = append(foodNameSlice, k)
	}
	sort.Slice(foodNameSlice, func(i, j int) bool {
		return foodNameSlice[i] < foodNameSlice[j]
	})

	tableNum := len(tables)
	res := make([][]string, tableNum+1)
	res[0] = append(res[0], "Table")
	res[0] = append(res[0], foodNameSlice...)

	foodNum := len(foodNameSlice)
	for i := 0; i < tableNum; i++ {
		res[i+1] = make([]string, foodNum+1)
		res[i+1][0] = strconv.Itoa(tables[i].no)
		j := 1
		for _, v := range foodNameSlice {
			if cnt, ok := tables[i].items[v]; ok {
				res[i+1][j] = strconv.Itoa(cnt)
			} else {
				res[i+1][j] = "0"
			}
			j++
		}
	}

	return res
}

// @lc code=end

