/*
 * @lc app=leetcode.cn id=2300 lang=rust
 *
 * [2300] 咒语和药水的成功对数
 */

// @lc code=start
impl Solution {
    pub fn successful_pairs(spells: Vec<i32>, potions: Vec<i32>, success: i64) -> Vec<i32> {
        let n = spells.len();
        let mut res = vec![0; n];
        let mut potions = potions;
        potions.sort_unstable();

        for i in 0..n {
            res[i] = Self::binary_find(spells[i], success, &potions);
        }

        res
    }
    fn binary_find(spell: i32, success: i64, potions: &Vec<i32>) -> i32 {
        let (mut left, mut right) = (0, potions.len());
        while left < right {
            let mid = left + (right - left) / 2;
            if potions[mid] as i64 * spell as i64 >= success {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        (potions.len() - left) as i32
    }
}
// @lc code=end
