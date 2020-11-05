/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (66.36%)
 * Likes:    327
 * Dislikes: 0
 * Total Accepted:    88.8K
 * Total Submissions: 133.2K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1] 0 1   1  4    2  6   3
 * [1,5,10,10,5,1]]
 * ]
 *
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := make([][]int, 0)

	for i:=0; i<numRows;i++{
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1

		for j:=1; j<i; j++{
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, row)
	}

	return res

}


// @lc code=end
