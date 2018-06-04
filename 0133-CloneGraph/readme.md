# Clone Graph

## Description

&emsp;&emsp;Clone an undirected graph. Each node in the graph contains a **label** and a list of its **neighbors**.

## OJ's undirected graph serialization

&emsp;&emsp;Nodes are labeled uniquely.

&emsp;&emsp;We use **#** as a separator for each node, and **,** as a separator for node label and each neighbor of 
the node.

&emsp;&emsp;As an example, consider the serialized graph **\{0,1,2\#1,2\#2,2\}**.

&emsp;&emsp;The graph has a total of three nodes, and therefore contains three parts as separated by **#**.

1. First node is labeled as 0. Connect node 0 to both nodes **1** and **2**.
2. Second node is labeled as 1. Connect node 1 to node **2**.
3. Third node is labeled as 2. Connect node 2 to node 2 \(itself\), thus forming a self-cycle.

&emsp;&emsp;Visually, the graph looks like the following:

```
       1
      / \
     /   \
    0 --- 2
         / \
         \_/
```

## Difficulty

&emsp;&emsp;Medium

## Other

&emsp;&emsp;todo，添加解题思路。