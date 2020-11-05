/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (60.66%)
 * Likes:    614
 * Dislikes: 0
 * Total Accepted:    159.2K
 * Total Submissions: 260.4K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 *
 *
 */

// @lc code=start
func moveZeroes(nums []int)  {
	if len(nums) < 2{
		return
	}

	// 快慢指针的方法
	var j int
	for i:=0; i<len(nums); i++{
		if nums[i] == 0{
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		j++
	}

}


func moveZeroes1(nums []int)  {

	for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++{
			if nums[i] == 0{
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}

}
// @lc code=end
