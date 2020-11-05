/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (60.93%)
 * Likes:    160
 * Dislikes: 0
 * Total Accepted:    58.8K
 * Total Submissions: 95.9K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */

// @lc code=start
func getRow(rowIndex int) []int {
	res := []int{1}

	for i:=1; i<=rowIndex; i++{
		res = append(res, 1)

		for j:=i-1; j>0; j--{
			res[j] += res[j-1]
		}
	}

	return res
}
// @lc code=end
