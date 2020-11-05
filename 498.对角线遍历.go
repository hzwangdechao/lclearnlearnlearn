/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 *
 * https://leetcode-cn.com/problems/diagonal-traverse/description/
 *
 * algorithms
 * Medium (40.50%)
 * Likes:    94
 * Dislikes: 0
 * Total Accepted:    15.8K
 * Total Submissions: 39K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
 *
 *
 *
 * 示例:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 *
 * 输出:  [1,2,4,7,5,3,6,8,9]
 *
 * 解释:
 *
 *
 *
 *
 *
 * 说明:
 *
 *
 * 给定矩阵中的元素总数不会超过 100000 。
 *
 *
 */

// @lc code=start
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0{
		return nil
	}
	rows, cols := len(matrix), len(matrix[0])

	res := []int{} // 存储最后的结果
	line := []int{} // 存储每一条对角线上的结果

	// 遍历每一条对角线
	for i:=0; i<cols+rows-1; i++{
		// 清空对角线数组
		line = line[:0]

		// 确定对角线起点的横坐标和纵坐标
		var row, col int
		if i<cols{
			row = 0
			col = i
		}else{
			row = i - cols + 1
			col = cols - 1
		}

		// 进行对角线遍历
		for row<rows && col > -1{
			line = append(line, matrix[row][col])
			 row ++
			 col --
		}

		// 如果为偶数对角线则对数组进行翻转
		if i % 2 == 0{
			line = reverseSlice(line)
		}

		res = append(res, line...)

	}
	return res


}


// 翻转数组
func reverseSlice(nums []int) []int {
	for i,j:=0,len(nums)-1; i<j; i, j = i+1, j-1{
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
// @lc code=end
