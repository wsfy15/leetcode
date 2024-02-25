/*
 * @lc app=leetcode.cn id=1631 lang=rust
 *
 * [1631] 最小体力消耗路径
 */

// @lc code=start
impl Solution {
    pub fn minimum_effort_path(heights: Vec<Vec<i32>>) -> i32 {
        let (n, m) = (heights.len(), heights[0].len());
        let (mut left, mut right) = (0, 1000 * 1000);

        fn abs(num: i32) -> i32 {
            if num < 0 {
                return -num;
            }
            num
        }

        fn bfs(mut visited: Vec<Vec<bool>>, heights: &Vec<Vec<i32>>, limit: i32) -> bool {
            let dir = vec![(0, 1), (0, -1), (1, 0), (-1, 0)];

            let mut queue = vec![(0, 0)];
            while queue.len() > 0 {
                let node = queue.remove(0);
                if node.0 == heights.len() - 1 && node.1 == heights[0].len() - 1 {
                    return true;
                }

                for &d in &dir {
                    let (x, y) = (node.0 as i32 + d.0, node.1 as i32 + d.1);
                    if 0 <= x
                        && x < heights.len() as i32
                        && 0 <= y
                        && y < heights[0].len() as i32
                        && !visited[x as usize][y as usize]
                        && abs(heights[x as usize][y as usize] - heights[node.0][node.1]) <= limit
                    {
                        visited[x as usize][y as usize] = true;
                        queue.push((x as usize, y as usize));
                    }
                }
            }

            false
        }

        while left < right {
            let mid = left + (right - left) / 2;
            let mut visited = vec![vec![false; m]; n];
            visited[0][0] = true;
            if bfs(visited, &heights, mid) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }
        left
    }
}
// @lc code=end
