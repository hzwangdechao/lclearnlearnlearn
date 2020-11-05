/*
 * @lc app=leetcode.cn id=1288 lang=golang
 *
 * [1288] 删除被覆盖区间
 *
 * https://leetcode-cn.com/problems/remove-covered-intervals/description/
 *
 * algorithms
 * Medium (54.47%)
 * Likes:    14
 * Dislikes: 0
 * Total Accepted:    3.4K
 * Total Submissions: 6.2K
 * Testcase Example:  '[[1,4],[3,6],[2,8]]'
 *
 * 给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
 *
 * 只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。
 *
 * 在完成所有删除操作后，请你返回列表中剩余区间的数目。
 *
 *
 *
 * 示例：
 *
 *
 * 输入：intervals = [[1,4],[3,6],[2,8]]
 * 输出：2
 * 解释：区间 [3,6] 被区间 [2,8] 覆盖，所以它被删除了。
 *
 *
 *
 *
 * 提示：​​​​​​
 *
 *
 * 1 <= intervals.length <= 1000
 * 0 <= intervals[i][0] < intervals[i][1] <= 10^5
 * 对于所有的 i != j：intervals[i] != intervals[j]
 *
 *
 */

// @lc code=start
func removeCoveredIntervals(intervals [][]int) int {
	// 如果区间列表长度为0的话直接返回结果
	if len(intervals) == 0 {
		return 0
	}

	// 对区间列表的起点进行升序排序, 如果起点相同的话， 按照区间的终点进行降序排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res := 0

	// 初始化区间的起点和终点
	start, end := intervals[0][0], intervals[0][1]

	// 从第二个位置开始遍历区间列表
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]  // 当前的区间

		// 当前区间能被上一个区间覆盖的情况
		if cur[0] >= start && cur[1] <= end{
			res ++
		}

		// 当前区间能和上一个区间合并成一个更大区间的情况
		if cur[0] <= end && cur[1] >= end{
			end = cur[1]
		}

		// 当前区间和上一个区间没有相交部分的情况， 更新区间的起点和终点
		if cur[0] > end{
			start = cur[0]
			end = cur[1]
		}
	}

	return len(intervals) - res
}


// @lc code=end
