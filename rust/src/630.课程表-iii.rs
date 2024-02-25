/*
 * @lc app=leetcode.cn id=630 lang=rust
 *
 * [630] 课程表 III
 */

// @lc code=start
use std::collections::BinaryHeap;

impl Solution {
    pub fn schedule_course(mut courses: Vec<Vec<i32>>) -> i32 {
        // 根据结束时间升序排序
        courses.sort_unstable_by_key(|course| course[1]);

        let mut heap = BinaryHeap::new();
        let mut total = 0;
        for course in courses.iter() {
            let (duration, lastday) = (course[0], course[1]);
            heap.push(duration);
            total += duration;
            if total > lastday {
                total -= heap.pop().unwrap();
            }
        }

        heap.len() as i32
    }
}
// @lc code=end
