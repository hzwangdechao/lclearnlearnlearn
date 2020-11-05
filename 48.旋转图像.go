/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 *
 * https://leetcode-cn.com/problems/rotate-image/description/
 *
 * algorithms
 * Medium (67.99%)
 * Likes:    469
 * Dislikes: 0
 * Total Accepted:    77.2K
 * Total Submissions: 112.7K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个 n × n 的二维矩阵表示一个图像。
 *
 * 将图像顺时针旋转 90 度。
 *
 * 说明：
 *
 * 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
 *
 * 示例 1:
 *
 * 给定 matrix =
 * [
 * ⁠ [1,2,3],
 * ⁠ [4,5,6],
 * ⁠ [7,8,9]
 * ],
 *
 * 原地旋转输入矩阵，使其变为:
 * [
 * ⁠ [7,4,1],
 * ⁠ [8,5,2],
 * ⁠ [9,6,3]
 * ]
 *
 *
 * 示例 2:
 *
 * 给定 matrix =
 * [
 * ⁠ [ 5, 1, 9,11],
 * ⁠ [ 2, 4, 8,10],
 * ⁠ [13, 3, 6, 7],
 * ⁠ [15,14,12,16]
 * ],
 *
 * 原地旋转输入矩阵，使其变为:
 * [
 * ⁠ [15,13, 2, 5],
 * ⁠ [14, 3, 4, 1],
 * ⁠ [12, 6, 8, 9],
 * ⁠ [16, 7,10,11]
 * ]
 *
 *
 */

// @lc code=start
func rotate(matrix [][]int)  {
	method1(matrix)


	// length := len(matrix)

	// // 对同一列上的元素进行替换
	// for i :=0; i< length; i++{
	// 	for j:=0; j<length/2; j++{
	// 		matrix[i][j], matrix[i][length-1-j] = matrix[i][length-1-j], matrix[i][j]
	// 	}
	// }

	// // 对角线替换
	// for i := 0; i < length -1; i++{
	// 	for j :=0; j < length -i -1; j++{
	// 		matrix[i][j], matrix[length-j-1][length-i-1] = matrix[length-j-1][length-i-1], matrix[i][j]
	// 	}
	// }

}


// 先转置， 在翻转
func method1(matrix [][]int)  {

	/*  先将图像进行转置 (第n行第n个元素和第n列第n个元素进行交换)
			[
				[1,4,7],
				[2,5,8],
				[3,6,9],
			]
		*/
	for row:=0; row<len(matrix); row++{
		for col:=row; col<len(matrix); col++{
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	/*
		对每一行上的元素的头和尾进行翻转
		[
			[7,4,1],
			[8,5,2],
			[9,6,3],
		]
	*/
	for row:=0; row<len(matrix); row++{
		for i:=0; i<len(matrix)/2; i++{
			matrix[row][i], matrix[row][len(matrix)-1-i] = matrix[row][len(matrix)-1-i], matrix[row][i]
		}
	}

}

// @lc code=end
