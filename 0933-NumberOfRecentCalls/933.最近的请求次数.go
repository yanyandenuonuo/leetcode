/*
 * @lc app=leetcode.cn id=933 lang=golang
 *
 * [933] 最近的请求次数
 */

// @lc code=start
type RecentCounter struct {
	reqList []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		reqList: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.reqList = append(this.reqList, t)

	for idx := 0; idx < len(this.reqList); {
		if t-this.reqList[idx] > 3000 {
			idx += 1
			continue
		}

		this.reqList = this.reqList[idx:]
		break
	}

	return len(this.reqList)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

