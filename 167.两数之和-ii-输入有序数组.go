/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 *
 * https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/description/
 *
 * algorithms
 * Easy (55.94%)
 * Likes:    353
 * Dislikes: 0
 * Total Accepted:    128.7K
 * Total Submissions: 230K
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
 *
 * 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
 *
 * 说明:
 *
 *
 * 返回的下标值（index1 和 index2）不是从零开始的。
 * 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
 *
 *
 * 示例:
 *
 * 输入: numbers = [2, 7, 11, 15], target = 9
 * 输出: [1,2]
 * 解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
 *
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
    return binarySearch(numbers, target)
}


func binarySearch(numbers []int, target int) []int{
    for i:=0; i<len(numbers); i++{
        remainder := target - numbers[i]

        left, right := i+1, len(numbers)-1
        for left <= right{
            mid := (right - left )/ 2+ left
            if numbers[mid] == remainder{
                return []int{i+1, mid+1}
            }else if numbers[mid] > remainder{
                right = mid -1
            }else{
                left = mid + 1
            }
        }
    }
    return []int{-1, -1}
}

func twoPointer(numbers []int, target int)[]int{
    left , right := 0, len(numbers)-1
    for left <= right{
        sum := numbers[left] + numbers[right]
        if sum == target{
            return []int{left+1, right+1}
        }else if sum > target{
            right -= 1
        }else{
            left += 1
        }
    }
    return  []int{-1, -1}
}


func hashMap(numbers []int, target int) []int {
    m := make(map[int]int, 0)
    for i, num := range numbers{
        remainder := target - num
        if v, ok := m[num]; ok {
            return []int{v+1, i+1}
        }else{
            m[remainder] = i
        }
    }
    return []int{-1, -1}
}
// @lc code=end
