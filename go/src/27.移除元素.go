/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
  j := 0 
  for _, v := range nums {
    if v != val {
      nums[j] = v
      j++
    } 
  }
  return j
}

func removeElement2(nums []int, val int) int {
  i, j := 0, len(nums) - 1 
  if i == j && nums[0] == val {
    return 0
  }
  
  for i <= j {
    for i <= j && nums[i] != val {
      i++
    }
    
    for i <= j && nums[j] == val {
      j--
    }
    
    if i < j {
        nums[i], nums[j] = nums[j], nums[i]
      i++
      j--
    }
  }
  return j + 1
}
// @lc code=end

