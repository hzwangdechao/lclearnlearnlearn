/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 *
 * https://leetcode-cn.com/problems/word-break-ii/description/
 *
 * algorithms
 * Hard (38.47%)
 * Likes:    214
 * Dislikes: 0
 * Total Accepted:    19.7K
 * Total Submissions: 51.3K
 * Testcase Example:  '"catsanddog"\n["cat","cats","and","sand","dog"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典
 * wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
 *
 * 说明：
 *
 *
 * 分隔时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 *
 *
 * 示例 1：
 *
 * 输入:
 * s = "catsanddog"
 * wordDict = ["cat", "cats", "and", "sand", "dog"]
 * 输出:
 * [
 * "cats and dog",
 * "cat sand dog"
 * ]
 *
 *
 * 示例 2：
 *
 * 输入:
 * s = "pineapplepenapple"
 * wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
 * 输出:
 * [
 * "pine apple pen apple",
 * "pineapple pen apple",
 * "pine applepen apple"
 * ]
 * 解释: 注意你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入:
 * s = "catsandog"
 * wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出:
 * []
 *
 *
 */

// @lc code=start
func wordBreak(s string, wordDict []string) []string {

	// 首先利用动态规划判断是否可以满足单词拆分的条件
	m := make(map[string]bool)
	for _, word := range wordDict{
		m[word] = true
	}
	n:=len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i:=1; i<=n; i++{
		for j:=0; j<i; j++{
			if dp[j] && m[s[j:i]]{
				dp[i] = true
				break
			}
		}
	}
	res := []string{}
	if !dp[n]{
		return res
	}

	// 回溯算法
	var DFS = func (string,[]string) {}
    DFS = func(ns string,r []string) {
		fmt.Println(ns, r)
        if len(ns) == 0 {
            res = append(res, strings.Join(r," "))
            return
        }
        for i:=1; i<=len(ns); i++ {
			fmt.Println(i)
            if m[ns[:i]] {
                DFS(ns[i:],append(r,ns[:i]))
            }
        }
    }
    DFS(s,[]string{})

	return res

}

// @lc code=end
