/*
 * @lc app=leetcode.cn id=382 lang=golang
 *
 * [382] 链表随机节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	// solution1: 使用数组记录链表数据，生成数组长度内的随机数返回即可
	// solution2: 蓄水池抽样
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() int {
	if this.head == nil {
		return 0
	}
	res := 0
	for scanNode, idx := this.head, 1; scanNode != nil; scanNode = scanNode.Next {
		if rand.Intn(idx) == 0 {
			res = scanNode.Val
		}
		idx += 1
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end

