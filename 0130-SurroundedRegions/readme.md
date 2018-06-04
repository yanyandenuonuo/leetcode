# Surrounded Regions

## Description

&emsp;&emsp;Given a 2D board containing **'X'** and **'O'** \(the letter O\), capture all regions surrounded by 
**'X'**.

&emsp;&emsp;A region is captured by flipping all **'O'** s into **'X'** s in that surrounded region.

## Example

### eg1

```
    X X X X
    X O O X
    X X O X
    X O X X
```

&emsp;&emsp;After running your function, the board should be:

```
    X X X X
    X X X X
    X X X X
    X O X X
```

## Explanation

&emsp;&emsp;Surrounded regions shouldn't be on the border, which means that any **'O'** on the border of the board are 
not flipped to **'X'**. Any **'O'** that is not on the border and it is not connected to an **'O'** on the border 
will be flipped to **'X'**. Two cells are connected if they are adjacent cells connected horizontally or vertically.

## Difficulty

&emsp;&emsp;Medium

## Other

&emsp;&emsp;todo，添加解题思路。