/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (42.70%)
 * Likes:    518
 * Dislikes: 0
 * Total Accepted:    120.7K
 * Total Submissions: 282.7K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */

// @lc code=start

func merge(intervals [][]int) [][]int {
	var res [][]int

	// 如果区间列表为空的情况， 直接返回结果
	if len(intervals) == 0 {
		return res
	}

	// 对区间列表的起点进行升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		cur := intervals[i]
		// 如果res为空或者当前区间的起点比res最后一个区间的终点大的话，直接添加到res中
		if len(res) == 0 || cur[0] > res[len(res)-1][1] {
			res = append(res, cur)
		}

		// 如果当前区间的起点比上一个区间的终点小的话， 对区间进行合并
		if cur[0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(cur[1], res[len(res)-1][1])
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
