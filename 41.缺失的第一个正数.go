/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode-cn.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (40.23%)
 * Likes:    794
 * Dislikes: 0
 * Total Accepted:    90.7K
 * Total Submissions: 225.3K
 * Testcase Example:  '[1,2,0]'
 *
 * 给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [1,2,0]
 * 输出: 3
 *
 *
 * 示例 2:
 *
 * 输入: [3,4,-1,1]
 * 输出: 2
 *
 *
 * 示例 3:
 *
 * 输入: [7,8,9,11,12]
 * 输出: 1
 *
 *
 *
 *
 * 提示：
 *
 * 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
 *
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	return methodWithHash(nums)
}



// hash表的方法
func methodWithHash(nums []int) int {
	n := len(nums)
	for i:=0; i<n; i++{
		if nums[i] <= 0{
			nums[i] = n+1
		}
	}
	fmt.Println(nums)

	for i:=0; i<n; i++{
		num := abs(nums[i])
		if num <= n{
			// 将元素值对应的位置的元素变成负数
			nums[num-1] = -abs(nums[num-1])
		}
	}
	fmt.Println(nums)

	for i:=0; i<n; i++{
		if nums[i] > 0{
			return i+1
		}
	}

	return n+1
}

func abs(x int) int {
	if x < 0{
		return -x
	}
	return x
}

// @lc code=end
n = 6
a = [3, 4, -1, 1, 9, -5]
b = [3, 4, 7, 1, 9, 7]
c = [-3, 4, -7, -1, 9, 7]

