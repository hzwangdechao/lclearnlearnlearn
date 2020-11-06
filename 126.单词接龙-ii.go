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
	// 用字符串在wordList中的位置来表示
	ids := make(map[string]int)
	for i, word := range wordList {
		ids[word] = i
	}
	// 将beginWord也添加到wordList中
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = len(wordList) - 1
	}
	// 如果endWord没有在wordList中的话直接返回
	if _, ok := ids[endWord]; !ok {
		return [][]string{}
	}
	n := len(wordList)
	edges := make([][]int, n)
	// 计算字符串相邻的字符串， 用字符串在wordList中索引表示
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	cost := make([]int, n) // cost[i]表示从beginWord转换到wordList[i]所需要转换的次数
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0

	paths := [][]int{} /// 存储转换的路径
	Q := make([][]int, 0)
	Q = append(Q, []int{ids[beginWord]}) // 队列中存放的是beginWord --> endWord的变化路径

	for i := 0; i < len(Q); i++ {
		now := Q[i]

		last := now[len(now)-1] // 遍历的终止条件是 last==ids[endWord] , 说明已经从beginWord转换到了endWord,
		if last == ids[endWord] {
			tmp := make([]int, len(now))
			copy(tmp, now)
			paths = append(paths, tmp)
		} else {
			// 对最后一个转换的字符串进行处理
			for _, to := range edges[last] {
				if cost[last]+1 <= cost[to] {
					cost[to] = cost[last] + 1
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)
					Q = append(Q, tmp)
				}

			}
		}

	}
	var res [][]string
	for _, path := range paths {
		tmp := make([]string, 0)
		for _, to := range path {
			tmp = append(tmp, wordList[to])
		}
		res = append(res, tmp)
	}

	return res
}

// 判断两个字符串是否可以通过改变一个字符后变成相同的字符串
func transformCheck(word1 string, word2 string) bool {
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			// true [dot dit] [loc log]
			// false [dob dac]
			return word1[i+1:] == word2[i+1:]
		}
	}
	// false [abc abc]
	return false
}

// @lc code=end
