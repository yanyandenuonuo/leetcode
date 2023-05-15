/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	// solution: 分离奇偶链表然后合并，将偶数链表拆分，最后拼接到奇数链表的结尾
	if head == nil {
		return nil
	}

	evenNodeHead, tailOddNode := head.Next, head

	for tailEvenNode := evenNodeHead; tailEvenNode != nil && tailEvenNode.Next != nil; {
		tailOddNode.Next = tailEvenNode.Next
		tailOddNode = tailOddNode.Next

		tailEvenNode.Next = tailOddNode.Next
		tailEvenNode = tailEvenNode.Next
	}

	tailOddNode.Next = evenNodeHead

	return head
}

// @lc code=end

