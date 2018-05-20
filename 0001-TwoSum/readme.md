# Two Sum

## Description

&emsp;&emsp;Given an array of integers, return indices of the two numbers such that they add up to a specific target.

&emsp;&emsp;You may assume that each input would have exactly one solution, and you may not use the same element twice.

## Example:

``` 
    Given nums = [2, 7, 11, 15], target = 9,
    
    Because nums[0] + nums[1] = 2 + 7 = 9,
    
    return [0, 1].
```

## Difficulty

&emsp;&emsp;Easy

## Solution

### eg1

&emsp;&emsp;两层遍历，复杂度为O\(n^2\)。

#### Complexity Analysis
- Time complexity : O\(n^2\). For each element, we try to find its complement by looping through the rest of array 
which takes O\(n\) time. Therefore, the time complexity is O\(n^2\)

- Space complexity : O\(1\).

### eg2

&emsp;&emsp;一层遍历，时间复杂度为O\(n\)。

&emsp;&emsp;遍历时得到一个value，使用target \- value得到一个needle值。

- PHP借助系统函数array_search\($array, $needle\)得到一个key，若key存在且不等于index，则返回返回\[index\, key\]即可即可。
- Java借助HashMap存放\(Value\, Key\)，添加之前判断是否存在needle值。若存在则判断得到的key是否与index值相等，若相等则继续
遍历，否则返回\[key\, index\]即可。

&emsp;&emsp;**Note:** 注意返回的key和index的顺序。

#### Complexity Analysis:

- Time complexity : O\(n\). We traverse the list containing n elements only once. Each look up in the table costs only 
O\(1\) time.

- Space complexity : O\(n\). The extra space required depends on the number of items stored in the hash table, which 
stores at most n elements.

## 测试

- PHP版本通过自测
- Java版本通过LeetCode官方测试

## Other

