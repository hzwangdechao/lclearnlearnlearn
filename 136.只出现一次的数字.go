/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 *
 * https://leetcode-cn.com/problems/single-number/description/
 *
 * algorithms
 * Easy (66.80%)
 * Likes:    1313
 * Dislikes: 0
 * Total Accepted:    223.1K
 * Total Submissions: 324.3K
 * Testcase Example:  '[2,2,1]'
 *
 * 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
 *
 * 说明：
 *
 * 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 * 示例 1:
 *
 * 输入: [2,2,1]
 * 输出: 1
 *
 *
 * 示例 2:
 *
 * 输入: [4,1,2,1,2]
 * 输出: 4
 *
 */

// @lc code=start
// func singleNumber(nums []int) int {
// 	m := make(map[int]int)
// 	for _, num := range nums{
// 		if _, ok := m[num]; !ok{
// 			m[num] = 1
// 		}else{
// 			m[num] ++
// 		}
// 	}
// 	fmt.Println(m)
// 	for k , v := range m{
// 		if v == 1{
// 			return k
// 		}
// 	}
// 	return 0
// }


func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums{
		res ^= num
	}
	return res

}


// @lc code=end
