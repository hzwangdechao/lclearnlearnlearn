/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 *
 * https://leetcode-cn.com/problems/word-ladder-ii/description/
 *
 * algorithms
 * Hard (38.60%)
 * Likes:    355
 * Dislikes: 0
 * Total Accepted:    25.2K
 * Total Submissions: 65.4K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord
 * 的最短转换序列。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换后得到的单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回一个空列表。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出:
 * [
 * ⁠ ["hit","hot","dot","dog","cog"],
 * ["hit","hot","lot","log","cog"]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: []
 *
 * 解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
 *
 */

// @lc code=start
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if beginWord == "" || endWord == "" || len(wordList) == 0 || !strContains(endWord, wordList) {
		return [][]string{}
	}

	var (
		res     [][]string
		L       = len(beginWord)
		m       = make(map[string][]string)
		visited = make(map[string]bool)
		Q       = make([][]interface{}, 0)
	)

	for _, word := range wordList {
		for i := 0; i < L; i++ {
			key := word[:i] + "*" + word[i+1:]
			m[key] = append(m[key], word)
		}
	}
	Q = append(Q, []interface{}{beginWord, []string{beginWord}})
	visited[beginWord] = true

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		curWord := cur[0].(string)
		curList := cur[1].([]string)

		for i := 0; i < L; i++ {
			key := curWord[:i] + "*" + curWord[i+1:]

			if words, ok := m[key]; ok {
				for _, word := range words {
					if word == endWord{
						res = append(res, append(curList, word))
						break
					}
					if !visited[word]{
						visited[word] = true

						Q = append(Q, []interface{}{word, append(curList, word)})
					}
				}
			}
			delete(m, key)
		}
	}

	return res
}
func strContains(word string, wordList []string) bool {
	for _, w := range wordList {
		if w == word {
			return true
		}
	}
	return false
}

// @lc code=end
