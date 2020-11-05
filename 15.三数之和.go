/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (26.76%)
 * Likes:    2069
 * Dislikes: 0
 * Total Accepted:    211.6K
 * Total Submissions: 789.5K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例：
 *
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// @lc code=start
func threeSum(nums []int) [][]int {
    // 双指针法

    // 如果元素不足三个的话则直接返回
    if len(nums) < 3{
        return [][]int{}
    }
    // 先对数组进行排序
    sort.Ints(nums)

    var res [][]int
    for i:=0; i<len(nums); i++{

        // 如果当前元素大于0的话， 则直接break掉循环， 说明后面的元素都比0大， 无法组成三数之和为0的数组
        if nums[i] > 0{
            break
        }
        // 过滤重复元素(当前元素和前一个元素相同)
        if i>0 && nums[i] == nums[i-1]{
            continue
        }
        left, right := i+1, len(nums)-1
        for left < right{
            sum := nums[i] + nums[left] + nums[right]

            if sum > 0{
                // 三数之和过大， 需要将right变小一点
                right --
                continue
            }else if sum < 0{
                // 三数之和过小， 需要将left变大一点
                left ++
                continue
            }else{
                // 三数之和相等的情况
                res = append(res, []int{nums[i], nums[left], nums[right]})

                // 过滤掉重复元素
                for right > left && nums[right-1] == nums[right]{
                    right --
                }

                for left < right && nums[left] == nums[left+1]{
                    left ++
                }
            }

            left++
            right--
        }

    }
    return res
}
// @lc code=end
