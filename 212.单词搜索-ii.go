/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 *
 * https://leetcode-cn.com/problems/word-search-ii/description/
 *
 * algorithms
 * Hard (42.21%)
 * Likes:    213
 * Dislikes: 0
 * Total Accepted:    19.2K
 * Total Submissions: 45.5K
 * Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n' +
  '["oath","pea","eat","rain"]'
 *
 * 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
 *
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
 *
 * 示例:
 *
 * 输入:
 * words = ["oath","pea","eat","rain"] and board =
 * [
 * ⁠ ['o','a','a','n'],
 * ⁠ ['e','t','a','e'],
 * ⁠ ['i','h','k','r'],
 * ⁠ ['i','f','l','v']
 * ]
 *
 * 输出: ["eat","oath"]
 *
 * 说明:
 * 你可以假设所有输入都由小写字母 a-z 组成。
 *
 * 提示:
 *
 *
 * 你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
 * 如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？
 * 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。
 *
 *
*/

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	var res []string

	for _, word := range words {
		if exist(board, word) {
			res = append(res, word)
		}
	}

	return res
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	var (
		rows, cols = len(board), len(board[0])
	)
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == word[0] {
				visited[row][col] = true
				if existDFS(board, visited, row, col, word[1:]) {
					return true
				}
				visited[row][col] = false
			}
		}
	}

	return false
}

func existDFS(board [][]byte, visited [][]bool, row, col int, word string) bool {
	if len(word) == 0 {
		return true
	}

	ds := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var r, c int
	for _, d := range ds {
		r, c = row+d[0], col+d[1]
		if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) && !visited[r][c] && word[0] == board[r][c] {
			visited[r][c] = true
			if existDFS(board, visited, r, c, word[1:]) {
				return true
			}
			visited[r][c] = false
		}
	}

	return false
}

// @lc code=end
