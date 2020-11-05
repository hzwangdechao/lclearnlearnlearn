/*
 * @lc app=leetcode.cn id=679 lang=golang
 *
 * [679] 24 点游戏
 *
 * https://leetcode-cn.com/problems/24-game/description/
 *
 * algorithms
 * Hard (44.65%)
 * Likes:    192
 * Dislikes: 0
 * Total Accepted:    14.4K
 * Total Submissions: 27.3K
 * Testcase Example:  '[4,1,8,7]'
 *
 * 你有 4 张写有 1 到 9 数字的牌。你需要判断是否能通过 *，/，+，-，(，) 的运算得到 24。
 *
 * 示例 1:
 *
 * 输入: [4, 1, 8, 7]
 * 输出: True
 * 解释: (8-4) * (7-1) = 24
 *
 *
 * 示例 2:
 *
 * 输入: [1, 2, 1, 2]
 * 输出: False
 *
 *
 * 注意:
 *
 *
 * 除法运算符 / 表示实数除法，而不是整数除法。例如 4 / (1 - 2/3) = 12 。
 * 每个运算符对两个数进行运算。特别是我们不能用 - 作为一元运算符。例如，[1, 1, 1, 1] 作为输入时，表达式 -1 - 1 - 1 - 1
 * 是不允许的。
 * 你不能将数字连接在一起。例如，输入为 [1, 2, 1, 2] 时，不能写成 12 + 12 。
 *
 *
 */

// @lc code=start
const (
	TARGET = 24
	EPSILON = 1e-6
	ADD, MULTIPLY, SUBTRACT, DIVIDE = 0, 1, 2, 3
)
func judgePoint24(nums []int) bool {
	// 因为 / 是实数相除， 所以需要将int 转换为 float进行计算
	list := []float64{}
	for _, num := range nums{
		list = append(list, float64(num))
	}
	return solve(list)

}
func solve(list []float64) bool {
	if len(list) == 0{
		return false
	}

	// 如果列表只剩下一个元素， 验证是否是24
	if len(list) == 1{
		return abs(list[0] - TARGET) < EPSILON
	}

	size := len(list)
	for i:=0; i<size; i++{
		for j:=0; j<size; j++{
			if i != j{  //  不能用相同的元素进行加减乘除
				list2 := []float64{}
				for k:=0; k<size; k++{
					if k != i && k!=j{
						list2 = append(list2, list[k])
					}
				}

				for k:=0; k<4; k++{
					if k<2 && i<j{
						// 加法和乘法满足交换律
						continue
					}

					switch k {
					case ADD:
						list2 = append(list2, list[i] + list[j])
					case MULTIPLY:
						list2 = append(list2, list[i]*list[j])
					case SUBTRACT:
						list2 = append(list2, list[i]-list[j])
					case DIVIDE:
						// 0 不能作为除数
						if abs(list[j]) < EPSILON{
							continue
						}else{
							list2 = append(list2, list[i]/list[j])
						}
					}

					if solve(list2){
						return true
					}
					list2 = list2[:len(list2)-1]
				}
			}
		}
	}
	return false
}

func abs(x float64) float64 {
	if x >= 0{
		return x
	}
	return -x
}
// @lc code=end
