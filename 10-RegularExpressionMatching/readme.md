# Regular Expression Matching

## Description

&emsp;&emsp;Given an input string (**s**) and a pattern (**p**), implement regular expression matching with support 
for '**\.**' and '__\*__'.

```
    '.' Matches any single character.
    
    '*' Matches zero or more of the preceding element.
```

&emsp;&emsp;The matching should cover the **entire** input string (not partial).

### Note

- **s** could be empty and contains only lowercase letters **a-z**.
- **p** could be empty and contains only lowercase letters **a-z**, and characters like **\.** or __\*__.

## Example:

### eg1

``` 
    Input:
    
    s = "aa"
    
    p = "a"
    
    Output: false
    
    Explanation: "a" does not match the entire string "aa".
```

### eg2

``` 
    Input:
    
    s = "aa"
    
    p = "a*"
    
    Output: true
    
    Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes 
                    "aa".
```

### eg3

``` 
    Input:
    
    s = "ab"
    
    p = ".*"
    
    Output: true
    
    Explanation: ".*" means "zero or more (*) of any character (.)".
```

### eg4

``` 
    Input:
    
    s = "aab"
    
    p = "c*a*b"
    
    Output: true
    
    Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
```

### eg5

``` 
    Input:
    
    s = "mississippi"
    
    p = "mis*is*p*."
    
    Output: false
```

### Difficulty

&emsp;&emsp;Hard

### Other

&emsp;&emsp;todo，添加解题思路。