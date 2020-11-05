/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (39.79%)
 * Likes:    363
 * Dislikes: 0
 * Total Accepted:    56.6K
 * Total Submissions: 142K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	return methodWithLayer(matrix)
}

// 模拟移动的方法
func methodWithSimulate(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, cols := len(matrix), len(matrix[0])
	// 构造一个体积和matrix一样的数组用来存储位置上的元素是否被访问过
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	var (
		total    = rows * cols                                                   // 总元素的数量， 可用此来判断是否全部遍历
		res      = make([]int, total)                                            // 返回最终的结果
		row, col = 0, 0                                                          //纵坐标和横坐标
		ds       = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}} // 路径的方向  → ↓ ← ↑
		dsIdx    = 0                                                             // 用来判断朝哪个方向移动
	)

	for i := 0; i < total; i++ {
		num := matrix[row][col]                                // 位置上的元素
		res[i] = num                                           // 将元素放置到res中
		visited[row][col] = true                               // 将坐标更改为已经访问的状态
		nextRow, nextCol := row+ds[dsIdx][0], col+ds[dsIdx][1] // 下一个要访问的坐标
		// 如果下一个要访问的坐标位置超出了界限或者已经访问过了， 修改接下来要移动的方向
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols || visited[nextRow][nextCol] {
			dsIdx = (dsIdx + 1) % 4
		}

		// 改变坐标
		row += ds[dsIdx][0]
		col += ds[dsIdx][1]
	}

	return res
}

// 按层进行遍历
func methodWithLayer(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, cols               = len(matrix), len(matrix[0]) // 行数和列数
		left, right, top, bottom = 0, cols - 1, 0, rows - 1    // 各个顶点的位置
		total                    = rows * cols                 // 元素的个数
		res                      = make([]int, total)          // 最终返回的结果
		index                    = 0                           // res的索引
	)

	for left <= right && bottom >= top {
		// 从左到右开始  (top, left) → (top, right)
		for col := left; col <= right; col++ {
			res[index] = matrix[top][col]
			index++
		}

		// 从上到下  (top+1, right) ↓ (bottom, right)
		for row := top + 1; row <= bottom; row++ {
			res[index] = matrix[row][right]
			index++
		}

		// 防止越界
		if left < right && top < bottom {
			// 从右到左 (bottom, right-1) ← (bottom, left+1)
			for col := right - 1; col >= left+1; col-- {
				res[index] = matrix[bottom][col]
				index++
			}

			// 从下到上  (bottom, left) ↑ (top+1, left)
			for row := bottom; row >= top+1; row-- {
				res[index] = matrix[row][left]
				index++
			}
		}

		// 访问里层元素
		left++
		right--
		top++
		bottom--
	}

	return res
}

// @lc code=end
