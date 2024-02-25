/*
 * @lc app=leetcode.cn id=2530 lang=rust
 *
 * [2530] 执行 K 次操作后的最大分数
 */

// @lc code=start
use std::collections::BinaryHeap;
use std::iter::FromIterator;

impl Solution {
    pub fn max_kelements(nums: Vec<i32>, k: i32) -> i64 {
        let mut heap = BinaryHeap::from_iter(nums);
        (0..k).fold(0, |mut acc, _| {
            let mut max = heap.peek_mut().unwrap();
            acc += *max as i64;
            *max = (*max + 2) / 3;
            acc
        })
    }
}
// @lc code=end
