/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 *
 * https://leetcode-cn.com/problems/reconstruct-itinerary/description/
 *
 * algorithms
 * Medium (38.82%)
 * Likes:    225
 * Dislikes: 0
 * Total Accepted:    18.7K
 * Total Submissions: 44.5K
 * Testcase Example:  '[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]'
 *
 * 给定一个机票的字符串二维数组 [from,
 * to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。所有这些机票都属于一个从
 * JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。
 *
 * 说明:
 *
 *
 * 如果存在多种有效的行程，你可以按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"]
 * 相比就更小，排序更靠前
 * 所有的机场都用三个大写字母表示（机场代码）。
 * 假定所有机票至少存在一种合理的行程。
 *
 *
 * 示例 1:
 *
 * 输入: [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
 * 输出: ["JFK", "MUC", "LHR", "SFO", "SJC"]
 *
 *
 * JFK MUC LHR SFO SJC
 *
 * 示例 2:
 *
 * 输入: [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
 * JFK ATL JFK
 * 输出: ["JFK","ATL","JFK","SFO","ATL","SFO"]
 * 解释: 另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。
 *
 */

// @lc code=start
/*
	1. 从起点出发，进行深度优先搜索。
	2. 每次沿着某条边从某个顶点移动到另外一个顶点的时候，都需要删除这条边。
	3.如果没有可移动的路径，则将所在节点加入到栈中，并返回。

*/
func findItinerary(tickets [][]string) []string {
	var (
		m = map[string][]string{}
		res []string
	)
	for _, ticket := range tickets{
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}

	for key := range m{
		sort.Strings(m[key])
	}
	fmt.Println(m)

	var dfs func(curr string)

	dfs = func(curr string)  {
		for {
			if v, ok := m[curr]; !ok || len(v) ==0{
				break
			}
			tmp := m[curr][0] // curr的第一个出发地
			m[curr] = m[curr][1:] // 防止重复遍历
			dfs(tmp)
		}
		res = append(res, curr)
	}

	// 从JFK作为起点开始计算
	dfs("JFK")

	fmt.Println(res)

	for l, r:=0, len(res)-1; l < r; l, r = l+1, r -1{
		res[l], res[r] = res[r], res[l]
	}

	return res
}
// @lc code=end
