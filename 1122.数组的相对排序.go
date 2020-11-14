/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 *
 * https://leetcode-cn.com/problems/relative-sort-array/description/
 *
 * algorithms
 * Easy (67.07%)
 * Likes:    116
 * Dislikes: 0
 * Total Accepted:    35.5K
 * Total Submissions: 51.1K
 * Testcase Example:  '[2,3,1,3,2,4,6,7,9,2,19]\n[2,1,4,3,9,6]'
 *
 * 给你两个数组，arr1 和 arr2，
 *
 *
 * arr2 中的元素各不相同
 * arr2 中的每个元素都出现在 arr1 中
 *
 *
 * 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1
 * 的末尾。
 *
 *
 *
 * 示例：
 *
 * 输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
 * 输出：[2,2,2,1,4,3,3,9,6,7,19]
 *
 *
 *
 *
 * 提示：
 *
 *
 * arr1.length, arr2.length <= 1000
 * 0 <= arr1[i], arr2[i] <= 1000
 * arr2 中的元素 arr2[i] 各不相同
 * arr2 中的每个元素 arr2[i] 都出现在 arr1 中
 *
 *
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	return method1(arr1, arr2)
	res := make([]int, 0)
	m1 := make(map[int]int, 0) // 记录arr1中每个元素的数量
	for _, num := range arr1 {
		m1[num] += 1
	}
	for _, num := range arr2 {
		for i := 0; i < m1[num]; i++ {
			res = append(res, num)
		}
		delete(m1, num)
	}
	residue := make([]int, 0)
	for k, v := range m1 {
		for i := 0; i < v; i++ {
			residue = append(residue, k)
		}
	}
	sort.Ints(residue)
	res = append(res, residue...)

	return res
}

func method1(arr1 []int, arr2 []int) []int {
	m := make(map[int]int) // 记录元素在arr2中出现的位置
	for i, num := range arr2 {
		m[num] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		// 元素出现的位置， 元素是否出现
		rankX, hasX := m[x]
		rankY, hasY := m[y]

		// 如果两个元素都出现的话， 按照rank进行排序
		if hasX && hasY {
			return rankX < rankY
		}

		// 如果只有一个元素出现话， 返回hasX(如果x出现的话， x就在前面， 否则y在前面)
		if hasX || hasY {
			return hasX
		}

		// 如果两个元素都没有出现的话， 升序排序
		return x < y
	})

	return arr1
}

// @lc code=end
