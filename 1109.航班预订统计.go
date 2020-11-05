/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 *
 * https://leetcode-cn.com/problems/corporate-flight-bookings/description/
 *
 * algorithms
 * Medium (46.10%)
 * Likes:    84
 * Dislikes: 0
 * Total Accepted:    13.6K
 * Total Submissions: 29K
 * Testcase Example:  '[[1,2,10],[2,3,20],[2,5,25]]\n5'
 *
 * 这里有 n 个航班，它们分别从 1 到 n 进行编号。
 *
 * 我们这儿有一份航班预订表，表中第 i 条预订记录 bookings[i] = [i, j, k] 意味着我们在从 i 到 j 的每个航班上预订了 k
 * 个座位。
 *
 * 请你返回一个长度为 n 的数组 answer，按航班编号顺序返回每个航班上预订的座位数。
 *
 *
 *
 * 示例：
 *
 * 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
 * 输出：[10,55,45,25,25]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= bookings.length <= 20000
 * 1 <= bookings[i][0] <= bookings[i][1] <= n <= 20000
 * 1 <= bookings[i][2] <= 10000
 *
 *
 */

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	// 构造差分数组
	diff := make([]int, n)

	for _, booking := range bookings{
		i := booking[0] -1
		j := booking[1] -1
		z := booking[2]
		diff[i] += z
		if j + 1 < n{
			diff[j+1] -= z
		}
	}

	fmt.Println(diff)

	// 从差分数组回复到结果数组
	res := make([]int, n)
	res[0] = diff[0]

	for i:=1; i<n; i++{
		res[i] = res[i-1] + diff[i]
	}

	return res
}
// @lc code=end
