# Course Schedule

## Description

&emsp;&emsp;There are a total of n courses you have to take, labeled from **0** to **n\-1**.

&emsp;&emsp;Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is 
expressed as a pair: **\[0,1\]**

&emsp;&emsp;Given the total number of courses and a list of prerequisite **pairs**, is it possible for you to finish 
all courses?

## Example

### eg1

```
    Input: 2, [[1,0]] 
    
    Output: true
    
    Explanation: There are a total of 2 courses to take. 
                 To take course 1 you should have finished course 0. So it is possible.
```

### eg2

```
    Input: 2, [[1,0],[0,1]]
    
    Output: false
    
    Explanation: There are a total of 2 courses to take. 
                 To take course 1 you should have finished course 0, and to take course 0 you should
                 also have finished course 1. So it is impossible.
```

## Note

- The input prerequisites is a graph represented by **a list of edges**, not adjacency matrices. Read more about 
[how a graph is represented](https://www.khanacademy.org/computing/computer-science/algorithms/graph-representation/a/representing-graphs).
- You may assume that there are no duplicate edges in the input prerequisites.

## Difficulty

&emsp;&emsp;Medium

## Other

&emsp;&emsp;todo，添加解题思路。