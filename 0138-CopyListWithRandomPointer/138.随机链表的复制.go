/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 随机链表的复制
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	// solution: copyWithMemory，需要记录节点对应关系，不然同样节点可能复制出多个不同节点
	return copyNodeWithMemory(head, map[*Node]*Node{})
}

func copyNodeWithMemory(node *Node, nodeMap map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if _, isExist := nodeMap[node]; !isExist {
		nodeMap[node] = &Node{
			Val: node.Val,
		}

		nodeMap[node].Next = copyNodeWithMemory(node.Next, nodeMap)

		nodeMap[node].Random = copyNodeWithMemory(node.Random, nodeMap)
	}

	return nodeMap[node]
}

// @lc code=end

