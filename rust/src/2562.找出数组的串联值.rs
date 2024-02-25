/*
 * @lc app=leetcode.cn id=2562 lang=rust
 *
 * [2562] 找出数组的串联值
 */

// @lc code=start
impl Solution {
    pub fn find_the_array_conc_val(nums: Vec<i32>) -> i64 {
        let n = nums.len();
        let mut res = 0;
        let (mut left, mut right) = (0, n - 1);
        while left < right {
            let mut left_num = nums[left] as i64;
            let mut num = nums[right];
            while num > 0 {
                num /= 10;
                left_num *= 10;
            }
    
            res += left_num + nums[right] as i64;
    
            left += 1;
            right -= 1;
        }
        
        if left == right {
            res += nums[left] as i64;
        }
    
        res
    }
}
// @lc code=end
