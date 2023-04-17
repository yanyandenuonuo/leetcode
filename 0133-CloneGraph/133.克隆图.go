/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	nodeMemory := make(map[*Node]*Node)
	return cloneNodeWithMemory(node, nodeMemory)
}

func cloneNodeWithMemory(node *Node, nodeMemory map[*Node]*Node) *Node {
	if cloneNode, isExist := nodeMemory[node]; isExist {
		return cloneNode
	}

	if node == nil {
		return nil
	}
	cloneNode := new(Node)
	cloneNode.Val = node.Val

	nodeMemory[node] = cloneNode

	cloneNeighbors := make([]*Node, 0, len(node.Neighbors))
	for _, neighbor := range node.Neighbors {
		cloneNeighbor := cloneNodeWithMemory(neighbor, nodeMemory)
		if cloneNeighbor == nil {
			continue
		}
		cloneNeighbors = append(cloneNeighbors, cloneNeighbor)
	}
	cloneNode.Neighbors = cloneNeighbors
	return cloneNode
}

// @lc code=end

