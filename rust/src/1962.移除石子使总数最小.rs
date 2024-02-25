/*
 * @lc app=leetcode.cn id=1962 lang=rust
 *
 * [1962] 移除石子使总数最小
 */

// @lc code=start
impl Solution {
    pub fn min_stone_sum(piles: Vec<i32>, k: i32) -> i32 {
        let mut heap = std::collections::BinaryHeap::from(piles);
        for _ in 0..k {
            let top = heap.pop().unwrap();
            heap.push((top + 1) / 2);
            if *heap.peek().unwrap() == 0 {
                break;
            }
        }

        heap.iter().sum()
    }
}
// @lc code=end
