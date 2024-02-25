/*
 * @lc app=leetcode.cn id=907 lang=rust
 *
 * [907] 子数组的最小值之和
 */

// @lc code=start
impl Solution {
    pub fn sum_subarray_mins(arr: Vec<i32>) -> i32 {
        const MOD: usize = 1e9 as usize + 7;
        let n = arr.len();
        let (mut left, mut right) = (vec![0i32; n], vec![0i32; n]);
        // 找到该元素左右两边 不比 该元素大的元素的下标
        let mut stack = vec![];
        for i in 0..n {
            while stack.len() > 0 && arr[*stack.last().unwrap() as usize] >= arr[i] {
                stack.pop();
            }

            left[i] = if stack.len() == 0 {
                -1
            } else {
                *stack.last().unwrap()
            };

            stack.push(i as i32);
        }

        let mut stack = vec![];
        for i in (0..n).rev() {
            while stack.len() > 0 && arr[*stack.last().unwrap() as usize] > arr[i] {
                stack.pop();
            }

            right[i] = if stack.len() == 0 {
                n as i32
            } else {
                *stack.last().unwrap()
            };

            stack.push(i as i32);
        }

        let mut res = 0;
        for i in 0..n {
            res = (res
                + arr[i] as usize * (i as i32 - left[i]) as usize * (right[i] - i as i32) as usize)
                % MOD;
        }

        res as i32
    }
}
// @lc code=end
