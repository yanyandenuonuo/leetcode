# Word Break

## Description

&emsp;&emsp;Given a **non-empty** string s and a dictionary wordDict containing a list of **non-empty** words, 
determine if s can be segmented into a space-separated sequence of one or more dictionary words.

## Note

- The same word in the dictionary may be reused multiple times in the segmentation.
- You may assume the dictionary does not contain duplicate words.

## Example

### eg1

```
    Input: s = "leetcode", wordDict = ["leet", "code"]
    
    Output: true
    
    Explanation: Return true because "leetcode" can be segmented as "leet code".
```

### eg2

```
    Input: s = "applepenapple", wordDict = ["apple", "pen"]
    
    Output: true
    
    Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
                 Note that you are allowed to reuse a dictionary word.
```

### eg3

```
    Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
    
    Output: false
```

## Difficulty

&emsp;&emsp;Medium

## Other

&emsp;&emsp;todo，添加解题思路。