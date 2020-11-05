/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (59.71%)
 * Likes:    204
 * Dislikes: 0
 * Total Accepted:    107.4K
 * Total Submissions: 178.5K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 *
 * 示例 1:
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 * 说明:
 * 你可以假设字符串只包含小写字母。
 *
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	// // 三次遍历的方法， 分别遍历s t m
	// // 长度不相等的情况下直接返回false
	// if len(s) != len(t){
	// 	return false
	// }

	// // hash方法进行遍历判断
	// m  := make(map[rune]int)

	// // 遍历s
	// for _, b := range s{
	// 	if _, ok := m[b]; !ok{
	// 		m[b] = 1
	// 	}else{
	// 		m[b] ++
	// 	}
	// }
	// fmt.Println(m)
	// // 遍历t
	// for _, b := range t{
	// 	if _, ok := m[b]; !ok{
	// 		// 出现了不同的字符
	// 		return false
	// 	}else{
	// 		m[b]--
	// 	}
	// }

	// // 遍历m
	// for _, v := range m{
	// 	if v != 0{
	// 		return false
	// 	}
	// }
	// return true

	// 两个数组进行比较的方法, 利用两个长度都为26的字符串数组， 统计每个字符串中小写字母出现的字数， 最后判断两个数组是否相等
	a := [26]int{}
	b := [26]int{}

	for _, r := range s{
		a[r-'a'] +=1
	}
	fmt.Println(a)
	for _, r := range t{
		b[r-'a'] +=1
	}
	fmt.Println(b)
	return a==b

}
// @lc code=end
