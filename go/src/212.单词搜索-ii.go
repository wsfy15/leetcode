/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
type Trie struct {
	children [26]*Trie
	isEnd bool
	ch byte
}

var set map[string]bool
var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func findWords(board [][]byte, words []string) []string {
	set = make(map[string]bool)
	if len(words) == 0 {
		return nil
	}

	m := len(board)
	if m == 0 {
		return nil
	}

	root := &Trie{}
	for _, word := range words {
		cur := root
		for i := range word {
			if cur.children[word[i] - 'a'] == nil {
				cur.children[word[i] - 'a'] = &Trie{
					ch: word[i],
				}
			}
			cur = cur.children[word[i] - 'a']
		}
		cur.isEnd = true
	}

	for i := range board {
		for j := range board[i] {
			if root.children[board[i][j] - 'a'] != nil {
				dfs(board, root, i, j, "")
			}
		}
	}

	var res []string
	for key := range set {
		res = append(res, key)
	}
	return res
}

func dfs(board [][]byte, cur *Trie, i, j int, str string) {
	str += string(board[i][j])
	cur = cur.children[board[i][j] - 'a']

	if cur.isEnd {
		set[str] = true
	}

	var tmp byte
	tmp, board[i][j] = board[i][j], '@' // 因为不允许重复，所以先置为@
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if 0 <= x && x < len(board) && 0 <= y && y < len(board[0]) && 
			board[x][y] != '@' && cur.children[board[x][y] - 'a'] != nil {
				dfs(board, cur, x, y, str)
			}
	}

	board[i][j] = tmp
}
// @lc code=end

