# Evaluate Reverse Polish Notation

## Description

&emsp;&emsp;Evaluate the value of an arithmetic expression in **Reverse Polish Notation**.
            
&emsp;&emsp;Valid operators are \+, \-, \*, \/. Each operand may be an integer or another expression.

## Note

&emsp;&emsp;Division between two integers should truncate toward zero.
&emsp;&emsp;The given RPN expression is always valid. That means the expression would always evaluate to a result and 
there won't be any divide by zero operation.

## Example

### eg1

```
    Input: ["2", "1", "+", "3", "*"]
    
    Output: 9
    
    Explanation: ((2 + 1) * 3) = 9
```

### eg2

```
    Input: ["4", "13", "5", "/", "+"]
    
    Output: 6
    
    Explanation: (4 + (13 / 5)) = 6
```

### eg3

```
    Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
    
    Output: 22
    
    Explanation: 
      ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
    = ((10 * (6 / (12 * -11))) + 17) + 5
    = ((10 * (6 / -132)) + 17) + 5
    = ((10 * 0) + 17) + 5
    = (0 + 17) + 5
    = 17 + 5
    = 22
```

## Difficulty

&emsp;&emsp;Medium

## Other

&emsp;&emsp;todo，添加解题思路。