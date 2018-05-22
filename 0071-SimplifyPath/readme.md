# Simplify Path

## Description

&emsp;&emsp;Given an absolute path for a file \(Unix-style\), simplify it.

## Example

### eg1

```
    path = "/home/", => "/home"
    
    path = "/a/./b/../../c/", => "/c"
```

## Corner Cases

- Did you consider the case where path = **"/../"**? In this case, you should return **"/"**.
- Another corner case is the path might contain multiple slashes **'/'** together, such as **"/home//foo/"**. In this 
case, you should ignore redundant slashes and return **"/home/foo"**.


## Difficulty

&emsp;&emsp;Medium

## Other

&emsp;&emsp;todo，添加解题思路。