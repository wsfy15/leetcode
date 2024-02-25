/*
 * @lc app=leetcode.cn id=823 lang=rust
 *
 * [823] 带因子的二叉树
 */

// @lc code=start
impl Solution {
    pub fn num_factored_binary_trees(arr: Vec<i32>) -> i32 {
        let mut arr: Vec<_> = arr.iter().map(|&x| x as i64).collect();
        arr.sort();
        static MOD: i64 = 1_000_000_007;
        let mut res = 0;
        let n = arr.len();
        let mut dp: Vec<i64> = vec![0; n];

        // 以arr[i]为根节点的二叉树数量
        for i in 0..n {
            dp[i] = 1;
            let (mut left, mut right) = (0, i as i32 - 1);
            while left <= right {
                while left <= right && arr[left as usize] * arr[right as usize] > arr[i] {
                    right -= 1;
                }

                if left <= right && arr[left as usize] * arr[right as usize] == arr[i] {
                    if arr[left as usize] != arr[right as usize] {
                        dp[i] = (dp[i] + 2 * dp[left as usize] * dp[right as usize]) % MOD;
                    } else {
                        dp[i] = (dp[i] + dp[left as usize] * dp[right as usize]) % MOD;
                    }
                }
                left += 1;
            }

            res = (res + dp[i]) % MOD;
        }

        res as i32
    }
}
// @lc code=end
