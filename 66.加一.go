/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 *
 * https://leetcode-cn.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (44.01%)
 * Likes:    465
 * Dislikes: 0
 * Total Accepted:    151.3K
 * Total Submissions: 343.3K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
 *
 * 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
 *
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 *
 * 示例 1:
 *
 * 输入: [1,2,3]
 * 输出: [1,2,4]
 * 解释: 输入数组表示数字 123。
 *
 *
 * 示例 2:
 *
 * 输入: [4,3,2,1]
 * 输出: [4,3,2,2]
 * 解释: 输入数组表示数字 4321。
 *
 *
 */

// @lc code=start
// func plusOne(digits []int) []int {

// 	a := 1
// 	var res []int
// 	for i:= len(digits)-1; i>=0;i--{
// 		r := digits[i] + a
// 		if r == 10{
// 			res = append(res, 0)
// 			a = 1
// 		}else{
// 			res = append(res, r)
// 			a = 0
// 		}
// 	}
// 	if a == 1{
// 		res = append(res, 1)
// 	}
// 	var res2 []int

// 	for i:= len(res)-1; i>=0;i--{
// 		res2 = append(res2, res[i])
// 	}
// 	return res2

// }

func plusOne(digits []int) []int {
	for i:= len(digits)-1; i >= 0; i--{
		if digits[i] < 9{
			digits[i]++
			fmt.Println(digits)
			return digits
		}else{
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

// @lc code=end
