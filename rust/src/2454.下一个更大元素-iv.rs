/*
 * @lc app=leetcode.cn id=2454 lang=rust
 *
 * [2454] 下一个更大元素 IV
 */

// @lc code=start
impl Solution {
    pub fn second_greater_element(nums: Vec<i32>) -> Vec<i32> {
        let mut ans = vec![-1; nums.len()];
        let mut s = Vec::new(); // 单调递减栈
        let mut t = Vec::new(); // 存储已经发现一个更大元素的 元素下标
        for (i, &x) in nums.iter().enumerate() {
            while !t.is_empty() && nums[*t.last().unwrap()] < x {
                ans[t.pop().unwrap()] = x; // t 栈顶的下下个更大元素是 x
            }
            let mut j = s.len();
            while j > 0 && nums[s[j - 1]] < x {
                j -= 1; // s 栈顶的下一个更大元素是 x
            }
            t.extend(s.drain(j..)); // 把从 s 弹出的这一整段元素加到 t
            s.push(i); // 当前元素（的下标）加到 s 栈顶
        }
        ans
    }
}
// @lc code=end
