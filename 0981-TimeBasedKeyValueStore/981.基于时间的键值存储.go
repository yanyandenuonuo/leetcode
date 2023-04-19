/*
 * @lc app=leetcode.cn id=981 lang=golang
 *
 * [981] 基于时间的键值存储
 */

// @lc code=start
type TimeMap struct {
	timeMap map[string][]valueObject
}

type valueObject struct {
	val       string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{
		timeMap: map[string][]valueObject{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, isExist := this.timeMap[key]; !isExist {
		this.timeMap[key] = make([]valueObject, 0)
	}
	this.timeMap[key] = append(this.timeMap[key],
		valueObject{
			val:       value,
			timestamp: timestamp,
		})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, isExist := this.timeMap[key]; !isExist || len(this.timeMap[key]) == 0 {
		return ""
	}

	if this.timeMap[key][0].timestamp > timestamp {
		return ""
	}

	if this.timeMap[key][len(this.timeMap[key])-1].timestamp <= timestamp {
		return this.timeMap[key][len(this.timeMap[key])-1].val
	}

	// 二分查找
	for leftIdx, rightIdx := 0, len(this.timeMap[key])-1; leftIdx <= rightIdx; {
		midIdx := leftIdx + (rightIdx-leftIdx)/2
		if timestamp < this.timeMap[key][midIdx].timestamp {
			rightIdx = midIdx - 1
		} else if this.timeMap[key][midIdx].timestamp < timestamp {
			leftIdx = midIdx + 1
		} else {
			return this.timeMap[key][midIdx].val
		}

		if leftIdx == rightIdx {
			if this.timeMap[key][leftIdx].timestamp <= timestamp {
				return this.timeMap[key][leftIdx].val
			} else {
				return this.timeMap[key][leftIdx-1].val
			}
		}
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end

