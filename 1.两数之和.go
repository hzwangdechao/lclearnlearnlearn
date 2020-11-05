/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start

// func twoSum(nums []int, target int) []int {
	// // make 一个 map， map[剩余值]index
	// m := make(map[int]int)
	// // 遍历数组
	// for i , v := range nums{
		// // 判断 map 中是否含有target-v这个键
		// k, ok := m[target - v]
		// if ok{
			// // 如果存在的话， 返回两个的index
			// return []int{k, i}
		// }
		// // 不存在的话， 给map赋值
		// m[v] = i
	// }
	// return []int{-1, -1}
// }

// func twoSum(nums []int, target int) []int {
// 	// 暴力求解
// 	for i, v := range nums{
// 		// 计算差值
// 		diff := target - v
// 		// 从i+1后面的数组中查找是否存在与diff相同的数字
// 		for j, k := range nums[i+1:]{
// 			if diff == k{
// 				return  []int{i, i+j+1}
// 			}
// 		}

// 	}
// 	return []int{-1, -1}
// }

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums{
		if k, ok := [target-v]; ok==true{
			return []int{k, i}
		}
		m[v] = i
	}
	return []int{-1, -1}
}

// @lc code=end
