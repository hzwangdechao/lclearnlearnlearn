/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (62.68%)
 * Likes:    409
 * Dislikes: 0
 * Total Accepted:    92.2K
 * Total Submissions: 146.6K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 *
 * 示例:
 *
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * 说明：
 *
 *
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 *
 *
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	hash := make(map[[26]int][]string)

	// 计算出字符串的字母列表， 并放到hash中
	for _, str := range strs{
		b := [26]int{}
		for _, word := range str{
			b[word-'a'] += 1
		}
		hash[b] = append(hash[b], str)
	}

	for _, v :=range hash{
		res = append(res, v)
	}

	return res

}
// @lc code=end
