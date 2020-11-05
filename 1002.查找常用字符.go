/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找常用字符
 *
 * https://leetcode-cn.com/problems/find-common-characters/description/
 *
 * algorithms
 * Easy (68.84%)
 * Likes:    118
 * Dislikes: 0
 * Total Accepted:    20K
 * Total Submissions: 28.1K
 * Testcase Example:  '["bella","label","roller"]'
 *
 * 给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3
 * 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。
 *
 * 你可以按任意顺序返回答案。
 *
 *
 *
 * 示例 1：
 *
 * 输入：["bella","label","roller"]
 * 输出：["e","l","l"]
 *
 *
 * 示例 2：
 *
 * 输入：["cool","lock","cook"]
 * 输出：["c","o"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= A.length <= 100
 * 1 <= A[i].length <= 100
 * A[i][j] 是小写字母
 *
 *
 */

// @lc code=start
func commonChars(A []string) []string {
	var (
		n       = len(A)             // 字符串列表的长度
		inspect = make([][26]int, n) // 每个字符串中每个字符出现的次数
		res     []string             // 返回最终的结果

	)

	// 统计每个字符串中每个字符出现的次数
	for i, str := range A {
		for j := 0; j < len(str); j++ {
			inspect[i][str[j]-'a']++
		}
	}

	// 依次遍历26个字符
	for i := 0; i < 26; i++ {
		var count = math.MaxInt32

		for j := 0; j < n; j++ { // 每个字符的最少出现次数
			count = min(inspect[j][i], count)
		}

		if count != 0 { // !=0代表该字符在每个字符串中都重复出现了
			for c := 0; c < count; c++ { // 出现了几次就往res中添加几次
				res = append(res, string('a'+i))
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
