#
# @lc app=leetcode.cn id=1379 lang=python3
#
# [1379] 找出克隆二叉树中的相同节点
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        if original == target:
            return cloned
        if original.left == None and original.right == None:
            return None
        if original.left != None:
            find_left = self.getTargetCopy(original.left, cloned.left, target)
            if find_left != None :
                return find_left
        if original.right != None:
            return self.getTargetCopy(original.right, cloned.right, target)

# @lc code=end

