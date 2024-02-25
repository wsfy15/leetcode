/*
 * @lc app=leetcode.cn id=1444 lang=rust
 *
 * [1444] 切披萨的方案数
 */

// @lc code=start
impl Solution {
    pub fn ways(pizza: Vec<String>, k: i32) -> i32 {
        const MOD: i32 = 1e9 as i32 + 7;
    
        let k = k as usize;
        let pizza: Vec<_> = pizza.into_iter().map(|v|v.into_bytes()).collect();
        let (row, col) = (pizza.len(), pizza[0].len());
    
        // apples[i][j] 表示以[i][j]为左上角,[row-1][col-1]为右下角的矩形内苹果数量
        let mut apples = vec![vec![0; col+1]; row+1];
    
        // dp[i][j][k] 表示以[i][j]为左上角，[row-1][col-1]为右下角的pizza 切成k块的方案数
        let mut dp =vec![vec![vec![0; k+1]; col+1]; row+1];
        (0..row).rev().for_each(|i| (0..col).rev().for_each(|j| {
            apples[i][j] = apples[i+1][j]+apples[i][j+1]-apples[i+1][j+1]+ if pizza[i][j]==b'A'{1}else{0};
            dp[i][j][1] = if apples[i][j] > 0 { 1 } else { 0 };
        }));
    
        (2..=k).for_each(|l| (0..row).for_each(|i| (0..col).for_each(|j| {
            // 横着切
            for ii in i+1..row {
                if apples[i][j] > apples[ii][j] {
                    dp[i][j][l] = (dp[i][j][l] + dp[ii][j][l-1]) % MOD;
                }
            }
            // 竖着切
            for jj in j+1..col {
                if apples[i][j] > apples[i][jj] {
                    dp[i][j][l] = (dp[i][j][l] + dp[i][jj][l-1]) % MOD;
                }
            }
        })));
    
        return dp[0][0][k];
    }
}
// @lc code=end

