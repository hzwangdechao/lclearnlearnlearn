/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有K个重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/description/
 *
 * algorithms
 * Medium (44.53%)
 * Likes:    223
 * Dislikes: 0
 * Total Accepted:    14.4K
 * Total Submissions: 32.2K
 * Testcase Example:  '"aaabb"\n3'
 *
 * 找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。
 *
 * 示例 1:
 *
 *
 * 输入:
 * s = "aaabb", k = 3
 *
 * 输出:
 * 3
 *
 * 最长子串为 "aaa" ，其中 'a' 重复了 3 次。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * s = "ababbc", k = 2
 *
 * 输出:
 * 5
 *
 * 最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
 *
 *
 */

// @lc code=start


/*
先用hash表统计s中每个字符出现的次数，显然如果字符c出现的次数小于k，c必然不在最长子串里面，
根据这个特性可以将原始s分割成多个子串递归地求解问题，
我们用一个split数组依次来存放每个分割点的索引，对每个分割区间同样求解该问题(多路的分治问题)，
并取结果的最大值保存在变量ans中，此处有一个小trick（如果当前求解的子串长度比已存在的ans还要小
，则没有必要求解该区间，这样可以减少不必要的计算），最后递归的结束点就是当前求解的字符串s符合最长子串的要求

*/

//  *********************    以下方法运行超时了    ***************************
func longestSubstring_wrong(s string, k int) int {
	n := len(s)
	dp := make([][]int, n)
	res := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			repeatNum := stringRepeatNum(s, i, j)

			//fmt.Println([]int{i, j}, repeatNum)
			dp[i][j] = repeatNum
			if repeatNum >= k {
				res = max(res, j-i+1)
			}
		}
	}
	// for _, v := range dp{
	// 	fmt.Println(v)
	// }
	return res
}

// 检查字符串开始位置到结束位置 重复字符的最少出现次数
func stringRepeatNum(s string, start, end int) int {
	res := math.MaxInt32
	m := make(map[byte]int)

	for i := start; i <= end; i++ {
		m[s[i]] ++
	}

	for _, v := range m {
		res = min(res, v)
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y{
		return x
	}
	return y
}
// @lc code=end
