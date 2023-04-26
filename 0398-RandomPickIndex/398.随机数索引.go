/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] 随机数索引
 */

// @lc code=start
// solution1: 使用hash表记录num => [idx]的映射关系
type Solution struct {
	numMap map[int][]int
}

func Constructor(nums []int) Solution {
	numMap := make(map[int][]int, len(nums))
	for idx, num := range nums {
		if _, isExist := numMap[num]; !isExist {
			numMap[num] = make([]int, 0)
		}
		numMap[num] = append(numMap[num], idx)
	}
	return Solution{
		numMap: numMap,
	}
}

func (this *Solution) Pick(target int) int {
	return this.numMap[target][rand.Intn(len(this.numMap[target]))]
}

// solution2: 蓄水池抽样，可能存在超时
/**
type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Pick(target int) int {
	res, matchCount := 0, 0
	for idx, val := range this.nums {
		if val == target {
			matchCount += 1
			if rand.Intn(matchCount) == 0 {
				res = idx
			}
		}
	}
	return res
}
*/

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end

