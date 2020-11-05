/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 *
 * https://leetcode-cn.com/problems/partition-labels/description/
 *
 * algorithms
 * Medium (72.55%)
 * Likes:    299
 * Dislikes: 0
 * Total Accepted:    28.8K
 * Total Submissions: 38.5K
 * Testcase Example:  '"ababcbacadefegdehijhklij"'
 *
 * 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段的长度的列表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：S = "ababcbacadefegdehijhklij"
 * 输出：[9,7,8]
 * 解释：
 * 划分结果为 "ababcbaca", "defegde", "hijhklij"。
 * 每个字母最多出现在一个片段中。
 * 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
 *
 *
 *
 *
 * 提示：
 *
 *
 * S的长度在[1, 500]之间。
 * S只包含小写字母 'a' 到 'z' 。
 *
 *
 */

// @lc code=start
func partitionLabels(S string) []int {
	l := make([][2]int, 26)
	m := make(map[byte]int) // 记录每个字符出现的次数
	for i := 0; i < len(S); i++ {
		last := l[S[i]-'a']
		start, end := last[0], last[1]
		if m[S[i]] == 0 {
			start = i
			end = i
		} else {
			end = i
		}
		m[S[i]]++
		l[S[i]-'a'] = [2]int{start, end}

	}

	// fmt.Println(l)
	// fmt.Println(m)

	var intervals [][2]int
	for k, _ := range m {
		intervals = append(intervals, l[k-'a'])
	}

	// fmt.Println(intervals)

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] == intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	// fmt.Println(intervals)

	// 接下来进行合并区间
	var merges [][2]int
	for i := 0; i < len(intervals); i++ {
		cur := intervals[i]
		// 如果res为空或者res的结束位置比当前的开始位置小的话
		if len(merges) == 0 || merges[len(merges)-1][1] < cur[0] {
			merges = append(merges, intervals[i])
		}

		// 如果当前区间的开始位置比上一个区间的
		if cur[0] <= merges[len(merges)-1][1] {
			merges[len(merges)-1][1] = max(merges[len(merges)-1][1], cur[1])
		}
	}

	var res []int

	for _, m := range merges {
		res = append(res, m[1]-m[0]+1)
	}

	// fmt.Println(res)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
