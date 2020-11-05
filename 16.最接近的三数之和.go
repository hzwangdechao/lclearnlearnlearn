/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.75%)
 * Likes:    525
 * Dislikes: 0
 * Total Accepted:    139.1K
 * Total Submissions: 303.9K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 *
 *
 * 示例：
 *
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 10^3
 * -10^3 <= nums[i] <= 10^3
 * -10^4 <= target <= 10^4
 *
 *
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	// 与三数之和为0的解法相似， 都是利用双指针的方法
	sort.Ints(nums)
	// 最终返回的结果
	ans := nums[0] + nums[1] + nums[2]
	n := len(nums)
	for i:=0; i<n; i++{
		left, right := i+1, len(nums)-1
		for left < right{
			sum := nums[i] + nums[left] + nums[right]
			if abs(target-sum) < abs(target-ans){
				ans = sum
			}

			// 如果当前的和比target小的话， 将left++， 否则将right--
			if sum < target{
				left ++
			}else if sum > target{
				right --
			}else{
				return ans
			}
		}
	}
	return ans
}

func abs(x int)int  {
	if x < 0{
		return -x
	}
	return x
}

// @lc code=end
