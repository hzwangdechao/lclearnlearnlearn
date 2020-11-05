/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 *
 * https://leetcode-cn.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (37.81%)
 * Likes:    211
 * Dislikes: 0
 * Total Accepted:    31.8K
 * Total Submissions: 83.9K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * 给出一个无重叠的 ，按照区间起始端点排序的区间列表。
 *
 * 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
 *
 *
 *
 * 示例 1：
 *
 * 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 输出：[[1,5],[6,9]]
 *
 *
 * 示例 2：
 *
 * 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 输出：[[1,2],[3,10],[12,16]]
 * 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 *
 *
 *
 *
 * 注意：输入类型已在 2019 年 4 月 15 日更改。请重置为默认代码定义以获取新的方法签名。
 *
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	// 从左向右找到 left    newInterval[0]<=intervals[i][1]  比较 left=min(newInterval[0], intervals[i][0])
	// 从右向左找到 right  newInterval[1]>=intervals[i][0]  比较 right=max(newInterval[1], intervals[i][1])
	// 如果没有找到right说明不需要合并区间
	// 如果找到了left和right  返回intervals[:start] + [left, right] + intervals[end+1:]

	var (
		start = 0
		end   = len(intervals) - 1
		left  = math.MinInt32
		right = math.MinInt32
	)

	for start <= end {
		if newInterval[0] <= intervals[start][1] {
			left = min(newInterval[0], intervals[start][0])
			break
		} else {
			start++
		}
	}
	if left == math.MinInt32 {
		intervals = append(intervals, newInterval)
		return intervals
	}

	for end >= start {
		if newInterval[1] >= intervals[end][0] {
			right = max(newInterval[1], intervals[end][1])
			break
		} else {
			end--
		}
	}
	fmt.Println(start, end)
	fmt.Println(left, right)
	// 有left但是没有right
	/* 这种情况
	[[3,5],[12,15]]
	[6,6]
	*/
	if right == math.MinInt32 {
		tmp := make([][]int, 0)
		tmp = append(tmp, intervals[:start]...)
		tmp = append(tmp, newInterval)
		tmp = append(tmp, intervals[start:]...)
		return tmp
	}
	return append(append(intervals[:start], []int{left, right}), intervals[end+1:]...)

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
