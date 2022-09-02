/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 *
 * https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (48.32%)
 * Likes:    3491
 * Dislikes: 169
 * Total Accepted:    367.1K
 * Total Submissions: 759.1K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 * Serialization is the process of converting a data structure or object into a
 * sequence of bits so that it can be stored in a file or memory buffer, or
 * transmitted across a network connection link to be reconstructed later in
 * the same or another computer environment.
 * 
 * Design an algorithm to serialize and deserialize a binary tree. There is no
 * restriction on how your serialization/deserialization algorithm should work.
 * You just need to ensure that a binary tree can be serialized to a string and
 * this string can be deserialized to the original tree structure.
 * 
 * Clarification: The input/output format is the same as how LeetCode
 * serializes a binary tree. You do not necessarily need to follow this format,
 * so please be creative and come up with different approaches yourself.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: root = [1,2,3,null,null,4,5]
 * Output: [1,2,3,null,null,4,5]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: root = []
 * Output: []
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: root = [1]
 * Output: [1]
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: root = [1,2]
 * Output: [1,2]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in the tree is in the range [0, 10^4].
 * -1000 <= Node.val <= 1000
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 import "strconv"
 import "strings"


type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	resList = make([]string, 0)

	if root == nil {
		return ""
	}
	resList := convert(root)

	resStr := ""

	for _, val := range resList {
		resStr += val + ","
	}

	return strings.TrimSuffix(resStr, ",")
}

func convert(root *TreeNode) []string {
	if root == nil {
		resList = append(resList, "nil")
		return resList
	}
	resList = append(resList, strconv.Itoa(root.Val))

	convert(root.Left)
	convert(root.Right)

	return resList
}

var resList = make([]string, 0)

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	resList = strings.Split(data, ",")
    return revert()
}

func revert() *TreeNode {
	if len(resList) == 0 {
		return nil
	}

	currentVal := resList[0]
	if len(resList) > 1 {
		resList = resList[1:]
	} else {
		resList = []string{}
	}

	if currentVal == "nil" {
		return nil
	}
	val, _ := strconv.Atoi(currentVal)

	root := new(TreeNode)
	root.Val = val

	root.Left = revert()
	root.Right = revert()

	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

