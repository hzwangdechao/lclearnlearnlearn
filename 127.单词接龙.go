/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (43.64%)
 * Likes:    447
 * Dislikes: 0
 * Total Accepted:    61.5K
 * Total Submissions: 140.8K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord
 * 的最短转换序列的长度。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回 0。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出: 5
 *
 * 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * ⁠    返回它的长度 5。
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: 0
 *
 * 解释: endWord "cog" 不在字典中，所以无法进行转换。
 *
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	return methodWithIndividualBfs(beginWord, endWord, wordList)
}

// 双向广度优先搜索
func methodWithBilateralBfs(beginWord string, endWord string, wordList []string) int {
	if beginWord == "" || endWord == "" || len(wordList) == 0 || !isStrContain(endWord, wordList) {
		return 0
	}
	L := len(beginWord)
	m := make(map[string][]string) // 用来记录h*t 对应能匹配到的字符串
	for _, word := range wordList {
		for i := 0; i < L; i++ {
			key := word[:i] + "*" + word[i+1:] // 生成word所有可能的模糊匹配符放到m对应的列表中
			m[key] = append(m[key], word)
		}
	}
	QueueBegin := make([][]interface{}, 0)
	QueueBegin = append(QueueBegin, []interface{}{beginWord, 1})
	QueueEnd := make([][]interface{}, 0)
	QueueEnd = append(QueueEnd, []interface{}{endWord, 1})
	ans := 0
	visitedBegin := make(map[string]int)
	visitedBegin[beginWord] = 1
	visitedEnd := make(map[string]int)
	visitedEnd[endWord] = 1

	visitWordNode := func(queue [][]interface{}, visited, otherVisited map[string]int) int {
		current := queue[0]
		queue = queue[1:]

		curWord := current[0].(string)
		curLevel := current[1].(int)

		for i := 0; i < L; i++ {
			key := curWord[:i] + "*" + curWord[i+1:]

			if matchWordList, ok := m[key]; ok {
				for _, matchWord := range matchWordList {
					if otherLevel, ok := otherVisited[matchWord]; ok {
						return curLevel + otherLevel
					}

					if _, ok := visited[matchWord]; !ok {
						visited[matchWord] = curLevel + 1
						queue = append(queue, []interface{}{matchWord, curLevel + 1})
					}
				}
			}
		}

		return -1
	}

	for len(QueueBegin) > 0 && len(QueueEnd) > 0 {
		ans = visitWordNode(QueueBegin, visitedBegin, visitedEnd)
		if ans != -1 {
			return ans
		}
		ans = visitWordNode(QueueEnd, visitedEnd, visitedBegin)
		if ans != -1 {
			return ans
		}
	}

	return 0

}

func methodWithIndividualBfs(beginWord string, endWord string, wordList []string) int {
	if beginWord == "" || endWord == "" || len(wordList) == 0 || !isStrContain(endWord, wordList) {
		return 0
	}
	L := len(beginWord)
	m := make(map[string][]string) // 用来记录h*t 对应能匹配到的字符串
	for _, word := range wordList {
		for i := 0; i < L; i++ {
			key := word[:i] + "*" + word[i+1:] // 生成word所有可能的模糊匹配符放到m对应的列表中
			m[key] = append(m[key], word)
		}
	}

	Q := make([][]interface{}, 0)
	Q = append(Q, []interface{}{beginWord, 1})

	visited := make(map[string]bool) // 记录字符串是否之前访问过
	visited[beginWord] = true

	for len(Q) > 0 {
		current := Q[0]
		Q = Q[1:]

		word := current[0].(string)
		level := current[1].(int)

		for i := 0; i < L; i++ {
			key := word[:i] + "*" + word[i+1:]

			if mList, ok := m[key]; ok {
				for _, matchWord := range mList {
					if matchWord == endWord {
						return level + 1
					}

					if !visited[matchWord] {
						visited[matchWord] = true
						Q = append(Q, []interface{}{matchWord, level + 1})
					}
				}
			}
			// 防止重复遍历
			delete(m, key)
		}
	}
	return 0

}

// 判断wordList中是否包含endWord
func isStrContain(endWord string, wordList []string) bool {
	for _, word := range wordList {
		if word == endWord {
			return true
		}
	}
	return false
}

// @lc code=end
