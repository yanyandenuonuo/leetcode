/*
 * @lc app=leetcode.cn id=385 lang=golang
 *
 * [385] 迷你语法分析器
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	// solution1: DFS
	/**
	idx := 0

	var dfs func() *NestedInteger
	dfs = func() *NestedInteger {
		nestedInteger := new(NestedInteger)

		if s[idx] == '[' {
			idx += 1
			for s[idx] != ']' {
				nestedInteger.Add(*dfs())
				if s[idx] == ',' {
					idx += 1
				}
			}

			idx += 1

			return nestedInteger
		}

		isNegative := s[idx] == '-'
		if isNegative {
			idx += 1
		}
		num := 0
		for ; idx < len(s) && s[idx] >= '0' && s[idx] <= '9'; idx += 1 {
			num = num*10 + int(s[idx]-'0')
		}

		if isNegative {
			num = -num
		}

		nestedInteger.SetInteger(num)

		return nestedInteger
	}

	return dfs()
	*/

	// solution2: 栈
	if s[0] != '[' {
		num, _ := strconv.Atoi(s)
		nestedInteger := new(NestedInteger)
		nestedInteger.SetInteger(num)
		return nestedInteger
	}

	stack := make([]*NestedInteger, 0)

	num, isNegative := 0, false

	for idx := range s {
		switch s[idx] {
		case '-':
			isNegative = true
		case '[':
			nestedInteger := new(NestedInteger)
			stack = append(stack, nestedInteger)
		case ',':
			if s[idx-1] >= '0' && s[idx-1] <= '9' {
				if isNegative {
					num = -num
					isNegative = false
				}

				nestedInteger := new(NestedInteger)
				nestedInteger.SetInteger(num)

				stack[len(stack)-1].Add(*nestedInteger)

				num = 0
			}
		case ']':
			if s[idx-1] >= '0' && s[idx-1] <= '9' {
				if isNegative {
					num = -num
					isNegative = false
				}

				nestedInteger := new(NestedInteger)
				nestedInteger.SetInteger(num)

				stack[len(stack)-1].Add(*nestedInteger)

				num = 0
			}

			if len(stack) > 1 {
				stack[len(stack)-2].Add(*stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		default:
			num = num*10 + int(s[idx]-'0')
		}
	}

	return stack[0]
}

// @lc code=end

