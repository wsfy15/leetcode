/*
 * @lc app=leetcode.cn id=1423 lang=rust
 *
 * [1423] 可获得的最大点数
 */

// @lc code=start
impl Solution {
    pub fn max_score(card_points: Vec<i32>, k: i32) -> i32 {
        let k = k as usize;
        let n = card_points.len();
        // 把首尾各k个元素 构造一个[n-k, n-1] + [0, k-1]
        let mut nums = vec![0; 2 * k as usize];
        for i in 0..k {
            nums[i] = card_points[n - k + i];
        }
        for i in k..k * 2 {
            nums[i] = card_points[i - k];
        }
        let mut res = 0;
        let mut cur = 0i64;
        for i in 0..k - 1 {
            cur += nums[i] as i64;
        }

        for i in (k - 1)..(k * 2) {
            cur += nums[i] as i64;
            res = res.max(cur as i32);
            cur -= nums[i + 1 - k] as i64;
        }

        res
    }
}
// @lc code=end
