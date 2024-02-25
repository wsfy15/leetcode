/*
 * @lc app=leetcode.cn id=1782 lang=rust
 *
 * [1782] 统计点对的数目
 */

// @lc code=start
use std::collections::HashMap;

impl Solution {
    // 与点对 a,b 相连的边数等于「与 a 相连的边数加上与 b 相连的边数，再减去同时与 a,b 同时 相连的边数」。
    // 与点 i 相连的边数为 deg[i]，通过一次遍历 edges 数组求出 deg 数组的取值。
    // 对于任意的点对，可以轻松地求出 deg[a]+deg[b] 的值

    // 与 a,b 同时相连的边数为 overlap[a][b],采用哈希表计算，而不是二维数组
    // key 为 a*(n+1)+b (a>b)
    pub fn count_pairs(n: i32, edges: Vec<Vec<i32>>, queries: Vec<i32>) -> Vec<i32> {
        let mut deg = vec![0; n as usize + 1];
        let mut overlap = HashMap::new();
        edges.into_iter().for_each(|edge| {
            let (a,b) = (edge[0], edge[1]);
            deg[a as usize] = deg[a as usize] + 1;
            deg[b as usize] = deg[b as usize] + 1;
    
            let index = a.max(b)*(n+1) + a.min(b);
            let cnt = overlap.get(&index).unwrap_or(&0);
            overlap.insert(index, cnt+1);
        });
    
        let mut sorted_deg = deg.clone();
        sorted_deg.sort();
        // println!("not sorted: {:?}", deg);
        // println!("sorted: {:?}", sorted_deg);
    
        let mut res = vec![];
        // sorted_deg[a]+sorted_deg[b] > query
        for query in queries {
            let mut cnt = 0;
            let  (mut l, mut r) = (1, n);
            while l <= n {
                while l < r && sorted_deg[l as usize] + sorted_deg[r as usize] > query {
                    r -= 1;
                }
                cnt += n - l.max(r);
                l += 1;
            }
    
            for (k, val) in overlap.iter() {
                let (a, b) = (k / (n+1), k % (n+1));
                // println!("({},{}): {}", a, b, val);
                if deg[a as usize] + deg[b as usize] > query && deg[a as usize] + deg[b as usize]  - val <= query {
                    // 求的是点对的数量，因此这里-1
                    cnt -= 1;
                    // println!("decrease {}", val);
                }
            }
    
            res.push(cnt);
        }
    
        res
    }
}
// @lc code=end

