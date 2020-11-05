/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 *
 * https://leetcode-cn.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (34.91%)
 * Likes:    135
 * Dislikes: 0
 * Total Accepted:    28.9K
 * Total Submissions: 82.1K
 * Testcase Example:  '[3,2,1]'
 *
 * 给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
 *
 * 示例 1:
 *
 *
 * 输入: [3, 2, 1]
 *
 * 输出: 1
 *
 * 解释: 第三大的数是 1.
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1, 2]
 *
 * 输出: 2
 *
 * 解释: 第三大的数不存在, 所以返回最大的数 2 .
 *
 *
 * 示例 3:
 *
 *
 * 输入: [2, 2, 3, 1]
 *
 * 输出: 1
 *
 * 解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
 * 存在两个值为2的数，它们都排第二。
 *
 *
 */

// @lc code=start


// func thirdMax(nums []int) int {
// 	sort.Ints(nums)
// 	if len(nums) == 0{
//         return 0
// 	}

//     f, s, t := nums[0], nums[0], nums[0]

//     for i:=1; i<len(nums); i++{

//         if nums[i] > f{
//             s, t = f, s
//             f = nums[i]
//         }else if nums[i] ==f{
// 			continue
// 		}else if nums[i] > s{
// 			t = s
// 			s = nums[i]
// 		}else if nums[i] == s {
// 			continue
// 		}else if nums[i] > t{
// 			t = nums[i]
// 		}
// 	}
// 	fmt.Println(f, s, t)

// 	if s == t{
// 		return f
// 	}else{
// 		return t
// 	}

// }


func thirdMax(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	first, second, third := nums[0], math.MinInt64, math.MinInt64

	for i:=1; i<len(nums); i++{
		if nums[i] == first || nums[i] == second || nums[i] == third{
			continue
		}

		if nums[i] > first{
			third = second
			second = first
			first = nums[i]
		}else if nums[i] > second{
			third = second
			second = nums[i]
		}else if nums[i] > third{
			third = nums[i]
		}
	}
	if third > math.MinInt64{
		return third
	}
	return first
}

// @lc code=end
