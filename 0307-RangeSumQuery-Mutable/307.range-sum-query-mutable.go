/*
 * @lc app=leetcode id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 */

// @lc code=start
// solution1: 线段树
// todo

// solution2: 树状数组
// todo

// solution3: 分块处理，存储原始数组和块的和，降低update影响的前缀和范围，每个size的大小取值可使用sqrt(len(nums))
type NumArray struct {
	nums []int
	sums []int
	size int
}

func Constructor(nums []int) NumArray {
	size := int(math.Sqrt(float64(len(nums))))
	sums := make([]int, 0, len(nums)/size+1)
	sum := 0
	for idx, num := range nums {
		if idx != 0 && idx%size == 0 {
			sums = append(sums, sum)
			sum = 0
		}
		sum += num
	}
	sums = append(sums, sum)

	return NumArray{
		nums: append([]int{}, nums...),
		sums: sums,
		size: size,
	}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.nums[index]
	this.nums[index] = val
	this.sums[index/this.size] += diff
}

func (this *NumArray) SumRange(left int, right int) int {
	// 存在几种case
	// [..., left, ...] [...] [..., right, ...]
	// [left, ...] [...] [..., right]
	// [..., left, ..., right,...]
	// [left, ..., right, ...]
	// [..., left, ..., right]
	// [..., left=right, ...]
	leftBlockSum := 0
	if left%this.size != 0 {
		for idx := left; idx < left/this.size*this.size+this.size && idx <= right; idx += 1 {
			leftBlockSum += this.nums[idx]
		}
	}

	rightBlockSum := 0
	if right%this.size != this.size-1 {
		for idx := right - right%this.size; idx >= left && idx <= right; idx += 1 {
			rightBlockSum += this.nums[idx]
		}
	}
	sum := leftBlockSum + rightBlockSum
	for idx := (left + this.size - 1) / this.size; idx < (right+1)/this.size; idx += 1 {
		sum += this.sums[idx]
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
// @lc code=end

