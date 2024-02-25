/*
 * @lc app=leetcode.cn id=260 lang=rust
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
impl Solution {
    pub fn single_number(nums: Vec<i32>) -> Vec<i32> {
        let mut xor = 0;
        nums.iter().for_each(|num| {
            xor ^= num;
        });

        // 两个不同元素异或，肯定有某一位是1，即在这一位上，一个数是1，另一个数是0
        // 找到最低位的1
        let flag = xor & (-xor);

        let (mut a, mut b) = (0, 0);
        nums.iter().for_each(|num| {
            if num & flag != 0 {
                a ^= num;
            } else {
                b ^= num;
            }
        });

        vec![a, b]
    }
}
// @lc code=end
