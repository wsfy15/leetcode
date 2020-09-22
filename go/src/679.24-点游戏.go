/*
 * @lc app=leetcode.cn id=679 lang=golang
 *
 * [679] 24 点游戏
 */

// @lc code=start
func judgePoint24(nums []int) bool {
	floats := make([]float64, 4)
	for i, v := range nums {
		floats[i] = float64(v)
	}
	return judge(floats)
}

func judge(nums []float64) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return math.Abs(nums[0] - 24.0) < 1e-6;
	}

	n := len(nums)
	// 每次取两个数进行运算
	for i := 0; i < n; i++ {
		// j也要从0开始，因为减法、除法不满足交换律
		for j := 0; j < n; j++ {
			if i != j {
				nums2 := []float64{}
				for k := 0; k < len(nums); k++ {
					// 把剩下的元素放进去
					if k != i && k != j {
						nums2 = append(nums2, nums[k])
					}
				}
	
				for k := 0; k < 4; k++ {
					// 0: +; 1: *; 2: -; 3: /
					if k < 2 && i > j {
						continue
					}

					switch k {
					case 0:
						nums2 = append(nums2, nums[i] + nums[j])
					case 1:
						nums2 = append(nums2, nums[i] * nums[j])
					case 2:
						nums2 = append(nums2, nums[i] - nums[j])
					case 3:
						if nums[j] != 0 { // abs(nums[j]) > 1e-6
							nums2 = append(nums2, nums[i] / nums[j])
						} else {
							continue
						}
					}

					if judge(nums2) {
						return true
					}
					nums2 = nums2[:len(nums2) - 1] // 回溯
				}
			}
		}
	}
	return false
}
// @lc code=end

