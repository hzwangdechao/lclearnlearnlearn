/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * [841] 钥匙和房间
 *
 * https://leetcode-cn.com/problems/keys-and-rooms/description/
 *
 * algorithms
 * Medium (62.35%)
 * Likes:    108
 * Dislikes: 0
 * Total Accepted:    20.3K
 * Total Submissions: 31.8K
 * Testcase Example:  '[[1],[2],[3],[]]'
 *
 * 有 N 个房间，开始时你位于 0 号房间。每个房间有不同的号码：0，1，2，...，N-1，并且房间里可能有一些钥匙能使你进入下一个房间。
 *
 * 在形式上，对于每个房间 i 都有一个钥匙列表 rooms[i]，每个钥匙 rooms[i][j] 由 [0,1，...，N-1] 中的一个整数表示，其中
 * N = rooms.length。 钥匙 rooms[i][j] = v 可以打开编号为 v 的房间。
 *
 * 最初，除 0 号房间外的其余所有房间都被锁住。
 *
 * 你可以自由地在房间之间来回走动。
 *
 * 如果能进入每个房间返回 true，否则返回 false。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入: [[1],[2],[3],[]] 一共有4个房间， 第一个房间有打开第二个房间的钥匙， 第二个房间有打开第三个房间的钥匙， 最终到达第三个最后一个房间
 * 输出: true
 * 解释:
 * 我们从 0 号房间开始，拿到钥匙 1。
 * 之后我们去 1 号房间，拿到钥匙 2。
 * 然后我们去 2 号房间，拿到钥匙 3。
 * 最后我们去了 3 号房间。
 * 由于我们能够进入每个房间，我们返回 true。
 *
 *
 * 示例 2：
 *
 * 输入：[[1,3],[3,0,1],[2],[0]]  一共有四个房间， 第一个房间有打开第二个房间和第四个房间的钥匙， 然而第二个房间和第四个房间没有打开第三个房间的钥匙
 * 输出：false
 * 解释：我们不能进入 2 号房间。
 *
 *
 * 提示：
 *
 *
 * 1 <= rooms.length <= 1000
 * 0 <= rooms[i].length <= 1000
 * 所有房间中的钥匙数量总计不超过 3000。
 *
 *
 */

// @lc code=start
var (
	num int // 遍历过的房间数量
	vis []bool  // 记录房间是否遍历过
)

func canVisitAllRooms(rooms [][]int) bool {
	return bfs(rooms)
	n := len(rooms)
	num = 0
	vis = make([]bool, n)

	dfs(rooms, 0)
	return num == n
}

// 深度遍历房间, x 代表第几号房间
func dfs(rooms [][]int, x int)  {
	vis[x] = true  // 将此房间设置为已经遍历过的状态
	num ++  // 增加遍历过的房间数量
	// 遍历房间的钥匙数量
	for _, y := range rooms[x]{
		// 如果房间没有遍历过的话， 就去遍历一下
		if !vis[y]{
			dfs(rooms, y)
		}
	}
}

// 广度遍历房间
func bfs(rooms [][]int) bool {
	n := len(rooms)
	num := 0  // 遍历过的房间数量
	vis := make([]bool, n)
	vis[0] = true

	queue := make([]int, 0)
	queue = append(queue, 0)

	for len(queue) > 0{
		room := queue[0]
		queue = queue[1:]
		num ++
		for _, y := range rooms[room]{
			if !vis[y]{
				vis[y] = true
				queue = append(queue, y)
			}
		}
	}

	return num == n
}

// @lc code=end
