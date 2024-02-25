/*
 * @lc app=leetcode.cn id=1690 lang=rust
 *
 * [1690] 石子游戏 VII
 */

// @lc code=start
impl Solution {
    pub fn stone_game_vii(stones: Vec<i32>) -> i32 {
        let n = stones.len();
        let mut s = vec![0; n + 1];
        for (i, &x) in stones.iter().enumerate() {
            s[i + 1] = s[i] + x;
        }
        let mut f = vec![0; n];
        for i in (0..n - 1).rev() {
            for j in i + 1..n {
                f[j] = (s[j + 1] - s[i + 1] - f[j]).max(s[j] - s[i] - f[j - 1]);
            }
        }
        f[n - 1]
    }
}
// @lc code=end
