/*
 * @lc app=leetcode.cn id=1094 lang=rust
 *
 * [1094] 拼车
 */

// @lc code=start
impl Solution {
    pub fn car_pooling(trips: Vec<Vec<i32>>, capacity: i32) -> bool {
        let mut nums = vec![0; 1001];
        trips.iter().for_each(|trip| {
            let passengers = trip[0];
            let from = trip[1] as usize;
            let to = trip[2] as usize;
            nums[from] += passengers;
            nums[to] -= passengers;
        });

        let mut cur = 0;
        for num in nums {
            cur += num;
            if cur > capacity {
                return false;
            }
        }

        true
    }
}
// @lc code=end
