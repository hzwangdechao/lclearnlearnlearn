/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 *
 * https://leetcode-cn.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (60.73%)
 * Likes:    324
 * Dislikes: 0
 * Total Accepted:    51.5K
 * Total Submissions: 84.1K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。
 *
 * 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4,
 * 2, 1, 1, 0, 0]。
 *
 * 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
 *
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	// 单纯遍历
    // res := []int{}

    // if len(T) == 0{
    //     return res
    // }
    // for i:=0; i<len(T); i++{
    //     r := i
    //     for j := i+1; j<len(T); j++{
    //         if T[j] > T[i]{
    //             r = j
    //             break
    //         }
    //     }
    //     res = append(res, r-i)
    // }

	// return res

	// 利用栈解题
	tempStack := make([]Stack, 0)
	res := make([]int, len(T))


	for i, v := range T{
		for len(tempStack) > 0{
			if v > tempStack[len(tempStack)-1].val{
				// 如果当前元素大于栈顶元素的话， 则计算相差距离
				res[tempStack[len(tempStack)-1].index] = i - tempStack[len(tempStack)-1].index
				// 出栈
				tempStack = tempStack[:len(tempStack)-1]
			}else{
				break
			}
		}
		tempStack = append(tempStack, Stack{v, i})
	}
	return res
}
type Stack struct{
	val int
	index int
}
// @lc code=end
