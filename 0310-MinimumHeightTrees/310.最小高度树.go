/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	// solution: 树越矮则越宽，此时距离最长的两个点肯定是经过根节点的两条最长边
	// 			  这时高度肯定为distance[left, right]/2
	//			  若为偶数，则根节点为length/2, length/2+1，注意取值时的索引值偏移
	//			  若为奇数，则根节点为length/2
	// solution1: 寻找距离最长的2个点用DFS
	// solution2: 寻找距离最长的2个点用BFS
	if n == 1 {
		return []int{0}
	}

	nodeRelation := make([][]int, n)
	for _, edge := range edges {
		nodeRelation[edge[0]] = append(nodeRelation[edge[0]], edge[1])
		nodeRelation[edge[1]] = append(nodeRelation[edge[1]], edge[0])
	}

	preNodeRelation := make([]int, n)
	var bfs func(startNode int) int
	bfs = func(startNode int) int {
		targetNode := -1

		memory := make([]bool, n)

		nodeQueue := []int{startNode}
		for len(nodeQueue) > 0 {
			targetNode = nodeQueue[0]
			memory[targetNode] = true
			nodeQueue = nodeQueue[1:]

			for _, node := range nodeRelation[targetNode] {
				if memory[node] {
					continue
				}
				memory[node] = true

				preNodeRelation[node] = targetNode

				nodeQueue = append(nodeQueue, node)
			}
		}

		return targetNode
	}

	leftNode := bfs(rand.Intn(n))
	rightNode := bfs(leftNode)

	preNodeRelation[leftNode] = -1

	nodePath := make([]int, 0)
	for startNode := rightNode; preNodeRelation[startNode] != -1; startNode = preNodeRelation[startNode] {
		nodePath = append(nodePath, startNode)
	}
	nodePath = append(nodePath, leftNode)

	if len(nodePath)%2 == 0 {
		// 注意索引偏移，所以中间值索引应该为length/2-1, length/2
		return []int{nodePath[len(nodePath)/2-1], nodePath[len(nodePath)/2]}
	}
	return []int{nodePath[len(nodePath)/2]}

	// solution3: 基于拓扑排序，计算每个节点的度，然后循环删除度较小的节点，直到剩余2个以内的节点，最后剩余的即为可选的根节点
	// todo
}

// @lc code=end

