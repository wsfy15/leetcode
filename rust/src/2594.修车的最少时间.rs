/*
 * @lc app=leetcode.cn id=2594 lang=rust
 *
 * [2594] 修车的最少时间
 */

// @lc code=start
impl Solution {
    pub fn repair_cars(ranks: Vec<i32>, cars: i32) -> i64 {
        let best = ranks.iter().min().unwrap();
        let cars = cars as i64;
        let best = *best as i64;
        let (mut left, mut right) = (0 as i64, cars * cars * best);
        while left < right {
            let mid = left + (right - left) / 2;
            let mut count = 0 as i64;
            ranks.iter().for_each(|r| {
                count += ((mid / (*r as i64)) as f64).sqrt() as i64;
            });

            if count >= cars {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        right as i64
    }
}
// @lc code=end
