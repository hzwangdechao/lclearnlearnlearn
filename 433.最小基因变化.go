/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 *
 * https://leetcode-cn.com/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (52.63%)
 * Likes:    60
 * Dislikes: 0
 * Total Accepted:    9.1K
 * Total Submissions: 17.3K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * 一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
 *
 * 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
 *
 * 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
 *
 * 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
 *
 * 现在给定3个参数 — start, end,
 * bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回
 * -1。
 *
 * 注意:
 *
 *
 * 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
 * 所有的目标基因序列必须是合法的。
 * 假定起始基因序列与目标基因序列是不一样的。
 *
 *
 * 示例 1:
 *
 *
 * start: "AACCGGTT"
 * end:   "AACCGGTA"
 * bank: ["AACCGGTA"]
 *
 * 返回值: 1
 *
 *
 * 示例 2:
 *
 *
 * start: "AACCGGTT"
 * end:   "AAACGGTA"
 * bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
 *
 * 返回值: 2
 *
 *
 * 示例 3:
 *
 *
 * start: "AAAAACCC"
 * end:   "AACCCCCC"
 * bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
 *
 * 返回值: 3
 *
 *
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	if start == "" || end == "" || len(bank) == 0 || !strContains(end, bank) {
		return -1
	}

	L := len(start)
	m := make(map[string][]string)
	for _, word := range bank {
		for i := 0; i < L; i++ {
			key := word[:i] + "*" + word[i+1:]
			m[key] = append(m[key], word)
		}
	}

	Q := make([][]interface{}, 0)
	Q = append(Q, []interface{}{start, 0})

	visited := make(map[string]bool)
	visited[start] = true
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		word := cur[0].(string)
		count := cur[1].(int)
		for i := 0; i < L; i++ {
			key := word[:i] + "*" + word[i+1:]

			if mList, ok := m[key]; ok {
				for _, matchWord := range mList {
					if matchWord == end {
						return count + 1
					}
					if !visited[matchWord] {
						visited[matchWord] = true
						Q = append(Q, []interface{}{matchWord, count + 1})
					}
				}
			}
			delete(m, key)
		}
	}

	return -1
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
