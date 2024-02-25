/*
 * @lc app=leetcode.cn id=2866 lang=rust
 *
 * [2866] 美丽塔 II
 */

// @lc code=start
impl Solution {
    pub fn maximum_sum_of_heights(max_heights: Vec<i32>) -> i64 {
        let n = max_heights.len();
        let mut stack = vec![n];
        let mut sum = 0;
        let mut suffix = vec![0; n + 1];
        max_heights
            .iter()
            .rev()
            .enumerate()
            .for_each(|(index, &val)| {
                let index = n - index - 1;
                while stack.len() > 1 && val <= max_heights[*stack.last().unwrap()] {
                    let j = stack.pop().unwrap();
                    sum -= max_heights[j] as i64 * (stack.last().unwrap() - j) as i64;
                }

                sum += val as i64 * (stack.last().unwrap() - index) as i64;
                suffix[index] = sum;
                stack.push(index);
            });

        let mut res = sum;
        let mut stack: Vec<i32> = vec![-1];
        let mut pre = 0;
        max_heights.iter().enumerate().for_each(|(index, &val)| {
            while stack.len() > 1 && val <= max_heights[*stack.last().unwrap() as usize] {
                let j = stack.pop().unwrap() as usize;
                pre -= max_heights[j] as i64 * (j as i32 - stack.last().unwrap()) as usize as i64;
            }
            pre += val as i64 * (index as i32 - stack.last().unwrap()) as i64;
            res = res.max(pre + suffix[index + 1]);
            stack.push(index as i32);
        });

        res
    }
}
// @lc code=end
