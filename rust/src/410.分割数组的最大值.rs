/*
 * @lc app=leetcode.cn id=410 lang=rust
 *
 * [410] 分割数组的最大值
 */

// @lc code=start
impl Solution {
    pub fn split_array(nums: Vec<i32>, k: i32) -> i32 {
        let total: i32 = nums.iter().sum::<i32>();
        let max = *nums.iter().max().unwrap();

        // 最大值的最小值 介于 最大元素 和 nums和 之间，通过二分查找
        let (mut low, mut high) = (max, total);
        while low < high {
            let mid = low + (high - low) / 2;
            let (mut tmp, mut count) = (0, 1);

            for &num in &nums {
                if tmp + num > mid {
                    count += 1;
                    tmp = num;
                    if count > k {
                        break;
                    }
                } else {
                    tmp += num;
                }
            }
            if count > k {
                low = mid + 1;
            } else {
                high = mid;
            }
        }

        low as i32
    }
}
// @lc code=end
