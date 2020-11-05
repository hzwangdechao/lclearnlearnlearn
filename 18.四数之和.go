/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (38.43%)
 * Likes:    645
 * Dislikes: 0
 * Total Accepted:    128.6K
 * Total Submissions: 328.3K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4{
		return nil
	}

	sort.Ints(nums)

	var res [][]int
	for i:=0; i<len(nums)-3; i++{

		// 第一个元素去重
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		for j:=i+1; j<len(nums)-2; j++{
			// 第二个元素去重
			if j>i+1 && nums[j] == nums[j-1]{
				continue
			}

			left, right := j+1, len(nums)-1
			for left < right{
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum > target{
					right --
					continue

				}else if sum < target{
					left ++
					continue

				}else{
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for right > left && nums[right] == nums[right-1]{
						right --
					}
					for left < right && nums[left] == nums[left+1]{
						left ++
					}
				}
				left ++
				right --
			}

		}
	}

	return res

}
// @lc code=end
