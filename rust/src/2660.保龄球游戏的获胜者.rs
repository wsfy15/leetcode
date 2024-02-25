/*
 * @lc app=leetcode.cn id=2660 lang=rust
 *
 * [2660] 保龄球游戏的获胜者
 */

// @lc code=start
impl Solution {
    pub fn is_winner(player1: Vec<i32>, player2: Vec<i32>) -> i32 {
        fn get_score(points: &Vec<i32>) -> i32 {
            let mut score = 0;
            let mut double_to = -1;
            for i in 0..points.len() {
                score += points[i];
                if double_to >= i as i32 {
                    score += points[i];
                }

                if points[i] == 10 {
                    double_to = i as i32 + 2;
                }
            }
            score
        }
        let (score1, score2) = (get_score(&player1), get_score(&player2));
        if score1 > score2 {
            1
        } else if score1 < score2 {
            2
        } else {
            0
        }
    }
}
// @lc code=end
