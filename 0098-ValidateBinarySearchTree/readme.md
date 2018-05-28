# Validate Binary Search Tree

## Description

&emsp;&emsp;Given a binary tree, determine if it is a valid binary search tree \(BST\).

&emsp;&emsp;Assume a BST is defined as follows:
            
- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

## Example

### eg1

```
    Input:
        2
       / \
      1   3
      
    Output: true
```

### eg2

```
    Input:
        5
       / \
      1   4
         / \
        3   6
        
    Output: false
    
    Explanation: The input is: [5,1,4,null,null,3,6]. The root node's value
                 is 5 but its right child's value is 4.
```

## Difficulty

&emsp;&emsp;Medium

## Other

&emsp;&emsp;todo，添加解题思路。