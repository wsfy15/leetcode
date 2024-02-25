/*
 * @lc app=leetcode.cn id=2251 lang=rust
 *
 * [2251] 花期内花的数目
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn full_bloom_flowers(flowers: Vec<Vec<i32>>, people: Vec<i32>) -> Vec<i32> {
        let mut times = HashMap::new();
        flowers.iter().for_each(|flower| {
            *times.entry(flower[0]).or_insert(0) += 1;
            *times.entry(flower[1] + 1).or_insert(0) -= 1;
        });

        // 得到花开花落时间点，升序
        let n = times.len();
        let mut order_times = Vec::with_capacity(n);
        times.iter().for_each(|(&t, _)| order_times.push(t));
        order_times.sort_unstable();

        // 对每个人按到达时间升序排序
        let m = people.len();
        let mut order_people: Vec<_> = (0..m).collect();
        order_people.sort_by(|&i, &j| people[i].cmp(&people[j]));

        let mut ans = vec![0; m];
        let (mut sum, mut index) = (0, 0);
        for p in order_people {
            while index < n && people[p] >= order_times[index] {
                sum += times.get(&order_times[index]).unwrap();
                index += 1;
            }
            ans[p] = sum;
        }

        ans
    }
}
// @lc code=end
