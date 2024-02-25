/*
 * @lc app=leetcode.cn id=1686 lang=rust
 *
 * [1686] 石子游戏 VI
 */

// @lc code=start
impl Solution {
    pub fn stone_game_vi(alice_values: Vec<i32>, bob_values: Vec<i32>) -> i32 {
        let mut pair = alice_values.into_iter().zip(bob_values).collect::<Vec<_>>();
        pair.sort_unstable_by(|(px, py), (qx, qy)| (qx + qy).cmp(&(px + py)));
        pair.iter()
            .enumerate()
            .map(|(i, &(x, y))| if i % 2 == 0 { x } else { -y })
            .sum::<i32>()
            .signum()
    }
}
// @lc code=end
