/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 *
 * https://leetcode-cn.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (55.64%)
 * Likes:    255
 * Dislikes: 0
 * Total Accepted:    47.2K
 * Total Submissions: 84.9K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [1,1,1],
 * [1,0,1],
 * [1,1,1]
 * ]
 * 输出:
 * [
 * [1,0,1],
 * [0,0,0],
 * [1,0,1]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * [0,1,2,0],
 * [3,4,5,2],
 * [1,3,1,5]
 * ]
 * 输出:
 * [
 * [0,0,0,0],
 * [0,4,5,0],
 * [0,3,1,0]
 * ]
 *
 * 进阶:
 *
 *
 * 一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
 * 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
 * 你能想出一个常数空间的解决方案吗？
 *
 *
 */

// @lc code=start
func setZeroes(matrix [][]int)  {
	traverse(matrix)
}

// 遍历法， 如果遇到了为0的元素， 则将当前元素的同行和同列的第一个元素设置成0； 再进行二次遍历
func traverse(matrix [][]int)  {
	row := len(matrix)
	col := len(matrix[0])

	isCol := false
	for i:=0; i<row; i++{
		if matrix[i][0] == 0{
			isCol = true
		}
		for j:=1; j<col; j++{
			if matrix[i][j] == 0{
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i:=1; i<row; i++{
		for j:=1; j<col; j++{
			if matrix[i][0] == 0 || matrix[0][j] == 0{
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0{
		for j:=0; j<col; j++{
			matrix[0][j] = 0
		}
	}

	if isCol{
		for i:=0; i<row; i++{
			matrix[i][0] = 0
		}
	}
}


// 暴力解法， 如果当前元素为0的话， 将当前元素的同行和同列上的非零元素变为math,MinInt64；  最后再遍历一遍数组将math.MinInt64变为0
func violence(matrix [][]int)  {
	row := len(matrix)
	col := len(matrix[0])

	for i:=0; i<row; i++{
		for j:=0; j<col; j++{
			if matrix[i][j] == 0{
				for r:=0; r<row; r++{
					if matrix[r][j] != 0{
						matrix[r][j] = math.MinInt64
					}
				}

				for c:=0; c<col; c++{
					if matrix[i][c] != 0{
						matrix[i][c] = math.MinInt64
					}
				}
			}
		}
	}

	for i, r := range matrix{
		for j, c := range r{
			if c ==  math.MinInt64{
				matrix[i][j] = 0
			}
		}
	}

}

// 使用额外的空间
func withExtra(matrix [][]int)  {
	// 使用两个map来记录行号和列号
	rows := make(map[int]bool, 0)
	cols := make(map[int]bool, 0)
	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[i]); j++{
			if matrix[i][j] == 0{
				rows[i] = true
				cols[j] = true
			}
		}
	}
	fmt.Println(rows)
	fmt.Println(cols)

	for i:=0; i<len(matrix); i++{
		for j:=0; j<len(matrix[i]); j++{
			_, colOk := cols[j]
			_, rowOk := rows[i]

			if colOk || rowOk{
				matrix[i][j] = 0
			}

		}
	}

}

// @lc code=end
