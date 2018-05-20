# Add Two Numbers

## Description

&emsp;&emsp;You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored 
in **reverse** order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

&emsp;&emsp;You may assume that each input would have exactly one solution, and you may not use the same element twice.

## Example:

``` 
    Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
    
    Output: 7 -> 0 -> 8
    
    Explanation: 342 + 465 = 807.
```

## Difficulty

&emsp;&emsp;Medium

## Solution

### eg1

&emsp;&emsp;按位进行处理
Complexity Analysis

#### Complexity Analysis

- Time complexity : O\(max\(m, n\)\). Assume that mm and nn represents the length of l1 and l2 respectively, the 
algorithm above iterates at most max\(m,n\) times.

- Space complexity : O\(max\(m, n\)\). The length of the new list is at most max\(m, n\) + 1.

## 测试

- Java版本通过LeetCode官方测试

## Other