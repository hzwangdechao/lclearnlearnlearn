/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (78.21%)
 * Likes:    247
 * Dislikes: 0
 * Total Accepted:    49.2K
 * Total Submissions: 62.8K
 * Testcase Example:  '3'
 *
 * 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	return methodWithLayer(n)
}

// 按层移动的方法  先对外层坐标进行赋值， 再对内层元素进行赋值
func methodWithLayer(n int) [][]int {
	var (
		rows, cols               = n, n                     // 行数和列数
		left, right, bottom, top = 0, cols - 1, rows - 1, 0 // 各个顶点的位置
		res   = make([][]int, rows)
		index = 1
	)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, cols)
	}

	for left <= right && top <= bottom {
		// 从左到右  (top, left) > (top, right)
		for col := left; col <= right; col++ {
			res[top][col] = index
			index++
		}

		// 从上到下 (top+1, right) > (bottom, right)
		for row := top + 1; row <= bottom; row++ {
			res[row][right] = index
			index++
		}
		if left < right && top < bottom {
			// 从右到左 (bottom, right-1) > (bottom, left+1)
			for col := right - 1; col >= left+1; col-- {
				res[bottom][col] = index
				index++
			}

			// 从下到上 (bottom, left) > (top+1, left)
			for row := bottom; row >= top+1; row-- {
				res[row][left] = index
				index++
			}
		}

		// 继续访问内层元素
		left++
		right--
		bottom--
		top++
	}

	return res
}

// 模拟移动的方法
func methodWithSimulate(n int) [][]int {
	if n == 1 {
		return [][]int{[]int{1}}
	}

	var (
		rows, cols = n, n                                                          // 行数和列数
		total      = n * n                                                         // 总元素的数量
		matrix     = make([][]int, rows)                                           // 辅助变量
		visited    = make([][]bool, rows)                                          // 记录坐标位置是否访问过
		res        [][]int                                                         // 返回最终的结果
		row, col   = 0, 0                                                          //纵坐标和横坐标
		ds         = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}} // 路径的方向  → ↓ ← ↑
		dsIdx      = 0                                                             // 用来判断朝哪个方向移动
	)

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
		matrix[i] = make([]int, cols)
	}

	res = matrix

	for i := 1; i <= total; i++ {
		// 将元素放置到res中
		res[row][col] = i
		// 将坐标设置成已经访问过的状态
		visited[row][col] = true
		// 下一个要访问的坐标位置
		nextRow, nextCol := row+ds[dsIdx][0], col+ds[dsIdx][1]
		// 如果下一个要访问的坐标位置超出了界线或已经访问过了则变化方向
		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols || visited[nextRow][nextCol] {
			dsIdx = (dsIdx + 1) % 4
		}

		// 改变坐标位置
		row += ds[dsIdx][0]
		col += ds[dsIdx][1]
	}
	return res
}

// @lc code=end
