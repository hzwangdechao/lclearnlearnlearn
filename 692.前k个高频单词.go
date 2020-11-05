/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 *
 * https://leetcode-cn.com/problems/top-k-frequent-words/description/
 *
 * algorithms
 * Medium (51.44%)
 * Likes:    125
 * Dislikes: 0
 * Total Accepted:    15.8K
 * Total Submissions: 30.7K
 * Testcase Example:  '["i", "love", "leetcode", "i", "love", "coding"]\n2'
 *
 * 给一非空的单词列表，返回前 k 个出现次数最多的单词。
 *
 * 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
 *
 * 示例 1：
 *
 *
 * 输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
 * 输出: ["i", "love"]
 * 解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
 * ⁠   注意，按字母顺序 "i" 在 "love" 之前。
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"],
 * k = 4
 * 输出: ["the", "is", "sunny", "day"]
 * 解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
 * ⁠   出现次数依次为 4, 3, 2 和 1 次。
 *
 *
 *
 *
 * 注意：
 *
 *
 * 假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。w
 * 输入的单词均由小写字母组成。
 *
 *
 *
 *
 * 扩展练习：
 *
 *
 * 尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
 *
 *
 */

// @lc code=start


type Word struct {
	Str string
	Num int
}

type Words []Word

func (t *Words) Len() int {
	return len(*t)
}

func (t *Words) Less(i, j int) bool {
	if (*t)[i].Num == (*t)[j].Num {
		return (*t)[i].Str < (*t)[j].Str
	}
	return (*t)[i].Num > (*t)[j].Num
}

func (t *Words) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *Words)Push(x interface{}) {
	*t = append(*t, x.(Word))
}

func (t *Words)Pop() interface{}  {
	n := len(*t)
	tmp := (*t)[n-1]
	*t = (*t)[:n-1]
	return tmp
}

func topKFrequent(words []string, k int) []string {
	// strs := []string{"i", "love", "leetcode", "i", "love", "coding"}
	// k := 2

	m := make(map[string]int)

	for _, s := range words{
		m[s] += 1
	}

	wordList := Words{}
	for k, v := range m{
		wordInfo := Word{
			Str: k,
			Num: v,
		}
		wordList = append(wordList, wordInfo)
	}

	heap.Init(&wordList)

	var res []string
	for wordList.Len()> 0 && k > 0{
		x := heap.Pop(&wordList)
		res = append(res, x.(Word).Str)
		k --
	}
	fmt.Println(res)
	return res
}

// type node struct {
// 	s   string
// 	num int
// }
// type nodes []node

// func (n nodes) Len() int {
// 	return len(n)
// }
// //返回true是要保持的模样
// func (n nodes) Swap(i, j int) {
// 	n[i], n[j] = n[j], n[i]
// }
// func (n nodes) Less(i, j int) bool {
// 	if n[i].num == n[j].num {
// 		return n[i].s < n[j].s
// 	}
// 	return n[i].num > n[j].num
// }
// func (n *nodes) Pop() interface{} {
// 	// 移除并返回最后一个元素
// 	var t = (*n)[0]
// 	old := *n
// 	old[0] = old[n.Len()-1]
// 	old = old[:n.Len()-1]
// 	*n = old
// 	return t
// }
// func (n *nodes) Push(x interface{})  {
// 	*n = append(*n, x.(node))
// }


// func topKFrequent(words []string, k int) []string {
// 	m := make(map[string]int)
// 	for _, word := range words{
// 		m[word] ++
// 	}

// 	var n nodes
// 	for k, v := range m{
// 		heap.Push(&n, node{s:k, num:v})
// 	}
// 	heap.Init(&n)

// 	var res []string
// 	for k != 0{
// 		var t = n.Pop().(node)
// 		heap.Fix(&n, 0)
// 		res = append(res, t.s)
// 		k--
// 	}
// 	return res

// }
// @lc code=end
