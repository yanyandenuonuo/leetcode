/*
 * @lc app=leetcode id=341 lang=golang
 *
 * [341] Flatten Nested List Iterator
 *
 * https://leetcode.com/problems/flatten-nested-list-iterator/description/
 *
 * algorithms
 * Medium (53.47%)
 * Likes:    1809
 * Dislikes: 711
 * Total Accepted:    184K
 * Total Submissions: 344K
 * Testcase Example:  '[[1,1],2,[1,1]]'
 *
 * Given a nested list of integers, implement an iterator to flatten it.
 *
 * Each element is either an integer, or a list -- whose elements may also be
 * integers or other lists.
 *
 * Example 1:
 *
 *
 *
 * Input: [[1,1],2,[1,1]]
 * Output: [1,1,2,1,1]
 * Explanation: By calling next repeatedly until hasNext returns
 * false,
 * the order of elements returned by next should be: [1,1,2,1,1].
 *
 *
 * Example 2:
 *
 *
 * Input: [1,[4,[6]]]
 * Output: [1,4,6]
 * Explanation: By calling next repeatedly until hasNext returns
 * false,
 * the order of elements returned by next should be: [1,4,6].
 *
 *
 *
 *
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// solution1: 构造函数中遍历然后存放入slice，然后next更新slice
/**
type NestedIterator struct {
	valList []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		valList: converNestedListToIntList(nestedList),
	}
}

func converNestedListToIntList(nestedList []*NestedInteger) []int {
	resList := make([]int, 0)
	for _, nestedInteger := range nestedList {
		if nestedInteger.IsInteger() {
			resList = append(resList, nestedInteger.GetInteger())
		} else {
			resList = append(resList, converNestedListToIntList(nestedInteger.GetList())...)
		}
	}
	return resList
}

func (this *NestedIterator) Next() int {
	val := this.valList[0]
	this.valList = this.valList[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	if len(this.valList) == 0 {
		return false
	}
	return true
}
*/

// solution2: 栈，直接保存nestedList，然后每次取栈顶元素的第一个元素，并将剩余元素放入栈顶
type NestedIterator struct {
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		stack: [][]*NestedInteger{nestedList},
	}
}

func (this *NestedIterator) Next() int {
	val := this.stack[len(this.stack)-1][0].GetInteger()
	this.stack[len(this.stack)-1] = this.stack[len(this.stack)-1][1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		topNestedList := this.stack[len(this.stack)-1]
		if len(topNestedList) == 0 {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}

		topNestedIterator := topNestedList[0]
		if topNestedIterator.IsInteger() {
			return true
		}
		this.stack[len(this.stack)-1] = this.stack[len(this.stack)-1][1:]
		this.stack = append(this.stack, topNestedIterator.GetList())
	}
	return false
}

// @lc code=end

