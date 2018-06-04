# Populating Next Right Pointers in Each Node II

## Description

&emsp;&emsp;Given a binary tree

```
    struct TreeLinkNode {
        TreeLinkNode *left;
        TreeLinkNode *right;
        TreeLinkNode *next;
    }
```

&emsp;&emsp;Populate each next pointer to point to its next right node. If there is no next right node, the next 
pointer should be set to **NULL**.

&emsp;&emsp;Initially, all next pointers are set to **NULL**.

## Note

- You may only use constant extra space.
- Recursive approach is fine, implicit stack space does not count as extra space for this problem.

## Example

### eg1

&emsp;&emsp;Given the following binary tree,

```
         1
       /  \
      2    3
     / \    \
    4   5    7
```

&emsp;&emsp;After calling your function, the tree should look like:

```
         1 -> NULL
       /  \
      2 -> 3 -> NULL
     / \    \
    4-> 5 -> 7 -> NULL
```

## Difficulty

&emsp;&emsp;Medium

## Other

&emsp;&emsp;todo，添加解题思路。