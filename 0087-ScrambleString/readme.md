# Partition List

## Description

&emsp;&emsp;Given a string s1, we may represent it as a binary tree by partitioning it to two non-empty substrings 
recursively.

&emsp;&emsp;Below is one possible representation of s1 = **"great"**:

```
        great
       /    \
      gr    eat
     / \    /  \
    g   r  e   at
               / \
              a   t
```

&emsp;&emsp;To scramble the string, we may choose any non-leaf node and swap its two children.

&emsp;&emsp;For example, if we choose the node **"gr"** and swap its two children, it produces a scrambled string 
**"rgeat"**.

```
        rgeat
       /    \
      rg    eat
     / \    /  \
    r   g  e   at
               / \
              a   t
```

&emsp;&emsp;We say that **"rgeat"** is a scrambled string of **"great"**.

&emsp;&emsp;Similarly, if we continue to swap the children of nodes **"eat"** and **"at"**, it produces a scrambled 
string **"rgtae"**.

```
        rgtae
       /    \
      rg    tae
     / \    /  \
    r   g  ta  e
           / \
          t   a
```

&emsp;&emsp;We say that **"rgtae"** is a scrambled string of **"great"**.

&emsp;&emsp;Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.
            
## Example

### eg1

```
    Input: s1 = "great", s2 = "rgeat"
    
    Output: true
```

### eg2

```
    Input: s1 = "abcde", s2 = "caebd"
    
    Output: false
```

## Difficulty

&emsp;&emsp;Hard

## Other

&emsp;&emsp;todo，添加解题思路。