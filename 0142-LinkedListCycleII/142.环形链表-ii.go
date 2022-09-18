/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// Solution1: hashmap
	// Solution2: 推导a+(n+1)b+nc=2(a+b) => a=c+(n−1)(b+c)

	// 为什么会在slow跑不完一圈就相遇？
	// 假设圈长L
	// slow进入环的时候，假设fast距离它n步（n<L）
	// fast一次走两步，slow一次走一步，所以每走一次，它们的距离会缩短1；
	// 那么它们的距离会从n变为n-1,n-2,n-2.....1,0;即不可能只超越而不相遇；
	// 这个过程种slow走了n步，而n<L，即slow用不了一圈就会被追上
	slowPointer, fastPointer := head, head
	for {
		if slowPointer == nil {
			return nil
		}
		slowPointer = slowPointer.Next

		if fastPointer.Next == nil || fastPointer.Next.Next == nil {
			return nil
		}
		fastPointer = fastPointer.Next.Next

		if slowPointer == fastPointer {
			break
		}
	}
	slowPointer = head
	for {
		if slowPointer == fastPointer {
			return slowPointer
		}
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next
	}
	return nil
}

// @lc code=end

